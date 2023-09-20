// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEcdsaSignMessage(t *testing.T) {
	message := []byte("the code must be like a piece of music")

	privKey, err := PrivateKeyFromFile("../../../examples/rangerguard/server/private-key.p8")
	assert.Equal(t, nil, err, "key could be loaded")

	r, s, err := SignMessage(privKey, message)
	assert.Equal(t, nil, err, "they should be equal")

	certificate, err := CertificateFromFile("../../../examples/rangerguard/server/cert.pem")
	assert.Equal(t, nil, err, "key could be loaded")

	pk, err := PublicKeyFromCert(certificate)
	require.NoError(t, err)
	valid := VerifyMessage(pk, message, r, s)
	assert.Equal(t, true, valid, "signature is valud")
}
