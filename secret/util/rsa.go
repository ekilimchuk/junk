package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func LoadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	priv, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v - %s", err, filePath))
	}
	privPem, _ := pem.Decode(priv)
	if privPem == nil {
		return nil, errors.New(fmt.Sprintf("The RSA private key is a wrong (Decode error) - %s", filePath))
	}
	if privPem.Type != "RSA PRIVATE KEY" {
		return nil, errors.New(fmt.Sprintf("The RSA private key is a wrong type (%s) - %s", privPem.Type, filePath))
	}
	parsedKey, err := x509.ParsePKCS1PrivateKey(privPem.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v - %s", err, filePath))
	}
	return parsedKey, nil
}

func LoadPublicKey(filePath string) (*rsa.PublicKey, error) {
	pub, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v - %s", err, filePath))
	}
	pubPem, _ := pem.Decode(pub)
	if pubPem == nil {
		return nil, errors.New(fmt.Sprintf("The RSA public key is a wrong (Decode         error) - %s", filePath))
	}
	if pubPem.Type != "RSA PUBLIC KEY" {
		return nil, errors.New(fmt.Sprintf("The RSA public key is a wrong type (%s) - %s", pubPem.Type, filePath))
	}
	parsedKey, err := x509.ParsePKCS1PublicKey(pubPem.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v - %s", err, filePath))
	}
	return parsedKey, nil
}

func EncryptOAEP(secret []byte, key *rsa.PublicKey) ([]byte, error) {
	label := []byte("orders")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, key, secret, label)
	if err != nil {
		return ciphertext, err
	}
	return ciphertext, nil
}

func DecryptOAEP(secret []byte, key *rsa.PrivateKey) (string, error) {
	label := []byte("orders")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, key, secret, label)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
