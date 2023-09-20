// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package statictoken

import (
	"fmt"
	"net/http"

	"go.mondoo.com/ranger-rpc"
)

func NewRangerPlugin(token string) ranger.ClientPlugin {
	return &statictokenClientPlugin{token: token}
}

type statictokenClientPlugin struct {
	token string
}

func (scp *statictokenClientPlugin) GetName() string {
	return "Ranger Guard Static Token Plugin"
}

func (scp *statictokenClientPlugin) GetHeader(serialzed []byte) http.Header {
	header := make(http.Header)
	header.Set("Authorization", fmt.Sprintf("Bearer %s", scp.token))
	return header
}
