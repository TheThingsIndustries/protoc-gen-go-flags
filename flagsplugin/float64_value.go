// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

func NewFloat64Flag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Float64Value{},
	}
}

func GetFloat64(fs *pflag.FlagSet, name string) (value float64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*Float64Value).Value, flag.Changed, nil
}

type Float64Value struct {
	Value float64
}

func (fv *Float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	fv.Value = v
	return err
}

func (*Float64Value) Type() string { return "float64" }

func (fv *Float64Value) String() string { return strconv.FormatFloat(fv.Value, 'g', -1, 64) }

func NewFloat64SliceFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Float64SliceValue{},
	}
}

func GetFloat64Slice(fs *pflag.FlagSet, name string) (value []float64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]float64, len(flag.Value.(*Float64SliceValue).Values))
	for i, v := range flag.Value.(*Float64SliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type Float64SliceValue struct {
	Values []Float64Value
}

func (fsv *Float64SliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var fv Float64Value
		if err := fv.Set(v); err != nil {
			return err
		}
		fsv.Values = append(fsv.Values, fv)
	}
	return nil
}

func (*Float64SliceValue) Type() string { return "float64Slice" }

func (fsv *Float64SliceValue) String() string {
	if len(fsv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(fsv.Values))
	for i, v := range fsv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

func NewStringFloat64MapFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringFloat64MapValue{},
	}
}

func GetStringFloat64Map(fs *pflag.FlagSet, name string) (value map[string]float64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]float64, len(flag.Value.(*StringFloat64MapValue).Values))
	for i, v := range flag.Value.(*StringFloat64MapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type StringFloat64MapValue struct {
	Values map[string]Float64Value
}

func (sfmv *StringFloat64MapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv Float64Value
		if err := fv.Set(v); err != nil {
			return err
		}
		if sfmv.Values == nil {
			sfmv.Values = make(map[string]Float64Value)
		}
		sfmv.Values[k] = fv
	}
	return nil
}

func (*StringFloat64MapValue) Type() string { return "stringToFloat64" }

func (sfmv *StringFloat64MapValue) String() string {
	if len(sfmv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(sfmv.Values))
	for k := range sfmv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(sfmv.Values))
	for _, k := range ks {
		v := sfmv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
