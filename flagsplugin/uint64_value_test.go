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

func TestUint64Value(t *testing.T) {
	flag := flagsplugin.NewUint64Flag("uint64-value", "uint64 value")

	if ft := flag.Value.Type(); ft != "uint64" {
		t.Errorf("flag type is %q (expected uint64)", ft)
	}

	if fv := flag.Value.String(); fv != "0" {
		t.Errorf("flag value is %v (expected 0)", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetUint64(fs, "uint64-value")
	if err == nil {
		t.Errorf("expected to get error from GetUint64 on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--uint64-value uint   uint64 value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetUint64(fs, "uint64-value")
	if err != nil {
		t.Errorf("unexpected error from GetUint64")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value != 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--uint64-value", "1234"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetUint64(fs, "uint64-value")
	if err != nil {
		t.Errorf("unexpected error from GetUint64")
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

	pvalue, err := fs.GetUint64("uint64-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetUint64: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetUint64 returned different value than GetUint64")
	}
}

func TestUint64SliceValue(t *testing.T) {
	flag := flagsplugin.NewUint64SliceFlag("uint64-values", "uint64 values")

	if ft := flag.Value.Type(); ft != "uint64Slice" {
		t.Errorf("flag type is %q (expected uint64Slice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetUint64Slice(fs, "uint64-values")
	if err == nil {
		t.Errorf("expected to get error from GetUint64Slice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--uint64-values uint64Slice   uint64 values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetUint64Slice(fs, "uint64-values")
	if err != nil {
		t.Errorf("unexpected error from GetUint64Slice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--uint64-values", "1234", "--uint64-values", "5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetUint64Slice(fs, "uint64-values")
	if err != nil {
		t.Errorf("unexpected error from GetUint64Slice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []uint64{1234, 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[1234,5678]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}

func TestStringUint64MapValue(t *testing.T) {
	flag := flagsplugin.NewStringUint64MapFlag("uint64-map", "uint64 map")

	if ft := flag.Value.Type(); ft != "stringToUint64" {
		t.Errorf("flag type is %q (expected stringToUint64)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringUint64Map(fs, "uint64-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringUint64Map on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--uint64-map stringToUint64   uint64 map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringUint64Map(fs, "uint64-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringUint64Map")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--uint64-map", "foo=1234", "--uint64-map", "bar=5678"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringUint64Map(fs, "uint64-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringUint64Map")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string]uint64{"foo": 1234, "bar": 5678}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=5678,foo=1234]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
