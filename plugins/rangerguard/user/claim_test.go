// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package user_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
)

func TestClaimParser(t *testing.T) {
	// create a real jwt claims object
	cl := jwt.Claims{
		Subject:   "subject",
		Issuer:    "issuer",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		// valid for 60 seconds
		Expiry: jwt.NewNumericDate(time.Now().Add(time.Duration(60) * time.Second)),
	}

	// marshal claims to json as done in JWT
	data, err := json.Marshal(cl)
	require.NoError(t, err)

	// unmarshal json to our claims object
	c := user.Claims{}
	err = json.Unmarshal(data, &c)
	require.NoError(t, err)

	identity, err := user.ParseClaims(&c)
	require.NoError(t, err)

	assert.Equal(t, "subject", identity.GetSubject())
	assert.Equal(t, "issuer", identity.GetIssuer())
}
