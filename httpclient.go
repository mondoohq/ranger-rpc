package ranger

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	DefaultHttpTimeout         = 30 * time.Second
	DefaultIdleConnTimeout     = 30 * time.Second
	DefaultTLSHandshakeTimeout = 10 * time.Second
)

func newHttpTransport(proxy string) (*http.Transport, error) {
	proxyFn, err := newProxyFunc(proxy)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{
		Proxy: proxyFn,
		DialContext: (&net.Dialer{
			Timeout:   DefaultHttpTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       DefaultIdleConnTimeout,
		TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return tr, nil
}

func newHttpClient(tr *http.Transport) *http.Client {
	return &http.Client{
		Transport: tr,
		Timeout:   DefaultHttpTimeout,
	}
}

// DefaultHttpClient will set up a basic client
// with default timeouts/proxies/etc.
func DefaultHttpClient() *http.Client {
	tr, err := newHttpTransport("")
	if err != nil {
		log.Fatal().Err(err).Msg("unexpectedly failed to create a default HttpClient")
	}

	return newHttpClient(tr)
}

type HttpClientOpts struct {
	// Proxy is the string representation of the proxy the client
	// should use for connections.
	Proxy string
}

func HttpClient(opts *HttpClientOpts) (*http.Client, error) {
	tr, err := newHttpTransport(opts.Proxy)
	if err != nil {
		return nil, err
	}

	return newHttpClient(tr), nil
}

type proxyFunc func(*http.Request) (*url.URL, error)

// newProxyFunc will use a standard ProxyFromEnvironment if no proxy
// information is provided. Otherwise, parse the proxy string, and use
// it as the proxy endpoint.
func newProxyFunc(proxy string) (proxyFunc, error) {
	proxyFn := http.ProxyFromEnvironment

	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}

		proxyFn = func(*http.Request) (*url.URL, error) {
			return proxyURL, nil
		}
	}

	return proxyFn, nil
}
