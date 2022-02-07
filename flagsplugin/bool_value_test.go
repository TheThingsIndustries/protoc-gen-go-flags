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

func TestBoolValue(t *testing.T) {
	flag := flagsplugin.NewBoolFlag("bool-value", "bool value")

	if ft := flag.Value.Type(); ft != "bool" {
		t.Errorf("flag type is %q (expected bool)", ft)
	}

	if fv := flag.Value.String(); fv != "false" {
		t.Errorf("flag value is %v (expected false)", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetBool(fs, "bool-value")
	if err == nil {
		t.Errorf("expected to get error from GetBool on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bool-value   bool value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetBool(fs, "bool-value")
	if err != nil {
		t.Errorf("unexpected error from GetBool")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bool-value"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetBool(fs, "bool-value")
	if err != nil {
		t.Errorf("unexpected error from GetBool")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !value {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "true" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetBool("bool-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetBool: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetBool returned different value than GetBool")
	}
}

func TestBoolSliceValue(t *testing.T) {
	flag := flagsplugin.NewBoolSliceFlag("bool-values", "bool values")

	if ft := flag.Value.Type(); ft != "boolSlice" {
		t.Errorf("flag type is %q (expected boolSlice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetBoolSlice(fs, "bool-values")
	if err == nil {
		t.Errorf("expected to get error from GetBoolSlice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bool-values bools   bool values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetBoolSlice(fs, "bool-values")
	if err != nil {
		t.Errorf("unexpected error from GetBoolSlice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bool-values", "true", "--bool-values", "false"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetBoolSlice(fs, "bool-values")
	if err != nil {
		t.Errorf("unexpected error from GetBoolSlice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []bool{true, false}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[true,false]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetBoolSlice("bool-values")
	if err != nil {
		t.Errorf("unexpected error from standard GetBoolSlice: %v", err)
	}
	if !reflect.DeepEqual(pvalue, value) {
		t.Errorf("standard GetBoolSlice returned different value than GetBoolSlice")
	}
}

func TestStringBoolMapValue(t *testing.T) {
	flag := flagsplugin.NewStringBoolMapFlag("bool-map", "bool map")

	if ft := flag.Value.Type(); ft != "stringToBool" {
		t.Errorf("flag type is %q (expected stringToBool)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringBoolMap(fs, "bool-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringBoolMap on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bool-map stringToBool   bool map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringBoolMap(fs, "bool-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringBoolMap")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bool-map", "foo=true", "--bool-map", "bar=false"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringBoolMap(fs, "bool-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringBoolMap")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]bool{"foo": true, "bar": false}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=false,foo=true]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
