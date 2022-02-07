// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
)

func TestInt32Value(t *testing.T) {
	flag := flagsplugin.NewInt32Flag("int32-value", "int32 value")

	if ft := flag.Value.Type(); ft != "int32" {
		t.Errorf("flag type is %q (expected int32)", ft)
	}

	if fv := flag.Value.String(); fv != "0" {
		t.Errorf("flag value is %v (expected 0)", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetInt32(fs, "int32-value")
	if err == nil {
		t.Errorf("expected to get error from GetInt32 on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--int32-value int32   int32 value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetInt32(fs, "int32-value")
	if err != nil {
		t.Errorf("unexpected error from GetInt32")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value != 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--int32-value", "-1234"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetInt32(fs, "int32-value")
	if err != nil {
		t.Errorf("unexpected error from GetInt32")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if value != -1234 {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "-1234" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetInt32("int32-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetInt32: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetInt32 returned different value than GetInt32")
	}
}

func TestInt32SliceValue(t *testing.T) {
	flag := flagsplugin.NewInt32SliceFlag("int32-values", "int32 values")

	if ft := flag.Value.Type(); ft != "int32Slice" {
		t.Errorf("flag type is %q (expected int32Slice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetInt32Slice(fs, "int32-values")
	if err == nil {
		t.Errorf("expected to get error from GetInt32Slice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--int32-values int32Slice   int32 values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetInt32Slice(fs, "int32-values")
	if err != nil {
		t.Errorf("unexpected error from GetInt32Slice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--int32-values", "-1234", "--int32-values", "5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetInt32Slice(fs, "int32-values")
	if err != nil {
		t.Errorf("unexpected error from GetInt32Slice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []int32{-1234, 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[-1234,5678]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetInt32Slice("int32-values")
	if err != nil {
		t.Errorf("unexpected error from standard GetInt32Slice: %v", err)
	}
	if !reflect.DeepEqual(pvalue, value) {
		t.Errorf("standard GetInt32Slice returned different value than GetInt32Slice")
	}
}

func TestStringInt32MapValue(t *testing.T) {
	flag := flagsplugin.NewStringInt32MapFlag("int32-map", "int32 map")

	if ft := flag.Value.Type(); ft != "stringToInt32" {
		t.Errorf("flag type is %q (expected stringToInt32)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringInt32Map(fs, "int32-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringInt32Map on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--int32-map stringToInt32   int32 map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringInt32Map(fs, "int32-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringInt32Map")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--int32-map", "foo=-1234", "--int32-map", "bar=5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringInt32Map(fs, "int32-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringInt32Map")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]int32{"foo": -1234, "bar": 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=5678,foo=-1234]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
