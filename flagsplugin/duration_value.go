// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

func NewDurationFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &DurationValue{},
	}
}

func GetDuration(fs *pflag.FlagSet, name string) (value time.Duration, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return 0, false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*DurationValue).Value, flag.Changed, nil
}

type DurationValue struct {
	Value time.Duration
}

func (dv *DurationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	dv.Value = v
	return err
}

func (*DurationValue) Type() string { return "duration" }

func (dv *DurationValue) String() string {
	if dv.Value == 0 {
		return ""
	}
	return dv.Value.String()
}

func NewDurationSliceFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:     name,
		Usage:    usage,
		Value:    &DurationSliceValue{},
		DefValue: "0",
	}
}

func GetDurationSlice(fs *pflag.FlagSet, name string) (value []time.Duration, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]time.Duration, len(flag.Value.(*DurationSliceValue).Values))
	for i, v := range flag.Value.(*DurationSliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type DurationSliceValue struct {
	Values []DurationValue
}

func (dsv *DurationSliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var dv DurationValue
		if err := dv.Set(v); err != nil {
			return err
		}
		dsv.Values = append(dsv.Values, dv)
	}
	return nil
}

func (*DurationSliceValue) Type() string { return "durationSlice" }

func (dsv *DurationSliceValue) String() string {
	if len(dsv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(dsv.Values))
	for i, v := range dsv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

func NewStringDurationMapFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringDurationMapValue{},
	}
}

func GetStringDurationMap(fs *pflag.FlagSet, name string) (value map[string]time.Duration, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]time.Duration, len(flag.Value.(*StringDurationMapValue).Values))
	for i, v := range flag.Value.(*StringDurationMapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type StringDurationMapValue struct {
	Values map[string]DurationValue
}

func (sdmv *StringDurationMapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv DurationValue
		if err := fv.Set(v); err != nil {
			return err
		}
		if sdmv.Values == nil {
			sdmv.Values = make(map[string]DurationValue)
		}
		sdmv.Values[k] = fv
	}
	return nil
}

func (*StringDurationMapValue) Type() string { return "stringToDuration" }

func (sdmv *StringDurationMapValue) String() string {
	if len(sdmv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(sdmv.Values))
	for k := range sdmv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(sdmv.Values))
	for _, k := range ks {
		v := sdmv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s=%s`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
