package jwt

import (
	"encoding/base64"
	"encoding/json"
)

// Keyfunc is used by Parse to supply the key for verification.
// The function receives the parsed but unverified Token (e.g. to read "alg" from header).
type Keyfunc func(*Token) (interface{}, error)

// Token represents a JWT.
type Token struct {
	Raw      string
	Method   SigningMethod
	Header   map[string]interface{}
	Claims   MapClaims
	Signature []byte
	Valid    bool
}

// NewWithClaims creates a new Token with the given signing method and claims.
func NewWithClaims(method SigningMethod, claims MapClaims) *Token {
	if claims == nil {
		claims = MapClaims{}
	}
	return &Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}

// SignedString signs the token and returns the full JWT string.
func (t *Token) SignedString(key interface{}) (string, error) {
	sstr, err := t.SigningString()
	if err != nil {
		return "", err
	}
	sig, err := t.Method.Sign(sstr, key)
	if err != nil {
		return "", err
	}
	t.Signature = sig
	return sstr + "." + t.EncodeSegment(sig), nil
}

// SigningString returns the base64url(header).base64url(claims) string.
func (t *Token) SigningString() (string, error) {
	h, err := json.Marshal(t.Header)
	if err != nil {
		return "", err
	}
	c, err := json.Marshal(t.Claims)
	if err != nil {
		return "", err
	}
	return t.EncodeSegment(h) + "." + t.EncodeSegment(c), nil
}

// EncodeSegment encodes bytes to base64url without padding.
func (t *Token) EncodeSegment(seg []byte) string {
	return base64.RawURLEncoding.EncodeToString(seg)
}
