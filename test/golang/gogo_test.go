package test_test

import (
	"encoding/hex"
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-flags/test/golang"
	types "github.com/TheThingsIndustries/protoc-gen-go-flags/test/types"
	pflag "github.com/spf13/pflag"
)

var testMessagesWithGoGoOptions = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithGoGoOptions
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithGoGoOptions{},
		expectedMask:    nil,
	},
	{
		name: "full",
		arguments: []string{
			"--eui-with-custom-name", "AQIDBAUGBwg=",
			"--eui-with-custom-name-and-type", "0102030405060708",
			"--non-nullable-eui-with-custom-name-and-type", "0102030405060708",
			"--euis-with-custom-name-and-type",
			"0102030405060708,0102030405060707,0102030405060709",
		},
		expectedMessage: MessageWithGoGoOptions{
			EuiWithCustomName:                   []byte{1, 2, 3, 4, 5, 6, 7, 8},
			EuiWithCustomNameAndType:            []byte{1, 2, 3, 4, 5, 6, 7, 8},
			NonNullableEuiWithCustomNameAndType: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			EuisWithCustomNameAndType: [][]byte{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{1, 2, 3, 4, 5, 6, 7, 7},
				{1, 2, 3, 4, 5, 6, 7, 9},
			},
		},
		expectedMask: []string{
			"eui_with_custom_name",
			"eui_with_custom_name_and_type",
			"non_nullable_eui_with_custom_name_and_type",
			"euis_with_custom_name_and_type",
		},
	},
}

var testMessagesWithNullable = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithNullable
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithNullable{},
		expectedMask:    nil,
	},
}

func unmarshalEUI64(s string) *types.EUI64 {
	var eui64 types.EUI64
	b, _ := hex.DecodeString(s)
	eui64.Unmarshal(b)
	return &eui64
}

func TestSetFlagsMessageWithGoGoOptions(t *testing.T) {
	for _, tt := range testMessagesWithGoGoOptions {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithGoGoOptions(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}

func TestSetFlagsMessageWithNullable(t *testing.T) {
	for _, tt := range testMessagesWithNullable {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithNullable(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}
