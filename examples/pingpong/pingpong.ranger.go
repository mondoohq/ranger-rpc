// Code generated by protoc-gen-rangerrpc version DO NOT EDIT.
// source: pingpong.proto

package pingpong

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

// client implementation

type PingPongClient struct {
	ranger.Client
	httpclient ranger.HTTPClient
	prefix     string
}

func NewPingPongClient(addr string, client ranger.HTTPClient, plugins ...ranger.ClientPlugin) (*PingPongClient, error) {
	base, err := url.Parse(ranger.SanitizeUrl(addr))
	if err != nil {
		return nil, err
	}

	u, err := url.Parse("./PingPong")
	if err != nil {
		return nil, err
	}

	serviceClient := &PingPongClient{
		httpclient: client,
		prefix:     base.ResolveReference(u).String(),
	}
	serviceClient.AddPlugin(plugins...)
	return serviceClient, nil
}
func (c *PingPongClient) Ping(ctx context.Context, in *PingRequest) (*PongReply, error) {
	out := new(PongReply)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/Ping"}, ""), in, out)
	return out, err
}
func (c *PingPongClient) NoPing(ctx context.Context, in *Empty) (*PongReply, error) {
	out := new(PongReply)
	err := c.DoClientRequest(ctx, c.httpclient, strings.Join([]string{c.prefix, "/NoPing"}, ""), in, out)
	return out, err
}

// service interface definition

type PingPong interface {
	Ping(context.Context, *PingRequest) (*PongReply, error)
	NoPing(context.Context, *Empty) (*PongReply, error)
}

// server implementation

type PingPongServerOption func(s *PingPongServer)

func WithUnknownFieldsForPingPongServer() PingPongServerOption {
	return func(s *PingPongServer) {
		s.allowUnknownFields = true
	}
}

func NewPingPongServer(handler PingPong, opts ...PingPongServerOption) http.Handler {
	srv := &PingPongServer{
		handler: handler,
	}

	for i := range opts {
		opts[i](srv)
	}

	service := ranger.Service{
		Name: "PingPong",
		Methods: map[string]ranger.Method{
			"Ping":   srv.Ping,
			"NoPing": srv.NoPing,
		},
	}
	return ranger.NewRPCServer(&service)
}

type PingPongServer struct {
	handler            PingPong
	allowUnknownFields bool
}

func (p *PingPongServer) Ping(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
	var req PingRequest
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
	return p.handler.Ping(ctx, &req)
}
func (p *PingPongServer) NoPing(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
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
	return p.handler.NoPing(ctx, &req)
}
