// Copyright Mondoo, Inc. 2026
// SPDX-License-Identifier: MPL-2.0

package ranger

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/status"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
)

// HttpError writes an error to the response.
func HttpError(span trace.Span, w http.ResponseWriter, req *http.Request, err error) {
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

	span.RecordError(err, trace.WithAttributes(
		attribute.Int("mondoo.ranger.status", status),
	))

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
	if secs, ok := retryAfterSeconds(s); ok {
		h.Set("Retry-After", strconv.FormatInt(secs, 10))
	}
	w.WriteHeader(status)
	w.Write(payload)
}

// retryAfterSeconds returns the delay of a google.rpc.RetryInfo detail in whole
// seconds, rounded up, so HTTP clients that don't parse the status body still
// learn when to retry.
func retryAfterSeconds(s *status.Status) (int64, bool) {
	for _, d := range s.Details() {
		ri, ok := d.(*errdetails.RetryInfo)
		if !ok || ri.GetRetryDelay() == nil {
			continue
		}
		delay := ri.GetRetryDelay().AsDuration()
		if delay <= 0 {
			return 0, false
		}
		return int64((delay + time.Second - 1) / time.Second), true
	}
	return 0, false
}

// parseStatus tries to parse the proto Status from the body of the response.
// It also returns the raw payload it read so callers can fall back to using the
// body as an error message when it is not a valid proto Status (e.g. an HTML or
// plain-text error page from an intermediary load balancer or proxy).
func parseStatus(reader io.Reader) (*spb.Status, []byte, error) {
	payload, err := io.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}
	var status spb.Status
	err = proto.Unmarshal(payload, &status)
	if err != nil {
		return nil, payload, err
	}
	return &status, payload, nil
}
