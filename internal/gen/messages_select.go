// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (g *generator) messageHasSelectFlags(message *protogen.Message) bool {
	// No code is generated for map entries, so we also don't need to generate select flags.
	if message.Desc.IsMapEntry() {
		return false
	}

	var generateSelectFlags bool

	for _, field := range message.Fields {
		// If the field has the (thethings.flags.field) option, and select is set, we need to generate select flags for the message.
		fieldOpts := field.Desc.Options().(*descriptorpb.FieldOptions)
		if proto.HasExtension(fieldOpts, annotations.E_Field) {
			if fieldExt, ok := proto.GetExtension(fieldOpts, annotations.E_Field).(*annotations.FieldOptions); ok {
				selectFlag := fieldExt.Select
				if selectFlag == nil || fieldExt.GetSelect() {
					generateSelectFlags = true
				}
			}
		}
	}

	// Finally, the select field can still override to true or false if explicitly set.
	messageOpts := message.Desc.Options().(*descriptorpb.MessageOptions)
	if proto.HasExtension(messageOpts, annotations.E_Message) {
		if messageExt, ok := proto.GetExtension(messageOpts, annotations.E_Message).(*annotations.MessageOptions); ok {
			if messageExt.Select != nil {
				generateSelectFlags = *messageExt.Select
			}
		}
	}

	return generateSelectFlags
}

func (g *generator) messageHasCycle(message *protogen.Message, visited ...*protogen.Message) bool {
	// Since we're going to be looking at the fields of this message, it's possible that there will be cycles.
	// If that's the case, we'll return false here so that the caller can continue with the next field.
	for _, visited := range visited {
		if message == visited {
			return true
		}
	}
	for _, field := range message.Fields {
		selectFlag, _, _ := g.getFieldFlagBoolOptions(field)
		if field.Message != nil && g.messageHasCycle(field.Message, append(visited, message)...) && g.messageHasSelectFlags(field.Message) && selectFlag != nil && !*selectFlag {
			return true
		}
	}
	return false
}

func (g *generator) genMessageSelectFlags(message *protogen.Message) {
	g.P("// AddSelectFlagsFor", message.GoIdent, " adds flags to select fields in ", message.GoIdent, ".")
	g.P("func AddSelectFlagsFor", message.GoIdent, "(flags *", pflagPackage.Ident("FlagSet"), ", prefix string, hidden bool) ", " {")
nextField:
	for _, field := range message.Fields {
		selectFlag, _, hidden := g.getFieldFlagBoolOptions(field)
		if selectFlag != nil && !*selectFlag {
			continue nextField
		}
		flagName := flagNameReplacer.Replace(string(field.Desc.Name()))
		if field.Oneof != nil {
			// If field is oneof, add oneof field name to the flag name as prefix.
			flagName = flagNameReplacer.Replace(string(field.Oneof.Desc.Name())) + "." + flagName
		}
		if field.Message == nil || field.Desc.IsList() || field.Desc.IsMap() || messageIsWKT(field.Message) {
			g.P("flags.AddFlag(", flagspluginPackage.Ident("NewBoolFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, flagspluginPackage.Ident("SelectDesc("), flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), false), `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
			continue nextField
		}
		// If field is non-repeated message field, add bool flag with description that selecting it selects all subfields.
		g.P("flags.AddFlag(", flagspluginPackage.Ident("NewBoolFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, flagspluginPackage.Ident("SelectDesc("), flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), true), `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
		// If message has no select flags or creates a cycle, don't recursively add subfield flags.
		if !g.messageHasSelectFlags(field.Message) || g.messageHasCycle(field.Message) {
			g.P("// NOTE: ", field.Desc.Name(), " (", field.Message.GoIdent, ") does not seem to have select flags.")
			continue nextField
		}
		// Add select flags for subfields of the message.
		g.P(field.Message.GoIdent.GoImportPath.Ident("AddSelectFlagsFor"+field.Message.GoIdent.GoName), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, ifThenElse(hidden, "true", "hidden"), ")")
	}
	g.P("}")
	g.P()
}

func (g *generator) genFieldMaskSelectFromFlags(message *protogen.Message) {
	g.P("// SelectFromFlags outputs the fieldmask paths for", message.GoIdent, " message from select flags.")
	g.P("func PathsFromSelectFlagsFor", message.GoIdent, "(flags *", pflagPackage.Ident("FlagSet"), ", prefix string) (paths []string, err error) {")
nextField:
	for _, field := range message.Fields {
		selectFlag, _, _ := g.getFieldFlagBoolOptions(field)
		if selectFlag != nil && !*selectFlag {
			continue nextField
		}
		flagName := field.Desc.Name()
		// If field is oneof, add oneof field name to the flag name as prefix.
		if field.Oneof != nil {
			flagName = field.Oneof.Desc.Name() + "." + flagName
		}
		g.P("if val, selected, err := ", flagspluginPackage.Ident("GetBool"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
		g.P("return nil, err")
		g.P("} else if selected && val {")
		g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
		g.P("}")
		if field.Message == nil || field.Desc.IsList() || field.Desc.IsMap() || messageIsWKT(field.Message) {
			continue nextField
		}
		if !g.messageHasSelectFlags(field.Message) || g.messageHasCycle(field.Message) {
			g.P("// NOTE: ", field.Desc.Name(), " (", field.Message.GoIdent, ") does not seem to have select flags.")
			continue nextField
		}
		// If message has subfields and select flags are generated, add all set subpaths to select paths.
		g.P("if selectPaths, err := ", "PathsFromSelectFlagsFor", field.Message.GoIdent.GoName, "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
		g.P("return nil, err")
		g.P("} else {")
		g.P("paths = append(paths, selectPaths...)")
		g.P("}")
	}
	g.P("return paths, nil")
	g.P("}")
}
