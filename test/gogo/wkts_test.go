package test_test

import (
	fmt "fmt"
	"strings"
	"testing"
	time "time"

	. "github.com/TheThingsIndustries/protoc-gen-go-flags/test/gogo"
	"github.com/gogo/protobuf/types"
	pflag "github.com/spf13/pflag"
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
			"--bytes-value", "abc123",
			"--timestamp-value", now.Format(time.RFC3339Nano),
			"--duration-value", "3723.123456789s",
			"--field-mask-value", "a.b.c,double-value.value,float-value.value",
		},
		expectedMessage: MessageWithWKTs{
			DoubleValue:    &types.DoubleValue{Value: 432.1423412},
			FloatValue:     &types.FloatValue{Value: 23.324132},
			Int32Value:     &types.Int32Value{Value: 32142},
			Int64Value:     &types.Int64Value{Value: 32142},
			Uint32Value:    &types.UInt32Value{Value: 32142},
			Uint64Value:    &types.UInt64Value{Value: 32142},
			BoolValue:      &types.BoolValue{Value: true},
			StringValue:    &types.StringValue{Value: "abc123"},
			BytesValue:     &types.BytesValue{Value: getBytesValue("abc123")},
			TimestampValue: getTimestampProto(now),
			DurationValue:  types.DurationProto(testDuration),
			FieldMaskValue: &types.FieldMask{Paths: []string{"a.b.c", "double-value.value", "float-value.value"}},
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
			DoubleValues: []*types.DoubleValue{{Value: 432.1423412}, {Value: 43}, {Value: 3123}},
			FloatValues:  []*types.FloatValue{{Value: 23.324132}, {Value: 43}, {Value: 3123}},
			Int32Values:  []*types.Int32Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			Int64Values:  []*types.Int64Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			Uint32Values: []*types.UInt32Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			Uint64Values: []*types.UInt64Value{{Value: 32142}, {Value: 43}, {Value: 3123}},
			BoolValues:   []*types.BoolValue{{Value: true}, {Value: true}, {Value: false}},
			StringValues: []*types.StringValue{
				{Value: "abc123"},
				{Value: "abc321"},
				{Value: "abc123"},
			},
			BytesValues: []*types.BytesValue{
				{Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
				{Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
			},
			TimestampValues: []*types.Timestamp{
				getTimestampProto(now),
				getTimestampProto(now.Add(testDuration)),
			},
			DurationValues: []*types.Duration{
				types.DurationProto(testDuration),
				types.DurationProto(testDuration.Truncate(10)),
				types.DurationProto(testDuration.Truncate(100)),
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
			"--value.bytes-value", "abc123",
			"--value.timestamp-value", now.Format(time.RFC3339Nano),
			"--value.duration-value", "3723.123456789s",
			"--value.field-mask-value", "a.b.c,double-value.value,float-value.value",
		},
		expectedMessage: MessageWithOneofWKTs{
			Value: &MessageWithOneofWKTs_FieldMaskValue{
				&types.FieldMask{Paths: []string{"a.b.c", "double-value.value", "float-value.value"}},
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
			StringDoubleMap: map[string]*types.DoubleValue{
				"foo": {Value: 432.1423412},
				"bar": {Value: 231.12},
			},
			StringFloatMap: map[string]*types.FloatValue{
				"foo": {Value: 432.1423412},
				"bar": {Value: 231.12},
			},
			StringInt32Map: map[string]*types.Int32Value{
				"foo": {Value: 1}, "bar": {Value: -2},
			},
			StringInt64Map: map[string]*types.Int64Value{
				"foo": {Value: 1}, "bar": {Value: -2},
			},
			StringUint32Map: map[string]*types.UInt32Value{
				"foo": {Value: 1}, "bar": {Value: 2},
			},
			StringUint64Map: map[string]*types.UInt64Value{
				"foo": {Value: 1}, "bar": {Value: 2},
			},
			StringBoolMap: map[string]*types.BoolValue{
				"foo": {Value: true},
				"bar": {Value: false},
			},
			StringStringMap: map[string]*types.StringValue{
				"foo": {Value: "bar"},
				"bar": {Value: "baz"},
			},
			StringBytesMap: map[string]*types.BytesValue{
				"foo": {Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
				"bar": {Value: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
			},
			StringTimestampMap: map[string]*types.Timestamp{
				"foo": getTimestampProto(now),
				"bar": getTimestampProto(now.Add(testDuration)),
			},
			StringDurationMap: map[string]*types.Duration{
				"foo": types.DurationProto(testDuration),
				"bar": types.DurationProto(testDuration.Truncate(10)),
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

func getTimestampProto(t time.Time) *types.Timestamp {
	timestamp, err := types.TimestampProto(t)
	if err != nil {
		return nil
	}
	return timestamp
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
