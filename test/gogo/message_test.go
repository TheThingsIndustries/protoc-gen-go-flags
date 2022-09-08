package test_test

import (
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-flags/test/gogo"
	pflag "github.com/spf13/pflag"
)

var testMessagesWithSemanticalMeaning = []struct {
	name            string
	arguments       []string
	expectedMessage SemanticalMessage
	expectedMask    []string
}{
	{
		name:      "exists",
		arguments: []string{"--empty"},
		expectedMessage: SemanticalMessage{
			Empty: &SemanticalMessage_Empty{},
		},
		expectedMask: []string{"empty"},
	},
	{
		name:      "doesn't exist",
		arguments: []string{},
		expectedMessage: SemanticalMessage{
			Empty:          nil,
			EmptyOverruled: nil,
			NonEmpty:       nil,
		},
		expectedMask: nil,
	},
	{
		name:      "overruled semantical",
		arguments: []string{"--empty_overruled"},
		expectedMessage: SemanticalMessage{
			Empty:          nil,
			EmptyOverruled: nil,
			NonEmpty:       nil,
		},
		expectedMask: nil,
	},
	{
		name:      "non_empty semantical",
		arguments: []string{"--non-empty"},
		expectedMessage: SemanticalMessage{
			NonEmpty: &SemanticalMessage_NonEmpty{},
		},
		expectedMask: []string{"non_empty"},
	},
	{
		name:      "non_empty with value",
		arguments: []string{"--non-empty.bool-value", "true"},
		expectedMessage: SemanticalMessage{
			NonEmpty: &SemanticalMessage_NonEmpty{
				BoolValue: true,
			},
		},
		expectedMask: []string{"non_empty.bool_value"},
	},
}

var testMessagesWithOneOfSemanticalMeaning = []struct {
	name            string
	arguments       []string
	expectedMessage SemanticalOneOfMessage
	expectedMask    []string
}{
	{
		name:      "alternative exists",
		arguments: []string{"--option.alternative"},
		expectedMessage: SemanticalOneOfMessage{
			Option: &SemanticalOneOfMessage_Alternative{
				Alternative: &SemanticalOneOfMessage_NonEmpty{},
			},
		},
		expectedMask: []string{"option.alternative"},
	},
	{
		name:      "doesn't exist",
		arguments: []string{},
		expectedMessage: SemanticalOneOfMessage{
			Option: nil,
		},
		expectedMask: nil,
	},
	{
		name:      "semantical exists",
		arguments: []string{"--option.semantical"},
		expectedMessage: SemanticalOneOfMessage{
			Option: &SemanticalOneOfMessage_Semantical{
				Semantical: &SemanticalOneOfMessage_Empty{},
			},
		},
		expectedMask: []string{"option.semantical"},
	},
	{
		name:      "alternative exists with value",
		arguments: []string{"--option.alternative.bool-value", "true"},
		expectedMessage: SemanticalOneOfMessage{
			Option: &SemanticalOneOfMessage_Alternative{
				Alternative: &SemanticalOneOfMessage_NonEmpty{
					BoolValue: true,
				},
			},
		},
		expectedMask: []string{"option.alternative.bool_value"},
	},
}

func TestSetFlagsMessageWithSemanticalMeaning(t *testing.T) {
	for _, tt := range testMessagesWithSemanticalMeaning {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForSemanticalMessage(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}

func TestSetFlagsMessageWithOneOfSemanticalMeaning(t *testing.T) {
	for _, tt := range testMessagesWithOneOfSemanticalMeaning {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForSemanticalOneOfMessage(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}
