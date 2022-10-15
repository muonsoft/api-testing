package assertjson

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

// JWTAssertion is used to build a chain of assertions for the JWT node.
type JWTAssertion struct {
	t       TestingT
	message string
	path    string
	token   *jwt.Token
}

// IsJWT asserts that the JSON node has a string value with JWT.
func (node *AssertNode) IsJWT(keyFunc jwt.Keyfunc, msgAndArgs ...interface{}) *JWTAssertion {
	node.t.Helper()
	return node.IsString().WithJWT(keyFunc, msgAndArgs...)
}

// WithJWT asserts that the JSON node has a string value with JWT.
func (a *StringAssertion) WithJWT(keyFunc jwt.Keyfunc, msgAndArgs ...interface{}) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	token, err := jwt.Parse(a.value, keyFunc)
	if err == nil {
		return &JWTAssertion{t: a.t, message: a.message, path: a.path, token: token}
	}

	a.fail(
		fmt.Sprintf(
			`failed asserting that JSON node "%s" is JWT: %s`,
			a.path,
			err.Error(),
		),
		msgAndArgs...,
	)

	return nil
}

// Algorithm asserts that the JWT is signed with expected algorithm ("alg" header).
func (a *JWTAssertion) Algorithm(alg string, msgAndArgs ...interface{}) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.token.Method.Alg() != alg {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is JWT with algorithm "%s", actual is "%s"`,
				a.path,
				alg,
				a.token.Method.Alg(),
			),
			msgAndArgs...,
		)
	}

	return a
}

// Header executes JSON assertion on JWT header.
func (a *JWTAssertion) Header(jsonAssert JSONAssertFunc) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	jsonAssert(&AssertJSON{
		t:       a.t,
		message: fmt.Sprintf(`failed asserting that JSON node "%s" is JWT with header`, a.path),
		data:    a.token.Header,
	})

	return a
}

// Payload executes JSON assertion on JWT payload.
func (a *JWTAssertion) Payload(jsonAssert JSONAssertFunc) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	jsonAssert(&AssertJSON{
		t:       a.t,
		message: fmt.Sprintf(`failed asserting that JSON node "%s" is JWT with payload`, a.path),
		data:    map[string]interface{}(a.token.Claims.(jwt.MapClaims)),
	})

	return a
}

// Token returns decoded jwt.Token.
func (a *JWTAssertion) Token() *jwt.Token {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.token
}

// Assert asserts that the JWT is satisfied by the user function assertFunc.
func (a *JWTAssertion) Assert(assertFunc func(tb testing.TB, token *jwt.Token)) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	assertFunc(a.t.(testing.TB), a.token)

	return a
}

func (a *JWTAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	if a.message != "" {
		message = a.message + ": " + message
	}
	assert.Fail(a.t, message, msgAndArgs...)
}
