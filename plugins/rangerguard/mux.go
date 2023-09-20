// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package rangerguard

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/ranger-rpc"
	"go.mondoo.com/ranger-rpc/codes"
	"go.mondoo.com/ranger-rpc/plugins/authentication"
	"go.mondoo.com/ranger-rpc/plugins/authorization"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/user"
	"go.mondoo.com/ranger-rpc/status"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("go.mondoo.com/ranger-rpc/plugins/rangerguard")

var (
	AUTHENTICATION_DENIED_ERROR = status.Error(codes.Unauthenticated, "request permission unauthenticated")
	PERMISSION_DENIED_ERROR     = status.Error(codes.PermissionDenied, "request permission denied")
)

type Options struct {
	Authenticators []authentication.Authenticator
	Authorizors    []authorization.Authorizor
	Hooks          []Hook
}

func New(opts Options, next http.Handler) *guardMux {
	if opts.Authenticators == nil || len(opts.Authenticators) == 0 {
		log.Warn().Str("component", "guard").Msg("no authenticator set, access will be denied")
	}
	if opts.Authorizors == nil || len(opts.Authorizors) == 0 {
		log.Warn().Str("component", "guard").Msg("no authorizer set, access will be denied")
	}
	return &guardMux{
		authenticators: opts.Authenticators,
		authorizors:    opts.Authorizors,
		hooks:          opts.Hooks,
		next:           next,
		middlewares:    []http.HandlerFunc{},
	}
}

type Hook interface {
	Name() string
	// Run is called a call has a validated identity
	// it can optionally return a new Context that will be attached to the incoming request
	Run(context.Context, user.User, *http.Request) (context.Context, error)
}

// guardMux secures a go http mux handler
// if no authenticator or authorizor is defined, the request will be denied
// use the AllowAuthenticator and AllowAuthorizor middleware to allow access to specific routes
type guardMux struct {
	next           http.Handler
	authenticators []authentication.Authenticator
	authorizors    []authorization.Authorizor
	hooks          []Hook
	// Slice of middlewares to be called after a match is found
	middlewares []http.HandlerFunc
}

// Use appends a http.HandlerFunc to the chain. Those functions can be used to intercept or modify requests/responses,
// and are executed in the order that they are added
func (r *guardMux) Use(mwf ...http.HandlerFunc) {
	for _, fn := range mwf {
		r.middlewares = append(r.middlewares, fn)
	}
}

func (gm guardMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	guardCtx, span := tracer.Start(r.Context(), "ranger.guard.ServeHTTP")
	defer span.End()
	// iterate over middle ware
	for i := len(gm.middlewares) - 1; i >= 0; i-- {
		fn := gm.middlewares[i]
		fn(w, r)
	}

	// handle authentication calls
	user, authenticated, authnReq := gm.authenticate(r)
	if !authenticated || user == nil {
		log.Info().Str("component", "guard").Str("client", r.RemoteAddr).Str("uri", r.RequestURI).Msg("Unauthenticated")
		ranger.HttpError(w, authnReq, AUTHENTICATION_DENIED_ERROR)
		span.End()
		return
	}

	// tell hooks about the user
	for i := range gm.hooks {
		_, span := tracer.Start(guardCtx, "ranger.guard.ServeHTTP/hook/"+gm.hooks[i].Name())
		newCtx, err := gm.hooks[i].Run(authnReq.Context(), user, r)
		span.End()
		if err != nil {
			log.Error().Err(err).Str("hook", gm.hooks[i].Name()).Msg("could not authenticate because ranger guard hook returned an error")
			ranger.HttpError(w, authnReq, AUTHENTICATION_DENIED_ERROR)
			span.End()
			return
		}
		// assign new context
		if newCtx != nil {
			authnReq = authnReq.WithContext(newCtx)
		}
	}

	authorized, authzReq := gm.authorize(authnReq, user)
	if !authorized {
		log.Info().Str("component", "guard").Str("client", r.RemoteAddr).Str("uri", r.RequestURI).Msg("Unauthorized")
		ranger.HttpError(w, authzReq, PERMISSION_DENIED_ERROR)
		span.End()
		return
	}

	span.End()
	gm.next.ServeHTTP(w, authzReq)
}

func (gm guardMux) authenticate(r *http.Request) (user.User, bool, *http.Request) {
	// log := logger.FromContext(r.Context())
	_, span := tracer.Start(r.Context(), "ranger.guard.authenticate")
	defer span.End()

	// iterate over the authenticators, once one is passing, we go forward
	errMsgs := []string{}
	for i := range gm.authenticators {
		authenticator := gm.authenticators[i]
		user, valid, err := authenticator.Verify(r)
		if err != nil {
			errMsgs = append(errMsgs, err.Error())
			continue
		}
		if valid {
			log.Debug().Str("component", "guard").
				Str("authenticator", authenticator.Name()).
				Str("client", r.RemoteAddr).
				Str("subject", user.GetSubject()).
				Str("issuer", user.GetIssuer()).
				Str("email", user.GetEmail()).
				Str("uri", r.RequestURI).
				Msg("request authenticated")

			// set user into request context
			ctx := NewUserContext(r.Context(), user)
			req := r.WithContext(ctx)

			return user, true, req
		}
	}

	log.Debug().
		Err(errors.New(strings.Join(errMsgs, ", "))).
		Str("component", "guard").
		Str("client", r.RemoteAddr).
		Str("uri", r.RequestURI).
		Msg("request not authenticated")
	return nil, false, r
}

func (gm guardMux) authorize(r *http.Request, user user.User) (bool, *http.Request) {
	// log := logger.FromContext(r.Context())
	spanCtx, span := tracer.Start(r.Context(), "ranger.guard.authorize")
	defer span.End()

	af := &authorization.AttributesRecord{
		User:   user,
		Action: r.Method,
		Path:   r.URL.Path,
	}

	for i := range gm.authorizors {
		authorizor := gm.authorizors[i]
		_, authSpan := tracer.Start(spanCtx, "ranger.guard.authorize/hook/"+authorizor.Name())
		decision, _, err := authorizor.Authorize(af)
		authSpan.End()
		if err != nil {
			continue
		}

		if decision == authorization.DecisionAllow {
			log.Debug().Str("component", "guard").
				Str("authorizor", authorizor.Name()).
				Str("client", r.RemoteAddr).
				Str("uri", r.RequestURI).
				Str("subject", user.GetSubject()).
				Str("issuer", user.GetIssuer()).
				Str("email", user.GetEmail()).
				Msg("request authorized")
			return true, r
		}
	}

	log.Debug().
		Str("component", "guard").
		Str("client", r.RemoteAddr).
		Str("subject", user.GetSubject()).
		Str("issuer", user.GetIssuer()).
		Str("email", user.GetEmail()).
		Str("uri", r.RequestURI).
		Msg("request not authorized")
	return false, r
}
