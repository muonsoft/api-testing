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
	if err := decodeHeader(parts[0], token); err != nil {
		return nil, err
	}
	if err := decodeClaims(parts[1], token); err != nil {
		return nil, err
	}
	if err := decodeSignature(parts[2], token); err != nil {
		return nil, err
	}
	if err := verifyToken(token, parts, keyFunc); err != nil {
		return nil, err
	}

	token.Valid = true
	return token, nil
}

func decodeHeader(seg string, token *Token) error {
	headerBytes, err := decodeSegment(seg)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrTokenMalformed, err)
	}
	if err := json.Unmarshal(headerBytes, &token.Header); err != nil {
		return fmt.Errorf("%w: %w", ErrTokenMalformed, err)
	}

	alg, _ := token.Header["alg"].(string)
	if alg == "" {
		return fmt.Errorf("%w: signing method (alg) is unspecified", ErrTokenUnverifiable)
	}
	token.Method = &methodByAlg{alg: alg}

	return nil
}

func decodeClaims(seg string, token *Token) error {
	claimBytes, err := decodeSegment(seg)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrTokenMalformed, err)
	}
	token.Claims = MapClaims{}
	if err := json.Unmarshal(claimBytes, &token.Claims); err != nil {
		return fmt.Errorf("%w: %w", ErrTokenMalformed, err)
	}

	return nil
}

func decodeSignature(seg string, token *Token) error {
	sig, err := decodeSegment(seg)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrTokenMalformed, err)
	}
	token.Signature = sig

	return nil
}

func verifyToken(token *Token, parts []string, keyFunc Keyfunc) error {
	if keyFunc == nil {
		return fmt.Errorf("%w: no keyfunc was provided", ErrTokenUnverifiable)
	}
	key, err := keyFunc(token)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrTokenUnverifiable, err)
	}

	signingString := strings.Join(parts[0:2], ".")
	if err := token.Method.Verify(signingString, token.Signature, key); err != nil {
		return fmt.Errorf("%w: %w", ErrTokenSignatureInvalid, err)
	}

	return nil
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
