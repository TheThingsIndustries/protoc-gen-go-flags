package test_test

import (
	"reflect"
	"testing"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/google/go-cmp/cmp"
	pflag "github.com/spf13/pflag"
	"google.golang.org/protobuf/encoding/prototext"
	proto "google.golang.org/protobuf/proto"
)

func setFromFlags(t *testing.T, msg proto.Message, flagSet *pflag.FlagSet) []string {
	t.Helper()

	message, ok := msg.(flagsplugin.SetterFromFlags)
	if !ok {
		t.Fatalf("message %T does not implement the flagsplugin.SetFromFlags", msg)
	}
	paths, err := message.SetFromFlags(flagSet, "")
	if err != nil {
		t.Fatalf("unexpected error from SetFromFlags for %T: %v", msg, err)
	}
	return paths
}

func expectMessageEqual(t *testing.T, flagSet *pflag.FlagSet, arguments []string, expectedMessage proto.Message, expectedMask []string) {
	t.Helper()
	if expectedMessage == nil {
		return
	}
	expectedMsgText := prototext.Format(expectedMessage)
	generatedMessage := reflect.New(reflect.ValueOf(expectedMessage).Elem().Type()).Interface().(proto.Message)

	flagSet.Parse(arguments)
	mask := setFromFlags(t, generatedMessage, flagSet)
	generatedMsgText := prototext.Format(generatedMessage)

	generatedDiff := cmp.Diff(expectedMsgText, generatedMsgText)
	maskDiff := cmp.Diff(expectedMask, mask)

	if generatedDiff != "" {
		t.Errorf("expected : %s", string(expectedMsgText))
		t.Errorf("generated: %s", string(generatedMsgText))
		if generatedDiff != "" {
			t.Errorf("  diff   : %s", generatedDiff)
		}
	}

	if maskDiff != "" {
		t.Errorf("mask diff: %s", maskDiff)
	}
}
