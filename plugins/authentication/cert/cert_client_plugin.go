// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package cert

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/crypto"
	jose "gopkg.in/square/go-jose.v2"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

type ClientConfig struct {
	Subject               string
	Issuer                string
	Kid                   string
	PrivateKey            *ecdsa.PrivateKey
	DisableContentHashing bool
	NoExpiration          bool
}

func NewRangerPlugin(cfg ClientConfig) (ranger.ClientPlugin, error) {
	return &certAuthenticationClientPlugin{cfg: cfg}, nil
}

type certAuthenticationClientPlugin struct {
	cfg ClientConfig
}

func (cap *certAuthenticationClientPlugin) GetName() string {
	return "Ranger Guard Signer Plugin"
}

func (cap *certAuthenticationClientPlugin) GetHeader(serialzed []byte) http.Header {
	header := make(http.Header)

	// generate jwt security header
	signature := crypto.HashPayload(serialzed)

	bearer, err := cap.Sign(signature)
	if err != nil {
		log.Error().Err(err).Str("component", "guard").Msg("could not generate bearer token")
		return header
	}

	header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer))
	return header
}

// Sign generates JWT bearer token
func (cap *certAuthenticationClientPlugin) Sign(signature []byte) (string, error) {
	var shash string

	if !cap.cfg.DisableContentHashing {
		shash = fmt.Sprintf("%x", signature)
	}
	issuedAt := time.Now()

	cl := jwt.Claims{
		Subject:   cap.cfg.Subject,
		Issuer:    cap.cfg.Issuer,
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		NotBefore: jwt.NewNumericDate(issuedAt),
	}

	if !cap.cfg.NoExpiration {
		// valid for 60 seconds
		cl.Expiry = jwt.NewNumericDate(issuedAt.Add(time.Duration(60) * time.Second))
	}

	customClaims := struct {
		ContentHash string `json:"chash,omitempty"`
	}{
		shash,
	}

	sig, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.ES384,
		Key:       cap.cfg.PrivateKey,
	}, (&jose.SignerOptions{}).WithHeader("kid", cap.cfg.Kid).WithType("JWT"))
	if err != nil {
		return "", err
	}

	bearer, err := jwt.Signed(sig).Claims(cl).Claims(customClaims).CompactSerialize()
	if err != nil {
		return "", err
	}

	return bearer, nil
}
