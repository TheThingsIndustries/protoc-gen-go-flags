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

func TestInt64Value(t *testing.T) {
	flag := flagsplugin.NewInt64Flag("int64-value", "int64 value")

	if ft := flag.Value.Type(); ft != "int64" {
		t.Errorf("flag type is %q (expected int64)", ft)
	}

	if fv := flag.Value.String(); fv != "0" {
		t.Errorf("flag value is %v (expected 0)", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetInt64(fs, "int64-value")
	if err == nil {
		t.Errorf("expected to get error from GetInt64 on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--int64-value int   int64 value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetInt64(fs, "int64-value")
	if err != nil {
		t.Errorf("unexpected error from GetInt64")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value != 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--int64-value", "-1234"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetInt64(fs, "int64-value")
	if err != nil {
		t.Errorf("unexpected error from GetInt64")
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

	pvalue, err := fs.GetInt64("int64-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetInt64: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetInt64 returned different value than GetInt64")
	}
}

func TestInt64SliceValue(t *testing.T) {
	flag := flagsplugin.NewInt64SliceFlag("int64-values", "int64 values")

	if ft := flag.Value.Type(); ft != "int64Slice" {
		t.Errorf("flag type is %q (expected int64Slice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetInt64Slice(fs, "int64-values")
	if err == nil {
		t.Errorf("expected to get error from GetInt64Slice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--int64-values int64Slice   int64 values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetInt64Slice(fs, "int64-values")
	if err != nil {
		t.Errorf("unexpected error from GetInt64Slice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--int64-values", "-1234", "--int64-values", "5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetInt64Slice(fs, "int64-values")
	if err != nil {
		t.Errorf("unexpected error from GetInt64Slice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []int64{-1234, 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[-1234,5678]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetInt64Slice("int64-values")
	if err != nil {
		t.Errorf("unexpected error from standard GetInt64Slice: %v", err)
	}
	if !reflect.DeepEqual(pvalue, value) {
		t.Errorf("standard GetInt64Slice returned different value than GetInt64Slice")
	}
}

func TestStringInt64MapValue(t *testing.T) {
	flag := flagsplugin.NewStringInt64MapFlag("int64-map", "int64 map")

	if ft := flag.Value.Type(); ft != "stringToInt64" {
		t.Errorf("flag type is %q (expected stringToInt64)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringInt64Map(fs, "int64-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringInt64Map on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--int64-map stringToInt64   int64 map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringInt64Map(fs, "int64-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringInt64Map")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--int64-map", "foo=-1234", "--int64-map", "bar=5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringInt64Map(fs, "int64-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringInt64Map")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]int64{"foo": -1234, "bar": 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=5678,foo=-1234]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
