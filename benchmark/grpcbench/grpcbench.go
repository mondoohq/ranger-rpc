// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package grpcbench

import (
	fmt "fmt"
	"net"

	"go.mondoo.com/ranger-rpc/benchmark/sample"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpcbench.proto

type BenchmarkServer struct {
	UnimplementedBenchmarkServiceServer
}

func (b *BenchmarkServer) RpcSmall(ctx context.Context, in *SmallQuery) (*SmallResponse, error) {
	return &SmallResponse{Id: in.Id, Message: in.Message, Name: in.Name}, nil
}
func (b *BenchmarkServer) RpcEmpty(ctx context.Context, in *Empty) (*DefaultResponse, error) {
	return &DefaultResponse{Message: sample.Message}, nil
}

func Serve(port int) int {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(fmt.Sprintf("GRPC error %v", err))
	}

	s := grpc.NewServer()
	RegisterBenchmarkServiceServer(s, &BenchmarkServer{})

	go func() {
		err := s.Serve(lis)
		if err != nil {
			fmt.Printf("GRPC error %v", err)
		}
	}()

	return lis.Addr().(*net.TCPAddr).Port
}
