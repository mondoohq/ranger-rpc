package always_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc/codes"
	helloworld "go.mondoo.com/ranger-rpc/examples/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/authentication"
	"go.mondoo.com/ranger-rpc/plugins/authentication/defaultuser"
	"go.mondoo.com/ranger-rpc/plugins/authorization"
	"go.mondoo.com/ranger-rpc/plugins/authorization/always"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
	"go.mondoo.com/ranger-rpc/status"
)

func TestDenyWithAuthnAndAuthz(t *testing.T) {
	server, err := createGuardServer([]authentication.Authenticator{defaultuser.New(user.Anonymous)}, []authorization.Authorizor{always.Deny()})
	require.NoError(t, err)

	// close the server when test finishes
	defer server.Close()

	// do client request with signed jwt token
	protoClient, err := helloworld.NewHelloWorldClient(server.URL+"/hello/", &http.Client{})
	require.NoError(t, err)

	data := &helloworld.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)
	assert.NotNil(t, err, "service returns with permission error")

	s, _ := status.FromError(err)
	assert.Equal(t, codes.PermissionDenied, s.Code())
	assert.Equal(t, "", protoResp.Text, "get expected service response")
}
