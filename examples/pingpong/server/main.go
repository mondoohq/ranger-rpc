// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"net/http"

	"go.mondoo.com/ranger-rpc/examples/pingpong"
)

func main() {
	fmt.Println("start pingpong example server")

	// init service implementation and attach it to the standard mux router for the /api/ prefix
	mux := http.DefaultServeMux
	p := &pingpong.PingPongServiceImpl{}
	mux.Handle("/api/", http.StripPrefix("/api", pingpong.NewPingPongServer(p)))

	// start the server on port 2155
	port := 2155
	fmt.Printf("listen on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil)
	if err != nil {
		fmt.Println("was not able to start the server")
		panic(err)
	}
}
