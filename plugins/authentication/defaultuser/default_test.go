// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaultuser_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	helloworld "go.mondoo.com/ranger-rpc/examples/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/authentication"
	"go.mondoo.com/ranger-rpc/plugins/authentication/defaultuser"
	"go.mondoo.com/ranger-rpc/plugins/authentication/statictoken"
	"go.mondoo.com/ranger-rpc/plugins/authorization"
	"go.mondoo.com/ranger-rpc/plugins/authorization/always"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard"
)

func createGuardServer(authenticators []authentication.Authenticator, authorizors []authorization.Authorizor) (*httptest.Server, error) {
	helloHandler := helloworld.NewHelloWorldServer(&helloworld.HelloWorldServerImpl{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	mux.Handle("/hello/", http.StripPrefix("/hello", helloHandler))

	// Start a local HTTP server
	server := httptest.NewServer(rangerguard.New(
		rangerguard.Options{
			Authenticators: authenticators,
			Authorizors:    authorizors,
		},
		mux))
	return server, nil
}

func TestGuardWithAnonymousAuth(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{defaultuser.Anonymous()}, []authorization.Authorizor{always.Allow()})
	if err != nil {
		t.Fatal(err)
	}

	// Close the server when test finishes
	defer server.Close()

	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{})
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)

	assert.Nil(t, err, "service returns without error")
	assert.Equal(t, "Hello World", protoResp.Text, "get expected service response")
}

// Default user only works if the authorization header is empty
func TestGuardWithAnonymousAuthWithInvalidCreds(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{defaultuser.Anonymous()}, []authorization.Authorizor{always.Allow()})
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{}, statictoken.NewRangerPlugin("abcdefg"))
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	_, err = protoClient.Hello(context.Background(), data)
	assert.Equal(t, "rpc error: code = Unauthenticated desc = request permission unauthenticated", err.Error(), "service returns without error")
}
