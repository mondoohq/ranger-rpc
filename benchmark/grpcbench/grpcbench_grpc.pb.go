// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: grpcbench.proto

package grpcbench

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BenchmarkServiceClient is the client API for BenchmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BenchmarkServiceClient interface {
	RpcSmall(ctx context.Context, in *SmallQuery, opts ...grpc.CallOption) (*SmallResponse, error)
	RpcEmpty(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type benchmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkServiceClient(cc grpc.ClientConnInterface) BenchmarkServiceClient {
	return &benchmarkServiceClient{cc}
}

func (c *benchmarkServiceClient) RpcSmall(ctx context.Context, in *SmallQuery, opts ...grpc.CallOption) (*SmallResponse, error) {
	out := new(SmallResponse)
	err := c.cc.Invoke(ctx, "/mondoo.ranger.grpcbench.BenchmarkService/RpcSmall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *benchmarkServiceClient) RpcEmpty(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/mondoo.ranger.grpcbench.BenchmarkService/RpcEmpty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BenchmarkServiceServer is the server API for BenchmarkService service.
// All implementations must embed UnimplementedBenchmarkServiceServer
// for forward compatibility
type BenchmarkServiceServer interface {
	RpcSmall(context.Context, *SmallQuery) (*SmallResponse, error)
	RpcEmpty(context.Context, *Empty) (*DefaultResponse, error)
	mustEmbedUnimplementedBenchmarkServiceServer()
}

// UnimplementedBenchmarkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBenchmarkServiceServer struct {
}

func (UnimplementedBenchmarkServiceServer) RpcSmall(context.Context, *SmallQuery) (*SmallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcSmall not implemented")
}
func (UnimplementedBenchmarkServiceServer) RpcEmpty(context.Context, *Empty) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RpcEmpty not implemented")
}
func (UnimplementedBenchmarkServiceServer) mustEmbedUnimplementedBenchmarkServiceServer() {}

// UnsafeBenchmarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BenchmarkServiceServer will
// result in compilation errors.
type UnsafeBenchmarkServiceServer interface {
	mustEmbedUnimplementedBenchmarkServiceServer()
}

func RegisterBenchmarkServiceServer(s grpc.ServiceRegistrar, srv BenchmarkServiceServer) {
	s.RegisterService(&BenchmarkService_ServiceDesc, srv)
}

func _BenchmarkService_RpcSmall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmallQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchmarkServiceServer).RpcSmall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mondoo.ranger.grpcbench.BenchmarkService/RpcSmall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchmarkServiceServer).RpcSmall(ctx, req.(*SmallQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _BenchmarkService_RpcEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BenchmarkServiceServer).RpcEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mondoo.ranger.grpcbench.BenchmarkService/RpcEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BenchmarkServiceServer).RpcEmpty(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BenchmarkService_ServiceDesc is the grpc.ServiceDesc for BenchmarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BenchmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mondoo.ranger.grpcbench.BenchmarkService",
	HandlerType: (*BenchmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcSmall",
			Handler:    _BenchmarkService_RpcSmall_Handler,
		},
		{
			MethodName: "RpcEmpty",
			Handler:    _BenchmarkService_RpcEmpty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcbench.proto",
}