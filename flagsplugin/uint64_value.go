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

func NewUint64Flag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Uint64Value{},
	}
}

func GetUint64(fs *pflag.FlagSet, name string) (value uint64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*Uint64Value).Value, flag.Changed, nil
}

type Uint64Value struct {
	Value uint64
}

func (uv *Uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	uv.Value = v
	return err
}

func (*Uint64Value) Type() string { return "uint64" }

func (uv *Uint64Value) String() string { return strconv.FormatUint(uv.Value, 10) }

func NewUint64SliceFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &Uint64SliceValue{},
	}
}

func GetUint64Slice(fs *pflag.FlagSet, name string) (value []uint64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]uint64, len(flag.Value.(*Uint64SliceValue).Values))
	for i, v := range flag.Value.(*Uint64SliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type Uint64SliceValue struct {
	Values []Uint64Value
}

func (usv *Uint64SliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var uv Uint64Value
		if err := uv.Set(v); err != nil {
			return err
		}
		usv.Values = append(usv.Values, uv)
	}
	return nil
}

func (*Uint64SliceValue) Type() string { return "uint64Slice" }

func (usv *Uint64SliceValue) String() string {
	if len(usv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(usv.Values))
	for i, v := range usv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

func NewStringUint64MapFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringUint64MapValue{},
	}
}

func GetStringUint64Map(fs *pflag.FlagSet, name string) (value map[string]uint64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]uint64, len(flag.Value.(*StringUint64MapValue).Values))
	for i, v := range flag.Value.(*StringUint64MapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type StringUint64MapValue struct {
	Values map[string]Uint64Value
}

func (sumv *StringUint64MapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv Uint64Value
		if err := fv.Set(v); err != nil {
			return err
		}
		if sumv.Values == nil {
			sumv.Values = make(map[string]Uint64Value)
		}
		sumv.Values[k] = fv
	}
	return nil
}

func (*StringUint64MapValue) Type() string { return "stringToUint64" }

func (sumv *StringUint64MapValue) String() string {
	if len(sumv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(sumv.Values))
	for k := range sumv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(sumv.Values))
	for _, k := range ks {
		v := sumv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
