// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeystore(t *testing.T) {
	assert := assert.New(t)

	privateKey, err := PrivateKeyFromFile("../../../examples/rangerguard/server/private-key.p8")
	assert.Equal(nil, err, "key could be loaded")
	assert.NotEqual(nil, privateKey, "key should be loaded")
}
