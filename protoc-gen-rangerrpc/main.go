// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"go.mondoo.com/ranger-rpc/protoc-gen-rangerrpc/generator"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).
		RegisterModule(generator.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
