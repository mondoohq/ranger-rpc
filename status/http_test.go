// Copyright (c) Mondoo, Inc.
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
}
