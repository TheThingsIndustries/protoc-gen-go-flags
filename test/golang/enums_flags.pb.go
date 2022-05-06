// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.20.1
// source: enums.proto

package test

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	enum "github.com/TheThingsIndustries/protoc-gen-go-flags/test/enum"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForCustomEnumValue adds flags to select fields in CustomEnumValue.
func AddSelectFlagsForCustomEnumValue(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("value", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("value", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forCustomEnumValue message from select flags.
func PathsFromSelectFlagsForCustomEnumValue(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("value", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("value", prefix))
	}
	return paths, nil
}

// AddSetFlagsForCustomEnumValue adds flags to select fields in CustomEnumValue.
func AddSetFlagsForCustomEnumValue(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("value", prefix), flagsplugin.EnumValueDesc(CustomEnum_value, enum.CustomEnum_customvalue), flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the CustomEnumValue message from flags.
func (m *CustomEnumValue) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("value", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, CustomEnum_value, enum.CustomEnum_customvalue)
		if err != nil {
			return nil, err
		}
		m.Value = CustomEnum(enumValue)
		paths = append(paths, flagsplugin.Prefix("value", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForMessageWithEnums adds flags to select fields in MessageWithEnums.
func AddSelectFlagsForMessageWithEnums(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("regular", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("regular", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("regulars", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("regulars", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("custom", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("custom", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("customs", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("customs", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("wrapped-custom", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("wrapped-custom", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForCustomEnumValue(flags, flagsplugin.Prefix("wrapped-custom", prefix), true)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("wrapped-customs", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("wrapped-customs", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forMessageWithEnums message from select flags.
func PathsFromSelectFlagsForMessageWithEnums(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("regular", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("regular", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("regulars", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("regulars", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("custom", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("custom", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("customs", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("customs", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("wrapped_custom", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("wrapped_custom", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForCustomEnumValue(flags, flagsplugin.Prefix("wrapped_custom", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("wrapped_customs", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("wrapped_customs", prefix))
	}
	return paths, nil
}

// AddSetFlagsForMessageWithEnums adds flags to select fields in MessageWithEnums.
func AddSetFlagsForMessageWithEnums(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("regular", prefix), flagsplugin.EnumValueDesc(RegularEnum_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("regulars", prefix), flagsplugin.EnumValueDesc(RegularEnum_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("custom", prefix), flagsplugin.EnumValueDesc(CustomEnum_value, enum.CustomEnum_customvalue), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("customs", prefix), flagsplugin.EnumValueDesc(CustomEnum_value, enum.CustomEnum_customvalue), flagsplugin.WithHidden(hidden)))
	AddSetFlagsForCustomEnumValue(flags, flagsplugin.Prefix("wrapped-custom", prefix), true)
	flagsplugin.AddAlias(flags, flagsplugin.Prefix("wrapped-custom.value", prefix), flagsplugin.Prefix("wrapped-custom", prefix), flagsplugin.WithHidden(hidden))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("wrapped-customs", prefix), flagsplugin.EnumValueDesc(CustomEnum_value, enum.CustomEnum_customvalue), flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the MessageWithEnums message from flags.
func (m *MessageWithEnums) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("regular", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, RegularEnum_value)
		if err != nil {
			return nil, err
		}
		m.Regular = RegularEnum(enumValue)
		paths = append(paths, flagsplugin.Prefix("regular", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("regulars", prefix)); err != nil {
		return nil, err
	} else if changed {
		for _, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, RegularEnum_value)
			if err != nil {
				return nil, err
			}
			m.Regulars = append(m.Regulars, RegularEnum(enumValue))
		}
		paths = append(paths, flagsplugin.Prefix("regulars", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("custom", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, CustomEnum_value, enum.CustomEnum_customvalue)
		if err != nil {
			return nil, err
		}
		m.Custom = CustomEnum(enumValue)
		paths = append(paths, flagsplugin.Prefix("custom", prefix))
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("customs", prefix)); err != nil {
		return nil, err
	} else if changed {
		for _, v := range val {
			enumValue, err := flagsplugin.SetEnumString(v, CustomEnum_value, enum.CustomEnum_customvalue)
			if err != nil {
				return nil, err
			}
			m.Customs = append(m.Customs, CustomEnum(enumValue))
		}
		paths = append(paths, flagsplugin.Prefix("customs", prefix))
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("wrapped_custom", prefix)); changed {
		m.WrappedCustom = &CustomEnumValue{}
		if setPaths, err := m.WrappedCustom.SetFromFlags(flags, flagsplugin.Prefix("wrapped_custom", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if val, changed, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("wrapped_customs", prefix)); err != nil {
		return nil, err
	} else if changed {
		for _, value := range val {
			enumValue, err := flagsplugin.SetEnumString(value, CustomEnum_value, enum.CustomEnum_customvalue)
			if err != nil {
				return nil, err
			}
			v := &CustomEnumValue{Value: CustomEnum(enumValue)}
			m.WrappedCustoms = append(m.WrappedCustoms, v)
		}
		paths = append(paths, flagsplugin.Prefix("wrapped_customs", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForMessageWithOneofEnums adds flags to select fields in MessageWithOneofEnums.
func AddSelectFlagsForMessageWithOneofEnums(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("value.regular", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("value.regular", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("value.custom", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("value.custom", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("value.wrapped-custom", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("value.wrapped-custom", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForCustomEnumValue(flags, flagsplugin.Prefix("value.wrapped-custom", prefix), true)
}

// SelectFromFlags outputs the fieldmask paths forMessageWithOneofEnums message from select flags.
func PathsFromSelectFlagsForMessageWithOneofEnums(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("value.regular", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("value.regular", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("value.custom", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("value.custom", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("value.wrapped_custom", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("value.wrapped_custom", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForCustomEnumValue(flags, flagsplugin.Prefix("value.wrapped_custom", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForMessageWithOneofEnums adds flags to select fields in MessageWithOneofEnums.
func AddSetFlagsForMessageWithOneofEnums(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("value.regular", prefix), flagsplugin.EnumValueDesc(RegularEnum_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("value.custom", prefix), flagsplugin.EnumValueDesc(CustomEnum_value, enum.CustomEnum_customvalue), flagsplugin.WithHidden(hidden)))
	AddSetFlagsForCustomEnumValue(flags, flagsplugin.Prefix("value.wrapped-custom", prefix), true)
	flagsplugin.AddAlias(flags, flagsplugin.Prefix("value.wrapped-custom.value", prefix), flagsplugin.Prefix("value.wrapped-custom", prefix), flagsplugin.WithHidden(hidden))
}

// SetFromFlags sets the MessageWithOneofEnums message from flags.
func (m *MessageWithOneofEnums) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("value.regular", prefix)); err != nil {
		return nil, err
	} else if changed {
		ov := &MessageWithOneofEnums_Regular{}
		enumValue, err := flagsplugin.SetEnumString(val, RegularEnum_value)
		if err != nil {
			return nil, err
		}
		ov.Regular = RegularEnum(enumValue)
		paths = append(paths, flagsplugin.Prefix("value.regular", prefix))
		m.Value = ov
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("value.custom", prefix)); err != nil {
		return nil, err
	} else if changed {
		ov := &MessageWithOneofEnums_Custom{}
		enumValue, err := flagsplugin.SetEnumString(val, CustomEnum_value, enum.CustomEnum_customvalue)
		if err != nil {
			return nil, err
		}
		ov.Custom = CustomEnum(enumValue)
		paths = append(paths, flagsplugin.Prefix("value.custom", prefix))
		m.Value = ov
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("value.wrapped_custom", prefix)); changed {
		ov := &MessageWithOneofEnums_WrappedCustom{}
		ov.WrappedCustom = &CustomEnumValue{}
		if setPaths, err := ov.WrappedCustom.SetFromFlags(flags, flagsplugin.Prefix("value.wrapped_custom", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
		m.Value = ov
	}
	return paths, nil
}

// AddSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum adds flags to select fields in MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum.
func AddSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("string-value", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("string-value", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("test-enum-value", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("test-enum-value", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("test-another-enum-value", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("test-another-enum-value", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum message from select flags.
func PathsFromSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("string_value", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("string_value", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("test_enum_value", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("test_enum_value", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("test_another_enum_value", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("test_another_enum_value", prefix))
	}
	return paths, nil
}

// AddSetFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum adds flags to select fields in MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum.
func AddSetFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("string-value", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("test-enum-value", prefix), flagsplugin.EnumValueDesc(MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_value), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("test-another-enum-value", prefix), flagsplugin.EnumValueDesc(MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_value), flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum message from flags.
func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("string_value", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.StringValue = val
		paths = append(paths, flagsplugin.Prefix("string_value", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("test_enum_value", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_value)
		if err != nil {
			return nil, err
		}
		m.TestEnumValue = MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum(enumValue)
		paths = append(paths, flagsplugin.Prefix("test_enum_value", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("test_another_enum_value", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_value)
		if err != nil {
			return nil, err
		}
		m.TestAnotherEnumValue = MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum(enumValue)
		paths = append(paths, flagsplugin.Prefix("test_another_enum_value", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums adds flags to select fields in MessageWithEmbeddedMessageDefinitionWithEnums.
func AddSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("test-message-field", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("test-message-field", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum(flags, flagsplugin.Prefix("test-message-field", prefix), hidden)
}

// SelectFromFlags outputs the fieldmask paths forMessageWithEmbeddedMessageDefinitionWithEnums message from select flags.
func PathsFromSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("test_message_field", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("test_message_field", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum(flags, flagsplugin.Prefix("test_message_field", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForMessageWithEmbeddedMessageDefinitionWithEnums adds flags to select fields in MessageWithEmbeddedMessageDefinitionWithEnums.
func AddSetFlagsForMessageWithEmbeddedMessageDefinitionWithEnums(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForMessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum(flags, flagsplugin.Prefix("test-message-field", prefix), hidden)
}

// SetFromFlags sets the MessageWithEmbeddedMessageDefinitionWithEnums message from flags.
func (m *MessageWithEmbeddedMessageDefinitionWithEnums) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("test_message_field", prefix)); changed {
		m.TestMessageField = &MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum{}
		if setPaths, err := m.TestMessageField.SetFromFlags(flags, flagsplugin.Prefix("test_message_field", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	return paths, nil
}
