// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package rangerguard

// To regenerate the protocol buffer output for this package, run
//	go generate

//go:generate protoc --go_out=. --go_opt=paths=source_relative --rangerrpc_out=. hello.proto
