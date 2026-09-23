package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
)

// SigningMethodHMAC implements HS256.
type SigningMethodHMAC struct {
	Name string
}

var signingMethodHS256 = &SigningMethodHMAC{Name: "HS256"}

// SigningMethodHS256 is the HMAC-SHA256 signing method.
var SigningMethodHS256 SigningMethod = signingMethodHS256

func (m *SigningMethodHMAC) Alg() string {
	return m.Name
}

func (m *SigningMethodHMAC) Verify(signingString string, sig []byte, key interface{}) error {
	keyBytes, ok := key.([]byte)
	if !ok {
		return ErrInvalidKeyType
	}
	hasher := hmac.New(sha256.New, keyBytes)
	hasher.Write([]byte(signingString))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return ErrSignatureInvalid
	}
	return nil
}

func (m *SigningMethodHMAC) Sign(signingString string, key interface{}) ([]byte, error) {
	keyBytes, ok := key.([]byte)
	if !ok {
		return nil, ErrInvalidKeyType
	}
	hasher := hmac.New(sha256.New, keyBytes)
	hasher.Write([]byte(signingString))
	return hasher.Sum(nil), nil
}
