package oneof

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	jsonpb "google.golang.org/protobuf/encoding/protojson"
)

// custom implementation
type oneOfServerImpl struct{}

func (s *oneOfServerImpl) Echo(ctx context.Context, in *OneOfRequest) (*OneOfReply, error) {
	resp := &OneOfReply{}
	if in.GetNumber() != 0 {
		resp.Options = &OneOfReply_Number{
			Number: in.GetNumber(),
		}
	}
	if in.GetText() != "" {
		resp.Options = &OneOfReply_Text{
			Text: in.GetText(),
		}
	}
	return resp, nil
}

func TestOneOf(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", NewOneOfServer(&oneOfServerImpl{})))
	server := httptest.NewServer(mux)
	defer server.Close()

	client, err := NewOneOfClient(server.URL+"/api/", &http.Client{})
	assert.Nil(t, err)

	t.Run("Echo text", func(t *testing.T) {
		resp, err := client.Echo(context.Background(), &OneOfRequest{Options: &OneOfRequest_Text{Text: "Example"}})
		require.Nil(t, err)
		assert.Equal(t, "Example", resp.GetText())
	})

	t.Run("Echo number", func(t *testing.T) {
		resp, err := client.Echo(context.Background(), &OneOfRequest{Options: &OneOfRequest_Number{Number: 42}})
		require.Nil(t, err)
		assert.Equal(t, int64(42), resp.GetNumber())
	})
}

func TestOneOfJSON(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", NewOneOfServer(&oneOfServerImpl{}))
	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}

	header := make(http.Header)
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")

	t.Run("Echo text", func(t *testing.T) {
		in := &OneOfRequest{Options: &OneOfRequest_Text{Text: "Example"}}
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
		r := &OneOfReply{}
		data, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		err = jsonpb.Unmarshal(data, r)
		assert.Nil(t, err)
		assert.Equal(t, "Example", r.GetText())
	})

	t.Run("Echo number", func(t *testing.T) {
		in := &OneOfRequest{Options: &OneOfRequest_Number{Number: 42}}
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
		r := &OneOfReply{}
		data, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		err = jsonpb.Unmarshal(data, r)
		assert.Nil(t, err)
		assert.Equal(t, int64(42), r.GetNumber())
	})
}
