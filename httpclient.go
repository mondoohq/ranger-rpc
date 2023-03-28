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

// DefaultHttpClient will set up a basic client
// with default timeouts/proxies/etc.
func DefaultHttpClient() *http.Client {
	return NewHttpClient()
}

type HttpClientOption func(c *http.Client)

func WithProxy(proxy *url.URL) HttpClientOption {
	return func(c *http.Client) {
		if proxy == nil {
			return
		}

		switch t := c.Transport.(type) {
		case *http.Transport:
			t.Proxy = http.ProxyURL(proxy)
		}
	}
}

func NewHttpClient(opts ...HttpClientOption) *http.Client {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   DefaultHttpTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       DefaultIdleConnTimeout,
		TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}

	c := &http.Client{
		Transport: tr,
		Timeout:   DefaultHttpTimeout,
	}

	for i := range opts {
		opts[i](c)
	}
	return c
}
