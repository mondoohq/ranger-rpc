// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRequestScopePlugin(t *testing.T) {
	t.Run("header is valid uuid", func(t *testing.T) {
		p := NewRequestIDRangerPlugin()
		header := p.GetHeader(nil)
		headerValue := header.Get(XRequestID)
		require.NotEmpty(t, headerValue)
		_, err := uuid.Parse(headerValue)
		require.NoError(t, err)
	})

	t.Run("multiple calls produce the same value", func(t *testing.T) {
		p := NewRequestIDRangerPlugin()
		headerValue1 := p.GetHeader(nil).Get(XRequestID)
		headerValue2 := p.GetHeader(nil).Get(XRequestID)
		require.Equal(t, headerValue1, headerValue2)
	})

	t.Run("uniqueness", func(t *testing.T) {
		p1 := NewRequestIDRangerPlugin()
		p2 := NewRequestIDRangerPlugin()
		headerValue1 := p1.GetHeader(nil).Get(XRequestID)
		headerValue2 := p2.GetHeader(nil).Get(XRequestID)
		require.NotEqual(t, headerValue1, headerValue2)
	})

	t.Run("custom uuid", func(t *testing.T) {
		p1 := NewRequestIDRangerPlugin(WithUuid("custom-uuid"))
		headerValue1 := p1.GetHeader(nil).Get(XRequestID)
		require.Equal(t, "custom-uuid", headerValue1)
	})
}
