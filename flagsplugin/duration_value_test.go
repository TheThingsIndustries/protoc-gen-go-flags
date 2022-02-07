// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
)

func TestDurationValue(t *testing.T) {
	flag := flagsplugin.NewDurationFlag("duration-value", "duration value")

	if ft := flag.Value.Type(); ft != "duration" {
		t.Errorf("flag type is %q (expected duration)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetDuration(fs, "duration-value")
	if err == nil {
		t.Errorf("expected to get error from GetDuration on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--duration-value duration   duration value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetDuration(fs, "duration-value")
	if err != nil {
		t.Errorf("unexpected error from GetDuration")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if value != 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--duration-value", "1h2m3s"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetDuration(fs, "duration-value")
	if err != nil {
		t.Errorf("unexpected error from GetDuration")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if value != time.Duration(time.Hour+2*time.Minute+3*time.Second) {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "1h2m3s" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetDuration("duration-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetDuration: %v", err)
	}
	if pvalue != value {
		t.Errorf("standard GetDuration returned different value than GetDuration")
	}
}

func TestDurationSliceValue(t *testing.T) {
	flag := flagsplugin.NewDurationSliceFlag("duration-values", "duration values")

	if ft := flag.Value.Type(); ft != "durationSlice" {
		t.Errorf("flag type is %q (expected durationSlice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetDurationSlice(fs, "duration-values")
	if err == nil {
		t.Errorf("expected to get error from GetDurationSlice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--duration-values durationSlice   duration values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetDurationSlice(fs, "duration-values")
	if err != nil {
		t.Errorf("unexpected error from GetDurationSlice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--duration-values", "1h", "--duration-values", "1m"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetDurationSlice(fs, "duration-values")
	if err != nil {
		t.Errorf("unexpected error from GetDurationSlice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, []time.Duration{time.Hour, time.Minute}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[1h0m0s,1m0s]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetDurationSlice("duration-values")
	if err != nil {
		t.Errorf("unexpected error from standard GetDurationSlice: %v", err)
	}
	if !reflect.DeepEqual(pvalue, value) {
		t.Errorf("standard GetDurationSlice returned different value than GetDurationSlice")
	}
}

func TestStringDurationMapValue(t *testing.T) {
	flag := flagsplugin.NewStringDurationMapFlag("duration-map", "duration map")

	if ft := flag.Value.Type(); ft != "stringToDuration" {
		t.Errorf("flag type is %q (expected stringToDuration)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringDurationMap(fs, "duration-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringDurationMap on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--duration-map stringToDuration   duration map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringDurationMap(fs, "duration-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringDurationMap")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--duration-map", "foo=10s", "--duration-map", "bar=1h"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringDurationMap(fs, "duration-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringDurationMap")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}

	duration_1, _ := time.ParseDuration("10s")
	duration_2, _ := time.ParseDuration("1h")
	if !reflect.DeepEqual(value, map[string]time.Duration{"foo": duration_1, "bar": duration_2}) {
		t.Errorf("unexpected map value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=1h0m0s,foo=10s]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
