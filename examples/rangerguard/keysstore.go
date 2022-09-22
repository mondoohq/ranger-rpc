package rangerguard

import (
	"crypto/x509"
	"errors"

	"go.mondoo.com/ranger-rpc/plugins/rangerguard/crypto"
)

type Keystore struct {
	entries map[string]*x509.Certificate
}

func NewKeystore() *Keystore {
	ks := &Keystore{}
	ks.entries = make(map[string]*x509.Certificate)
	return ks
}

func (ks *Keystore) Get(key string) (*x509.Certificate, error) {
	e, ok := ks.entries[key]
	if !ok {
		return nil, errors.New("could not find key")
	}
	return e, nil
}

func (ks *Keystore) Load(file string) error {
	// key id 1
	cert, err := crypto.CertificateFromFile(file)
	if err != nil {
		return err
	}
	ks.entries["1"] = cert
	return nil
}
