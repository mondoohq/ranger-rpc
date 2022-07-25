package ranger

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/status"
	"google.golang.org/protobuf/proto"
	"moul.io/http2curl"
)

// Client is the client for the Ranger service. It is used as the base client for all service calls.
type Client struct{}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// DoClientRequest makes a request to the Ranger service.
// It will marshal the proto.Message into the request body, do the request and then parse the response into the
// response proto.Message.
func (c *Client) DoClientRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return errors.Wrap(err, "failed to marshal proto request")
	}

	if err = ctx.Err(); err != nil {
		return errors.Wrap(err, "aborted because context was done")
	}

	header := make(http.Header)
	header.Set("Accept", "application/protobuf")
	header.Set("Content-Type", "application/protobuf")

	log.Debug().Str("url", url).Msg("call service")
	reader := bytes.NewReader(reqBodyBytes)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header = header

	// trace curl request
	if log.Trace().Enabled() {
		c.PrintTraceCurlCommand(url, in)
	}

	// do http call
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = errors.Wrap(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return errors.Wrap(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		// TODO wrap body in error
		spb, err := parseStatus(resp.Body)
		if err == nil {
			log.Debug().Str("body", spb.Message).Int("status", resp.StatusCode).Msg("non-ok http request")
			return status.FromProto(spb).Err()
		} else {
			payload, err := ioutil.ReadAll(reader)
			if err != nil {
				log.Error().Err(err).Msg("could not parse http body")
			}
			return status.New(status.CodeFromHTTPStatus(resp.StatusCode), string(payload)).Err()
		}
	}

	if err = ctx.Err(); err != nil {
		return errors.Wrap(err, "aborted because context was done")
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}
	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return errors.Wrap(err, "failed to unmarshal proto response")
	}
	return nil
}

func (c *Client) PrintTraceCurlCommand(url string, in proto.Message) {
	// for better debuggability we try to construct an equivalent json request
	jsonBytes, err := json.Marshal(in)
	if err != nil {
		log.Error().Err(err).Msg("could not generate trace http log")
	}

	header := make(http.Header)
	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")
	header.Set("Content-Length", strconv.Itoa(len(jsonBytes)))

	// create http request
	reader := bytes.NewReader(jsonBytes)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return
	}
	req.Header = header

	// convert request to curl command
	command, _ := http2curl.GetCurlCommand(req)
	log.Trace().Msg(command.String())
}
