// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (g *generator) getFieldFlagBoolOptions(field *protogen.Field) (genSelect *bool, genSet *bool, hidden bool) {
	fieldOpts := field.Desc.Options().(*descriptorpb.FieldOptions)
	if proto.HasExtension(fieldOpts, annotations.E_Field) {
		if fieldExt, ok := proto.GetExtension(fieldOpts, annotations.E_Field).(*annotations.FieldOptions); ok {
			// If field has a skip_set flag set specifically, override message annotation.
			return fieldExt.Select, fieldExt.Set, fieldExt.GetHidden()
		}
	}
	return nil, nil, false
}
