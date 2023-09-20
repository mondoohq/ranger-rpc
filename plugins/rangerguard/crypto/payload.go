// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func HashPayload(message []byte) []byte {
	pssh := sha256.New()
	pssh.Write(message)
	return pssh.Sum(nil)
}

func SignMessage(priv *ecdsa.PrivateKey, message []byte) (*big.Int, *big.Int, error) {
	hashed := HashPayload(message)
	r, s, err := ecdsa.Sign(
		rand.Reader,
		priv,
		hashed,
	)
	if err != nil {
		return nil, nil, err
	}

	return r, s, nil
}

func VerifyMessage(pub *ecdsa.PublicKey, message []byte, r, s *big.Int) bool {
	hashed := HashPayload(message)
	return ecdsa.Verify(
		pub,
		hashed,
		r,
		s,
	)
}
