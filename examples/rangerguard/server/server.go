package main

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	helloworld "go.mondoo.com/ranger-rpc/examples/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/authentication"
	"go.mondoo.com/ranger-rpc/plugins/authentication/cert"
	"go.mondoo.com/ranger-rpc/plugins/authorization"
	"go.mondoo.com/ranger-rpc/plugins/authorization/always"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard"
)

func main() {
	serve()
}

func serve() {
	ks := helloworld.NewKeystore()
	err := ks.Load("./cert.pem")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot not load public key")
	}

	// inefficient, just for testing
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// https server
	cfg := &tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	helloHandler := helloworld.NewHelloWorldServer(&helloworld.HelloWorldServerImpl{})
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	mux.Handle("/hello/", http.StripPrefix("/hello", helloHandler))

	ca := cert.New(cert.Config{
		KeyStore: ks,
	})
	authenticators := []authentication.Authenticator{ca}
	authorizors := []authorization.Authorizor{always.Allow()}

	srv := &http.Server{
		Addr: ":8443",
		// use a wrapper handler to verify authentication
		Handler: rangerguard.New(rangerguard.Options{
			Authenticators: authenticators,
			Authorizors:    authorizors,
		}, mux),
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Info().Msgf("start server on %s", srv.Addr)
	err = srv.ListenAndServeTLS("./cert.pem", "./private-key.pem")
	log.Fatal().
		Err(err).
		Str("service", "http").
		Msg("cannot start service")
}
