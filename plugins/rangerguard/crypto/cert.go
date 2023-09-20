// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: MPL-2.0

package crypto

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

// .p8 errors file.
var (
	ErrAuthKeyNotPem   = errors.New("AuthKey must be a valid .p8 PEM file")
	ErrAuthKeyNotECDSA = errors.New("AuthKey must be of type ecdsa.PrivateKey")
	ErrAuthKeyNil      = errors.New("AuthKey was nil")
)

func PublicKeyFromCert(cert *x509.Certificate) (*ecdsa.PublicKey, error) {
	key := cert.PublicKey
	switch pk := key.(type) {
	case *ecdsa.PublicKey:
		return pk, nil
	default:
		return nil, ErrAuthKeyNotECDSA
	}
}

func CertificateFromFile(filename string) (*x509.Certificate, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return CertificateFromBytes(bytes)
}

func CertificateFromBytes(bytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, ErrAuthKeyNotPem
	}

	return x509.ParseCertificate(block.Bytes)
}

func PrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromBytes(bytes)
}

// AuthKeyFromBytes loads a .p8 certificate from an in memory byte array and
// returns an *ecdsa.PrivateKey.
func PrivateKeyFromBytes(bytes []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, ErrAuthKeyNotPem
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, ErrAuthKeyNotECDSA
	}
}
