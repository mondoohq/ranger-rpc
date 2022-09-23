package scope

import (
	"net/http"

	"github.com/google/uuid"
	"go.mondoo.com/ranger-rpc"
)

const XRequestID = "X-Request-ID"

type RequestIdOption func(s *requestIdPlugin)

func WithUuid(requestId string) RequestIdOption {
	return func(s *requestIdPlugin) {
		s.requestID = requestId
	}
}

// NewRequestIDPlugin creates a RangerPlugin that adds the
// X-Request-ID header to outgoing requests. All RPCs
// that share the same RangerPlugin will get the same ID
func NewRequestIDRangerPlugin(opts ...RequestIdOption) ranger.ClientPlugin {
	p := &requestIdPlugin{}

	for i := range opts {
		opts[i](p)
	}

	if p.requestID == "" {
		p.requestID = uuid.New().String()
	}

	return p
}

type requestIdPlugin struct {
	requestID string
}

func (p *requestIdPlugin) GetName() string {
	return "Request Scope Plugin"
}

func (p *requestIdPlugin) GetHeader(data []byte) http.Header {
	header := make(http.Header)
	header.Set(XRequestID, p.requestID)
	return header
}
