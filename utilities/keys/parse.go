package utilkey

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
)

func ParsePrivateKey(byteBlock []byte) (*rsa.PrivateKey, error) {
	key, err := x509.ParsePKCS1PrivateKey(byteBlock)
	if err == nil {
		return key, err
	}

	key2, err := x509.ParsePKCS8PrivateKey(byteBlock)
	if err == nil {
		return key2.(*rsa.PrivateKey), err
	}

	return nil, errors.New("cannot parse private key")
}

func ParsePublicKey(byteBlock []byte) (*rsa.PublicKey, error) {
	key, err := x509.ParsePKCS1PublicKey(byteBlock)
	if err == nil {
		return key, err
	}

	key2, err := x509.ParsePKIXPublicKey(byteBlock)
	if err == nil {
		return key2.(*rsa.PublicKey), err
	}

	return nil, errors.New("cannot parse public key")
}
