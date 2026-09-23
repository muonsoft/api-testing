package assertions

import (
	"github.com/muonsoft/api-testing/assertjson"
	ijwt "github.com/muonsoft/api-testing/internal/jwt"
)

func adaptAssertionsKeyFunc(keyFunc JWTKeyFunc) ijwt.Keyfunc {
	if keyFunc == nil {
		return nil
	}

	return func(token *ijwt.Token) (interface{}, error) {
		return keyFunc(assertjson.WrapJWTToken(token))
	}
}
