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

func TestFloat32Value(t *testing.T) {
	flag := flagsplugin.NewFloat32Flag("float32-value", "float32 value")

	if ft := flag.Value.Type(); ft != "float32" {
		t.Errorf("flag type is %q (expected float32)", ft)
	}

	if fv := flag.Value.String(); fv != "0" {
		t.Errorf("flag value is %v (expected 0)", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetFloat32(fs, "float32-value")
	if err == nil {
		t.Errorf("expected to get error from GetFloat32 on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--float32-value float32   float32 value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetFloat32(fs, "float32-value")
	if err != nil {
		t.Errorf("unexpected error from GetFloat32")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value != 0.0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--float32-value", "12.34"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetFloat32(fs, "float32-value")
	if err != nil {
		t.Errorf("unexpected error from GetFloat32")
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

	pvalue, err := fs.GetFloat32("float32-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetFloat32: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetFloat32 returned different value than GetFloat32")
	}
}

func TestFloat32SliceValue(t *testing.T) {
	flag := flagsplugin.NewFloat32SliceFlag("float32-values", "float32 values")

	if ft := flag.Value.Type(); ft != "float32Slice" {
		t.Errorf("flag type is %q (expected float32Slice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetFloat32Slice(fs, "float32-values")
	if err == nil {
		t.Errorf("expected to get error from GetFloat32Slice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--float32-values float32Slice   float32 values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetFloat32Slice(fs, "float32-values")
	if err != nil {
		t.Errorf("unexpected error from GetFloat32Slice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--float32-values", "12.34", "--float32-values", "56.78"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetFloat32Slice(fs, "float32-values")
	if err != nil {
		t.Errorf("unexpected error from GetFloat32Slice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []float32{12.34, 56.78}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[12.34,56.78]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetFloat32Slice("float32-values")
	if err != nil {
		t.Errorf("unexpected error from standard GetFloat32Slice: %v", err)
	}
	if !reflect.DeepEqual(pvalue, value) {
		t.Errorf("standard GetFloat32Slice returned different value than GetFloat32Slice")
	}
}

func TestStringFloat32MapValue(t *testing.T) {
	flag := flagsplugin.NewStringFloat32MapFlag("float32-map", "float32 map")

	if ft := flag.Value.Type(); ft != "stringToFloat32" {
		t.Errorf("flag type is %q (expected stringToFloat32)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringFloat32Map(fs, "float32-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringFloat32Map on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--float32-map stringToFloat32   float32 map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringFloat32Map(fs, "float32-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringFloat32Map")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--float32-map", "foo=12.34", "--float32-map", "bar=56.78"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringFloat32Map(fs, "float32-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringFloat32Map")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]float32{"foo": 12.34, "bar": 56.78}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=56.78,foo=12.34]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
