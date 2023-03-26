package ranger

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	DefaultHttpTimeout         = 30 * time.Second
	DefaultIdleConnTimeout     = 30 * time.Second
	DefaultTLSHandshakeTimeout = 10 * time.Second
)

func newHttpTransport(proxy *url.URL) *http.Transport {
	proxyFn := http.ProxyFromEnvironment
	if proxy != nil {
		proxyFn = http.ProxyURL(proxy)
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

	return tr
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
	tr := newHttpTransport(nil)

	return newHttpClient(tr)
}

type HttpClientOpts struct {
	// Proxy is the string representation of the proxy the client
	// should use for connections.
	Proxy string
}

func NewHttpClient(opts *HttpClientOpts) (*http.Client, error) {
	var proxy *url.URL
	if opts.Proxy != "" {
		var err error
		proxy, err = url.Parse(opts.Proxy)
		if err != nil {
			return nil, err
		}
	}
	tr := newHttpTransport(proxy)

	return newHttpClient(tr), nil
}
