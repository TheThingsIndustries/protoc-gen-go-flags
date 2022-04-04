// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"fmt"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	"github.com/TheThingsIndustries/protoc-gen-go-flags/internal/gogoproto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (g *generator) messageHasSetFlags(message *protogen.Message, visited ...*protogen.Message) bool {
	// No code is generated for map entries, so we also don't need to generate set flags.
	if message.Desc.IsMapEntry() {
		return false
	}

	var generateSetFlags bool

	for _, field := range message.Fields {
		// If the field has the (thethings.flags.field) option, and set is set, we need to generate set flags for the message.
		fieldOpts := field.Desc.Options().(*descriptorpb.FieldOptions)
		if proto.HasExtension(fieldOpts, annotations.E_Field) {
			if fieldExt, ok := proto.GetExtension(fieldOpts, annotations.E_Field).(*annotations.FieldOptions); ok {
				if fieldExt.Set == nil || fieldExt.GetSet() {
					generateSetFlags = true
				}
			}
		}
	}

	// Finally, the set field can still override to true or false if explicitly set.
	messageOpts := message.Desc.Options().(*descriptorpb.MessageOptions)
	if proto.HasExtension(messageOpts, annotations.E_Message) {
		if messageExt, ok := proto.GetExtension(messageOpts, annotations.E_Message).(*annotations.MessageOptions); ok {
			if messageExt.Set != nil {
				generateSetFlags = *messageExt.Set
			}
		}
	}

	return generateSetFlags
}

func (g *generator) genMessageSetFlags(message *protogen.Message) {
	g.P("// AddSetFlagsFor", message.GoIdent, " adds flags to select fields in ", message.GoIdent, ".")
	g.P("func AddSetFlagsFor", message.GoIdent, "(flags *", pflagPackage.Ident("FlagSet"), ", prefix string, hidden bool) ", " {")
nextField:
	for _, field := range message.Fields {
		_, setFlag, hidden := g.getFieldFlagBoolOptions(field)
		if setFlag != nil && !*setFlag {
			continue nextField
		}
		var customFlagType *protogen.GoIdent
		fieldOpts := field.Desc.Options()
		if proto.HasExtension(fieldOpts, annotations.E_Field) {
			customFlagType = parseGoIdent(proto.GetExtension(field.Desc.Options(), annotations.E_Field).(*annotations.FieldOptions).GetSetFlagNewFunc())
		}
		// Convert field name to flag name (underscore to dash).
		flagName := flagNameReplacer.Replace(string(field.Desc.Name()))
		// If field is oneof, add oneof field name to flag name.
		if field.Oneof != nil {
			flagName = flagNameReplacer.Replace(string(field.Oneof.Desc.Name())) + "." + flagName
		}
		if field.Desc.IsMap() {
			// If the field is a map, the field type is a MapEntry message.
			// In the MapEntry message, the first field is the key, and the second field is the value.
			key := field.Message.Fields[0]
			value := field.Message.Fields[1]

			if key.Desc.Kind() != protoreflect.StringKind {
				g.P("// FIXME: Skipping ", field.GoName, " because maps with ", key.Desc.Kind(), " key types are currently not supported.")
				continue nextField
			}

			switch value.Desc.Kind() {
			default:
				g.P("flags.AddFlag(", flagspluginPackage.Ident("NewString"+g.libNameForField(value)+"MapFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
			case protoreflect.EnumKind:
				g.P("// FIXME: Skipping ", field.GoName, " because maps with enum value types are currently not supported.")
			case protoreflect.MessageKind:
				switch {
				case messageIsWrapper(value.Message):
					wrappedField := value.Message.Fields[0]
					if wrappedField.Desc.Kind() == protoreflect.EnumKind {
						g.P("// FIXME: Skipping ", field.GoName, " because maps with enum value types are currently not supported.")
					} else {
						g.P("flags.AddFlag(", flagspluginPackage.Ident("NewString"+g.libNameForField(wrappedField)+"MapFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
					}
				// Add only flags for supported WKTs.
				case isSupportedWKTSliceOrMap(value.Message):
					g.P("flags.AddFlag(", flagspluginPackage.Ident("NewString"+g.libNameForWKT(value.Message)+"MapFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
				default:
					g.P("// FIXME: Skipping ", field.GoName, " because maps with message value types are currently not supported.")
				}
			}
			continue nextField
		}

		if field.Desc.IsList() {
			if customFlagType != nil {
				// If flag has a custom new flag definition, add this and continue with next field.
				g.P("flags.AddFlag(", *customFlagType, "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
				continue nextField
			}
			switch field.Desc.Kind() {
			default:
				g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(field)+"SliceFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
			case protoreflect.EnumKind:
				g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(field)+"SliceFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix),`, flagspluginPackage.Ident("EnumValueDesc("), field.Enum.GoIdent, "_value), ", flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
			case protoreflect.MessageKind:
				switch {
				case messageIsWrapper(field.Message):
					wrappedField := field.Message.Fields[0]
					if wrappedField.Desc.Kind() == protoreflect.EnumKind {
						// If a wrapped field is enum, include enum value description.
						g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(wrappedField)+"SliceFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, flagspluginPackage.Ident("EnumValueDesc("), wrappedField.Enum.GoIdent, "_value), ", flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
					} else {
						g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(wrappedField)+"SliceFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
					}

				case messageIsWKT(field.Message):
					if !isSupportedWKTSliceOrMap(field.Message) {
						g.P("// FIXME: Skipping ", field.GoName, " because this repeated WKT is currently not supported.")
						continue nextField
					}
					g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForWKT(field.Message)+"SliceFlag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")

				default:
					g.P("// FIXME: Skipping ", field.GoName, " because repeated messages are currently not supported.")
				}
			}
			continue nextField
		}
		if customFlagType != nil {
			// If flag has a custom new flag definition, add this and continue with next field.
			g.P("flags.AddFlag(", *customFlagType, "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
			continue nextField
		}
		switch field.Desc.Kind() {
		default:
			g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(field)+"Flag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
		case protoreflect.EnumKind:
			// If a field is enum, include enum value description.
			g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(field)+"Flag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, flagspluginPackage.Ident("EnumValueDesc("), field.Enum.GoIdent, "_value), ", flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
		case protoreflect.MessageKind:
			switch {
			case g.messageHasSetFlags(field.Message):
				// If the field is of type message, and the message has set flags, add those.
				g.P(field.Message.GoIdent.GoImportPath.Ident("AddSetFlagsFor"+field.Message.GoIdent.GoName), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, ifThenElse(hidden, "true", "hidden"), ")")
				if messageIsWrapper(field.Message) {
					// If the message is a wrapper, include the parent flag as an alias that points to the wrapped flag value.
					g.P(flagspluginPackage.Ident("AddAlias"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `.value", prefix), `, flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), ")")
				}
			case messageIsWrapper(field.Message):
				wrappedField := field.Message.Fields[0]
				if wrappedField.Desc.Kind() == protoreflect.EnumKind {
					// If a wrapped field is enum, include enum value description.
					g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(wrappedField)+"Flag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `.value", prefix), `, flagspluginPackage.Ident("EnumValueDesc("), field.Enum.GoIdent, "_value), ", flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
					// If the message is a wrapper, include the parent flag as an alias that points to the wrapped flag value.
					g.P(flagspluginPackage.Ident("AddAlias"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `.value", prefix), `, flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), ")")
					continue nextField
				}
				g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForField(wrappedField)+"Flag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `.value", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), `))`)
				// If the message is a wrapper, include the parent flag as an alias that points to the wrapped flag value.
				g.P(flagspluginPackage.Ident("AddAlias"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `.value", prefix), `, flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), ")")
			case messageIsWKT(field.Message):
				if !isSupportedWKT(field.Message) {
					g.P("// FIXME: Skipping ", field.GoName, " because this WKT is currently not supported.")
					continue nextField
				}
				g.P("flags.AddFlag(", flagspluginPackage.Ident("New"+g.libNameForWKT(field.Message)+"Flag"), "(", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix), "", `, flagspluginPackage.Ident("WithHidden"), ifThenElse(hidden, "(true)", "(hidden)"), "))")
			default:
				g.P("// FIXME: Skipping ", field.GoName, " because it does not seem to implement AddSetFlags.")
			}
		}
	}
	g.P("}")
	g.P()
}

func (g *generator) genMessageSetterFromFlags(message *protogen.Message) {
	g.P("// SetFromFlags sets the ", message.GoIdent, " message from flags.")
	g.P("func (m *", message.GoIdent, ") SetFromFlags(flags *", pflagPackage.Ident("FlagSet"), ", prefix string) (paths []string, err error) {")
nextField:
	for _, field := range message.Fields {
		_, setFlag, _ := g.getFieldFlagBoolOptions(field)
		if setFlag != nil && !*setFlag {
			continue nextField
		}
		var (
			fieldGoName  interface{} = fieldGoName(field)
			customtype               = fieldCustomType(field)
			nullable                 = fieldIsNullable(field)
			customGetter *protogen.GoIdent
		)
		fieldOpts := field.Desc.Options()
		// If customtype annotation is set, the getter must be of format Get{CustomFlag} (with Slice appended if list).
		if customtype != nil {
			customGetter = flagFromCustomType(field)
		} else if proto.HasExtension(fieldOpts, annotations.E_Field) {
			// Otherwise if custom getter set, use the custom getter instead of underlying proto field type.
			customGetter = parseGoIdent(proto.GetExtension(field.Desc.Options(), annotations.E_Field).(*annotations.FieldOptions).GetSetFlagGetterFunc())
		}
		flagName := field.Desc.Name()
		if field.Oneof != nil {
			// If field is oneof, add oneof field name to flag name.
			flagName = field.Oneof.Desc.Name() + "." + flagName
		}
		if field.Desc.IsMap() {
			// If custom getter is set for the field, use it to retrieve the flag value.
			if customGetter != nil {
				g.P("if val, changed, err := ", *customGetter, "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
				g.P("return nil, err")
				g.P("} else if ", "changed", "{")
				g.P("m", ".", fieldGoName, " =  val")
				// If flag is set, append the path to the paths to be returned.
				g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
				g.P("}")
				continue nextField
			}
			// If customtype is set for the field, and there is no custom getter, fail immediately.
			// This shouldn't happen because the getters are always generated for custom types whether they are there or not.
			if customtype != nil {
				g.gen.Error(fmt.Errorf("field with customtype %q doesn't have custom_flag_type set", message.Desc.FullName()))
				return
			}
			// If the field is a map, the field type is a MapEntry message.
			// In the MapEntry message, the first field is the key, and the second field is the value.
			key := field.Message.Fields[0]
			value := field.Message.Fields[1]

			if key.Desc.Kind() != protoreflect.StringKind {
				g.P("// FIXME: Skipping ", field.GoName, " because maps with ", key.Desc.Kind(), " key types are currently not supported.")
				continue nextField
			}
			switch value.Desc.Kind() {
			default:
				g.P("if val, changed, err := ", flagspluginPackage.Ident("GetString"+g.libNameForField(value)+"Map"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
				g.P("return nil, err")
				g.P("} else if ", "changed", "{")
				g.P("m", ".", fieldGoName, " = val")
				g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
				g.P("}")
			case protoreflect.MessageKind:
				switch {
				case messageIsWrapper(value.Message):
					// If value is a wrapped message, consider wrapped value type.
					wrappedField := value.Message.Fields[0]
					if wrappedField.Desc.Kind() == protoreflect.EnumKind {
						g.P("// FIXME: Skipping ", field.GoName, " because maps with enum value types are currently not supported.")
						continue nextField
					}
					g.P("if val, changed, err := ", flagspluginPackage.Ident("GetString"+g.libNameForField(wrappedField)+"Map"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
					g.P("return nil, err")
					g.P("} else if changed", "{")
					// If nullable is false, create a map without a pointer value.
					g.P("m.", fieldGoName, " = make(map[", g.goTypeForField(key), "]", ifThenElse(fieldIsNullable(value), "*", ""), g.goTypeForField(value), ")")
					g.P("for key, value := range val {")
					// If field is a wrapper, every value needs to be converted to the go type for field and assigned to the map.
					g.P("m", ".", fieldGoName, "[key] = ", ifThenElse(fieldIsNullable(value), "&", ""), g.goTypeForField(value), "{Value: value}")
					g.P("}")
					g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
					g.P("}")
				case isSupportedWKTSliceOrMap(value.Message):
					g.P("if val, changed, err := ", flagspluginPackage.Ident("GetString"+g.libNameForWKT(value.Message)+"Map"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
					g.P("return nil, err")
					g.P("} else if changed", "{")
					g.P("m.", fieldGoName, " = make(map[", g.goTypeForField(key), "]", ifThenElse(fieldIsNullable(value), "*", ""), g.goTypeForField(value), ")")
					g.P("for key, value := range val {")
					// If value is not a wrapper, but a supported WKT, convert the value to the proto type.
					g.P("m", ".", fieldGoName, "[key] = ", g.readWKTValue(value, value.Message, "value"))
					g.P("}")
					g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
					g.P("}")
				default:
					g.P("// FIXME: Skipping ", field.GoName, " because maps with message value types are currently not supported.")
				}
			case protoreflect.EnumKind:
				g.P("// FIXME: Skipping ", field.GoName, " because maps with enum value types are currently not supported.")
			}
			continue nextField
		}

		if field.Desc.IsList() {
			// If custom getter is set for the field, use it to retrieve the flag value.
			if customGetter != nil {
				g.P("if val, changed, err := ", *customGetter, "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
				g.P("return nil, err")
				g.P("} else if changed", "{")
				g.P("m", ".", fieldGoName, " = val")
				g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
				g.P("}")
				continue nextField
			}
			// If customtype is set for the field, and there is no custom getter, fail immediately.
			// This shouldn't happen because the getters are always generated for custom types whether they are there or not.
			if customtype != nil {
				g.gen.Error(fmt.Errorf("field with customtype %q doesn't have custom_flag_type set", message.Desc.FullName()))
				return
			}
			switch field.Desc.Kind() {
			default:
				// When getting slice flags, append `Slice` to the go flag getter.
				g.P("if val, changed, err := ", flagspluginPackage.Ident("Get"+g.libNameForField(field)+"Slice"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
				g.P("return nil, err")
				g.P("} else if changed", "{")
				if field.Desc.Kind() == protoreflect.EnumKind {
					g.P("for _, v := range val {")
					// If field is enum slice, we first obtain the string representation for every value,
					// then use `SetEnumString` and pass value map to return the int32 identifier for the enum.
					g.P("enumValue, err :=", flagspluginPackage.Ident("SetEnumString"), "(v, ", field.Enum.GoIdent, "_value)")
					g.ifErrNotNil()
					// Pass the int32 identifier to proto generated function to get the enum value.
					g.P("m", ".", fieldGoName, " = ", "append(", "m", ".", fieldGoName, ", ", field.Enum.GoIdent, "(enumValue))")
					g.P("}")
				} else {
					// Otherwise just assign the slice value to the struct field.
					g.P("m", ".", fieldGoName, " = val")
				}
				g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
				g.P("}")
			case protoreflect.MessageKind:
				switch {
				case messageIsWrapper(field.Message):
					// If message is a wrapper, we consider the underlying wrapped field type.
					wrappedField := field.Message.Fields[0]
					g.P("if val, changed, err := ", flagspluginPackage.Ident("Get"+g.libNameForField(wrappedField)+"Slice"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
					g.P("return nil, err")
					g.P("} else if changed", "{")
					g.P("for _, value := range val {")
					if wrappedField.Desc.Kind() == protoreflect.EnumKind {
						// If field is enum slice, we first obtain the string representation for every value,
						// then use `SetEnumString` and pass value map to return the int32 identifier for the enum.
						g.P("enumValue, err :=", flagspluginPackage.Ident("SetEnumString"), "(value, ", wrappedField.Enum.GoIdent, "_value)")
						g.ifErrNotNil()
						// For wrapped message we need to assign the value to the wrapped struct field `Value`.
						g.P("v := &", field.Message.GoIdent, "{Value: ", wrappedField.Enum.GoIdent, "(enumValue)}")
					} else {
						// For wrapped message we need to assign the value to the wrapped struct field `Value`.
						g.P("v := &", field.Message.GoIdent, "{Value: value}")
					}
					// We append each struct to a struct slice.
					g.P("m.", fieldGoName, " = ", "append(", "m", ".", fieldGoName, ", v)")
					g.P("}")
					g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
					g.P("}")
				case messageIsWKT(field.Message):
					// Currently we support only FieldMask, Timestamp, and Duration WKTs (other than the wrapped WKTs).
					if !isSupportedWKTSliceOrMap(field.Message) {
						g.P("// FIXME: Skipping ", field.GoName, " because this repeated WKT is not supported")
						continue nextField
					}
					g.P("if val, changed, err := ", flagspluginPackage.Ident("Get"+g.libNameForWKT(field.Message)+"Slice"), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
					g.P("return nil, err")
					g.P("} else if changed", "{")
					g.P("for _, value := range val {")
					// For every WKT value in WKT slice we convert it to the timestamp proto type.
					g.P("v := ", g.readWKTValue(field, field.Message, "value"))
					g.P("m.", fieldGoName, " = ", "append(", "m", ".", fieldGoName, ", v)")
					g.P("}")
					g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
					g.P("}")
				default:
					g.P("// FIXME: Skipping ", field.GoName, " because it does not seem to implement AddSetFlags.")
				}
			}
			continue nextField
		}

		// The identifier of the message is m, but in case of a oneof, we'll be operating on ov.
		messageOrOneofIdent := "m"

		// If this field is in a oneof, allocate a new oneof value wrapper.
		if field.Oneof != nil {
			messageOrOneofIdent = "ov"
		}
		// If custom getter is set for the field, use it to retrieve the flag value.
		if customGetter != nil {
			g.P("if val, changed, err := ", *customGetter, "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
			g.P("return nil, err")
			g.P("} else if changed", "{")
			if field.Oneof != nil {
				// If field is in a oneof, initialize an appropriate proto oneof type.
				g.P(messageOrOneofIdent, " := &", field.GoIdent.GoName, "{}")
			}
			// Assign value to the underlying field of the proto oneof type.
			if nullable {
				g.P(messageOrOneofIdent, ".", fieldGoName, " = &val")
			} else {
				g.P(messageOrOneofIdent, ".", fieldGoName, " = val")
			}
			// Append set paths.
			g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)
		} else if customtype != nil {
			g.gen.Error(fmt.Errorf("field with customtype %q doesn't have custom_flag_type set", message.Desc.FullName()))
			return
		} else {
			switch field.Desc.Kind() {
			default:
				g.P("if val, changed, err := ", flagspluginPackage.Ident("Get"+g.libNameForField(field)), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
				g.P("return nil, err")
				g.P("} else if changed", "{")
				// If field is in a oneof, initialize an appropriate proto oneof type.
				if field.Oneof != nil {
					g.P(messageOrOneofIdent, " := &", field.GoIdent.GoName, "{}")
				}
				// If field is enum , we first obtain the string representation for the value,
				// then use `SetEnumString` and pass value map to return the int32 identifier for the enum.
				if field.Desc.Kind() == protoreflect.EnumKind {
					g.P("enumValue, err :=", flagspluginPackage.Ident("SetEnumString"), "(val, ", field.Enum.GoIdent, "_value)")
					g.ifErrNotNil()
					g.P(messageOrOneofIdent, ".", fieldGoName, " = ", field.Enum.GoIdent, "(enumValue)")
				} else {
					g.P(messageOrOneofIdent, ".", fieldGoName, " = val")
				}
				g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)

			case protoreflect.MessageKind:
				// If message is a wrapper, we consider wrapped field type.
				fieldName := field
				if messageIsWrapper(field.Message) {
					fieldName = field.Message.Fields[0]
				}
				switch {
				case g.messageHasSetFlags(field.Message):
					// If field message has flag setter, we first check if any flags for the message are set.
					g.P("if changed := ", flagspluginPackage.Ident("IsAnyPrefixSet(flags, "), flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); changed {`)
					if field.Oneof != nil {
						g.P(messageOrOneofIdent, " := &", field.GoIdent.GoName, "{}")
					}
					if nullable {
						g.P(messageOrOneofIdent, ".", fieldGoName, " =&", field.Message.GoIdent.GoName, "{}")
					}
					// If any flags are set, we use the message setter to set the field, and obtain the set paths to return.
					g.P("if setPaths, err := ", messageOrOneofIdent, ".", fieldGoName, ".SetFromFlags(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
					g.P("return nil, err")
					g.P("} else {")
					g.P("paths = append(paths, setPaths...)")
					g.P("}")
				case messageIsWrapper(field.Message):
					// If the message is wrapper we get value directly from the {fieldName}.value flag.
					g.P("if val, changed, err := ", flagspluginPackage.Ident("Get"+g.libNameForField(fieldName)), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `.value", prefix)); err != nil {`)
					g.P("return nil, err")
					g.P("} else if changed", "{")
					if field.Oneof != nil {
						g.P(messageOrOneofIdent, " := &", field.GoIdent.GoName, "{}")
					}
					g.P(messageOrOneofIdent, ".", fieldGoName, " = &", field.Message.GoIdent, "{Value: val}")
					g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)

				case messageIsWKT(field.Message):
					// Currently we only support FieldMask, Duration and Timestamp WKT's, apart from wrapper types.
					if !isSupportedWKT(field.Message) {
						g.P("// FIXME: Skipping ", field.GoName, " because this WKT is not supported.")
						continue nextField
					}
					g.P("if val, changed, err := ", flagspluginPackage.Ident("Get"+g.libNameForWKT(fieldName.Message)), "(flags, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix)); err != nil {`)
					g.P("return nil, err")
					g.P("} else if changed", "{")
					if field.Oneof != nil {
						g.P(messageOrOneofIdent, " := &", field.GoIdent.GoName, "{}")
					}
					g.P(messageOrOneofIdent, ".", fieldGoName, " = ", g.readWKTValue(field, field.Message, "val"))
					g.P("paths = append(paths, ", flagspluginPackage.Ident("Prefix"), `("`, flagName, `", prefix))`)

				default:
					g.P("// FIXME: Skipping ", field.GoName, " because it does not seem to implement AddSetFlags.")
					continue nextField
				}
			}
		}
		if field.Oneof != nil {
			// Set the message field to the oneof wrapper.
			g.P("m.", field.Oneof.GoName, " = ov")
		} // end field if statement
		g.P("}")
	}
	g.P("return paths, nil")
	g.P("}")
}

// readWKTValue assigns a different proto field type to a WKT value based on the plugin used.
func (g *generator) readWKTValue(field *protogen.Field, message *protogen.Message, fieldName string) string {
	pluginPackage := golangPluginPackage
	if Params.Lang == "gogo" {
		pluginPackage = gogoPluginPackage
	}
	switch message.Desc.FullName() {
	case "google.protobuf.FieldMask":
		return g.QualifiedGoIdent(pluginPackage.Ident("SetFieldMask(" + fieldName + ")"))
	case "google.protobuf.Timestamp":
		if Params.Lang == "gogo" && proto.HasExtension(field.Desc.Options(), gogoproto.E_Stdtime) && proto.GetExtension(field.Desc.Options(), gogoproto.E_Stdtime).(bool) {
			return ifThenElse(fieldIsNullable(field), "&", "") + fieldName
		}
		return g.QualifiedGoIdent(pluginPackage.Ident("SetTimestamp(" + fieldName + ")"))
	case "google.protobuf.Duration":
		if Params.Lang == "gogo" && proto.HasExtension(field.Desc.Options(), gogoproto.E_Stdduration) && proto.GetExtension(field.Desc.Options(), gogoproto.E_Stdduration).(bool) {
			return ifThenElse(fieldIsNullable(field), "&", "") + fieldName
		}
		return g.QualifiedGoIdent(pluginPackage.Ident("SetDuration(" + fieldName + ")"))
	default:
		g.gen.Error(fmt.Errorf("unsupported WKT %q", message.Desc.FullName()))
		return ""
	}
}

func (g *generator) ifErrNotNil() {
	g.P("if err != nil {")
	g.P("return nil, err")
	g.P("}")
}
