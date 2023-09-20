// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package always

import (
	"go.mondoo.com/ranger-rpc/plugins/authorization"
)

func Allow() *allowAuthorizor {
	return &allowAuthorizor{}
}

type allowAuthorizor struct{}

func (da *allowAuthorizor) Name() string {
	return "Allow Authorizer"
}

func (da *allowAuthorizor) Authorize(a authorization.AuthorizationFacts) (authorized authorization.Decision, reason string, err error) {
	return authorization.DecisionAllow, "allow authorizer allows this requests", nil
}
