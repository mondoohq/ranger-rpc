// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package ranger

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/status"
	"google.golang.org/protobuf/proto"
	"moul.io/http2curl"
)

// ClientPlugin is a plugin that can be used to modify the http headers of a request.
type ClientPlugin interface {
	GetName() string
	GetHeader(content []byte) http.Header
}

// Client is the client for the Ranger service. It is used as the base client for all service calls.
type Client struct {
	plugins []ClientPlugin
}

// AddPlugin adds a client plugin.
// Deprecated: use AddPlugins instead.
func (c *Client) AddPlugin(plugin ClientPlugin) {
	c.AddPlugins(plugin)
}

// AddPlugins adds one or many client plugins.
func (c *Client) AddPlugins(plugins ...ClientPlugin) {
	for i := range plugins {
		c.plugins = append(c.plugins, plugins[i])
	}
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type vtprotoMessage interface {
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
}

// DoClientRequest makes a request to the Ranger service.
// It will marshal the proto.Message into the request body, do the request and then parse the response into the
// response proto.Message.
func (c *Client) DoClientRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	var reqBodyBytes []byte
	if m, ok := in.(vtprotoMessage); ok {
		log.Debug().Msgf("using vtproto for input")
		reqBodyBytes, err = m.MarshalVT()
	} else {
		reqBodyBytes, err = proto.Marshal(in)
	}
	if err != nil {
		return errors.Wrap(err, "failed to marshal proto request")
	}

	if err = ctx.Err(); err != nil {
		return errors.Wrap(err, "aborted because context was done")
	}

	header := make(http.Header)
	header.Set("Accept", "application/protobuf")
	header.Set("Content-Type", "application/protobuf")

	for i := range c.plugins {
		p := c.plugins[i]
		pluginHeader := p.GetHeader(reqBodyBytes)
		for k, v := range pluginHeader {
			// check if we overwrite an existing header
			header[k] = v
		}
	}

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
			payload, err := io.ReadAll(reader)
			if err != nil {
				log.Error().Err(err).Msg("could not parse http body")
			}
			return status.New(status.CodeFromHTTPStatus(resp.StatusCode), string(payload)).Err()
		}
	}

	if err = ctx.Err(); err != nil {
		return errors.Wrap(err, "aborted because context was done")
	}

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}
	if m, ok := out.(vtprotoMessage); ok {
		log.Debug().Msgf("using vtproto for output")
		if err := m.UnmarshalVT(respBodyBytes); err != nil {
			return errors.Wrap(err, "failed to unmarshal proto response")
		}
	} else if err = proto.Unmarshal(respBodyBytes, out); err != nil {
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
