package authentication

import (
	"net/http"

	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
)

type Authenticator interface {
	Name() string
	Verify(req *http.Request) (user.User, bool, error)
}
