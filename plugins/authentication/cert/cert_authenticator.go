package cert

import (
	"bytes"
	"crypto/subtle"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/crypto"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/header"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
	jwt "gopkg.in/square/go-jose.v2/jwt"
)

type KeyStore interface {
	Get(keyid string) (*x509.Certificate, error)
}

type Config struct {
	KeyStore KeyStore
	Expected jwt.Expected
}

func New(cfg Config) *certificateAuthenticator {
	return &certificateAuthenticator{
		ks: cfg.KeyStore,
		verifier: &CertJWTVerifier{
			Ks:       cfg.KeyStore,
			Expected: cfg.Expected,
		},
	}
}

type certificateAuthenticator struct {
	ks       KeyStore
	verifier *CertJWTVerifier
}

func (ca *certificateAuthenticator) Name() string {
	return "Certificate Authenticator"
}

func (ca *certificateAuthenticator) Verify(r *http.Request) (user.User, bool, error) {
	verifier := ca.verifier
	payload := ca.readBody(r)
	hash := crypto.HashPayload(payload)

	bearerToken, err := header.ExtractToken(r)
	if err != nil {
		return nil, false, err
	}

	// abort checking if the bearer token is empty
	if len(bearerToken) == 0 {
		return nil, false, nil
	}

	var c user.Claims
	err = verifier.Validate(bearerToken, &c)
	if err != nil {
		// log.Debug().Err(err).Str("component", "guard").Msg("validate token")
		return nil, false, err
	}

	// hash from jwt claim
	// NOTE: at this point we make content signing optional to experiment with the public api
	// Unfortunately there is no easy standard that it implemented in all API tooling
	identical := true
	var chash string
	if c.HasClaim("chash") {
		err = c.UnmarshalClaim("chash", &chash)
		if err != nil {
			log.Debug().Err(err).Str("component", "guard").Msg("could not extract content hash")
			return nil, false, err
		}

		// hash from serialized content
		shash := fmt.Sprintf("%x", hash)

		identical = subtle.ConstantTimeCompare([]byte(shash), []byte(chash)) == 1
		if !identical {
			log.Debug().Str("component", "guard").Str("hash generated", shash).Str("hash encoded", chash).Msg("hash do not match")
		}
	}

	ui, err := user.ParseClaims(&c)
	if err != nil {
		return nil, false, err
	}

	return ui, identical, nil
}

func (ca *certificateAuthenticator) readBody(request *http.Request) []byte {
	if request.Body == nil {
		return []byte{}
	}
	payload, _ := io.ReadAll(request.Body)
	request.Body = io.NopCloser(bytes.NewReader(payload))
	return payload
}

type CertJWTVerifier struct {
	Ks       KeyStore
	Expected jwt.Expected
}

func (v *CertJWTVerifier) Validate(rawToken string, dest *user.Claims) error {
	tok, err := jwt.ParseSigned(rawToken)
	if err != nil {
		return err
	}

	// determine the public key to verify a bearer token
	if len(tok.Headers) != 1 {
		return errors.New("multiple headers not supported")
	}

	certificate, err := v.Ks.Get(tok.Headers[0].KeyID)
	if err != nil {
		return errors.Wrap(err, "invalid keyid in jwt")
	}

	// verify the signing algo
	if len(tok.Headers) == 0 || tok.Headers[0].Algorithm != "ES384" {
		return errors.New("unsupported signature algorithm token, only ES384 is supported")
	}

	// check that the token is valid
	out := jwt.Claims{}
	pk, err := crypto.PublicKeyFromCert(certificate)
	if err != nil {
		return err
	}
	if err := tok.Claims(pk, &out); err != nil {
		return err
	}

	err = out.Validate(v.Expected.WithTime(time.Now()))
	if err != nil {
		return err
	}

	// parse claims into user struct
	if err := tok.Claims(pk, dest); err != nil {
		return err
	}

	return nil
}
