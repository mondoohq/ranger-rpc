// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package authentication

import (
	"net/http"

	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
)

type Authenticator interface {
	Name() string
	Verify(req *http.Request) (user.User, bool, error)
}
