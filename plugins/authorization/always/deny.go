package always

import (
	"go.mondoo.com/ranger-rpc/plugins/authorization"
)

func Deny() *denyAuthorizor {
	return &denyAuthorizor{}
}

type denyAuthorizor struct{}

func (da *denyAuthorizor) Name() string {
	return "Deny Authorizer"
}

func (da *denyAuthorizor) Authorize(a authorization.AuthorizationFacts) (authorized authorization.Decision, reason string, err error) {
	return authorization.DecisionDeny, "deny authorizer denies this requests", nil
}
