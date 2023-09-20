// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package header_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/header"
)

func TestExtractTokenFromBearer(t *testing.T) {
	for name, tc := range map[string]struct {
		bearer string
		want   string
	}{
		"success": {
			bearer: "Bearer token",
			want:   "token",
		},
		"empty bearer": {
			bearer: "Bearer",
			want:   "",
		},
		"empty bearer with space": {
			bearer: "Bearer ",
			want:   "",
		},
	} {
		t.Run(name, func(t *testing.T) {
			if got := header.ExtractTokenFromBearer(tc.bearer); got != tc.want {
				t.Errorf("ExtractTokenFromBearer() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestExtractToken(t *testing.T) {
	t.Run("with garbage from uri", func(t *testing.T) {
		req := httptest.NewRequest("POST", "localhost:80", strings.NewReader("test"))
		req.RequestURI = "/;'!@#$%^&*(_+{}[]\\|"
		token, err := header.ExtractToken(req)
		require.NoError(t, err)
		require.Empty(t, token)
	})
}
