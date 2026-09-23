package assertjson

import ijwt "github.com/muonsoft/api-testing/internal/jwt"

func adaptKeyFunc(keyFunc JWTKeyFunc) ijwt.Keyfunc {
	if keyFunc == nil {
		return nil
	}

	return func(token *ijwt.Token) (interface{}, error) {
		return keyFunc(WrapJWTToken(token))
	}
}

func parseJWT(value string, keyFunc JWTKeyFunc) (*ijwt.Token, error) {
	return ijwt.Parse(value, adaptKeyFunc(keyFunc))
}

// WrapJWTToken maps an internal parsed token to the public assertion type.
func WrapJWTToken(token *ijwt.Token) *JWTToken {
	if token == nil {
		return &JWTToken{}
	}

	alg := ""
	if token.Method != nil {
		alg = token.Method.Alg()
	}

	return &JWTToken{
		Raw:    token.Raw,
		Header: token.Header,
		Claims: JWTMapClaims(token.Claims),
		alg:    alg,
	}
}

func signHS256JWT(claims JWTMapClaims, secret []byte) (string, error) {
	token := ijwt.NewWithClaims(ijwt.SigningMethodHS256, ijwt.MapClaims(claims))

	return token.SignedString(secret)
}
