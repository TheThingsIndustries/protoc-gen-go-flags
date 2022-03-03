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

func TestStringValue(t *testing.T) {
	flag := flagsplugin.NewStringFlag("string-value", "string value")

	if ft := flag.Value.Type(); ft != "string" {
		t.Errorf("flag type is %q (expected string)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %q (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetString(fs, "string-value")
	if err == nil {
		t.Errorf("expected to get error from GetString on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--string-value string   string value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetString(fs, "string-value")
	if err != nil {
		t.Errorf("unexpected error from GetString")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--string-value", "foo"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetString(fs, "string-value")
	if err != nil {
		t.Errorf("unexpected error from GetString")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if value != "foo" {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "foo" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetString("string-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetString: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetString returned different value than GetString")
	}
}

func TestStringSliceValue(t *testing.T) {
	flag := flagsplugin.NewStringSliceFlag("string-values", "string values")

	if ft := flag.Value.Type(); ft != "stringSlice" {
		t.Errorf("flag type is %q (expected stringSlice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringSlice(fs, "string-values")
	if err == nil {
		t.Errorf("expected to get error from GetStringSlice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--string-values strings   string values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringSlice(fs, "string-values")
	if err != nil {
		t.Errorf("unexpected error from GetStringSlice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--string-values", "foo", "--string-values", "bar", "--string-values", `"Hello, world!"`}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringSlice(fs, "string-values")
	if err != nil {
		t.Errorf("unexpected error from GetStringSlice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []string{"foo", "bar", "Hello, world!"}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != `[foo,bar,"Hello, world!"]` {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetStringSlice("string-values")
	if err != nil {
		t.Errorf("unexpected error from standard GetStringSlice: %v", err)
	}
	if !reflect.DeepEqual(pvalue, value) {
		t.Errorf("standard GetStringSlice returned different value than GetStringSlice: %v instead of %v", pvalue, value)
	}
}

func TestStringStringMapValue(t *testing.T) {
	flag := flagsplugin.NewStringStringMapFlag("string-map", "string map")

	if ft := flag.Value.Type(); ft != "stringToString" {
		t.Errorf("flag type is %q (expected stringToString)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringStringMap(fs, "string-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringStringMap on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--string-map stringToString   string map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringStringMap(fs, "string-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringStringMap")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--string-map", "foo=bar", "--string-map", "bar=baz", "--string-map", "greeting=Hello, world!"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringStringMap(fs, "string-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringStringMap")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]string{"foo": "bar", "bar": "baz", "greeting": "Hello, world!"}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != `[bar="baz",foo="bar",greeting="Hello, world!"]` {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
