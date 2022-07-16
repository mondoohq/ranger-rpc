package ranger_test

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/examples/pingpong"
	"go.mondoo.com/ranger-rpc/metadata"
	"go.mondoo.com/ranger-rpc/status"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
	pb "google.golang.org/protobuf/proto"
)

func TestRangerHttpServer(t *testing.T) {
	service := ranger.Service{
		Name: "PingPong",
		Methods: map[string]ranger.Method{
			"Ping": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				test := &pingpong.PongReply{}
				return test, nil
			},
		},
	}

	srv := ranger.NewRPCServer(&service)

	var req *http.Request
	var w *httptest.ResponseRecorder
	var resp *http.Response

	// get 404 since the default content-type is protobuf
	req = httptest.NewRequest("POST", "http://example.com/Ping", nil)
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp = w.Result()
	assert.Equal(t, 404, resp.StatusCode, "correct status code")

	// get 404 since the route is not registered
	req = httptest.NewRequest("POST", "http://example.com/Unknown", nil)
	req.Header.Set("Content-Type", "application/protobuf")
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp = w.Result()
	assert.Equal(t, 404, resp.StatusCode, "correct status code")

	// get 404 since the route is not registered
	req = httptest.NewRequest("POST", "http://example.com/PingPong/Ping", nil)
	req.Header.Set("Content-Type", "application/protobuf")
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp = w.Result()
	assert.Equal(t, 200, resp.StatusCode, "correct status code")

}

func TestRangerHttpHeader(t *testing.T) {
	service := ranger.Service{
		Name: "PingPong",
		Methods: map[string]ranger.Method{
			"Ping": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				md, ok := metadata.FromIncomingContext(ctx)
				if !ok {
					log.Debug().Msg("no header")
					return nil, errors.New("could not access header")
				}

				test := &pingpong.PongReply{Message: md.First("X-Custom-Header")}
				return test, nil
			},
		},
	}

	srv := ranger.NewRPCServer(&service)

	var req *http.Request
	var w *httptest.ResponseRecorder
	var resp *http.Response

	// test that the handler has access to http header
	req = httptest.NewRequest("POST", "http://example.com/PingPong/Ping", nil)
	req.Header.Set("Content-Type", "application/protobuf")
	req.Header.Set("X-Custom-Header", "my-header")
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp = w.Result()
	assert.Equal(t, 200, resp.StatusCode, "correct status code")

	content, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err, "should return protobuf content")

	var msg pingpong.PongReply
	if err := pb.Unmarshal(content, &msg); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "my-header", msg.Message, "correct header value")
}

func runStatusCall(srv http.Handler, path string) *http.Response {
	req := httptest.NewRequest("POST", path, nil)
	req.Header.Set("Content-Type", "application/protobuf")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	return w.Result()
}

func parseStatus(reader io.Reader) (*spb.Status, error) {
	payload, err := ioutil.ReadAll(reader)
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

func TestRangerErrorHandling(t *testing.T) {
	service := ranger.Service{
		Name: "Error",
		Methods: map[string]ranger.Method{
			"Ping": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				test := &pingpong.PongReply{}
				return test, nil
			},
			"NilResponseWithStatusError": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				return nil, status.Error(codes.PermissionDenied, "you really have no permission")
			},
			"NilResponseWithError": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				return nil, errors.New("my error message")
			},
			"NilResponseWithoutError": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				return nil, nil
			},
		},
	}

	srv := ranger.NewRPCServer(&service)

	var req *http.Request
	var w *httptest.ResponseRecorder
	var resp *http.Response

	// check argument error for wrong content type
	req = httptest.NewRequest("POST", "http://example.com/Error/Ping", nil)
	// this will result in a json error response
	req.Header.Set("Content-Type", "unknown/content-type")
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp = w.Result()
	assert.Equal(t, 400, resp.StatusCode, "correct status code")
	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	type errMsg struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	respDecoded := errMsg{}
	if err := json.Unmarshal(payload, &respDecoded); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, errMsg{
		Code:    3,
		Message: "unexpected content-type: \"unknown/content-type\"",
	}, respDecoded)

	// assume protobuf content-type if no content-type is there
	req = httptest.NewRequest("POST", "http://example.com/Error/Ping", nil)
	w = httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	resp = w.Result()
	assert.Equal(t, 200, resp.StatusCode, "correct status code")

	// test status codes are respected
	resp = runStatusCall(srv, "http://example.com/Error/NilResponseWithStatusError")
	assert.Equal(t, 403, resp.StatusCode, "correct status code")
	status, err := parseStatus(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, "you really have no permission", status.Message, "correct error message")

	resp = runStatusCall(srv, "http://example.com/Error/NilResponseWithError")
	assert.Equal(t, 500, resp.StatusCode, "correct status code")
	status, err = parseStatus(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, "my error message", status.Message, "correct error message")

	resp = runStatusCall(srv, "http://example.com/Error/NilResponseWithoutError")
	assert.Equal(t, 200, resp.StatusCode, "correct status code")
	status, err = parseStatus(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, "", status.Message, "correct error message")
}

func TestRangerStatusCodes(t *testing.T) {
	service := ranger.Service{
		Name: "Status",
		Methods: map[string]ranger.Method{
			"NotFound": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				return nil, status.Error(codes.NotFound, "id was not found")
			},
			"InvalidArgument": func(ctx context.Context, reqBytes *[]byte) (pb.Message, error) {
				return nil, status.Errorf(codes.InvalidArgument, "Value of `Input` is not allowed")
			},
		},
	}

	srv := ranger.NewRPCServer(&service)
	var resp *http.Response

	// test 404
	resp = runStatusCall(srv, "http://example.com/Status/NotFound")
	assert.Equal(t, 404, resp.StatusCode, "correct status code")
	status, err := parseStatus(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, codes.NotFound, codes.Code(status.GetCode()), "correct error message")
	assert.Equal(t, "id was not found", status.GetMessage(), "correct error message")

	// test 400
	resp = runStatusCall(srv, "http://example.com/Status/InvalidArgument")
	assert.Equal(t, 400, resp.StatusCode, "correct status code")
	status, err = parseStatus(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, codes.InvalidArgument, codes.Code(status.GetCode()), "correct error message")
	assert.Equal(t, "Value of `Input` is not allowed", status.GetMessage(), "correct error message")
}