package status

import (
	"net/http"

	"go.mondoo.com/ranger-rpc/codes"
)

var codeMap map[codes.Code]int
var statusMap map[int]codes.Code

func init() {
	codeMap = map[codes.Code]int{
		codes.OK:                 http.StatusOK,
		codes.Canceled:           http.StatusRequestTimeout,
		codes.Unknown:            http.StatusInternalServerError,
		codes.InvalidArgument:    http.StatusBadRequest,
		codes.DeadlineExceeded:   http.StatusGatewayTimeout,
		codes.NotFound:           http.StatusNotFound,
		codes.AlreadyExists:      http.StatusConflict,
		codes.PermissionDenied:   http.StatusForbidden,
		codes.Unauthenticated:    http.StatusUnauthorized,
		codes.ResourceExhausted:  http.StatusTooManyRequests,
		codes.FailedPrecondition: http.StatusPreconditionFailed,
		codes.Aborted:            http.StatusConflict,
		codes.OutOfRange:         http.StatusBadRequest,
		codes.Unimplemented:      http.StatusNotImplemented,
		codes.Internal:           http.StatusInternalServerError,
		codes.Unavailable:        http.StatusServiceUnavailable,
		codes.DataLoss:           http.StatusInternalServerError,
	}

	statusMap = map[int]codes.Code{
		// 2xx
		http.StatusOK:        codes.OK,
		http.StatusCreated:   codes.OK,
		http.StatusAccepted:  codes.OK,
		http.StatusNoContent: codes.OK,

		// 4xx
		http.StatusBadRequest:     codes.InvalidArgument,
		http.StatusUnauthorized:   codes.Unauthenticated,
		http.StatusForbidden:      codes.PermissionDenied,
		http.StatusNotFound:       codes.NotFound,
		http.StatusRequestTimeout: codes.Canceled,

		// 5xx
		http.StatusInternalServerError: codes.Internal,
		http.StatusServiceUnavailable:  codes.Unavailable,
		http.StatusNotImplemented:      codes.Unimplemented,
	}
}

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code codes.Code) int {
	status, ok := codeMap[code]
	if ok {
		return status
	}
	return http.StatusInternalServerError
}

func CodeFromHTTPStatus(status int) codes.Code {
	code, ok := statusMap[status]
	if ok {
		return code
	}

	return codes.Unknown
}
