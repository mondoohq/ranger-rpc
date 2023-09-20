// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomHeaderRangerPlugin(t *testing.T) {
	expected := map[string]string{
		"version": "1.1.4",
		"build":   "abc",
	}

	h := http.Header{}
	h.Set("X-Info", XInfoHeader(expected))

	userAgent := NewCustomHeaderRangerPlugin(h)
	pluginHeader := userAgent.GetHeader(nil)
	xinfoHeader := pluginHeader.Get("X-Info")
	assert.Equal(t, "build/abc version/1.1.4", xinfoHeader)

	value, err := ParseXInfoHeader(xinfoHeader)
	require.NoError(t, err)
	assert.Equal(t, expected, value)
}

func TestXInfoHeader(t *testing.T) {
	for _, test := range []struct {
		kv       map[string]string
		expected string
	}{
		{nil, ""},
		{map[string]string{"abc": "123"}, "abc/123"},
		{map[string]string{"abc": "123", "xyz": "567", "boo": ""}, "abc/123 boo/ xyz/567"},
		{map[string]string{"a b c": "1 2 3"}, "a-b-c/1-2-3"}, // spaces are replaced with dashes
	} {
		got := XInfoHeader(test.kv)
		if got != test.expected {
			t.Errorf("header(%q) = %q, want %q", test.kv, got, test.expected)
		}
	}
}
