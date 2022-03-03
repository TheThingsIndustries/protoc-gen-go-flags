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

func TestUint32Value(t *testing.T) {
	flag := flagsplugin.NewUint32Flag("uint32-value", "uint32 value")

	if ft := flag.Value.Type(); ft != "uint32" {
		t.Errorf("flag type is %q (expected uint32)", ft)
	}

	if fv := flag.Value.String(); fv != "0" {
		t.Errorf("flag value is %v (expected 0)", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetUint32(fs, "uint32-value")
	if err == nil {
		t.Errorf("expected to get error from GetUint32 on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--uint32-value uint32   uint32 value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetUint32(fs, "uint32-value")
	if err != nil {
		t.Errorf("unexpected error from GetUint32")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value != 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--uint32-value", "1234"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetUint32(fs, "uint32-value")
	if err != nil {
		t.Errorf("unexpected error from GetUint32")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if value != 1234 {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "1234" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetUint32("uint32-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetUint32: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetUint32 returned different value than GetUint32")
	}
}

func TestUint32SliceValue(t *testing.T) {
	flag := flagsplugin.NewUint32SliceFlag("uint32-values", "uint32 values")

	if ft := flag.Value.Type(); ft != "uint32Slice" {
		t.Errorf("flag type is %q (expected uint32Slice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetUint32Slice(fs, "uint32-values")
	if err == nil {
		t.Errorf("expected to get error from GetUint32Slice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--uint32-values uint32Slice   uint32 values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetUint32Slice(fs, "uint32-values")
	if err != nil {
		t.Errorf("unexpected error from GetUint32Slice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--uint32-values", "1234", "--uint32-values", "5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetUint32Slice(fs, "uint32-values")
	if err != nil {
		t.Errorf("unexpected error from GetUint32Slice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []uint32{1234, 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[1234,5678]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}

func TestStringUint32MapValue(t *testing.T) {
	flag := flagsplugin.NewStringUint32MapFlag("uint32-map", "uint32 map")

	if ft := flag.Value.Type(); ft != "stringToUint32" {
		t.Errorf("flag type is %q (expected stringToUint32)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringUint32Map(fs, "uint32-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringUint32Map on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--uint32-map stringToUint32   uint32 map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringUint32Map(fs, "uint32-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringUint32Map")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--uint32-map", "foo=1234", "--uint32-map", "bar=5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringUint32Map(fs, "uint32-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringUint32Map")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]uint32{"foo": 1234, "bar": 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=5678,foo=1234]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
