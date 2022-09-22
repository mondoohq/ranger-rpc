package statictoken_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	helloworld "go.mondoo.com/ranger-rpc/examples/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/authentication"
	"go.mondoo.com/ranger-rpc/plugins/authentication/statictoken"
	"go.mondoo.com/ranger-rpc/plugins/authorization"
	"go.mondoo.com/ranger-rpc/plugins/authorization/always"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard"
)

func createGuardServer(authenticators []authentication.Authenticator, authorizors []authorization.Authorizor) (*httptest.Server, error) {
	helloHandler := helloworld.NewHelloWorldServer(&helloworld.HelloWorldServerImpl{})
	// You can use any mux you like - NewHelloWorldServer gives you a http.Handler.
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

func TestGuardWithStaticTokenAuth(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{statictoken.New("abcdefg", statictoken.WithUser("johndoe", "John Doe", "john@example.com"))}, []authorization.Authorizor{always.Allow()})
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{}, statictoken.NewRangerPlugin("abcdefg"))
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)

	assert.Nil(t, err, "service returns without error")
	assert.Equal(t, "Hello World", protoResp.Text, "get expected service response")
}

type fakeComparer struct {
	wasCalled bool
}

func (fc *fakeComparer) CheckPasswordHash(password, hash string) bool {
	fc.wasCalled = true
	return true
}

func TestGuardWithMissingToken(t *testing.T) {
	fc := &fakeComparer{}
	server, err := createGuardServer([]authentication.Authenticator{statictoken.New("abcdefg", statictoken.WithUser("johndoe", "John Doe", "john@example.com"), statictoken.WithComparer(fc))}, []authorization.Authorizor{always.Allow()})
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{}, statictoken.NewRangerPlugin(""))
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	_, err = protoClient.Hello(context.Background(), data)

	assert.NotNil(t, err, "service returns error")
	// verify that the bcrypt compare call was not run
	assert.Equal(t, fc.wasCalled, false)
}

func TestGuardWithFakeComparer(t *testing.T) {
	fc := &fakeComparer{}
	server, err := createGuardServer([]authentication.Authenticator{statictoken.New("abcdefg", statictoken.WithUser("johndoe", "John Doe", "john@example.com"), statictoken.WithComparer(fc))}, []authorization.Authorizor{always.Allow()})
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{}, statictoken.NewRangerPlugin("abcdefg"))
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)

	assert.Nil(t, err, "service returns without error")
	assert.Equal(t, "Hello World", protoResp.Text, "get expected service response")
	assert.Equal(t, fc.wasCalled, true) // make sure that our comparer was called
}
