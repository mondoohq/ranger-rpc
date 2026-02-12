// Copyright Mondoo, Inc. 2026
// SPDX-License-Identifier: MPL-2.0

package optional

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

type optionalServiceImpl struct{}

func (s *optionalServiceImpl) Check(ctx context.Context, in *OptionalRequest) (*OptionalReply, error) {
	return &OptionalReply{
		Text:   in.Text,
		Flag:   in.Flag,
		Number: in.Number,
		Custom: in.Custom,
	}, nil
}

func TestOptionalFields(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", NewOptionalServiceServer(&optionalServiceImpl{})))
	server := httptest.NewServer(mux)
	defer server.Close()

	client, err := NewOptionalServiceClient(server.URL+"/api/", &http.Client{})
	require.NoError(t, err)

	t.Run("with optional fields set", func(t *testing.T) {
		req := &OptionalRequest{
			Text:   proto.String("test"),
			Flag:   proto.Bool(true),
			Number: proto.Uint32(42),
			Custom: &CustomType{},
		}
		resp, err := client.Check(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, req.Text, resp.Text)
		require.Equal(t, req.Flag, resp.Flag)
		require.Equal(t, req.Number, resp.Number)
		require.True(t, proto.Equal(req.Custom, resp.Custom))
	})

	t.Run("with optional fields unset", func(t *testing.T) {
		req := &OptionalRequest{}
		resp, err := client.Check(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, req.Text, resp.Text)
		require.Equal(t, req.Flag, resp.Flag)
		require.Equal(t, req.Number, resp.Number)
		require.True(t, proto.Equal(req.Custom, resp.Custom))
	})

	t.Run("only some members are set", func(t *testing.T) {
		req := &OptionalRequest{
			Text: proto.String("test"),
		}
		resp, err := client.Check(context.Background(), req)
		require.NoError(t, err)
		require.Equal(t, req.Text, resp.Text)
		require.Equal(t, req.Flag, resp.Flag)
		require.Equal(t, req.Number, resp.Number)
		require.True(t, proto.Equal(req.Custom, resp.Custom))
	})
}
