// Copyright Mondoo, Inc. 2026
// SPDX-License-Identifier: MPL-2.0

package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"go.mondoo.com/ranger-rpc/protoc-gen-rangerrpc/generator"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	features := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
		pgs.SupportedFeatures(&features),
	).
		RegisterModule(generator.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
