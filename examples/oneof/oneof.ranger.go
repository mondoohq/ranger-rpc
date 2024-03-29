// Code generated by protoc-gen-rangerrpc version DO NOT EDIT.
// source: oneof.proto

package oneof

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

type OneOf interface {
	Echo(context.Context, *OneOfRequest) (*OneOfReply, error)
}

// client implementation

type OneOfClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewOneOfClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*OneOfClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./OneOf")
	if err != nil {
		return nil, err
	}

	serviceClient := &OneOfClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugins(plugins...)
	return serviceClient, nil
}
func (c *OneOfClient) Echo(ctx context.Context, in *OneOfRequest) (*OneOfReply, error) {
	out := new(OneOfReply)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Echo"}, ""), in, out)
	return out, err
}

// server implementation

type OneOfServerOption func(s *OneOfServer)

func WithUnknownFieldsForOneOfServer() OneOfServerOption {
	return func(s *OneOfServer) {
		s.allowUnknownFields = true
	}
}

func NewOneOfServer(handler OneOf, opts ...OneOfServerOption) http.Handler {
	srv := &OneOfServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "OneOf",
		Methods: map[string]ranger.Method{
			"Echo": srv.Echo,
		},
	}
	return ranger.NewRPCServer(&service)
}

type OneOfServer struct {
	handler            OneOf
	allowUnknownFields bool
}

func (p *OneOfServer) Echo(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req OneOfRequest
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
	return p.handler.Echo(ctx, &req)
}
