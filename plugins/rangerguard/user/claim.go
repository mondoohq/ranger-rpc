package user

import (
	"encoding/json"
	"strings"

	"github.com/cockroachdb/errors"
)

// default claim names
const (
	ClaimIssuer        = "iss"
	ClaimSubject       = "sub"
	ClaimEmail         = "email"
	ClaimEmailVerified = "email_verified"
	ClaimName          = "name"
	ClaimGivenName     = "given_name"
	ClaimFamilyName    = "family_name"
	ClaimGroups        = "groups"
	ClaimPicture       = "picture"
)

type Claims map[string]json.RawMessage

func (c Claims) UnmarshalClaim(id string, v interface{}) error {
	val, ok := c[id]
	if !ok {
		return errors.New("claim " + id + " not present")
	}
	return json.Unmarshal([]byte(val), v)
}

func (c Claims) HasClaim(id string) bool {
	if _, ok := c[id]; !ok {
		return false
	}
	return true
}

// ParseClaims extracts basic information from the claims
// standard claims are defined in https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
// we follow the required fields from google https://developers.google.com/identity/protocols/OpenIDConnect#server-flow
// required claims: iss, sub
// optional claims: name, email, email_verified
func ParseClaims(c *Claims) (User, error) {
	user := UserInfo{}

	// the default claim is "sub"
	if c.HasClaim(ClaimIssuer) {
		var issuer string
		if err := c.UnmarshalClaim(ClaimIssuer, &issuer); err != nil {
			return nil, err
		}
		user.Issuer = issuer
	}

	if c.HasClaim(ClaimSubject) {
		var subject string
		if err := c.UnmarshalClaim(ClaimSubject, &subject); err != nil {
			return nil, err
		}
		user.Subject = subject
	}

	// if a human, we may have a name field
	if c.HasClaim(ClaimName) {
		var name string
		if err := c.UnmarshalClaim(ClaimName, &name); err != nil {
			return nil, err
		}
		user.Name = name
	}

	// If the email claim is used, email_verified claim needs to be present
	// that is not true for all implementations, lets follow kubernetes approach
	// https://github.com/kubernetes/kubernetes/issues/59496
	if c.HasClaim(ClaimEmail) {
		var email string
		if err := c.UnmarshalClaim(ClaimEmail, &email); err != nil {
			return nil, err
		}
		user.Email = strings.ToLower(email)
	}

	if c.HasClaim(ClaimEmailVerified) {
		var emailVerified bool
		if err := c.UnmarshalClaim(ClaimEmailVerified, &emailVerified); err != nil {
			return nil, errors.Wrap(err, "guuard oidc> could not parse 'email_verified' claim")
		}
	}

	// if the claims include groups, we parse them too
	if c.HasClaim(ClaimGroups) {
		var groups []string
		if err := c.UnmarshalClaim(ClaimGroups, &groups); err != nil {
			return nil, err
		}
		user.Groups = groups
	}

	return &user, nil
}
