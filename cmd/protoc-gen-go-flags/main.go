// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"

	plugin "github.com/TheThingsIndustries/protoc-gen-go-flags/internal/gen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-flags %v\n", plugin.Version)
		return
	}

	var flags flag.FlagSet
	flags.StringVar(&plugin.Params.Lang, "lang", "go", "language (go)")
	flags.StringVar(&plugin.Params.CustomTypeGetterPrefix, "customtype.getter-prefix", "Get", "prefix for customtype getter func")
	flags.StringVar(&plugin.Params.CustomTypeGetterSuffix, "customtype.getter-suffix", "", "suffix for customtype getter func")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if f.Desc.Syntax() != protoreflect.Proto3 {
				return fmt.Errorf("the protoc-gen-go-flags plugin only supports proto3 syntax")
			}
			plugin.GenerateFile(gen, f)
		}
		return nil
	})
}
