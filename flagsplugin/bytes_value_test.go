// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
)

func TestBase64BytesValue(t *testing.T) {
	flag := flagsplugin.NewBase64BytesFlag("bytes-value", "bytes value")

	if ft := flag.Value.Type(); ft != "bytesBase64" {
		t.Errorf("flag type is %q (expected bytesBase64)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %q (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetBytes(fs, "bytes-value")
	if err == nil {
		t.Errorf("expected to get error from GetBytes on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bytes-value bytesBase64   bytes value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetBytes(fs, "bytes-value")
	if err != nil {
		t.Errorf("unexpected error from GetBytes")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bytes-value", "Zm9v"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetBytes(fs, "bytes-value")
	if err != nil {
		t.Errorf("unexpected error from GetBytes")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !bytes.Equal(value, []byte("foo")) {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "Zm9v" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetBytesBase64("bytes-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetBytesBase64: %v", err)
	}
	if !bytes.Equal(pvalue, value) {
		t.Errorf("standard GetBytesBase64 returned different value than GetBytesBase64")
	}
}

func TestHexBytesValue(t *testing.T) {
	flag := flagsplugin.NewHexBytesFlag("bytes-value", "bytes value")

	if ft := flag.Value.Type(); ft != "bytesHex" {
		t.Errorf("flag type is %q (expected bytesHex)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %q (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetBytes(fs, "bytes-value")
	if err == nil {
		t.Errorf("expected to get error from GetBytes on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bytes-value bytesHex   bytes value" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetBytes(fs, "bytes-value")
	if err != nil {
		t.Errorf("unexpected error from GetBytes")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bytes-value", "666f6f"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetBytes(fs, "bytes-value")
	if err != nil {
		t.Errorf("unexpected error from GetBytes")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !bytes.Equal(value, []byte("foo")) {
		t.Errorf("unexpected %v value in parsed flag set", value)
	}

	if strVal := flag.Value.String(); strVal != "666f6f" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}

	pvalue, err := fs.GetBytesHex("bytes-value")
	if err != nil {
		t.Errorf("unexpected error from standard GetBytesHex: %v", err)
	}
	if !bytes.Equal(pvalue, value) {
		t.Errorf("standard GetBytesHex returned different value than GetBytesHex")
	}
}

func TestBase64BytesSliceValue(t *testing.T) {
	flag := flagsplugin.NewBase64BytesSliceFlag("bytes-values", "bytes values")

	if ft := flag.Value.Type(); ft != "bytesBase64" {
		t.Errorf("flag type is %q (expected bytesBase64)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetBytesSlice(fs, "bytes-values")
	if err == nil {
		t.Errorf("expected to get error from GetBytesSlice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bytes-values bytesBase64   bytes values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetBytesSlice(fs, "bytes-values")
	if err != nil {
		t.Errorf("unexpected error from GetBytesSlice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bytes-values", "Zm9v", "--bytes-values", "YmFy"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetBytesSlice(fs, "bytes-values")
	if err != nil {
		t.Errorf("unexpected error from GetBytesSlice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, [][]byte{[]byte("foo"), []byte("bar")}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[Zm9v,YmFy]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}

func TestHexBytesSliceValue(t *testing.T) {
	flag := flagsplugin.NewHexBytesSliceFlag("bytes-values", "bytes values")

	if ft := flag.Value.Type(); ft != "bytesHex" {
		t.Errorf("flag type is %q (expected bytesHex)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetBytesSlice(fs, "bytes-values")
	if err == nil {
		t.Errorf("expected to get error from GetBytesSlice on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bytes-values bytesHex   bytes values" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetBytesSlice(fs, "bytes-values")
	if err != nil {
		t.Errorf("unexpected error from GetBytesSlice")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bytes-values", "666f6f", "--bytes-values", "626172"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetBytesSlice(fs, "bytes-values")
	if err != nil {
		t.Errorf("unexpected error from GetBytesSlice")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, [][]byte{[]byte("foo"), []byte("bar")}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[666f6f,626172]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}

func TestStringBase64BytesMapValue(t *testing.T) {
	flag := flagsplugin.NewStringBase64BytesMapFlag("bytes-map", "bytes map")

	if ft := flag.Value.Type(); ft != "stringToBytesBase64" {
		t.Errorf("flag type is %q (expected stringToBytesBase64)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringBytesMap(fs, "bytes-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringBytesMap on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bytes-map stringToBytesBase64   bytes map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringBytesMap(fs, "bytes-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringBytesMap")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bytes-map", "foo=Zm9v", "--bytes-map", "bar=YmFy"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringBytesMap(fs, "bytes-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringBytesMap")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string][]byte{"foo": []byte("foo"), "bar": []byte("bar")}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=YmFy,foo=Zm9v]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}

func TestStringHexBytesMapValue(t *testing.T) {
	flag := flagsplugin.NewStringHexBytesMapFlag("bytes-map", "bytes map")

	if ft := flag.Value.Type(); ft != "stringToBytesHex" {
		t.Errorf("flag type is %q (expected stringToBytesHex)", ft)
	}

	if fv := flag.Value.String(); fv != "" {
		t.Errorf("flag value is %v (expected \"\")", fv)
	}

	fs := &pflag.FlagSet{}

	_, _, err := flagsplugin.GetStringBytesMap(fs, "bytes-map")
	if err == nil {
		t.Errorf("expected to get error from GetStringBytesMap on empty flag set")
	}

	fs.AddFlag(flag)

	if fd := strings.TrimSpace(fs.FlagUsages()); fd != "--bytes-map stringToBytesHex   bytes map" {
		t.Errorf("unexpected flag description: %q", fd)
	}

	value, changed, err := flagsplugin.GetStringBytesMap(fs, "bytes-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringBytesMap")
	}
	if changed {
		t.Errorf("unexpected changed value in unparsed flag set")
	}
	if len(value) > 0 {
		t.Errorf("unexpected non-default value in unparsed flag set")
	}

	if err := fs.Parse([]string{"--bytes-map", "foo=666f6f", "--bytes-map", "bar=626172"}); err != nil {
		t.Errorf("error in FlagSet.Parse: %v", err)
	}

	value, changed, err = flagsplugin.GetStringBytesMap(fs, "bytes-map")
	if err != nil {
		t.Errorf("unexpected error from GetStringBytesMap")
	}
	if !changed {
		t.Errorf("unexpected unchanged value in parsed flag set")
	}
	if !reflect.DeepEqual(value, map[string][]byte{"foo": []byte("foo"), "bar": []byte("bar")}) {
		t.Errorf("unexpected slice value (%v)", value)
	}

	if strVal := flag.Value.String(); strVal != "[bar=626172,foo=666f6f]" {
		t.Errorf("unexpected String value %q for parsed flag value", strVal)
	}
}
