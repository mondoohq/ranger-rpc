// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

// User describes the authenticated user
type User interface {
	// GetIssuer returns the issuer of the subject id
	GetIssuer() string

	// GetSubject returns a unique user id, it is expected to stay stable
	GetSubject() string

	// GetName returns a human-readable name of the user
	GetName() string

	// GetEmail returns the users email, only if available and verified
	GetEmail() string

	// GetGroups returns the names of the groups the user is a member of
	GetGroups() []string

	// GetSignInProvider returns the sign-in provider that authenticated the user
	GetSignInProvider() string

	// GetLabels returns customer information for the user
	GetLabels() map[string]string

	// GetMrn returns the mrn set in the claims, or an empty string if non is found
	GetMrn() string
}
