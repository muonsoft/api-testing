package jwt

import "errors"

var (
	ErrTokenMalformed        = errors.New("token is malformed")
	ErrTokenUnverifiable     = errors.New("token is unverifiable")
	ErrTokenSignatureInvalid = errors.New("token signature is invalid")
	ErrSignatureInvalid      = errors.New("signature is invalid")
	ErrInvalidKeyType        = errors.New("key is of invalid type")
)
