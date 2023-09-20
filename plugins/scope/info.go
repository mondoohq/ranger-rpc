// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package scope

import (
	"bytes"
	"errors"
	"net/http"
	"sort"
	"strings"

	"go.mondoo.com/ranger-rpc"
)

func NewCustomHeaderRangerPlugin(header http.Header) ranger.ClientPlugin {
	return &customHeaderPlugin{
		header: header,
	}
}

type customHeaderPlugin struct {
	header http.Header
}

func (u *customHeaderPlugin) GetName() string {
	return "Custom Header Plugin"
}

func (u *customHeaderPlugin) GetHeader(data []byte) http.Header {
	return u.header
}

// sanitizeString takes an input and prepares it for use in User-Client header
func sanitizeString(k string) string {
	return strings.ReplaceAll(k, " ", "-")
}

// XInfoHeader encodes key/value pairs to a string that can be used as a HTTP header.
// The keys and values are sanitized to be used as HTTP header values. All spaces in
// keys and values are replaced with dashes.
func XInfoHeader(keyValuePairs map[string]string) string {
	if len(keyValuePairs) == 0 {
		return ""
	}

	keys := make([]string, 0, len(keyValuePairs))
	for k := range keyValuePairs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for i := range keys {
		buf.WriteByte(' ')
		buf.WriteString(sanitizeString(keys[i]))
		buf.WriteByte('/')
		buf.WriteString(sanitizeString(keyValuePairs[keys[i]]))
	}
	// remove first whitespace on first entry
	return buf.String()[1:]
}

func ParseXInfoHeader(value string) (map[string]string, error) {
	res := make(map[string]string)
	entries := strings.Split(value, " ")
	for i := range entries {
		key_value := strings.SplitN(entries[i], "/", 2)
		if len(key_value) != 2 {
			return nil, errors.New("invalid info header")
		}
		key := key_value[0]
		value := key_value[1]
		res[key] = value
	}
	return res, nil
}
