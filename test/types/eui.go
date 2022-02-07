// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
)

type EUI64 [8]byte

func (t EUI64) Marshal() ([]byte, error) {
	return t[:], nil
}

func (t *EUI64) MarshalTo(data []byte) (n int, err error) {
	return copy(data, t[:]), nil
}

func (t *EUI64) Unmarshal(data []byte) error {
	if len(data) != 8 {
		return fmt.Errorf("invalid data length: got %d, want 8", len(data))
	}
	var dto EUI64
	copy(dto[:], data)
	*t = dto
	return nil
}

func (t *EUI64) Size() int { return 8 }

func GetEUI64(fs *pflag.FlagSet, name string) (value EUI64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return EUI64{}, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	var eui64 EUI64
	if !flag.Changed {
		return eui64, flag.Changed, nil
	}
	if err := eui64.Unmarshal(flag.Value.(*ExactBytesValue).Value); err != nil {
		return eui64, false, err
	}
	return eui64, flag.Changed, nil
}

func GetEUI64Slice(fs *pflag.FlagSet, name string) (value []EUI64, set bool, err error) {
	name = toDash.Replace(name)
	flag := fs.Lookup(name)
	if flag == nil {
		return nil, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	value = make([]EUI64, len(flag.Value.(*ExactBytesSliceValue).Values))
	for i, v := range flag.Value.(*ExactBytesSliceValue).Values {
		var eui64 EUI64
		if err := eui64.Unmarshal(v.Value); err != nil {
			return nil, false, err
		}
		value[i] = eui64
	}
	return value, flag.Changed, nil
}
