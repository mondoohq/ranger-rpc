package ranger

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/metadata"
	"go.mondoo.com/ranger-rpc/status"
)

const (
	requestTimeout = "Request-Timeout"
	xForwardedFor  = "X-Forwarded-For"
	xForwardedHost = "X-Forwarded-Host"
)

var (
	// DefaultContextTimeout is used for gRPC call context.WithTimeout whenever a Grpc-Timeout inbound
	// header isn't present. If the value is 0 the sent `context` will not have a timeout.
	DefaultContextTimeout = 0 * time.Second
)

/*
AnnotateContext adds context information such as metadata from the request. It is designed to work similar
to the gRPC AnnotateContext function. At a minimum, the RemoteAddr is included in the fashion of "X-Forwarded-For" in
the metadata.

The implementation is derived from GRPC-Gateway AnnotateContext (BSD3-License). See:
https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/context.go#L78
*/
func AnnotateContext(ctx context.Context, req *http.Request) (context.Context, context.CancelFunc, error) {
	var pairs []string

	cancelFunc := func() {}

	// handle timeouts
	timeout := DefaultContextTimeout
	if tm := req.Header.Get(requestTimeout); tm != "" {
		var err error
		timeout, err = timeoutDecode(tm)
		if err != nil {
			return nil, nil, status.Errorf(codes.InvalidArgument, "invalid rpc-timeout: %s", tm)
		}
	}

	// pass all headers through
	for key, vals := range req.Header {
		for _, val := range vals {
			key = textproto.CanonicalMIMEHeaderKey(key)

			// skip forwarded key here since it needs special handling
			if key == xForwardedHost {
				continue
			}
			pairs = append(pairs, key, val)
		}
	}

	// handle forward key
	if host := req.Header.Get(xForwardedHost); host != "" {
		pairs = append(pairs, strings.ToLower(xForwardedHost), host)
	} else if req.Host != "" {
		pairs = append(pairs, strings.ToLower(xForwardedHost), req.Host)
	}

	// handle remotr addr key
	if addr := req.RemoteAddr; addr != "" {
		if remoteIP, _, err := net.SplitHostPort(addr); err == nil {
			if fwd := req.Header.Get(xForwardedFor); fwd == "" {
				pairs = append(pairs, strings.ToLower(xForwardedFor), remoteIP)
			} else {
				pairs = append(pairs, strings.ToLower(xForwardedFor), fmt.Sprintf("%s, %s", fwd, remoteIP))
			}
		} else {
			log.Info().Str("remote_addr", addr).Msg("invalid remote addr")
		}
	}

	if timeout != 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, timeout)
	}
	if len(pairs) == 0 {
		return ctx, cancelFunc, nil
	}
	md := metadata.Pairs(pairs...)
	return metadata.NewIncomingContext(ctx, md), cancelFunc, nil
}

func timeoutDecode(s string) (time.Duration, error) {
	size := len(s)
	if size < 2 {
		return 0, fmt.Errorf("timeout string is too short: %q", s)
	}
	d, ok := timeoutUnitToDuration(s[size-1])
	if !ok {
		return 0, fmt.Errorf("timeout unit is not recognized: %q", s)
	}
	t, err := strconv.ParseInt(s[:size-1], 10, 64)
	if err != nil {
		return 0, err
	}
	return d * time.Duration(t), nil
}

func timeoutUnitToDuration(u uint8) (d time.Duration, ok bool) {
	switch u {
	case 'H':
		return time.Hour, true
	case 'M':
		return time.Minute, true
	case 'S':
		return time.Second, true
	case 'm':
		return time.Millisecond, true
	case 'u':
		return time.Microsecond, true
	case 'n':
		return time.Nanosecond, true
	default:
	}
	return
}
