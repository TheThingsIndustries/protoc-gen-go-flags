package test_test

import (
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-flags/test/golang"
	pflag "github.com/spf13/pflag"
)

var testMessagesWithSemanticalMeaning = []struct {
	name            string
	arguments       []string
	expectedMessage NonSemantic
	expectedMask    []string
}{
	{
		name:      "exists",
		arguments: []string{"--semantic"},
		expectedMessage: NonSemantic{
			Semantic: &NonSemantic_Semantic{},
		},
		expectedMask: []string{"semantic"},
	},
	{
		name:      "doesn't exist",
		arguments: []string{},
		expectedMessage: NonSemantic{
			Semantic: nil,
		},
		expectedMask: nil,
	},
}

func TestSetFlagsMessageWithSemanticalMeaning(t *testing.T) {
	for _, tt := range testMessagesWithSemanticalMeaning {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForNonSemantic(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}
