package header

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
)

// ExtractTokenFromBearer extracts the access token
func ExtractTokenFromBearer(token string) string {
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:]
	}
	if strings.ToUpper(token) == "BEARER" {
		return ""
	}
	return token
}

func ExtractToken(req *http.Request) (string, error) {
	// let's try to read the bearer token from the authorization header
	token := ExtractTokenFromBearer(req.Header.Get("Authorization"))

	// fallback to query-param
	// http://self-issued.info/docs/draft-ietf-oauth-v2-bearer.html#query-param
	if len(token) == 0 {
		u, err := url.Parse(req.RequestURI)
		if err != nil {
			log.Warn().Err(err).Msg("could not extract bearer token from url query parameter")
			return "", nil
		}

		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			return "", nil
		}
		token = m.Get("access_token")
	}

	return token, nil
}
