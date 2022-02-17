package test_test

import (
	"fmt"
	"strings"
	"testing"
	time "time"

	. "github.com/TheThingsIndustries/protoc-gen-go-flags/test/golang"
	pflag "github.com/spf13/pflag"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	now          = time.Now()
	testDuration = time.Hour + 2*time.Minute + 3*time.Second + 123456789
)

var testMessagesWithWKTs = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithWKTs
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithWKTs{},
		expectedMask:    nil,
	},
	{
		name: "single-values",
		arguments: []string{
			"--double-value", "432.1423412",
			"--float-value", "23.324132",
			"--int32-value", "32142",
			"--int64-value", "32142",
			"--uint32-value", "32142",
			"--uint64-value", "32142",
			"--bool-value",
			"--string-value", "abc123",
			"--bytes-value", "AQIDBAUGBwg=",
			"--timestamp-value", now.Format(time.RFC3339Nano),
			"--duration-value", "3723.123456789s",
			"--field-mask-value", "a.b.c,double-value.value,float-value.value",
		},
		expectedMessage: MessageWithWKTs{
			DoubleValue:    &wrapperspb.DoubleValue{Value: 432.1423412},
			FloatValue:     &wrapperspb.FloatValue{Value: 23.324132},
			Int32Value:     &wrapperspb.Int32Value{Value: 32142},
			Int64Value:     &wrapperspb.Int64Value{Value: 32142},
			Uint32Value:    &wrapperspb.UInt32Value{Value: 32142},
			Uint64Value:    &wrapperspb.UInt64Value{Value: 32142},
			BoolValue:      &wrapperspb.BoolValue{Value: true},
			StringValue:    &wrapperspb.StringValue{Value: "abc123"},
			BytesValue:     &wrapperspb.BytesValue{Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
			TimestampValue: timestampMust(now),
			DurationValue:  durationMust(testDuration),
			FieldMaskValue: &fieldmaskpb.FieldMask{Paths: []string{"a.b.c", "double-value.value", "float-value.value"}},
		},
		expectedMask: []string{
			"double_value",
			"float_value",
			"int32_value",
			"int64_value",
			"uint32_value",
			"uint64_value",
			"bool_value",
			"string_value",
			"bytes_value",
			"timestamp_value",
			"duration_value",
			"field_mask_value",
		},
	},
	{
		name: "slice-values",
		arguments: []string{
			"--double-values", "432.1423412,43,3123",
			"--float-values", "23.324132,43,3123",
			"--int32-values", "32142,43,3123",
			"--int64-values", "32142,43,3123",
			"--uint32-values", "32142,43,3123",
			"--uint64-values", "32142,43,3123",
			"--bool-values", "true,true,false",
			"--string-values", "abc123,abc321,abc123",
			"--bytes-values", "AQIDBAUGBwg=,AQIDBAUGBwg=",
			"--timestamp-values", now.Format(time.RFC3339Nano),
			"--timestamp-values", now.Add(testDuration).Format(time.RFC3339Nano),
			"--duration-values", strings.Join([]string{"3723.123456789s", "3723.123456780s", "3723.123456700s"}, ","),
		},
		expectedMessage: MessageWithWKTs{
			DoubleValues: []*wrapperspb.DoubleValue{{Value: 432.1423412}, {Value: 43}, {Value: 3123}},
			FloatValues:  []*wrapperspb.FloatValue{{Value: 23.324132}, {Value: 43}, {Value: 3123}},
			Int32Values:  []*wrapperspb.Int32Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			Int64Values:  []*wrapperspb.Int64Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			Uint32Values: []*wrapperspb.UInt32Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			Uint64Values: []*wrapperspb.UInt64Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			BoolValues:   []*wrapperspb.BoolValue{{Value: true}, {Value: true}, {Value: false}},
			StringValues: []*wrapperspb.StringValue{
				{Value: "abc123"},
				{Value: "abc321"},
				{Value: "abc123"},
			},
			BytesValues: []*wrapperspb.BytesValue{
				{Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
				{Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
			},
			TimestampValues: []*timestamppb.Timestamp{
				timestampMust(now),
				timestampMust(now.Add(testDuration)),
			},
			DurationValues: []*durationpb.Duration{
				durationMust(testDuration),
				durationMust(testDuration.Truncate(10)),
				durationMust(testDuration.Truncate(100)),
			},
		},
		expectedMask: []string{
			"double_values",
			"float_values",
			"int32_values",
			"int64_values",
			"uint32_values",
			"uint64_values",
			"bool_values",
			"string_values",
			"bytes_values",
			"timestamp_values",
			"duration_values",
		},
	},
}

var testMessagesWithOneofWKTs = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithOneofWKTs
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithOneofWKTs{},
		expectedMask:    nil,
	},
	{
		name: "single-values",
		arguments: []string{
			"--value.double-value", "432.1423412",
			"--value.float-value", "23.324132",
			"--value.int32-value", "32142",
			"--value.int64-value", "32142",
			"--value.uint32-value", "32142",
			"--value.uint64-value", "32142",
			"--value.bool-value",
			"--value.string-value", "abc123",
			"--value.bytes-value", "AQIDBAUGBwg=",
			"--value.timestamp-value", now.Format(time.RFC3339Nano),
			"--value.duration-value", "3723.123456789s",
			"--value.field-mask-value", "a.b.c,double-value.value,float-value.value",
		},
		expectedMessage: MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_FieldMaskValue{
				&fieldmaskpb.FieldMask{Paths: []string{"a.b.c", "double-value.value", "float-value.value"}},
			},
		},
		expectedMask: []string{
			"value.double_value",
			"value.float_value",
			"value.int32_value",
			"value.int64_value",
			"value.uint32_value",
			"value.uint64_value",
			"value.bool_value",
			"value.string_value",
			"value.bytes_value",
			"value.timestamp_value",
			"value.duration_value",
			"value.field_mask_value",
		},
	},
}

var testMessagesWithWKTMaps = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithWKTMaps
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithWKTMaps{},
		expectedMask:    nil,
	},
	{
		name: "single-values",
		arguments: []string{
			"--string-double-map", "foo=432.1423412,bar=231.12",
			"--string-float-map", "foo=432.1423412,bar=231.12",
			"--string-int32-map", "foo=1,bar=-2",
			"--string-int64-map", "foo=1,bar=-2",
			"--string-uint32-map", "foo=1,bar=2",
			"--string-uint64-map", "foo=1,bar=2",
			"--string-bool-map", "foo=true,bar=false",
			"--string-string-map", "foo=bar,bar=baz",
			"--string-bytes-map", "foo=AQIDBAUGBwg=,bar=AQIDBAUGBwg=",
			"--string-timestamp-map", fmt.Sprintf("foo=%s,bar=%s", now.Format(time.RFC3339Nano), now.Add(testDuration).Format(time.RFC3339Nano)),
			"--string-duration-map", "foo=3723.123456789s,bar=3723.123456780s",
		},
		expectedMessage: MessageWithWKTMaps{
			StringDoubleMap: map[string]*wrapperspb.DoubleValue{
				"foo": {Value: 432.1423412},
				"bar": {Value: 231.12},
			},
			StringFloatMap: map[string]*wrapperspb.FloatValue{
				"foo": {Value: 432.1423412},
				"bar": {Value: 231.12},
			},
			StringInt32Map: map[string]*wrapperspb.Int32Value{
				"foo": {Value: 1}, "bar": {Value: -2},
			},
			StringInt64Map: map[string]*wrapperspb.Int64Value{
				"foo": {Value: 1}, "bar": {Value: -2},
			},
			StringUint32Map: map[string]*wrapperspb.UInt32Value{
				"foo": {Value: 1}, "bar": {Value: 2},
			},
			StringUint64Map: map[string]*wrapperspb.UInt64Value{
				"foo": {Value: 1}, "bar": {Value: 2},
			},
			StringBoolMap: map[string]*wrapperspb.BoolValue{
				"foo": {Value: true},
				"bar": {Value: false},
			},
			StringStringMap: map[string]*wrapperspb.StringValue{
				"foo": {Value: "bar"},
				"bar": {Value: "baz"},
			},
			StringBytesMap: map[string]*wrapperspb.BytesValue{
				"foo": {Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
				"bar": {Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
			},
			StringTimestampMap: map[string]*timestamppb.Timestamp{
				"foo": timestampMust(now),
				"bar": timestampMust(now.Add(testDuration)),
			},
			StringDurationMap: map[string]*durationpb.Duration{
				"foo": durationMust(testDuration),
				"bar": durationMust(testDuration.Truncate(10)),
			},
		},
		expectedMask: []string{
			"string_double_map",
			"string_float_map",
			"string_int32_map",
			"string_int64_map",
			"string_uint32_map",
			"string_uint64_map",
			"string_bool_map",
			"string_string_map",
			"string_bytes_map",
			"string_timestamp_map",
			"string_duration_map",
		},
	},
}

func timestampMust(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func durationMust(t time.Duration) *durationpb.Duration {
	return durationpb.New(t)
}

func TestSetFlagsMessageWithWKTs(t *testing.T) {
	for _, tt := range testMessagesWithWKTs {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithWKTs(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}

func TestSetFlagsMessageWithOneofWKTs(t *testing.T) {
	for _, tt := range testMessagesWithOneofWKTs {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithOneofWKTs(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}

func TestSetFlagsMessageWithWKTMaps(t *testing.T) {
	for _, tt := range testMessagesWithWKTMaps {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithWKTMaps(fs, "", false)
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}
