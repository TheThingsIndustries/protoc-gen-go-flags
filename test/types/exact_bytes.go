// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
)

// GetExactBytes returns a value from a bytes flag.
func GetExactBytes(fs *pflag.FlagSet, name string) (value []byte, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*ExactBytesValue).Value, flag.Changed, nil
}

// ExactBytesValue implements pflag.Value interface.
type ExactBytesValue struct {
	length int
	Value  []byte
}

// Set implements pflag.Value interface.
func (ebv *ExactBytesValue) Set(s string) error {
	trimmed := strings.TrimSuffix(s, "=")
	switch len(trimmed) {
	case hex.EncodedLen(ebv.length):
		b, err := hex.DecodeString(trimmed)
		if err != nil {
			return err
		}
		ebv.Value = b
	case base64.RawStdEncoding.EncodedLen(ebv.length):
		b, err := base64.RawStdEncoding.DecodeString(flagsplugin.Base64Replacer.Replace(trimmed))
		if err != nil {
			return err
		}
		ebv.Value = b
	default:
		return fmt.Errorf("Invalid bytes length: want %d got %d", ebv.length, len(trimmed))
	}
	return nil
}

// Type implements pflag.Value interface.
func (ebv *ExactBytesValue) Type() string {
	return fmt.Sprintf("%d-bytes", ebv.length)
}

// String implements pflag.Value interface.
func (ebv *ExactBytesValue) String() string {
	return hex.EncodeToString(ebv.Value)
}

// New8BytesFlag defines a new flag that holds a byte array of length 8.
func New8BytesFlag(name, usage string, opts ...flagsplugin.FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesValue{length: 8},
	}
	flagsplugin.ApplyOptions(flag, opts...)
	return flag
}

// GetExactBytesSlice returns a value from a byte flag.
func GetExactBytesSlice(fs *pflag.FlagSet, name string) (value [][]byte, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	value = make([][]byte, len(flag.Value.(*ExactBytesSliceValue).Values))
	for i, v := range flag.Value.(*ExactBytesSliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

// ExactBytesSliceValue implements pflag.Value interface.
type ExactBytesSliceValue struct {
	length int
	Values []ExactBytesValue
}

// Set implements pflag.Value interface.
func (ebv *ExactBytesSliceValue) Set(s string) error {
	vs, err := flagsplugin.SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		ev := ExactBytesValue{length: ebv.length}
		if err := ev.Set(v); err != nil {
			return err
		}
		ebv.Values = append(ebv.Values, ev)
	}
	return nil
}

// Type implements pflag.Value interface.
func (ebv *ExactBytesSliceValue) Type() string {
	return fmt.Sprintf("%d-bytes", ebv.length)
}

// String implements pflag.Value interface.
func (ebv *ExactBytesSliceValue) String() string {
	if len(ebv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(ebv.Values))
	for i, v := range ebv.Values {
		vs[i] = v.String()
	}
	return "[" + flagsplugin.JoinSliceElements(vs) + "]"
}

// New8BytesSliceFlag defines a new flag that holds a slice of byte arrays of length 8.
func New8BytesSliceFlag(name, usage string, opts ...flagsplugin.FlagOption) *pflag.Flag {
	flag := &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesSliceValue{length: 8},
	}
	flagsplugin.ApplyOptions(flag, opts...)
	return flag
}
