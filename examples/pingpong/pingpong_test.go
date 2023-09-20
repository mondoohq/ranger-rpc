// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package pingpong

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPingPong(t *testing.T) {
	// create http test server and attack the pingpong service under /api/ prefix
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", NewPingPongServer(&PingPongServiceImpl{})))
	server := httptest.NewServer(mux)
	defer server.Close()

	// create a client to call the pingpong service
	client, err := NewPingPongClient(server.URL+"/api/", &http.Client{})
	assert.Nil(t, err)

	t.Run("Ping", func(t *testing.T) {
		resp, err := client.Ping(context.Background(), &PingRequest{Sender: "Example"})
		require.NoError(t, err)
		assert.Equal(t, "Hello Example", resp.Message)
	})

	t.Run("NoPing", func(t *testing.T) {
		resp, err := client.NoPing(context.Background(), &Empty{})
		require.NoError(t, err)
		assert.Equal(t, "HelloPong", resp.Message)
	})
}
