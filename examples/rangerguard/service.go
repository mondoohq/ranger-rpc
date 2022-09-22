package rangerguard

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard"
)

type HelloWorldServerImpl struct{}

func (s *HelloWorldServerImpl) Hello(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	return &HelloResp{Text: "Hello " + req.Subject}, nil
}

func (s *HelloWorldServerImpl) Info(ctx context.Context, req *Empty) (*Tags, error) {
	tags := make(map[string]string)

	val, ok := rangerguard.UserFromContext(ctx)
	if !ok {
		log.Error().Msg("could not extract falcon:user from context")
	} else {
		tags["issuer"] = val.GetIssuer()
		tags["subject"] = val.GetSubject()
		tags["name"] = val.GetName()
		tags["email"] = val.GetEmail()
		tags["groups"] = strings.Join(val.GetGroups(), ",")
	}

	return &Tags{Tags: tags}, nil
}
