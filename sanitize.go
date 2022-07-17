package ranger

import (
	"net/url"
)

// SanitizeUrl parses the given url and returns a sanitized url. If no scheme is given, it will be set to https.
func SanitizeUrl(addr string) string {
	url, err := url.Parse(addr)
	if err != nil {
		// url.Parse fails on it, return it unchanged.
		return addr
	}

	if url.Scheme == "" {
		url.Scheme = "https"
	}
	return url.String()
}
