package ranger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/status"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	ContentTypeProtobuf      = "application/protobuf"
	ContentTypeOctetProtobuf = "application/octet-stream"
	ContentTypeGrpcProtobuf  = "application/grpc+proto"
	ContentTypeJson          = "application/json"
)

var validContentTypes = []string{
	ContentTypeProtobuf,
	ContentTypeOctetProtobuf,
	ContentTypeGrpcProtobuf,
	ContentTypeJson,
}

// Method represents a RPC method and is used by protoc-gen-rangerrpc
type Method func(ctx context.Context, reqBytes *[]byte) (proto.Message, error)

// Service is the struct that holds all available methods. The protoc-gen-rangerrpc will generate the
// correct client and service definition to be used.
type Service struct {
	Name    string
	Methods map[string]Method
}

// NewServer creates a new server. This function is used by the protoc-gen-rangerrpc generated code and
// should not be used directly.
func NewRPCServer(service *Service) *server {
	var b strings.Builder
	b.WriteString("/")
	b.WriteString(service.Name)
	b.WriteString("/")
	return &server{service: service, prefix: b.String()}
}

type server struct {
	service *Service
	prefix  string
}

// ServeHTTP is the main entry point for the http server.
func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	// verify content type
	err := verifyContentType(req)
	if err != nil {
		httpError(w, req, err)
		return
	}

	if !strings.HasPrefix(req.URL.Path, s.prefix) {
		httpError(w, req, status.Error(codes.NotFound, req.URL.Path+" is not available"))
		return
	}

	// extract the rpc method name and invoke the method
	name := strings.TrimPrefix(req.URL.Path, s.prefix)

	method := s.service.Methods[name]
	if method == nil {
		err := status.Error(codes.NotFound, "method not defined")
		httpError(w, req, err)
		return
	}

	rctx, rcancel, body, err := PreProcessRequest(ctx, req)
	if err != nil {
		httpError(w, req, err)
		return
	}
	defer rcancel()

	// invoke method and send the response
	resp, err := method(rctx, &body)
	s.sendResponse(w, req, resp, err)
}

// PreProcessRequest is used to preprocess the incoming request.
// It returns the context, a cancel function and the body of the request. The cancel function can be used to cancel
// the context. It also adds the http headers to the context.
func PreProcessRequest(ctx context.Context, req *http.Request) (context.Context, context.CancelFunc, []byte, error) {
	// read body content
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		return nil, nil, nil, status.Error(codes.DataLoss, "unrecoverable data loss or corruption")
	}

	// pass-through the http headers
	rctx, rcancel, err := AnnotateContext(ctx, req)
	if err != nil {
		return nil, rcancel, nil, err
	}

	return rctx, rcancel, body, nil
}

func (s *server) sendResponse(w http.ResponseWriter, req *http.Request, resp proto.Message, err error) {
	if err != nil {
		httpError(w, req, err)
		return
	}

	// check if the accept header is set, otherwise use the incoming content type
	accept := determineResponseType(req.Header.Get("Content-Type"), req.Header.Get("Accept"))
	payload, contentType, err := convertProtoToPayload(resp, accept)

	if err != nil {
		httpError(w, req, status.Error(codes.Internal, "error encoding response"))
		return
	}

	h := w.Header()
	h.Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

// convertProtoToPayload converts a proto message to the approaptiate formatted payload.
// Depending on the accept header it will return the payload as marshalled protobuf or json.
func convertProtoToPayload(resp proto.Message, accept string) ([]byte, string, error) {
	var err error
	var payload []byte
	contentType := accept
	switch accept {
	case ContentTypeProtobuf, ContentTypeGrpcProtobuf, ContentTypeOctetProtobuf:
		contentType = ContentTypeProtobuf
		payload, err = proto.Marshal(resp)
	// as default, we return json to be compatible with browsers, since they do not
	// request as application/json as default
	default:
		payload, err = jsonpb.MarshalOptions{UseProtoNames: true}.Marshal(resp)
	}

	return payload, contentType, err
}

// verifyContentType validates the content type of the request is known.
func verifyContentType(req *http.Request) error {
	header := req.Header.Get("Content-Type")

	// we assume "application/protobuf" if no content-type is set
	if len(header) == 0 {
		return nil
	}

	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	ct := strings.TrimSpace(strings.ToLower(header[:i]))

	// check that the incoming request has a valid content type
	for _, a := range validContentTypes {
		if a == ct {
			return nil
		}
	}

	// if we reached here, we have to handle an unexpected incoming type
	return status.Error(codes.InvalidArgument, fmt.Sprintf("unexpected content-type: %q", req.Header.Get("Content-Type")))
}

// determineResponseType returns the content type based on the Content-Type and Accept header.
func determineResponseType(contenttype string, accept string) string {
	// use provided content type if no accept header was provided
	if len(accept) == 0 {
		accept = contenttype
	}

	switch accept {
	case ContentTypeProtobuf, ContentTypeGrpcProtobuf, ContentTypeOctetProtobuf:
		return ContentTypeProtobuf
	default:
		return ContentTypeJson
	}
}
