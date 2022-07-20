package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"go.mondoo.com/ranger-rpc/protoc-gen-rangerrpc-swagger/swagger"
)

var version string

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).
		RegisterModule(swagger.New()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
