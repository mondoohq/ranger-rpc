package authorization

import (
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
)

type Authorizor interface {
	Name() string
	// The default behavior for Authorize is to deny access
	Authorize(a AuthorizationFacts) (authorized Decision, reason string, err error)
}

// This AuthorizationFacts interface provides all the facts that the authorization engine can use to
// determine if a user has access or not
type AuthorizationFacts interface {
	GetUser() user.User

	// GetAction returns the action associated with API requests e.g get, create, update, patch, delete, list
	GetAction() string

	// The kind of object, that is affected by the request
	GetResource() string

	// GetAPIGroup returns the api group
	GetAPIGroup() string

	// GetAPIVersion returns the api version
	GetAPIVersion() string

	// GetPath returns the request path
	GetPath() string
}

// Decision is the response from an Authorizor
type Decision int

const (
	// DecisionAllow means that an Authorizor decided that the user is allowed to use the API.
	DecisionAllow Decision = iota

	// DecisionDeny means that an Authorizor decided to deny the request
	DecisionDeny

	// DecisionAbstention means that an Authorizor is not voting at all
	DecisionAbstention
)

type AttributesRecord struct {
	User       user.User
	Action     string
	Resource   string
	APIGroup   string
	APIVersion string
	Path       string
}

func (ar *AttributesRecord) GetUser() user.User {
	return ar.User
}

func (ar *AttributesRecord) GetAction() string {
	return ar.Action
}

func (ar *AttributesRecord) GetResource() string {
	return ar.Resource
}

func (ar *AttributesRecord) GetAPIGroup() string {
	return ar.APIGroup
}

func (ar *AttributesRecord) GetAPIVersion() string {
	return ar.APIVersion
}

func (ar *AttributesRecord) GetPath() string {
	return ar.Path
}
