// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package cert_test

import (
	"context"
	"crypto/x509"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc"
	helloworld "go.mondoo.com/ranger-rpc/examples/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/authentication"
	"go.mondoo.com/ranger-rpc/plugins/authentication/cert"
	"go.mondoo.com/ranger-rpc/plugins/authorization"
	"go.mondoo.com/ranger-rpc/plugins/authorization/always"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/crypto"
)

func createCertHelloworldServer() (*httptest.Server, *helloworld.Keystore) {
	helloHandler := helloworld.NewHelloWorldServer(&helloworld.HelloWorldServerImpl{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	mux.Handle("/hello/", http.StripPrefix("/hello", helloHandler))

	ks := helloworld.NewKeystore()
	// init auth middleware
	ca := cert.New(cert.Config{
		KeyStore: ks,
	})
	authenticators := []authentication.Authenticator{ca}
	authorizors := []authorization.Authorizor{always.Allow()}

	// Start a local HTTP server
	server := httptest.NewServer(rangerguard.New(
		rangerguard.Options{
			Authenticators: authenticators,
			Authorizors:    authorizors,
		},
		mux))
	return server, ks
}

func TestGuardCertAuthentication(t *testing.T) {
	server, ks := createCertHelloworldServer()
	err := ks.Load("../../../examples/rangerguard/server/cert.pem")
	require.NoError(t, err)
	// Close the server when test finishes
	defer server.Close()

	// key for signing the requests
	privateKey, err := crypto.PrivateKeyFromFile("../../../examples/rangerguard/server/private-key.p8")
	require.NoError(t, err)

	plugin, err := cert.NewRangerPlugin(cert.ClientConfig{
		PrivateKey: privateKey,
		Issuer:     "testissuer",
		Subject:    "testsubject",
		Kid:        "1",
	})
	require.NoError(t, err)

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", ranger.DefaultHttpClient(), plugin)
	require.NoError(t, err)

	// check that the client is authenticated
	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)
	assert.Nil(t, err, "service returns without error")
	assert.Equal(t, "Hello World", protoResp.Text, "get expected service response")

	// check that guard detects the user
	tagResp, err := protoClient.Info(context.Background(), &helloworld.Empty{})
	assert.Nil(t, err, "service returns without error")
	assert.Equal(t, "testsubject", tagResp.Tags["subject"], "get expected user subject")
	assert.Equal(t, "testissuer", tagResp.Tags["issuer"], "get expected issuer")
	assert.Equal(t, "", tagResp.Tags["name"], "get expected user name")
	assert.Equal(t, "", tagResp.Tags["email"], "get expected user email")
	assert.Equal(t, "", tagResp.Tags["groups"], "get expected user email")
}

func TestDenyGuardCertAuthentication(t *testing.T) {
	server, ks := createCertHelloworldServer()
	err := ks.Load("../../../examples/rangerguard/server/cert.pem")
	require.NoError(t, err)

	// Close the server when test finishes
	defer server.Close()

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", ranger.DefaultHttpClient())
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	_, err = protoClient.Hello(context.Background(), data)
	assert.NotNil(t, err, "service returns with error")
}

// implement fake keystore
type fakeKeyStore struct{}

func (ks *fakeKeyStore) Get(key string) (*x509.Certificate, error) {
	cert, err := crypto.CertificateFromFile("../../example/server/cert.pem")
	if err != nil {
		return nil, err
	}

	if key == "1" {
		return cert, nil
	}

	return nil, errors.New("invalid kid")
}
