package assertjson

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"

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

// WithAlgorithm asserts that the JWT is signed with expected algorithm ("alg" header).
func (a *JWTAssertion) WithAlgorithm(alg string, msgAndArgs ...interface{}) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.token.Method.Alg() != alg {
		a.fail(
			fmt.Sprintf(
				`is JWT with algorithm "%s", actual is "%s"`,
				alg,
				a.token.Method.Alg(),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithHeader executes JSON assertion on JWT header.
func (a *JWTAssertion) WithHeader(jsonAssert JSONAssertFunc) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	jsonAssert(&AssertJSON{
		t:       a.t,
		message: a.message + `is JWT with header: `,
		data:    a.token.Header,
	})

	return a
}

// WithPayload executes JSON assertion on JWT payload.
func (a *JWTAssertion) WithPayload(jsonAssert JSONAssertFunc) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	jsonAssert(&AssertJSON{
		t:       a.t,
		message: a.message + `is JWT with payload: `,
		data:    map[string]interface{}(a.token.Claims.(jwt.MapClaims)),
	})

	return a
}

// WithID asserts that the JWT is has id field ("jti" field in payload) with the expected value.
func (a *JWTAssertion) WithID(expected string, msgAndArgs ...interface{}) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.assertStringField("id", "jti", expected, msgAndArgs...)
}

// WithIssuer asserts that the JWT is has issuer field ("iss" field in payload) with the expected value.
func (a *JWTAssertion) WithIssuer(expected string, msgAndArgs ...interface{}) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.assertStringField("issuer", "iss", expected, msgAndArgs...)
}

// WithSubject asserts that the JWT is has subject field ("sub" field in payload) with the expected value.
func (a *JWTAssertion) WithSubject(expected string, msgAndArgs ...interface{}) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.assertStringField("subject", "sub", expected, msgAndArgs...)
}

// WithAudience asserts that the JWT is has audience field ("aud" field in payload) with the expected values.
func (a *JWTAssertion) WithAudience(expected []string, msgAndArgs ...interface{}) *JWTAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.assertStringsField("audience", "aud", expected, msgAndArgs...)
}

// WithExpiresAt asserts that the JWT is has expired at field ("exp" field in payload).
// It runs TimeAssertion on its value.
func (a *JWTAssertion) WithExpiresAt() *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.assertTimeField("expires at", "exp")
}

// WithNotBefore asserts that the JWT is has not before field ("nbf" field in payload).
// It runs TimeAssertion on its value.
func (a *JWTAssertion) WithNotBefore() *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.assertTimeField("not before", "nbf")
}

// WithIssuedAt asserts that the JWT is has issued at field ("iat" field in payload).
// It runs TimeAssertion on its value.
func (a *JWTAssertion) WithIssuedAt() *TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return a.assertTimeField("issued at", "iat")
}

// Value returns decoded jwt.Token. If parsing fails it will return empty struct.
func (a *JWTAssertion) Value() *jwt.Token {
	if a == nil {
		return &jwt.Token{}
	}
	a.t.Helper()

	return a.token
}

// JWT asserts that the JSON node is JWT and returns decoded jwt.Token. If value is not a valid JWT,
// then it will return empty struct. It is an alias for IsJWT().Value().
func (node *AssertNode) JWT(keyFunc jwt.Keyfunc) *jwt.Token {
	return node.IsJWT(keyFunc).Value()
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

func (a *JWTAssertion) assertStringField(title string, name string, expected string, msgAndArgs ...interface{}) *JWTAssertion {
	a.t.Helper()

	raw, exist := a.token.Claims.(jwt.MapClaims)[name]
	if !exist {
		return a.failOnMissingField(title, name, strconv.Quote(expected), msgAndArgs...)
	}

	value, ok := raw.(string)
	if !ok {
		return a.failOnUnexpectedType(title, name, strconv.Quote(expected), "string is expected", msgAndArgs...)
	}

	if value != expected {
		return a.failOnNotEqual(title, name, strconv.Quote(expected), strconv.Quote(value), msgAndArgs...)
	}

	return a
}

func (a *JWTAssertion) assertStringsField(title string, name string, expected []string, msgAndArgs ...interface{}) *JWTAssertion {
	a.t.Helper()

	raw, exist := a.token.Claims.(jwt.MapClaims)[name]
	if !exist {
		return a.failOnMissingField(title, name, wrapArray(formatStrings(expected)), msgAndArgs...)
	}

	actual, ok := castToStrings(raw)
	if !ok {
		return a.failOnUnexpectedType(title, name, wrapArray(formatStrings(expected)), "string or array of strings expected", msgAndArgs...)
	}

	if !areStringsEqual(actual, expected) {
		return a.failOnNotEqual(title, name, wrapArray(formatStrings(expected)), wrapArray(formatStrings(actual)), msgAndArgs...)
	}

	return a
}

func (a *JWTAssertion) assertTimeField(title string, name string) *TimeAssertion {
	raw, exist := a.token.Claims.(jwt.MapClaims)[name]
	if !exist {
		a.failOnMissingField(title, name, "")
		return nil
	}

	value, ok := raw.(float64)
	if !ok {
		a.failOnUnexpectedType(title, name, "", "number is expected")
		return nil
	}

	return &TimeAssertion{
		t:       a.t,
		message: fmt.Sprintf(`%sis JWT with %s ("%s"): `, a.message, title, name),
		path:    a.path,
		layout:  time.RFC3339,
		value:   timeFromFloat(value),
	}
}

func (a *JWTAssertion) failOnMissingField(title, name, expected string, msgAndArgs ...interface{}) *JWTAssertion {
	a.t.Helper()

	if expected != "" {
		expected = " " + expected
	}

	a.fail(
		fmt.Sprintf(
			`is JWT with %s ("%s")%s: field does not exist`,
			title,
			name,
			expected,
		),
		msgAndArgs...,
	)

	return a
}

func (a *JWTAssertion) failOnUnexpectedType(title, name, expected, expectedType string, msgAndArgs ...interface{}) *JWTAssertion {
	a.t.Helper()

	a.fail(
		fmt.Sprintf(
			`is JWT with %s ("%s") %s: %s`,
			title,
			name,
			expected,
			expectedType,
		),
		msgAndArgs...,
	)

	return a
}

func (a *JWTAssertion) failOnNotEqual(title, name, expected, actual string, msgAndArgs ...interface{}) *JWTAssertion {
	a.t.Helper()

	a.fail(
		fmt.Sprintf(
			`is JWT with %s ("%s") %s, actual is %s`,
			title,
			name,
			expected,
			actual,
		),
		msgAndArgs...,
	)

	return a
}

func (a *JWTAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	assert.Fail(a.t, a.message+message, msgAndArgs...)
}

func castToStrings(raw interface{}) ([]string, bool) {
	var actual []string
	switch v := raw.(type) {
	case string:
		actual = append(actual, v)
	case []string:
		actual = v
	case []interface{}:
		for _, vv := range v {
			vs, ok := vv.(string)
			if !ok {
				return nil, false
			}
			actual = append(actual, vs)
		}
	default:
		return nil, false
	}

	return actual, true
}

func timeFromFloat(value float64) time.Time {
	round, frac := math.Modf(value)

	return time.Unix(int64(round), int64(frac*1e9))
}
