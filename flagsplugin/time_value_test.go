// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
)

func TestTimestampValue(t *testing.T) {
	flag := flagsplugin.NewTimestampFlag("timestamp-value", "timestamp value")

	if ft := flag.Value.Type(); ft != "timestamp" {
		t.Errorf("flag type is %q (expected timestamp)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetTimestamp(fs, "timestamp-value")
	if err == nil {
		t.Errorf("expected to get error from GetTimestamp on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--timestamp-value timestamp   timestamp value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetTimestamp(fs, "timestamp-value")
	if err != nil {
		t.Errorf("unexpected error from GetTimestamp")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if !value.IsZero() {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	now := time.Now()

	if err := fs.Parse([]string{"--timestamp-value", now.Format(time.RFC3339Nano)}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetTimestamp(fs, "timestamp-value")
	if err != nil {
		t.Errorf("unexpected error from GetTimestamp")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !value.Equal(now) {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != now.Format(time.RFC3339Nano) {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	// pvalue, err := fs.GetTimestamp("time-value")
	// if err != nil {
	// 	t.Errorf("unexpected error from standard GetTimestamp: %v", err)
	// }
	// if pvalue != value {
	// 	t.Errorf("standard GetTimestamp returned different value than GetTimestamp")
	// }
}

func TestTimeSliceValue(t *testing.T) {
	flag := flagsplugin.NewTimestampSliceFlag("timestamp-values", "timestamp values")

	if ft := flag.Value.Type(); ft != "timestampSlice" {
		t.Errorf("flag type is %q (expected timeSlice)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetTimestampSlice(fs, "timestamp-values")
	if err == nil {
		t.Errorf("expected to get error from GetTimestampSlice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--timestamp-values timestampSlice   timestamp values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetTimestampSlice(fs, "timestamp-values")
	if err != nil {
		t.Errorf("unexpected error from GetTimestampSlice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	if err := fs.Parse([]string{"--timestamp-values", now.Format(time.RFC3339Nano), "--timestamp-values", tomorrow.Format(time.RFC3339Nano)}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetTimestampSlice(fs, "timestamp-values")
	if err != nil {
		t.Errorf("unexpected error from GetTimestampSlice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if len(value) != 2 {
		t.Errorf("unexpected slice length of %d", len(value))
	} else if !value[0].Equal(now) || !value[1].Equal(tomorrow) {
		t.Errorf("unexpected value in slice %v", value)
	}

	if strVal := flag.Value.String(); strVal != fmt.Sprintf("[%s,%s]", now.Format(time.RFC3339Nano), tomorrow.Format(time.RFC3339Nano)) {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	// pvalue, err := fs.GetTimestampSlice("time-values")
	// if err != nil {
	// 	t.Errorf("unexpected error from standard GetTimestampSlice: %v", err)
	// }
	// if !reflect.DeepEqual(pvalue, value) {
	// 	t.Errorf("standard GetTimestampSlice returned different value than GetTimestampSlice")
	// }
}

func TestStringTimestampMapValue(t *testing.T) {
	flag := flagsplugin.NewStringTimestampMapFlag("timestamp-map", "timestamp map")

	if ft := flag.Value.Type(); ft != "stringToTimestamp" {
		t.Errorf("flag type is %q (expected stringToTimestamp)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringTimestampMap(fs, "timestamp-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringTimestampMap on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--timestamp-map stringToTimestamp   timestamp map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringTimestampMap(fs, "timestamp-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringTimestampMap")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	if err := fs.Parse([]string{"--timestamp-map", fmt.Sprintf("foo=%s", now.Format(time.RFC3339Nano)), "--timestamp-map", fmt.Sprintf("bar=%s", tomorrow.Format(time.RFC3339Nano))}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringTimestampMap(fs, "timestamp-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringTimestampMap")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if len(value) != 2 {
		t.Errorf("unexpected slice length of %d", len(value))
	} else if !value["bar"].Equal(tomorrow) || !value["foo"].Equal(now) {
		t.Errorf("unexpected value in map %v", value)
	}

	if strVal := flag.Value.String(); strVal != fmt.Sprintf("[bar=%s,foo=%s]", tomorrow.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano)) {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
