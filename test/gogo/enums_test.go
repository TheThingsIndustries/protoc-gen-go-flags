package test_test

import (
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-flags/test/gogo"
	pflag "github.com/spf13/pflag"
)

var testMessagesWithEnums = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithEnums
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithEnums{},
		expectedMask:    nil,
	},
	{
		name:      "one field set",
		arguments: []string{"--regular", "1", ""},
		expectedMessage: MessageWithEnums{
			Regular: REGULAR_A,
		},
		expectedMask: []string{"regular"},
	},
	{
		name: "all fields set",
		arguments: []string{
			"--regular", "1",
			"--regulars", "1",
			"--regulars", "REGULAR_B",
			"--custom", "CUSTOM_V1_0",
			"--customs", "0,1",
			"--wrapped-custom", "1",
			"--wrapped-customs", "1,2",
		},
		expectedMessage: MessageWithEnums{
			Regular:        REGULAR_A,
			Regulars:       []RegularEnum{REGULAR_A, REGULAR_B},
			Custom:         CustomEnum_CUSTOM_V1_0,
			Customs:        []CustomEnum{CustomEnum_CUSTOM_UNKNOWN, CustomEnum_CUSTOM_V1_0},
			WrappedCustom:  &CustomEnumValue{Value: CustomEnum_CUSTOM_V1_0},
			WrappedCustoms: []*CustomEnumValue{{Value: CustomEnum_CUSTOM_V1_0}, {Value: CustomEnum_CUSTOM_V1_0_1}},
		},
		expectedMask: []string{
			"regular",
			"regulars",
			"custom",
			"customs",
			"wrapped_custom.value",
			"wrapped_customs",
		},
	},
	{
		name: "alias custom enum value set",
		arguments: []string{
			"--custom", "1.0.0",
			"--customs", "1.0.1,V1_0_1,1.0,unknown",
			"--wrapped-custom", "UNKNOWN",
			"--wrapped-customs", "unknown,1.0",
		},
		expectedMessage: MessageWithEnums{
			Custom:         CustomEnum_CUSTOM_V1_0,
			Customs:        []CustomEnum{CustomEnum_CUSTOM_V1_0_1, CustomEnum_CUSTOM_V1_0_1, CustomEnum_CUSTOM_V1_0, CustomEnum_CUSTOM_UNKNOWN},
			WrappedCustom:  &CustomEnumValue{Value: CustomEnum_CUSTOM_UNKNOWN},
			WrappedCustoms: []*CustomEnumValue{{Value: CustomEnum_CUSTOM_UNKNOWN}, {Value: CustomEnum_CUSTOM_V1_0}},
		},
		expectedMask: []string{
			"custom",
			"customs",
			"wrapped_custom.value",
			"wrapped_customs",
		},
	},
}

var testMessagesWithOneofEnums = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithOneofEnums
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithOneofEnums{},
		expectedMask:    nil,
	},
	{
		name: "field set",
		arguments: []string{
			"--value.regular", "1",
		},
		expectedMessage: MessageWithOneofEnums{
			Value: &MessageWithOneofEnums_Regular{Regular: REGULAR_A},
		},
		expectedMask: []string{
			"value.regular",
		},
	},
	{
		name: "all fields set",
		arguments: []string{
			"--value.wrapped-custom", "1",
		},
		expectedMessage: MessageWithOneofEnums{
			Value: &MessageWithOneofEnums_WrappedCustom{WrappedCustom: &CustomEnumValue{Value: CustomEnum_CUSTOM_V1_0}},
		},
		expectedMask: []string{
			"value.wrapped_custom.value",
		},
	},
	{
		name: "multiple fields set, last one overwriting others",
		arguments: []string{
			"--value.regular", "1",
			"--value.custom", "1",
			"--value.wrapped-custom", "1",
		},
		expectedMessage: MessageWithOneofEnums{
			Value: &MessageWithOneofEnums_WrappedCustom{WrappedCustom: &CustomEnumValue{Value: CustomEnum_CUSTOM_V1_0}},
		},
		expectedMask: []string{
			"value.regular",
			"value.custom",
			"value.wrapped_custom.value",
		},
	},
}

func TestSetFlagsMessageWithEnums(t *testing.T) {
	for _, tt := range testMessagesWithEnums {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithEnums(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}

func TestSetFlagsMessageWithOneofEnums(t *testing.T) {
	for _, tt := range testMessagesWithOneofEnums {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithOneofEnums(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}
