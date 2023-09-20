// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package statictoken

import (
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/header"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
	"golang.org/x/crypto/bcrypt"
)

var defaultBcryptCost = 14

var defaultUser = &user.UserInfo{
	Subject: "statictoken-authenticated",
	Name:    "Statictoken Authenticated User",
	Issuer:  "system/statictoken",
}

type StaticTokenOption func(c *staticTokenAuthenticator)

func WithUser(subject string, name string, email string) StaticTokenOption {
	return func(a *staticTokenAuthenticator) {
		a.user = &user.UserInfo{
			Subject: subject,
			Name:    name,
			Email:   strings.ToLower(email),
			Issuer:  "system/statictoken",
		}
	}
}

type StaticTokenComparer interface {
	CheckPasswordHash(password, hash string) bool
}

type bcryptComparer struct{}

func (bc *bcryptComparer) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func WithComparer(comparer StaticTokenComparer) StaticTokenOption {
	return func(a *staticTokenAuthenticator) {
		a.comparer = comparer
	}
}

func WithBcryptCost(bcryptCost int) StaticTokenOption {
	if bcryptCost < 9 || bcryptCost > 32 {
		log.Warn().Msgf("Invalid bcrypt cost %d, setting to default %d", bcryptCost, defaultBcryptCost)
		bcryptCost = defaultBcryptCost
	}
	return func(a *staticTokenAuthenticator) {
		a.bcryptCost = bcryptCost
	}
}

func New(token string, opts ...StaticTokenOption) *staticTokenAuthenticator {
	ta := &staticTokenAuthenticator{
		bcryptCost: defaultBcryptCost,
		comparer:   &bcryptComparer{},
	}

	// apply opts
	for i := range opts {
		opts[i](ta)
	}

	hash, err := hashPassword(token, ta.bcryptCost)
	if err != nil {
		log.Error().Err(err).Msg("could not hash static token")
	}
	ta.hashedToken = hash

	// fallback to default user
	if ta.user == nil {
		ta.user = defaultUser
	}

	return ta
}

type staticTokenAuthenticator struct {
	user        *user.UserInfo
	comparer    StaticTokenComparer
	hashedToken string
	bcryptCost  int
}

func (sa *staticTokenAuthenticator) Name() string {
	return "Static Token Authenticator"
}

func (sa *staticTokenAuthenticator) Verify(req *http.Request) (user.User, bool, error) {
	bearerToken, err := header.ExtractToken(req)
	if err != nil {
		return nil, false, err
	}

	if bearerToken == "" {
		// Don't compare a token if it has not been supplied
		return nil, false, nil
	}

	if sa.comparer.CheckPasswordHash(bearerToken, sa.hashedToken) {
		return sa.user, true, nil
	}

	return nil, false, nil
}

func hashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
