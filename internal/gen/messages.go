// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"fmt"
	"strings"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	jsonannotations "github.com/TheThingsIndustries/protoc-gen-go-flags/internal/jsonplugin"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (g *generator) messageHasAnyFlags(message *protogen.Message) bool {
	if g.messageHasSelectFlags(message) || g.messageHasSetFlags(message) {
		return true
	}
	return false
}

func (g *generator) genMessage(message *protogen.Message) {
	// Generate flags for all sub-messages defined in the message.
	for _, message := range message.Messages {
		g.genMessage(message)
	}

	// Generate select flags for the message itself, if it has them.
	if g.messageHasSelectFlags(message) {
		g.genMessageSelectFlags(message)
		g.genFieldMaskSelectFromFlags(message)
	}

	// Generate set flags for the message itself, if it has them.
	if g.messageHasSetFlags(message) {
		g.genMessageSetFlags(message)
		g.genMessageSetterFromFlags(message)
	}
}

func fieldIsNullable(field *protogen.Field) bool {
	// Typically, only message fields are nullable (use pointers).
	nullable := field.Desc.Kind() == protoreflect.MessageKind
	return nullable
}

func fieldGoName(field *protogen.Field) interface{} {
	var fieldGoName interface{} = field.GoName
	return fieldGoName
}

func fieldCustomType(field *protogen.Field) *protogen.GoIdent {
	return nil
}

func flagFromCustomType(field *protogen.Field) *protogen.GoIdent {
	return nil
}

func ifThenElse(condition bool, ifTrue, ifFalse string) string {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// goTypeForField returns the name of the Go type that corresponds to the type of a given field.
func (g *generator) goTypeForField(field *protogen.Field) interface{} {
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return "bool"
	case protoreflect.EnumKind:
		return field.Enum.GoIdent
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64"
	case protoreflect.FloatKind:
		return "float32"
	case protoreflect.DoubleKind:
		return "float64"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BytesKind:
		return "[]byte"
	case protoreflect.MessageKind:
		return field.Message.GoIdent
	default:
		g.gen.Error(fmt.Errorf("unsupported field kind %q", field.Desc.Kind()))
		return ""
	}
}

// libNameForField returns the name used in the protojson func that corresponds to the type of a given field.
func (g *generator) libNameForField(field *protogen.Field) string {
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		return "Bool"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "Int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "Uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "Int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "Uint64"
	case protoreflect.FloatKind:
		return "Float32"
	case protoreflect.DoubleKind:
		return "Float64"
	case protoreflect.StringKind:
		return "String"
	case protoreflect.BytesKind:
		return "Bytes"
	case protoreflect.EnumKind:
		return "String"
	default:
		g.gen.Error(fmt.Errorf("unsupported field kind %q", field.Desc.Kind()))
		return ""
	}
}

// libNameForWKT strips google.protobuf prefix to get the custom flag type
func (g *generator) libNameForWKT(message *protogen.Message) string {
	switch message.Desc.FullName() {
	case "google.protobuf.FieldMask":
		return "StringSlice"
	default:
		return strings.TrimPrefix(string(message.Desc.FullName()), "google.protobuf.")
	}
}

// parseGoIdent parses a custom type and returns a GoIdent for it.
// If it's unable to parse the custom type, it returns nil.
func parseGoIdent(customtype string) *protogen.GoIdent {
	if customtype == "" {
		return nil
	}
	i := strings.LastIndex(customtype, ".")
	ident := protogen.GoImportPath(customtype[:i]).Ident(customtype[i+1:])
	return &ident
}

// messageIsWrapper returns true if the given message is a wrapper type.
// This is the case for well known wrapper types (google.protobuf.XXXValue)
// and for messages that have the (thethings.json.message) option with wrapper = true.
func messageIsWrapper(message *protogen.Message) bool {
	switch message.Desc.FullName() {
	case "google.protobuf.DoubleValue",
		"google.protobuf.FloatValue",
		"google.protobuf.Int64Value",
		"google.protobuf.UInt64Value",
		"google.protobuf.Int32Value",
		"google.protobuf.UInt32Value",
		"google.protobuf.BoolValue",
		"google.protobuf.StringValue",
		"google.protobuf.BytesValue":
		return true
	}
	opts := message.Desc.Options().(*descriptorpb.MessageOptions)
	if ext, hasExt := proto.GetExtension(opts, annotations.E_Message).(*annotations.MessageOptions); hasExt {
		return ext.GetWrapper() && len(message.Fields) == 1 && message.Fields[0].GoName == "Value"
	}
	if ext, hasExt := proto.GetExtension(opts, jsonannotations.E_Message).(*jsonannotations.MessageOptions); hasExt {
		return ext.GetWrapper() && len(message.Fields) == 1 && message.Fields[0].GoName == "Value"
	}
	return false
}

// messageIsWKT returns true if the given message is a well-known type.
func messageIsWKT(message *protogen.Message) bool {
	return strings.HasPrefix(string(message.Desc.FullName()), "google.protobuf.")
}

func isSupportedWKT(message *protogen.Message) bool {
	switch message.Desc.FullName() {
	case "google.protobuf.Duration",
		"google.protobuf.Timestamp",
		"google.protobuf.FieldMask":
		return true
	default:
		return false
	}
}

func isSupportedWKTSliceOrMap(message *protogen.Message) bool {
	switch message.Desc.FullName() {
	case "google.protobuf.Duration",
		"google.protobuf.Timestamp":
		return true
	default:
		return false
	}
}
