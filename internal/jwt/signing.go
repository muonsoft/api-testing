package jwt

// SigningMethod is used to sign and verify tokens.
type SigningMethod interface {
	Verify(signingString string, sig []byte, key interface{}) error
	Sign(signingString string, key interface{}) ([]byte, error)
	Alg() string
}

// methodByAlg holds algorithm name from token header; only HS256 is verified.
type methodByAlg struct {
	alg string
}

func (m *methodByAlg) Alg() string {
	return m.alg
}

func (m *methodByAlg) Verify(signingString string, sig []byte, key interface{}) error {
	if m.alg != "HS256" {
		return ErrTokenSignatureInvalid
	}
	return signingMethodHS256.Verify(signingString, sig, key)
}

func (m *methodByAlg) Sign(signingString string, key interface{}) ([]byte, error) {
	return nil, ErrTokenUnverifiable
}
