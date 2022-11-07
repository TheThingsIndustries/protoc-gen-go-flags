// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin_test

import (
	"testing"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
)

// EnumValueDesc returns a string with possible values for enum.
func TestEnumValueDesc(t *testing.T) {
	tests := []struct {
		enum     map[string]int32
		expected string
	}{
		{
			enum: map[string]int32{
				"A": 1,
				"B": 1,
				"C": 1,
				"D": 1,
				"E": 1,
			},
			expected: "allowed values: A, B, C, D, E",
		},
		{
			enum: map[string]int32{
				"V1.0":  1,
				"1.0.0": 1,
				"V2.0":  2,
				"2.0.0": 2,
				"V3.0":  3,
				"3.0.0": 3,
			},
			expected: "allowed values: 1.0.0, V1.0, 2.0.0, V2.0, 3.0.0, V3.0",
		},
	}

	for _, test := range tests {
		test := test

		result := flagsplugin.EnumValueDesc(test.enum)
		if result != test.expected {
			t.Errorf("expected outcome: %s, created output: %s", test.expected, result)
		}
	}
}
