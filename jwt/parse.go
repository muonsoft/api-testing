package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

const tokenDelimiter = "."

// Parse parses and verifies the JWT and returns the token.
// Only HS256 signature verification is supported.
func Parse(tokenString string, keyFunc Keyfunc) (*Token, error) {
	parts, ok := splitToken(tokenString)
	if !ok {
		return nil, fmt.Errorf("%w: token contains an invalid number of segments", ErrTokenMalformed)
	}

	token := &Token{Raw: tokenString}

	// Decode header
	headerBytes, err := decodeSegment(parts[0])
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTokenMalformed, err)
	}
	if err := json.Unmarshal(headerBytes, &token.Header); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTokenMalformed, err)
	}

	// Decode claims
	claimBytes, err := decodeSegment(parts[1])
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTokenMalformed, err)
	}
	token.Claims = MapClaims{}
	if err := json.Unmarshal(claimBytes, &token.Claims); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTokenMalformed, err)
	}

	// Resolve signing method from header
	alg, _ := token.Header["alg"].(string)
	if alg == "" {
		return nil, fmt.Errorf("%w: signing method (alg) is unspecified", ErrTokenUnverifiable)
	}
	token.Method = &methodByAlg{alg: alg}

	// Decode signature
	token.Signature, err = decodeSegment(parts[2])
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTokenMalformed, err)
	}

	if keyFunc == nil {
		return nil, fmt.Errorf("%w: no keyfunc was provided", ErrTokenUnverifiable)
	}
	key, err := keyFunc(token)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTokenUnverifiable, err)
	}

	signingString := strings.Join(parts[0:2], ".")
	if err := token.Method.Verify(signingString, token.Signature, key); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTokenSignatureInvalid, err)
	}

	token.Valid = true
	return token, nil
}

func splitToken(s string) ([]string, bool) {
	parts := strings.SplitN(s, tokenDelimiter, 4)
	if len(parts) != 3 {
		return nil, false
	}
	if parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return nil, false
	}
	return parts, true
}

func decodeSegment(seg string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(seg)
}
