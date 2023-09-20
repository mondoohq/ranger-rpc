// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package ranger

// Ranger RPC is a simple and fast proto-based RPC framework
//
// It is designed to stay as close as possible to GRPC without the complexity that GRPC brings. Essentially the support
// for plain json endpoints require the use of the GRPC and GRPC Gateway.
//
// To make the use of Ranger RPC as easy as possible, the following features are provided:
//
// We use the same GRPC codes
// https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
//
// We use GRPC-compatible HTTP status code mapping
// https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
//
// GRPC Status Blog
// https://jbrandhorst.com/post/grpc-errors/
