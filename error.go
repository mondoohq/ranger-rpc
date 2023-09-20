// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package ranger

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/status"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
)

// HttpError writes an error to the response.
func HttpError(w http.ResponseWriter, req *http.Request, err error) {
	// check if the accept header is set, otherwise use the incoming content type
	accept := determineResponseType(req.Header.Get("Content-Type"), req.Header.Get("Accept"))

	// check that we got a status.Error package
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	h := w.Header()
	// write status code
	status := status.HTTPStatusFromCode(s.Code())

	if status >= 500 {
		log.Error().
			Err(err).
			Str("url", req.URL.Path).
			Int("status", status).
			Msg("returned internal error")
	} else {
		log.Debug().
			Err(err).
			Str("url", req.URL.Path).
			Int("status", status).
			Msg("non 5xx error")
	}

	// add body and content
	// NOTE: we ignore the error here since its already an error that we return
	payload, contentType, _ := convertProtoToPayload(s.Proto(), accept)
	h.Set("Content-Type", contentType)
	w.WriteHeader(status)
	w.Write(payload)
}

// parseStatus tries to parse the proto Status from the body of the response.
func parseStatus(reader io.Reader) (*spb.Status, error) {
	payload, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var status spb.Status
	err = proto.Unmarshal(payload, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}
