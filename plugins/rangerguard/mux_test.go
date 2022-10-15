package rangerguard_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc/codes"
	helloworld "go.mondoo.com/ranger-rpc/examples/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/authentication"
	"go.mondoo.com/ranger-rpc/plugins/authentication/defaultuser"
	"go.mondoo.com/ranger-rpc/plugins/authorization"
	"go.mondoo.com/ranger-rpc/plugins/authorization/always"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard"
	"go.mondoo.com/ranger-rpc/status"
)

func createGuardServer(authenticators []authentication.Authenticator, authorizors []authorization.Authorizor) (*httptest.Server, error) {
	helloHandler := helloworld.NewHelloWorldServer(&helloworld.HelloWorldServerImpl{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	mux.Handle("/hello/", http.StripPrefix("/hello", helloHandler))

	// Start a local HTTP server
	server := httptest.NewServer(rangerguard.New(rangerguard.Options{
		Authenticators: authenticators,
		Authorizors:    authorizors,
	}, mux))
	return server, nil
}

func TestGuardWithNoNilAuth(t *testing.T) {
	server, err := createGuardServer(nil, nil)
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{})
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)
	assert.NotNil(t, err, "non-ok http request")

	s, _ := status.FromError(err)
	assert.Equal(t, codes.Unauthenticated, s.Code(), "has correct error code")

	assert.Equal(t, "", protoResp.Text, "get expected service response")
}

func TestGuardWithNoAuth(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{}, []authorization.Authorizor{})
	if err != nil {
		t.Fatal(err)
	}

	// Close the server when test finishes
	defer server.Close()

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{})
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)
	assert.NotNil(t, err, "non-ok http request")

	s, _ := status.FromError(err)
	assert.Equal(t, codes.Unauthenticated, s.Code(), "has correct error code")

	assert.Equal(t, "", protoResp.Text, "get expected service response")
}

func TestGuardWithNoAuthentication(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{}, []authorization.Authorizor{always.Allow()})
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{})
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)

	assert.NotNil(t, err, "non-ok http request")

	s, _ := status.FromError(err)
	assert.Equal(t, codes.Unauthenticated, s.Code(), "has correct error code")
	assert.Equal(t, "", protoResp.Text, "get expected service response")
}

func TestGuardWithNoAuthorization(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{defaultuser.Anonymous()}, []authorization.Authorizor{})
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{})
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)
	assert.NotNil(t, err, "non-ok http request")

	s, _ := status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())

	assert.Equal(t, "", protoResp.Text, "get expected service response")
}

func TestGuardWithAuthnAndAuthz(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{defaultuser.Anonymous()}, []authorization.Authorizor{always.Allow()})
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{})
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)

	assert.Nil(t, err, "service returns without error")
	assert.Equal(t, "Hello World", protoResp.Text, "get expected service response")
}