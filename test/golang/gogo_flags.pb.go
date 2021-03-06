// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.20.1
// source: gogo.proto

package test

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	golang "github.com/TheThingsIndustries/protoc-gen-go-flags/golang"
	types "github.com/TheThingsIndustries/protoc-gen-go-flags/test/types"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForMessageWithGoGoOptions adds flags to select fields in MessageWithGoGoOptions.
func AddSelectFlagsForMessageWithGoGoOptions(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("eui-with-custom-name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("eui-with-custom-name", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("eui-with-custom-name-and-type", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("eui-with-custom-name-and-type", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("non-nullable-eui-with-custom-name-and-type", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("non-nullable-eui-with-custom-name-and-type", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("euis-with-custom-name-and-type", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("euis-with-custom-name-and-type", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("duration", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("duration", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("non-nullable-duration", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("non-nullable-duration", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("timestamp", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("timestamp", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("non-nullable-timestamp", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("non-nullable-timestamp", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forMessageWithGoGoOptions message from select flags.
func PathsFromSelectFlagsForMessageWithGoGoOptions(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("eui_with_custom_name", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("eui_with_custom_name", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("eui_with_custom_name_and_type", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("eui_with_custom_name_and_type", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("non_nullable_eui_with_custom_name_and_type", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("non_nullable_eui_with_custom_name_and_type", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("euis_with_custom_name_and_type", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("euis_with_custom_name_and_type", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("duration", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("duration", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("non_nullable_duration", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("non_nullable_duration", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("timestamp", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("timestamp", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("non_nullable_timestamp", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("non_nullable_timestamp", prefix))
	}
	return paths, nil
}

// AddSetFlagsForMessageWithGoGoOptions adds flags to select fields in MessageWithGoGoOptions.
func AddSetFlagsForMessageWithGoGoOptions(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBytesFlag(flagsplugin.Prefix("eui-with-custom-name", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(types.New8BytesFlag(flagsplugin.Prefix("eui-with-custom-name-and-type", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(types.New8BytesFlag(flagsplugin.Prefix("non-nullable-eui-with-custom-name-and-type", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(types.New8BytesSliceFlag(flagsplugin.Prefix("euis-with-custom-name-and-type", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewDurationFlag(flagsplugin.Prefix("duration", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewDurationFlag(flagsplugin.Prefix("non-nullable-duration", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("timestamp", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("non-nullable-timestamp", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the MessageWithGoGoOptions message from flags.
func (m *MessageWithGoGoOptions) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetBytes(flags, flagsplugin.Prefix("eui_with_custom_name", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.EuiWithCustomName = val
		paths = append(paths, flagsplugin.Prefix("eui_with_custom_name", prefix))
	}
	if val, changed, err := types.GetExactBytes(flags, flagsplugin.Prefix("eui_with_custom_name_and_type", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.EuiWithCustomNameAndType = val
		paths = append(paths, flagsplugin.Prefix("eui_with_custom_name_and_type", prefix))
	}
	if val, changed, err := types.GetExactBytes(flags, flagsplugin.Prefix("non_nullable_eui_with_custom_name_and_type", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NonNullableEuiWithCustomNameAndType = val
		paths = append(paths, flagsplugin.Prefix("non_nullable_eui_with_custom_name_and_type", prefix))
	}
	if val, changed, err := types.GetExactBytesSlice(flags, flagsplugin.Prefix("euis_with_custom_name_and_type", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.EuisWithCustomNameAndType = val
		paths = append(paths, flagsplugin.Prefix("euis_with_custom_name_and_type", prefix))
	}
	if val, changed, err := flagsplugin.GetDuration(flags, flagsplugin.Prefix("duration", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Duration = golang.SetDuration(val)
		paths = append(paths, flagsplugin.Prefix("duration", prefix))
	}
	if val, changed, err := flagsplugin.GetDuration(flags, flagsplugin.Prefix("non_nullable_duration", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NonNullableDuration = golang.SetDuration(val)
		paths = append(paths, flagsplugin.Prefix("non_nullable_duration", prefix))
	}
	if val, changed, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("timestamp", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Timestamp = golang.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("timestamp", prefix))
	}
	if val, changed, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("non_nullable_timestamp", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.NonNullableTimestamp = golang.SetTimestamp(val)
		paths = append(paths, flagsplugin.Prefix("non_nullable_timestamp", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForSubMessage adds flags to select fields in SubMessage.
func AddSelectFlagsForSubMessage(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("field", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("field", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forSubMessage message from select flags.
func PathsFromSelectFlagsForSubMessage(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("field", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("field", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSubMessage adds flags to select fields in SubMessage.
func AddSetFlagsForSubMessage(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("field", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SubMessage message from flags.
func (m *SubMessage) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("field", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.Field = val
		paths = append(paths, flagsplugin.Prefix("field", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForSubMessageWithoutMarshalers adds flags to select fields in SubMessageWithoutMarshalers.
func AddSelectFlagsForSubMessageWithoutMarshalers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("other-field", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("other-field", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forSubMessageWithoutMarshalers message from select flags.
func PathsFromSelectFlagsForSubMessageWithoutMarshalers(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("other_field", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("other_field", prefix))
	}
	return paths, nil
}

// AddSetFlagsForSubMessageWithoutMarshalers adds flags to select fields in SubMessageWithoutMarshalers.
func AddSetFlagsForSubMessageWithoutMarshalers(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("other-field", prefix), "", flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the SubMessageWithoutMarshalers message from flags.
func (m *SubMessageWithoutMarshalers) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("other_field", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.OtherField = val
		paths = append(paths, flagsplugin.Prefix("other_field", prefix))
	}
	return paths, nil
}

// AddSelectFlagsForMessageWithNullable adds flags to select fields in MessageWithNullable.
func AddSelectFlagsForMessageWithNullable(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("sub", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("sub", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("subs", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("subs", prefix), false), flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("other-sub", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("other-sub", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other-sub", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("other-subs", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("other-subs", prefix), false), flagsplugin.WithHidden(hidden)))
}

// SelectFromFlags outputs the fieldmask paths forMessageWithNullable message from select flags.
func PathsFromSelectFlagsForMessageWithNullable(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("sub", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("subs", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("subs", prefix))
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("other_sub", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("other_subs", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("other_subs", prefix))
	}
	return paths, nil
}

// AddSetFlagsForMessageWithNullable adds flags to select fields in MessageWithNullable.
func AddSetFlagsForMessageWithNullable(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix), hidden)
	// FIXME: Skipping Subs because repeated messages are currently not supported.
	AddSetFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other-sub", prefix), hidden)
	// FIXME: Skipping OtherSubs because repeated messages are currently not supported.
}

// SetFromFlags sets the MessageWithNullable message from flags.
func (m *MessageWithNullable) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("sub", prefix)); changed {
		if m.Sub == nil {
			m.Sub = &SubMessage{}
		}
		if setPaths, err := m.Sub.SetFromFlags(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	// FIXME: Skipping Subs because it does not seem to implement AddSetFlags.
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("other_sub", prefix)); changed {
		if m.OtherSub == nil {
			m.OtherSub = &SubMessageWithoutMarshalers{}
		}
		if setPaths, err := m.OtherSub.SetFromFlags(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	// FIXME: Skipping OtherSubs because it does not seem to implement AddSetFlags.
	return paths, nil
}

// AddSelectFlagsForMessageWithEmbedded adds flags to select fields in MessageWithEmbedded.
func AddSelectFlagsForMessageWithEmbedded(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("sub", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("sub", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("other-sub", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("other-sub", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other-sub", prefix), hidden)
}

// SelectFromFlags outputs the fieldmask paths forMessageWithEmbedded message from select flags.
func PathsFromSelectFlagsForMessageWithEmbedded(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("sub", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("other_sub", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForMessageWithEmbedded adds flags to select fields in MessageWithEmbedded.
func AddSetFlagsForMessageWithEmbedded(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix), hidden)
	AddSetFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other-sub", prefix), hidden)
}

// SetFromFlags sets the MessageWithEmbedded message from flags.
func (m *MessageWithEmbedded) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("sub", prefix)); changed {
		if m.Sub == nil {
			m.Sub = &SubMessage{}
		}
		if setPaths, err := m.Sub.SetFromFlags(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("other_sub", prefix)); changed {
		if m.OtherSub == nil {
			m.OtherSub = &SubMessageWithoutMarshalers{}
		}
		if setPaths, err := m.OtherSub.SetFromFlags(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	return paths, nil
}

// AddSelectFlagsForMessageWithNullableEmbedded adds flags to select fields in MessageWithNullableEmbedded.
func AddSelectFlagsForMessageWithNullableEmbedded(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("sub", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("sub", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix), hidden)
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("other-sub", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("other-sub", prefix), true), flagsplugin.WithHidden(hidden)))
	AddSelectFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other-sub", prefix), hidden)
}

// SelectFromFlags outputs the fieldmask paths forMessageWithNullableEmbedded message from select flags.
func PathsFromSelectFlagsForMessageWithNullableEmbedded(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("sub", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	if val, selected, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
		return nil, err
	} else if selected && val {
		paths = append(paths, flagsplugin.Prefix("other_sub", prefix))
	}
	if selectPaths, err := PathsFromSelectFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
		return nil, err
	} else {
		paths = append(paths, selectPaths...)
	}
	return paths, nil
}

// AddSetFlagsForMessageWithNullableEmbedded adds flags to select fields in MessageWithNullableEmbedded.
func AddSetFlagsForMessageWithNullableEmbedded(flags *pflag.FlagSet, prefix string, hidden bool) {
	AddSetFlagsForSubMessage(flags, flagsplugin.Prefix("sub", prefix), hidden)
	AddSetFlagsForSubMessageWithoutMarshalers(flags, flagsplugin.Prefix("other-sub", prefix), hidden)
}

// SetFromFlags sets the MessageWithNullableEmbedded message from flags.
func (m *MessageWithNullableEmbedded) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("sub", prefix)); changed {
		if m.Sub == nil {
			m.Sub = &SubMessage{}
		}
		if setPaths, err := m.Sub.SetFromFlags(flags, flagsplugin.Prefix("sub", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	if changed := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("other_sub", prefix)); changed {
		if m.OtherSub == nil {
			m.OtherSub = &SubMessageWithoutMarshalers{}
		}
		if setPaths, err := m.OtherSub.SetFromFlags(flags, flagsplugin.Prefix("other_sub", prefix)); err != nil {
			return nil, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	return paths, nil
}
