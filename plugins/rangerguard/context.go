package rangerguard

import (
	"context"

	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
)

type guardUserIdentifier struct{}

func UserFromContext(ctx context.Context) (u user.User, ok bool) {
	u, ok = ctx.Value(guardUserIdentifier{}).(user.User)
	return
}

func NewUserContext(ctx context.Context, u user.User) context.Context {
	return context.WithValue(ctx, guardUserIdentifier{}, u)
}
