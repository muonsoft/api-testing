package assertjson

// JWTMapClaims is the decoded JWT payload used in assertions and test helpers.
type JWTMapClaims map[string]interface{}

// JWTKeyFunc supplies the verification key while parsing a JWT string.
type JWTKeyFunc func(token *JWTToken) (interface{}, error)

// JWTToken is a parsed and verified JWT exposed to test code.
type JWTToken struct {
	Raw    string
	Header map[string]interface{}
	Claims JWTMapClaims
	alg    string
}

// Algorithm returns the JWT "alg" header value.
func (t *JWTToken) Algorithm() string {
	if t == nil {
		return ""
	}

	return t.alg
}

// SignHS256JWT builds a compact HS256 JWT for tests.
func SignHS256JWT(claims JWTMapClaims, secret []byte) (string, error) {
	return signHS256JWT(claims, secret)
}
