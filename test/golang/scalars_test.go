package test_test

import (
	"encoding/base64"
	"testing"

	. "github.com/TheThingsIndustries/protoc-gen-go-flags/test/golang"
	pflag "github.com/spf13/pflag"
)

var testMessagesWithScalars = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithScalars
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithScalars{},
		expectedMask:    nil,
	},
	{
		name: "full",
		arguments: []string{
			"--double-value", "2.0",
			"--double-values", "1.0,2.0,3",
			"--float-value", "1.32432",
			"--float-values", "1.213213,3.34214312",
			"--int32-value", "43214123",
			"--int32-values", "848348348,237148293",
			"--int64-value", "73458975481",
			"--int64-values", "9321894,4632112387634,6747823",
			"--uint32-value", "213123",
			"--uint32-values", "213123,7342859",
			"--uint64-value", "213123",
			"--uint64-values", "213123,7342859",
			"--sint32-value", "213123",
			"--sint32-values", "213123,7342859",
			"--sint64-value", "213123",
			"--sint64-values", "213123,7342859",
			"--fixed32-value", "213123",
			"--fixed32-values", "213123,7342859",
			"--fixed64-value", "213123",
			"--fixed64-values", "213123,7342859",
			"--sfixed32-value", "213123",
			"--sfixed32-values", "213123,7342859",
			"--sfixed64-value", "213123",
			"--sfixed64-values", "213123,7342859",
			"--bool-value",
			"--bool-values", "true,false",
			"--string-value", "string1",
			"--string-values", "string1,string2,string3",
			"--bytes-value", "AQIDBAUGBwg=",
			"--bytes-values", "AQIDBAUGBwg=,AQIDBAUGBwg=,AQIDBAUGBwg=",
			"--hex-bytes-value", "01020304",
			"--hex-bytes-values", "01020304",
			"--hex-bytes-values", "0102030405,010203040506",
		},
		expectedMessage: MessageWithScalars{
			DoubleValue:    2.0,
			DoubleValues:   []float64{1.0, 2.0, 3.0},
			FloatValue:     1.32432,
			FloatValues:    []float32{1.213213, 3.34214312},
			Int32Value:     43214123,
			Int32Values:    []int32{848348348, 237148293},
			Int64Value:     73458975481,
			Int64Values:    []int64{9321894, 4632112387634, 6747823},
			Uint32Value:    213123,
			Uint32Values:   []uint32{213123, 7342859},
			Uint64Value:    213123,
			Uint64Values:   []uint64{213123, 7342859},
			Sint32Value:    213123,
			Sint32Values:   []int32{213123, 7342859},
			Sint64Value:    213123,
			Sint64Values:   []int64{213123, 7342859},
			Fixed32Value:   213123,
			Fixed32Values:  []uint32{213123, 7342859},
			Fixed64Value:   213123,
			Fixed64Values:  []uint64{213123, 7342859},
			Sfixed32Value:  213123,
			Sfixed32Values: []int32{213123, 7342859},
			Sfixed64Value:  213123,
			Sfixed64Values: []int64{213123, 7342859},
			BoolValue:      true,
			BoolValues:     []bool{true, false},
			StringValue:    "string1",
			StringValues:   []string{"string1", "string2", "string3"},
			BytesValue:     []byte{1, 2, 3, 4, 5, 6, 7, 8},
			BytesValues:    [][]byte{{1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 8}, {1, 2, 3, 4, 5, 6, 7, 8}},
			HexBytesValue:  []byte{1, 2, 3, 4},
			HexBytesValues: [][]byte{{1, 2, 3, 4}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5, 6}},
		},
		expectedMask: []string{
			"double_value",
			"double_values",
			"float_value",
			"float_values",
			"int32_value",
			"int32_values",
			"int64_value",
			"int64_values",
			"uint32_value",
			"uint32_values",
			"uint64_value",
			"uint64_values",
			"sint32_value",
			"sint32_values",
			"sint64_value",
			"sint64_values",
			"fixed32_value",
			"fixed32_values",
			"fixed64_value",
			"fixed64_values",
			"sfixed32_value",
			"sfixed32_values",
			"sfixed64_value",
			"sfixed64_values",
			"bool_value",
			"bool_values",
			"string_value",
			"string_values",
			"bytes_value",
			"bytes_values",
			"hex_bytes_value",
			"hex_bytes_values",
		},
	},
}

var testMessagesWithOneofScalars = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithOneofScalars
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithOneofScalars{},
		expectedMask:    nil,
	},
	{
		name: "one value set",
		arguments: []string{
			"--value.sint32-value", "-213123",
		},
		expectedMessage: MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_Sint32Value{
				Sint32Value: -213123,
			},
		},
		expectedMask: []string{
			"value.sint32_value",
		},
	},
	{
		name: "all values set, last value overwrites others",
		arguments: []string{
			"--value.double-value", "2.0",
			"--value.float-value", "1.32432",
			"--value.int32-value", "43214123",
			"--value.int64-value", "73458975481",
			"--value.uint32-value", "213123",
			"--value.uint64-value", "213123",
			"--value.sint32-value", "213123",
			"--value.sint64-value", "213123",
			"--value.fixed32-value", "213123",
			"--value.fixed64-value", "213123",
			"--value.sfixed32-value", "213123",
			"--value.sfixed64-value", "213123",
			"--value.bool-value",
			"--value.string-value", "string1",
			"--value.bytes-value", "AQIDBAUGBwg=",
		},
		expectedMessage: MessageWithOneofScalars{
			Value: &MessageWithOneofScalars_BytesValue{
				BytesValue: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		expectedMask: []string{
			"value.double_value",
			"value.float_value",
			"value.int32_value",
			"value.int64_value",
			"value.uint32_value",
			"value.uint64_value",
			"value.sint32_value",
			"value.sint64_value",
			"value.fixed32_value",
			"value.fixed64_value",
			"value.sfixed32_value",
			"value.sfixed64_value",
			"value.bool_value",
			"value.string_value",
			"value.bytes_value",
		},
	},
}

var testMessagesWithScalarMaps = []struct {
	name            string
	arguments       []string
	expectedMessage MessageWithScalarMaps
	expectedMask    []string
}{
	{
		name:            "empty",
		arguments:       []string{},
		expectedMessage: MessageWithScalarMaps{},
		expectedMask:    nil,
	},
	{
		name: "full",
		arguments: []string{
			"--string-double-map", "foo=12.34,bar=23.21",
			"--string-double-map", "baz=10.11",
			"--string-float-map", "foo=10.3221534",
			"--string-int32-map", "bar=-234",
			"--string-int64-map", "foo=10,bar=12",
			"--string-uint32-map", "foo=10,bar=12",
			"--string-uint64-map", "foo=10,bar=12",
			"--string-sint32-map", "foo=10,bar=12",
			"--string-sint64-map", "foo=10,bar=12",
			"--string-fixed32-map", "foo=10,bar=12",
			"--string-fixed64-map", "foo=10,bar=12",
			"--string-sfixed32-map", "foo=10,bar=12",
			"--string-sfixed64-map", "foo=10,bar=12",
			"--string-bool-map", "foo=true,bar=false",
			"--string-string-map", "foo=bar,bar=baz",
			"--string-bytes-map", "foo=AQIDBAUGBwg=,bar=AQIDBAUGBwg=",
		},
		expectedMessage: MessageWithScalarMaps{
			StringDoubleMap: map[string]float64{
				"foo": 12.34, "bar": 23.21, "baz": 10.11,
			},
			StringFloatMap:    map[string]float32{"foo": 10.3221534},
			StringInt32Map:    map[string]int32{"bar": -234},
			StringInt64Map:    map[string]int64{"foo": 10, "bar": 12},
			StringUint32Map:   map[string]uint32{"foo": 10, "bar": 12},
			StringUint64Map:   map[string]uint64{"foo": 10, "bar": 12},
			StringSint32Map:   map[string]int32{"foo": 10, "bar": 12},
			StringSint64Map:   map[string]int64{"foo": 10, "bar": 12},
			StringFixed32Map:  map[string]uint32{"foo": 10, "bar": 12},
			StringFixed64Map:  map[string]uint64{"foo": 10, "bar": 12},
			StringSfixed32Map: map[string]int32{"foo": 10, "bar": 12},
			StringSfixed64Map: map[string]int64{"foo": 10, "bar": 12},
			StringBoolMap:     map[string]bool{"foo": true, "bar": false},
			StringStringMap:   map[string]string{"foo": "bar", "bar": "baz"},
			StringBytesMap: map[string][]byte{
				"foo": {1, 2, 3, 4, 5, 6, 7, 8}, "bar": {1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		expectedMask: []string{
			"string_double_map",
			"string_float_map",
			"string_int32_map",
			"string_int64_map",
			"string_uint32_map",
			"string_uint64_map",
			"string_sint32_map",
			"string_sint64_map",
			"string_fixed32_map",
			"string_fixed64_map",
			"string_sfixed32_map",
			"string_sfixed64_map",
			"string_bool_map",
			"string_string_map",
			"string_bytes_map",
		},
	},
}

func getBytesValue(s string) []byte {
	val, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return val
}

func TestSetFlagsMessageWithScalars(t *testing.T) {
	for _, tt := range testMessagesWithScalars {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithScalars(fs, "")
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}

func TestSetFlagsMessageWithOneofScalars(t *testing.T) {
	for _, tt := range testMessagesWithOneofScalars {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithOneofScalars(fs, "")
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}

func TestSetFlagsMessageWithScalarMaps(t *testing.T) {
	for _, tt := range testMessagesWithScalarMaps {
		t.Run(tt.name, func(t *testing.T) {
			fs := &pflag.FlagSet{}
			AddSetFlagsForMessageWithScalarMaps(fs, "")
			expectMessageEqual(t, fs, tt.arguments, &tt.expectedMessage, tt.expectedMask)
		})
	}
}
