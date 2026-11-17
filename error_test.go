// Copyright Mondoo, Inc. 2026
// SPDX-License-Identifier: MPL-2.0

package ranger_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	pb "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

func TestHttpErrorRetryAfter(t *testing.T) {
	busy, err := status.New(codes.ResourceExhausted, "busy").
		WithDetails(&errdetails.RetryInfo{RetryDelay: durationpb.New(2500 * time.Millisecond)})
	require.NoError(t, err)

	tests := []struct {
		name string
		err  error
		code int
		want string
	}{
		{"RetryInfo rounds up to whole seconds", busy.Err(), http.StatusTooManyRequests, "3"},
		{"no RetryInfo, no header", status.Error(codes.ResourceExhausted, "busy"), http.StatusTooManyRequests, ""},
		{"other errors unchanged", status.Error(codes.NotFound, "nope"), http.StatusNotFound, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := ranger.NewRPCServer(&ranger.Service{
				Name: "PingPong",
				Methods: map[string]ranger.Method{
					"Ping": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
						return nil, tc.err
					},
				},
			})
			req := httptest.NewRequest("POST", "http://example.com/PingPong/Ping", nil)
			req.Header.Set("Content-Type", "application/protobuf")
			w := httptest.NewRecorder()
			srv.ServeHTTP(w, req)

			resp := w.Result()
			assert.Equal(t, tc.code, resp.StatusCode)
			assert.Equal(t, tc.want, resp.Header.Get("Retry-After"))
		})
	}
}
