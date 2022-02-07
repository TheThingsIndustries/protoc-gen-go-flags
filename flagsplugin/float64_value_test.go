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

func TestFloat64Value(t *testing.T) {
	flag := flagsplugin.NewFloat64Flag("float64-value", "float64 value")

	if ft := flag.Value.Type(); ft != "float64" {
		t.Errorf("flag type is %q (expected float64)", ft)
	}

	if fv := flag.Value.String(); fv != "0" {
		t.Errorf("flag value is %v (expected 0)", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetFloat64(fs, "float64-value")
	if err == nil {
		t.Errorf("expected to get error from GetFloat64 on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--float64-value float   float64 value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetFloat64(fs, "float64-value")
	if err != nil {
		t.Errorf("unexpected error from GetFloat64")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value != 0.0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--float64-value", "12.34"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetFloat64(fs, "float64-value")
	if err != nil {
		t.Errorf("unexpected error from GetFloat64")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if value != 12.34 {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "12.34" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetFloat64("float64-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetFloat64: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetFloat64 returned different value than GetFloat64")
	}
}

func TestFloat64SliceValue(t *testing.T) {
	flag := flagsplugin.NewFloat64SliceFlag("float64-values", "float64 values")

	if ft := flag.Value.Type(); ft != "float64Slice" {
		t.Errorf("flag type is %q (expected float64Slice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetFloat64Slice(fs, "float64-values")
	if err == nil {
		t.Errorf("expected to get error from GetFloat64Slice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--float64-values float64Slice   float64 values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetFloat64Slice(fs, "float64-values")
	if err != nil {
		t.Errorf("unexpected error from GetFloat64Slice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--float64-values", "12.34", "--float64-values", "56.78"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetFloat64Slice(fs, "float64-values")
	if err != nil {
		t.Errorf("unexpected error from GetFloat64Slice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []float64{12.34, 56.78}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[12.34,56.78]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetFloat64Slice("float64-values")
	if err != nil {
		t.Errorf("unexpected error from standard GetFloat64Slice: %v", err)
	}
	if !reflect.DeepEqual(pvalue, value) {
		t.Errorf("standard GetFloat64Slice returned different value than GetFloat64Slice")
	}
}

func TestStringFloat64MapValue(t *testing.T) {
	flag := flagsplugin.NewStringFloat64MapFlag("float64-map", "float64 map")

	if ft := flag.Value.Type(); ft != "stringToFloat64" {
		t.Errorf("flag type is %q (expected stringToFloat64)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringFloat64Map(fs, "float64-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringFloat64Map on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--float64-map stringToFloat64   float64 map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringFloat64Map(fs, "float64-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringFloat64Map")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--float64-map", "foo=12.34", "--float64-map", "bar=56.78"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringFloat64Map(fs, "float64-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringFloat64Map")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]float64{"foo": 12.34, "bar": 56.78}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=56.78,foo=12.34]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
