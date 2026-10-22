// Copyright Mondoo, Inc. 2026
// SPDX-License-Identifier: MPL-2.0

package status_test

import (
	"net/http"
	"testing"

	"go.mondoo.com/ranger-rpc/status"

	"github.com/stretchr/testify/assert"
	"go.mondoo.com/ranger-rpc/codes"
)

func TestHTTPStatusFromCode(t *testing.T) {
	assert.Equal(t, http.StatusOK, status.HTTPStatusFromCode(codes.OK))
	assert.Equal(t, http.StatusInternalServerError, status.HTTPStatusFromCode(codes.Unknown))
	assert.Equal(t, http.StatusInternalServerError, status.HTTPStatusFromCode(codes.Code(1000)))
}

func TestCodeFromHTTPStatus(t *testing.T) {
	assert.Equal(t, codes.OK, status.CodeFromHTTPStatus(http.StatusOK))
	assert.Equal(t, codes.Internal, status.CodeFromHTTPStatus(http.StatusInternalServerError))
	assert.Equal(t, codes.Unknown, status.CodeFromHTTPStatus(1000))
	// Proxies and load balancers answer 502/503/504 with HTML; callers need a
	// retryable code to branch on.
	assert.Equal(t, codes.Unavailable, status.CodeFromHTTPStatus(http.StatusBadGateway))
	assert.Equal(t, codes.Unavailable, status.CodeFromHTTPStatus(http.StatusServiceUnavailable))
	assert.Equal(t, codes.Unavailable, status.CodeFromHTTPStatus(http.StatusGatewayTimeout))
}
