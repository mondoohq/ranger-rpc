package user

var Anonymous = &UserInfo{
	Subject: "anonymous",
	Name:    "anonymous",
	Issuer:  "system/anonymous",
}

var SystemAdmin = &UserInfo{
	Subject: "admin",
	Issuer:  "system/admin",
	Name:    "system-admin",
}

// DefaultInfo provides a simple user information exchange object
type UserInfo struct {
	Mrn            string
	Issuer         string
	Subject        string
	Name           string
	Email          string
	Groups         []string
	SignInProvider string
	Labels         map[string]string
}

func (i *UserInfo) GetIssuer() string {
	return i.Issuer
}

func (i *UserInfo) GetSubject() string {
	return i.Subject
}

func (i *UserInfo) GetName() string {
	return i.Name
}

func (i *UserInfo) GetEmail() string {
	return i.Email
}

func (i *UserInfo) GetGroups() []string {
	return i.Groups
}

func (i *UserInfo) GetSignInProvider() string {
	return i.SignInProvider
}

func (i *UserInfo) GetLabels() map[string]string {
	return i.Labels
}

func (i *UserInfo) GetMrn() string {
	return i.Mrn
}
