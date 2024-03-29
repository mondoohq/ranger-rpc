// Code generated by protoc-gen-rangerrpc version DO NOT EDIT.
// source: hello.proto

package rangerguard

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	ranger "go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/metadata"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
)

// service interface definition

type HelloWorld interface {
	Hello(context.Context, *HelloReq) (*HelloResp, error)
	Info(context.Context, *Empty) (*Tags, error)
}

// client implementation

type HelloWorldClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewHelloWorldClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*HelloWorldClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./HelloWorld")
	if err != nil {
		return nil, err
	}

	serviceClient := &HelloWorldClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *HelloWorldClient) Hello(ctx context.Context, in *HelloReq) (*HelloResp, error) {
	out := new(HelloResp)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Hello"}, ""), in, out)
	return out, err
}
func (c *HelloWorldClient) Info(ctx context.Context, in *Empty) (*Tags, error) {
	out := new(Tags)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Info"}, ""), in, out)
	return out, err
}

// server implementation

type HelloWorldServerOption func(s *HelloWorldServer)

func WithUnknownFieldsForHelloWorldServer() HelloWorldServerOption {
	return func(s *HelloWorldServer) {
		s.allowUnknownFields = true
	}
}

func NewHelloWorldServer(handler HelloWorld, opts ...HelloWorldServerOption) http.Handler {
	srv := &HelloWorldServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "HelloWorld",
		Methods: map[string]ranger.Method{
			"Hello": srv.Hello,
			"Info":  srv.Info,
		},
	}
	return ranger.NewRPCServer(&service)
}

type HelloWorldServer struct {
	handler            HelloWorld
	allowUnknownFields bool
}

func (p *HelloWorldServer) Hello(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req HelloReq
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Hello(ctx, &req)
}
func (p *HelloWorldServer) Info(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req Empty
	var err error

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not access header")
	}

	switch md.First("Content-Type") {
	case "application/protobuf", "application/octet-stream", "application/grpc+proto":
		err = pb.Unmarshal(*reqBytes, &req)
	default:
		// handle case of empty object
		if len(*reqBytes) > 0 {
			err = jsonpb.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(*reqBytes, &req)
		}
	}

	if err != nil {
		return nil, err
	}
	return p.handler.Info(ctx, &req)
}
