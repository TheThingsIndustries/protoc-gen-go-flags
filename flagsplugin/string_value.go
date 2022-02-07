// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/pflag"
)

func NewStringFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringValue{},
	}
}

func GetString(fs *pflag.FlagSet, name string) (value string, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return "", false, &ErrFlagNotFound{FlagName: name}
	}
	return flag.Value.(*StringValue).Value, flag.Changed, nil
}

type StringValue struct {
	Value string
}

func (sv *StringValue) Set(s string) error {
	sv.Value = s
	return nil
}

func (*StringValue) Type() string { return "string" }

func (sv *StringValue) String() string { return sv.Value }

func NewStringSliceFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringSliceValue{},
	}
}

func GetStringSlice(fs *pflag.FlagSet, name string) (value []string, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make([]string, len(flag.Value.(*StringSliceValue).Values))
	for i, v := range flag.Value.(*StringSliceValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type StringSliceValue struct {
	Values []StringValue
}

func (ssv *StringSliceValue) Set(s string) error {
	vs, err := SplitSliceElements(s)
	if err != nil {
		return err
	}
	for _, v := range vs {
		var sv StringValue
		if err := sv.Set(v); err != nil {
			return err
		}
		ssv.Values = append(ssv.Values, sv)
	}
	return nil
}

func (*StringSliceValue) Type() string { return "stringSlice" }

func (ssv *StringSliceValue) String() string {
	if len(ssv.Values) == 0 {
		return ""
	}
	vs := make([]string, len(ssv.Values))
	for i, v := range ssv.Values {
		vs[i] = v.String()
	}
	return "[" + JoinSliceElements(vs) + "]"
}

func NewStringStringMapFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &StringStringMapValue{},
	}
}

func GetStringStringMap(fs *pflag.FlagSet, name string) (value map[string]string, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &ErrFlagNotFound{FlagName: name}
	}
	value = make(map[string]string, len(flag.Value.(*StringStringMapValue).Values))
	for i, v := range flag.Value.(*StringStringMapValue).Values {
		value[i] = v.Value
	}
	return value, flag.Changed, nil
}

type StringStringMapValue struct {
	Values map[string]StringValue
}

func (ssmv *StringStringMapValue) Set(s string) error {
	kv, err := splitStringMapElements(s)
	if err != nil {
		return err
	}
	for k, v := range kv {
		var fv StringValue
		if err := fv.Set(v); err != nil {
			return err
		}
		if ssmv.Values == nil {
			ssmv.Values = make(map[string]StringValue)
		}
		ssmv.Values[k] = fv
	}
	return nil
}

func (*StringStringMapValue) Type() string { return "stringToString" }

func (ssmv *StringStringMapValue) String() string {
	if len(ssmv.Values) == 0 {
		return ""
	}
	ks := make([]string, 0, len(ssmv.Values))
	for k := range ssmv.Values {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vs := make([]string, 0, len(ssmv.Values))
	for _, k := range ks {
		v := ssmv.Values[k]
		vs = append(vs, fmt.Sprintf(`%s="%s"`, k, v.String()))
	}
	return "[" + strings.Join(vs, ",") + "]"
}
