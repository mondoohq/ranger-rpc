package oneof_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mondoo.com/ranger-rpc/examples/oneof"
	oneof_client "go.mondoo.com/ranger-rpc/examples/oneof/client"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
)

// custom implementation
type oneOfServerImpl struct{}

func (s *oneOfServerImpl) Echo(ctx context.Context, in *oneof.OneOfRequest) (*oneof.OneOfReply, error) {
	resp := &oneof.OneOfReply{}
	if in.GetNumber() != 0 {
		resp.Options = &oneof.OneOfReply_Number{
			Number: in.GetNumber(),
		}
	}
	if in.GetText() != "" {
		resp.Options = &oneof.OneOfReply_Text{
			Text: in.GetText(),
		}
	}
	return resp, nil
}

func TestOneOf(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", oneof.NewOneOfServer(&oneOfServerImpl{})))
	server := httptest.NewServer(mux)
	defer server.Close()

	client, err := oneof.NewOneOfClient(server.URL+"/api/", &http.Client{})
	assert.Nil(t, err)

	t.Run("Echo text", func(t *testing.T) {
		resp, err := client.Echo(context.Background(), &oneof.OneOfRequest{Options: &oneof.OneOfRequest_Text{Text: "Example"}})
		require.Nil(t, err)
		assert.Equal(t, "Example", resp.GetText())
	})

	t.Run("Echo number", func(t *testing.T) {
		resp, err := client.Echo(context.Background(), &oneof.OneOfRequest{Options: &oneof.OneOfRequest_Number{Number: 42}})
		require.Nil(t, err)
		assert.Equal(t, int64(42), resp.GetNumber())
	})
}

func TestOneOfJSON(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", oneof.NewOneOfServer(&oneOfServerImpl{}))
	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	header := make(http.Header)
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")

	t.Run("Echo text", func(t *testing.T) {
		in := &oneof.OneOfRequest{Options: &oneof.OneOfRequest_Text{Text: "Example"}}
		buf, err := jsonpb.Marshal(in)
		require.Nil(t, err)

		reader := bytes.NewReader(buf)
		req, err := http.NewRequest("POST", server.URL+"/OneOf/Echo", reader)
		require.Nil(t, err)
		req.Header = header

		resp, err := client.Do(req)
		require.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, "200 OK", resp.Status)

		// check response
		r := &oneof.OneOfReply{}
		data, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		err = jsonpb.Unmarshal(data, r)
		assert.Nil(t, err)
		assert.Equal(t, "Example", r.GetText())
	})

	t.Run("Echo number", func(t *testing.T) {
		in := &oneof.OneOfRequest{Options: &oneof.OneOfRequest_Number{Number: 42}}
		buf, err := jsonpb.Marshal(in)
		assert.Nil(t, err)

		reader := bytes.NewReader(buf)
		req, err := http.NewRequest("POST", server.URL+"/OneOf/Echo", reader)
		require.Nil(t, err)
		req.Header = header

		resp, err := client.Do(req)
		require.Nil(t, err)
		defer resp.Body.Close()
		assert.Equal(t, "200 OK", resp.Status)

		// check response
		r := &oneof.OneOfReply{}
		data, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		err = jsonpb.Unmarshal(data, r)
		assert.Nil(t, err)
		assert.Equal(t, int64(42), r.GetNumber())
	})
}

func TestOneOfExternalClient(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", oneof.NewOneOfServer(&oneOfServerImpl{})))
	server := httptest.NewServer(mux)
	defer server.Close()

	client, err := oneof_client.NewOneOfClient(server.URL+"/api/", &http.Client{})
	assert.Nil(t, err)

	t.Run("Echo text", func(t *testing.T) {
		resp, err := client.Echo(context.Background(), &oneof_client.OneOfRequest{Options: &oneof_client.OneOfRequest_Text{Text: "Example"}})
		require.Nil(t, err)
		assert.Equal(t, "Example", resp.GetText())
	})

	t.Run("Echo number", func(t *testing.T) {
		resp, err := client.Echo(context.Background(), &oneof_client.OneOfRequest{Options: &oneof_client.OneOfRequest_Number{Number: 42}})
		require.Nil(t, err)
		assert.Equal(t, int64(42), resp.GetNumber())
	})
}
