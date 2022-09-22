package defaultuser

import (
	"net/http"
	"regexp"
	"strings"

	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
)

func Anonymous() *defaultUserAuthenticator {
	u := &user.UserInfo{
		Name:    "anonymous",
		Subject: "anonymous",
		Issuer:  "system/anonymous",
		Groups:  []string{"anonymous"},
	}
	return New(u)
}

// passes request with a default user, if no other authentication is set
func New(defaultUser user.User) *defaultUserAuthenticator {
	return &defaultUserAuthenticator{
		defaultUser: &user.UserInfo{
			Name:    defaultUser.GetName(),
			Subject: defaultUser.GetSubject(),
			Issuer:  defaultUser.GetIssuer(),
			Email:   strings.ToLower(defaultUser.GetEmail()),
			Groups:  defaultUser.GetGroups(),
		},
	}
}

type defaultUserAuthenticator struct {
	defaultUser user.User
	allowed     []*regexp.Regexp
}

func (aa *defaultUserAuthenticator) Name() string {
	return "Default User Authenticator"
}

func (aa *defaultUserAuthenticator) Verify(req *http.Request) (user.User, bool, error) {
	// check that the authorization header is not set
	headerVal := req.Header.Get("Authorization")
	if len(headerVal) > 0 {
		return nil, false, nil
	}

	return aa.defaultUser, true, nil
}
