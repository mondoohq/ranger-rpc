// Copyright Mondoo, Inc. 2026
// SPDX-License-Identifier: MPL-2.0

package ranger_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	ranger "go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/examples/pingpong"
	"go.mondoo.com/ranger-rpc/status"
	"google.golang.org/protobuf/proto"
)

// TestDoClientRequestHappyPath verifies the normal success flow: the request
// proto is marshalled and sent, and a 200 response whose body is a valid
// protobuf message is unmarshalled back into the output message. The server
// echoes the request's Sender to prove the round trip end to end.
func TestDoClientRequestHappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var ping pingpong.PingRequest
		require.NoError(t, proto.Unmarshal(reqBody, &ping))

		respBody, err := proto.Marshal(&pingpong.PongReply{Message: "pong: " + ping.Sender})
		require.NoError(t, err)

		w.Header().Set("Content-Type", "application/protobuf")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respBody)
	}))
	defer srv.Close()

	c := &ranger.Client{}
	out := &pingpong.PongReply{}
	err := c.DoClientRequest(
		context.Background(),
		srv.Client(),
		srv.URL,
		&pingpong.PingRequest{Sender: "ping"},
		out,
	)

	require.NoError(t, err)
	assert.Equal(t, "pong: ping", out.Message, "the response body must be unmarshalled into the output message")
}

// TestDoClientRequestNonProtoErrorBody reproduces the failure mode where an
// intermediary in front of the API (load balancer / ingress / proxy) returns a
// non-200 response whose body is NOT a protobuf google.rpc.Status — for example
// an HTML or plain-text error page accompanying a 502 Bad Gateway.
//
// The client is supposed to fall back to using that raw body as the error
// message. Before the fix it instead read the already-consumed *request* body
// reader, so the message was always empty and the error surfaced as the
// famously uninformative "rpc error: code = Unknown desc = " (empty
// description).
func TestDoClientRequestNonProtoErrorBody(t *testing.T) {
	const body = "<html><head><title>502 Bad Gateway</title></head><body>upstream connect error</body></html>"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 502 is not in the reverse status map, so it maps to codes.Unknown,
		// and an HTML body cannot be decoded as a proto Status.
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()

	c := &ranger.Client{}
	err := c.DoClientRequest(
		context.Background(),
		srv.Client(),
		srv.URL,
		&pingpong.PingRequest{},
		&pingpong.PongReply{},
	)

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.Unknown, st.Code(), "502 with a non-proto body maps to Unknown")
	assert.NotEmpty(t, st.Message(), "the upstream error body must be surfaced, not dropped as an empty description")
	assert.Contains(t, st.Message(), "502 Bad Gateway", "the description should carry the upstream error body")
}

// TestDoClientRequestEmptyErrorBody is the sibling case to
// TestDoClientRequestNonProtoErrorBody: a non-200 response with an *empty*
// body. An empty body unmarshals without error into an OK-coded proto Status,
// so status.FromProto(...).Err() is nil. Without a guard this made
// DoClientRequest silently return nil (success) for a failed request, leaving
// the output message unpopulated. The client must instead surface a real error,
// falling back to the HTTP status line since there is no body to report.
func TestDoClientRequestEmptyErrorBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway) // 502 with no body
	}))
	defer srv.Close()

	c := &ranger.Client{}
	err := c.DoClientRequest(
		context.Background(),
		srv.Client(),
		srv.URL,
		&pingpong.PingRequest{},
		&pingpong.PongReply{},
	)

	require.Error(t, err, "a non-200 with an empty body must not be swallowed as success")
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.Unknown, st.Code(), "502 maps to Unknown")
	assert.NotEmpty(t, st.Message(), "an empty body should fall back to the HTTP status line, not an empty description")
	assert.Contains(t, st.Message(), "502", "the HTTP status line should be surfaced when the body is empty")
}

// TestDoClientRequestProtoStatusError verifies the primary error path is
// unchanged: a non-200 response whose body IS a valid protobuf
// google.rpc.Status is decoded and returned with its original code and message,
// not the HTTP-status-derived fallback. This is the case the earlier fix wraps
// in a guard, so it must still work end to end.
func TestDoClientRequestProtoStatusError(t *testing.T) {
	const msg = "you shall not pass"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Build the body exactly as the server side does: marshal the proto
		// Status carried by a status error.
		body, err := proto.Marshal(status.New(codes.PermissionDenied, msg).Proto())
		require.NoError(t, err)

		w.Header().Set("Content-Type", "application/protobuf")
		w.WriteHeader(http.StatusForbidden) // 403, consistent with PermissionDenied
		_, _ = w.Write(body)
	}))
	defer srv.Close()

	c := &ranger.Client{}
	err := c.DoClientRequest(
		context.Background(),
		srv.Client(),
		srv.URL,
		&pingpong.PingRequest{},
		&pingpong.PongReply{},
	)

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.PermissionDenied, st.Code(), "the code must come from the proto Status body")
	assert.Equal(t, msg, st.Message(), "the message must come from the proto Status body, not the HTTP fallback")
}

// TestDoClientRequestLargeErrorBodyTruncated verifies that a large non-proto
// error body (e.g. a big HTML error page from an intermediary) is capped rather
// than surfaced in full, so the returned error message stays bounded.
func TestDoClientRequestLargeErrorBodyTruncated(t *testing.T) {
	// A body far larger than the cap, all valid UTF-8.
	body := strings.Repeat("A", 64*1024)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()

	c := &ranger.Client{}
	err := c.DoClientRequest(
		context.Background(),
		srv.Client(),
		srv.URL,
		&pingpong.PingRequest{},
		&pingpong.PongReply{},
	)

	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	// The message must be bounded well below the body size, and flagged as cut.
	assert.Less(t, len(st.Message()), len(body), "the oversized body must not be surfaced in full")
	assert.LessOrEqual(t, len(st.Message()), 4*1024, "the message must stay bounded")
	assert.Contains(t, st.Message(), "[truncated]", "a truncated message should be marked as such")
}
