package apitest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/muonsoft/api-testing/assertions"
	"github.com/stretchr/testify/assert"
)

// AssertCookie asserts cookie.
func AssertCookie(t TestingT, cookie *http.Cookie) *CookieAssertion {
	t.Helper()

	return &CookieAssertion{t: t, cookie: cookie}
}

// CookieAssertion is used to build assertions for cookies.
type CookieAssertion struct {
	t      TestingT
	cookie *http.Cookie
}

// WithValue asserts cookie value with fluent string assertions.
func (a *CookieAssertion) WithValue() *assertions.StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return assertions.NewStringAssertion(
		a.t,
		fmt.Sprintf("failed asserting that cookie %q value ", a.cookie.Name),
		a.cookie.Value,
	)
}

// WithPath asserts cookie path with fluent string assertions.
func (a *CookieAssertion) WithPath() *assertions.StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return assertions.NewStringAssertion(
		a.t,
		fmt.Sprintf("failed asserting that cookie %q path ", a.cookie.Name),
		a.cookie.Path,
	)
}

// WithDomain asserts cookie domain with fluent string assertions.
func (a *CookieAssertion) WithDomain() *assertions.StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return assertions.NewStringAssertion(
		a.t,
		fmt.Sprintf("failed asserting that cookie %q domain ", a.cookie.Name),
		a.cookie.Domain,
	)
}

// Expires asserts cookie expiration with fluent time assertions.
func (a *CookieAssertion) Expires() *assertions.TimeAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	return assertions.NewTimeAssertion(
		a.t,
		fmt.Sprintf("failed asserting that cookie %q expires ", a.cookie.Name),
		a.cookie.Expires,
		time.RFC3339,
	)
}

// IsSecure asserts that cookie is secure.
func (a *CookieAssertion) IsSecure() *CookieAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.cookie.Secure {
		assert.Fail(a.t, fmt.Sprintf("failed asserting that cookie %q is secure", a.cookie.Name))
	}

	return a
}

// IsNotSecure asserts that cookie is not secure.
func (a *CookieAssertion) IsNotSecure() *CookieAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.cookie.Secure {
		assert.Fail(a.t, fmt.Sprintf("failed asserting that cookie %q is not secure", a.cookie.Name))
	}

	return a
}

// IsHTTPOnly asserts that cookie is HTTP only.
func (a *CookieAssertion) IsHTTPOnly() *CookieAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !a.cookie.HttpOnly {
		assert.Fail(a.t, fmt.Sprintf("failed asserting that cookie %q is HTTP only", a.cookie.Name))
	}

	return a
}

// IsNotHTTPOnly asserts that cookie is not HTTP only.
func (a *CookieAssertion) IsNotHTTPOnly() *CookieAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.cookie.HttpOnly {
		assert.Fail(a.t, fmt.Sprintf("failed asserting that cookie %q is not HTTP only", a.cookie.Name))
	}

	return a
}

// WithSameSite asserts that cookie has same site value.
func (a *CookieAssertion) WithSameSite(wantSameSite http.SameSite) *CookieAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.cookie.SameSite != wantSameSite {
		assert.Fail(a.t, fmt.Sprintf(
			"failed asserting that cookie %q has same site %q, actual is %q",
			a.cookie.Name,
			sameSiteToString(wantSameSite),
			sameSiteToString(a.cookie.SameSite),
		))
	}

	return a
}

func sameSiteToString(site http.SameSite) string {
	switch site {
	case http.SameSiteDefaultMode:
		return "DefaultMode"
	case http.SameSiteLaxMode:
		return "LaxMode"
	case http.SameSiteStrictMode:
		return "StrictMode"
	case http.SameSiteNoneMode:
		return "NoneMode"
	}

	return ""
}
