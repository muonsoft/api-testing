package assertjson

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/muonsoft/api-testing/internal/is"
	"github.com/stretchr/testify/assert"
)

// IsEmail asserts that the JSON node has a string value with email.
// Validation is based on simplified pattern. It allows all values
// with an "@" symbol in, and a "." in the second host part of the email address.
func (node *AssertNode) IsEmail(msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().WithEmail(msgAndArgs...)
}

// WithEmail asserts that the JSON node has a string value with email.
// Validation is based on simplified pattern. It allows all values
// with an "@" symbol in, and a "." in the second host part of the email address.
func (a *StringAssertion) WithEmail(msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if !is.Email(a.value) {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is email, actual is "%s"`,
				a.path,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// IsHTML5Email asserts that the JSON node has a string value with email. Validation is based on
// pattern for HTML5 (see https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address).
func (node *AssertNode) IsHTML5Email(msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().WithHTML5Email(msgAndArgs...)
}

// WithHTML5Email asserts that the JSON node has a string value with email. Validation is based on
// pattern for HTML5 (see https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address).
func (a *StringAssertion) WithHTML5Email(msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if !is.HTML5Email(a.value) {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is email (HTML5 format), actual is "%s"`,
				a.path,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// IsURL asserts that the JSON node has a string value with URL.
func (node *AssertNode) IsURL(msgAndArgs ...interface{}) *URLAssertion {
	node.t.Helper()
	return node.IsString().WithURL(msgAndArgs...)
}

// WithURL asserts that the JSON node has a string value with URL.
func (a *StringAssertion) WithURL(msgAndArgs ...interface{}) *URLAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	u, err := url.Parse(a.value)
	if err == nil && is.URL(a.value) {
		return &URLAssertion{t: a.t, path: a.path, url: u}
	}

	assert.Fail(
		a.t,
		fmt.Sprintf(
			`failed asserting that JSON node "%s" is URL, actual is "%s"`,
			a.path,
			a.value,
		),
		msgAndArgs...,
	)

	return nil
}

// URLAssertion is used to build a chain of assertions for the URL node.
type URLAssertion struct {
	t    TestingT
	path string
	url  *url.URL
}

// WithSchemas additionally asserts than URL schema contains one of the given values.
func (a *URLAssertion) WithSchemas(schemas ...string) *URLAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	for _, schema := range schemas {
		if a.url.Scheme == schema {
			return nil
		}
	}

	assert.Fail(
		a.t,
		fmt.Sprintf(
			`failed asserting that JSON node "%s" is URL with schemas %s, actual is "%s"`,
			a.path,
			strings.Join(quoteAll(schemas), ", "),
			a.url.Scheme,
		),
	)

	return a
}

// WithHosts additionally asserts than URL host contains one of the given values.
func (a *URLAssertion) WithHosts(hosts ...string) *URLAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	for _, host := range hosts {
		if a.url.Host == host {
			return nil
		}
	}

	assert.Fail(
		a.t,
		fmt.Sprintf(
			`failed asserting that JSON node "%s" is URL with hosts %s, actual is "%s"`,
			a.path,
			strings.Join(quoteAll(hosts), ", "),
			a.url.Host,
		),
	)

	return a
}

// That asserts that the JSON node has a URL value that is satisfied by callback function.
func (a *URLAssertion) That(f func(u *url.URL) error, msgAndArgs ...interface{}) *URLAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if err := f(a.url); err != nil {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting JSON node "%s": %s`,
				a.path,
				err.Error(),
			),
			msgAndArgs...,
		)
	}

	return a
}
