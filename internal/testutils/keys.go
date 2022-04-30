package testutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

const (
	bitSize = 4096
)

type rsaKeyPair struct {
	PrivateKey    *rsa.PrivateKey
	PrivateKeyPEM []byte
	PublicKey     []byte
}

func GenerateRSAKeyPair() (*rsaKeyPair, error) {
	privKey, err := generatePrivateKey()
	if err != nil {
		return nil, err
	}

	pubKey, err := generatePublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, err
	}

	keyPair := &rsaKeyPair{
		PrivateKey:    privKey,
		PrivateKeyPEM: encodePrivateKeyToPEM(privKey),
		PublicKey:     pubKey,
	}

	return keyPair, nil
}

func generatePrivateKey() (*rsa.PrivateKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = privKey.Validate()
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

func generatePublicKey(privkey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privkey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

func encodePrivateKeyToPEM(privKey *rsa.PrivateKey) []byte {
	privBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	return pem.EncodeToMemory(&privBlock)
}
