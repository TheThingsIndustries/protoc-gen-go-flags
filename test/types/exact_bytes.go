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

func GetExactBytes(fs *pflag.FlagSet, name string) (value []byte, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*ExactBytesValue).Value, flag.Changed, nil
}

type ExactBytesValue struct {
	length int
	Value  []byte
}

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

func (ebv *ExactBytesValue) Type() string {
	return fmt.Sprintf("%d-bytes", ebv.length)
}

func (ebv *ExactBytesValue) String() string {
	return hex.EncodeToString(ebv.Value)
}

func New8BytesFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesValue{length: 8},
	}
}

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

type ExactBytesSliceValue struct {
	length int
	Values []ExactBytesValue
}

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

func (ebv *ExactBytesSliceValue) Type() string {
	return fmt.Sprintf("%d-bytes", ebv.length)
}

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

func New8BytesSliceFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesSliceValue{length: 8},
	}
}
