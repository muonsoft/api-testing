package apitest_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/muonsoft/api-testing/apitest"
	"github.com/muonsoft/api-testing/internal/mock"
)

func TestCookieAssertion(t *testing.T) {
	tests := []struct {
		name         string
		cookie       *http.Cookie
		assert       func(cookie *apitest.CookieAssertion)
		wantMessages []string
	}{
		{
			name: "WithValue passed",
			cookie: &http.Cookie{
				Name:  "name",
				Value: "value",
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithValue().EqualTo("value")
			},
		},
		{
			name: "WithValue failed",
			cookie: &http.Cookie{
				Name:  "name",
				Value: "invalid",
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithValue().EqualTo("value")
			},
			wantMessages: []string{
				`failed asserting that cookie "name" value equal to "value", actual is "invalid"`,
			},
		},
		{
			name: "WithPath passed",
			cookie: &http.Cookie{
				Name: "name",
				Path: "/path",
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithPath().EqualTo("/path")
			},
		},
		{
			name: "WithPath failed",
			cookie: &http.Cookie{
				Name: "name",
				Path: "/invalid",
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithPath().EqualTo("/path")
			},
			wantMessages: []string{
				`failed asserting that cookie "name" path equal to "/path", actual is "/invalid"`,
			},
		},
		{
			name: "WithDomain passed",
			cookie: &http.Cookie{
				Name:   "name",
				Domain: "example.com",
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithDomain().EqualTo("example.com")
			},
		},
		{
			name: "WithDomain failed",
			cookie: &http.Cookie{
				Name:   "name",
				Domain: "invalid.com",
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithDomain().EqualTo("example.com")
			},
			wantMessages: []string{
				`failed asserting that cookie "name" domain equal to "example.com", actual is "invalid.com"`,
			},
		},
		{
			name: "Expires passed",
			cookie: &http.Cookie{
				Name:    "name",
				Expires: time.Now().Add(time.Hour),
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.Expires().After(time.Now())
			},
		},
		{
			name: "Expires failed",
			cookie: &http.Cookie{
				Name:    "name",
				Expires: time.Now().Add(time.Hour),
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.Expires().After(time.Now().Add(2 * time.Hour))
			},
			wantMessages: []string{
				`failed asserting that cookie "name" expires after`,
			},
		},
		{
			name: "IsSecure passed",
			cookie: &http.Cookie{
				Name:   "name",
				Secure: true,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsSecure()
			},
		},
		{
			name: "IsSecure failed",
			cookie: &http.Cookie{
				Name:   "name",
				Secure: false,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsSecure()
			},
			wantMessages: []string{
				`failed asserting that cookie "name" is secure`,
			},
		},
		{
			name: "IsNotSecure passed",
			cookie: &http.Cookie{
				Name:   "name",
				Secure: false,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsNotSecure()
			},
		},
		{
			name: "IsNotSecure failed",
			cookie: &http.Cookie{
				Name:   "name",
				Secure: true,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsNotSecure()
			},
			wantMessages: []string{
				`failed asserting that cookie "name" is not secure`,
			},
		},
		{
			name: "IsHTTPOnly passed",
			cookie: &http.Cookie{
				Name:     "name",
				HttpOnly: true,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsHTTPOnly()
			},
		},
		{
			name: "IsHTTPOnly failed",
			cookie: &http.Cookie{
				Name:     "name",
				HttpOnly: false,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsHTTPOnly()
			},
			wantMessages: []string{
				`failed asserting that cookie "name" is HTTP only`,
			},
		},
		{
			name: "IsNotHTTPOnly passed",
			cookie: &http.Cookie{
				Name:     "name",
				HttpOnly: false,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsNotHTTPOnly()
			},
		},
		{
			name: "IsNotHTTPOnly failed",
			cookie: &http.Cookie{
				Name:     "name",
				HttpOnly: true,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.IsNotHTTPOnly()
			},
			wantMessages: []string{
				`failed asserting that cookie "name" is not HTTP only`,
			},
		},
		{
			name: "WithSameSite passed",
			cookie: &http.Cookie{
				Name:     "name",
				SameSite: http.SameSiteStrictMode,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithSameSite(http.SameSiteStrictMode)
			},
		},
		{
			name: "WithSameSite failed",
			cookie: &http.Cookie{
				Name:     "name",
				SameSite: http.SameSiteDefaultMode,
			},
			assert: func(cookie *apitest.CookieAssertion) {
				cookie.WithSameSite(http.SameSiteStrictMode)
			},
			wantMessages: []string{
				`failed asserting that cookie "name" has same site "StrictMode", actual is "DefaultMode"`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tester := &mock.Tester{}
			cookie := apitest.AssertCookie(tester, test.cookie)

			test.assert(cookie)

			tester.AssertContains(t, test.wantMessages)
		})
	}
}
