// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	pb "go.mondoo.com/ranger-rpc/examples/rangerguard"
	"go.mondoo.com/ranger-rpc/plugins/authentication/cert"
	"go.mondoo.com/ranger-rpc/plugins/rangerguard/crypto"
)

func main() {
	// inefficient, just for testing
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    tlsconfig,
	}

	// key for signing the requests
	privateKey, err := crypto.PrivateKeyFromFile("../server/private-key.p8")
	if err != nil {
		log.Error().Err(err).Msg("could not read private key")
	}

	plugin, err := cert.NewRangerPlugin(cert.ClientConfig{
		PrivateKey: privateKey,
		Issuer:     "ranger_guard",
		Subject:    "ranger_guard_client",
		Kid:        "1",
	})
	if err != nil {
		log.Error().Err(err).Msg("could not create signer plugin")
	}

	log.Info().Msgf("start proto cLient")
	protoClient, err := pb.NewHelloWorldClient("https://localhost:8443/hello/", &http.Client{Transport: tr}, plugin)
	if err != nil {
		log.Error().Err(err).Msg("could not create hello world client")
	}

	data := &pb.HelloReq{Subject: "World"}
	protoResp, err := protoClient.Hello(context.Background(), data)
	if err == nil {
		log.Info().Msgf("Response %s", protoResp.Text) // prints "Hello World"
	} else {
		log.Error().Err(err).Msg("Could not get the response")
	}
}
