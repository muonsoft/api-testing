package assertjson

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

// IsString asserts that the JSON node has a string value.
// It returns StringAssertion to execute a chain of assertions for the node value.
func (node *AssertNode) IsString(msgAndArgs ...interface{}) *StringAssertion {
	node.t.Helper()
	if node.exists() {
		if s, ok := node.value.(string); ok {
			return &StringAssertion{t: node.t, path: node.Path(), value: s}
		}
		assert.Fail(
			node.t,
			fmt.Sprintf(`failed asserting that JSON node "%s" is string`, node.Path()),
			msgAndArgs...,
		)
	}

	return nil
}

// EqualToTheString asserts that the JSON node has a string value equals to the given value.
// Deprecated: use IsString().EqualTo() instead.
func (node *AssertNode) EqualToTheString(expectedValue string, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().EqualTo(expectedValue, msgAndArgs...)
}

// Matches asserts that the JSON node has a string value that matches the regular expression.
func (node *AssertNode) Matches(regexp interface{}, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().Matches(regexp, msgAndArgs...)
}

// DoesNotMatch asserts that the JSON node has a string value that does not match the regular expression.
func (node *AssertNode) DoesNotMatch(regexp interface{}, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().NotMatches(regexp, msgAndArgs...)
}

// Contains asserts that the JSON node has a string value that contains a string.
func (node *AssertNode) Contains(contain string, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().Contains(contain, msgAndArgs...)
}

// DoesNotContain asserts that the JSON node has a string value that does not contain a string.
func (node *AssertNode) DoesNotContain(contain string, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().NotContains(contain, msgAndArgs...)
}

// IsStringWithLength asserts that the JSON node has a string value with length equal to the given value.
// Deprecated: use IsString().WithLength() instead.
func (node *AssertNode) IsStringWithLength(length int, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().WithLength(length, msgAndArgs...)
}

// IsStringWithLengthInRange asserts that the JSON node has a string value with length in a given range.
// Deprecated: use IsString().WithLengthGreaterThanOrEqual().WithLengthLessThanOrEqual() instead.
func (node *AssertNode) IsStringWithLengthInRange(min int, max int, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsString().WithLengthGreaterThanOrEqual(min, msgAndArgs...).WithLengthLessThanOrEqual(max, msgAndArgs...)
}

// AssertString asserts that the JSON node has a string value and it is satisfied by the user function assertFunc.
// Deprecated: use IsString().Assert() instead.
func (node *AssertNode) AssertString(assertFunc func(t testing.TB, value string)) {
	node.t.Helper()
	node.IsString().Assert(assertFunc)
}

// StringAssertion is used to build a chain of assertions for the string node.
type StringAssertion struct {
	t     TestingT
	path  string
	value string
}

// EqualTo asserts that the JSON node has a string value equals to the given value.
func (a *StringAssertion) EqualTo(expectedValue string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value != expectedValue {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" equal to "%s", actual is "%s"`,
				a.path,
				expectedValue,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// NotEqualTo asserts that the JSON node has a string value not equals to the given value.
func (a *StringAssertion) NotEqualTo(expectedValue string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value == expectedValue {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" not equal to "%s", actual is "%s"`,
				a.path,
				expectedValue,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// EqualToOneOf asserts that the JSON node has a string value equals to one of the given values.
func (a *StringAssertion) EqualToOneOf(expectedValues ...string) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !isOneOf(a.value, expectedValues) {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" equal to one of values (%s), actual is "%s"`,
				a.path,
				formatStrings(expectedValues),
				a.value,
			),
		)
	}

	return a
}

// Matches asserts that the JSON node has a string value that matches the regular expression.
func (a *StringAssertion) Matches(regexp interface{}, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if !matchRegexp(regexp, a.value) {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" matches "%v", actual is "%s"`,
				a.path,
				regexp,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// NotMatches asserts that the JSON node has a string value that does not match the regular expression.
func (a *StringAssertion) NotMatches(regexp interface{}, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if matchRegexp(regexp, a.value) {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" not matches "%v", actual is "%s"`,
				a.path,
				regexp,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// Contains asserts that the JSON node has a string value that contains a string.
func (a *StringAssertion) Contains(value string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if !strings.Contains(a.value, value) {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" contains "%s", actual is "%s"`,
				a.path,
				value,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// NotContains asserts that the JSON node has a string value that does not contain a string.
func (a *StringAssertion) NotContains(value string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if strings.Contains(a.value, value) {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" not contains "%s", actual is "%s"`,
				a.path,
				value,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLength asserts that the JSON node has a string value with length equal to the given value.
func (a *StringAssertion) WithLength(length int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	actual := utf8.RuneCountInString(a.value)
	if actual != length {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is string with length is %d, actual is %d`,
				a.path,
				length,
				actual,
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThan asserts that the JSON node has a string value with length greater than the value.
func (a *StringAssertion) WithLengthGreaterThan(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length <= expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is string with length greater than %d, actual is %d`,
				a.path,
				expected,
				length,
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThanOrEqual asserts that the JSON node has a string value
// with length greater than or equal to the value.
func (a *StringAssertion) WithLengthGreaterThanOrEqual(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length < expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is string with length greater than or equal to %d, actual is %d`,
				a.path,
				expected,
				length,
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThan asserts that the JSON node has a string value with length less than the value.
func (a *StringAssertion) WithLengthLessThan(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length >= expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is string with length less than %d, actual is %d`,
				a.path,
				expected,
				length,
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThanOrEqual asserts that the JSON node has a string value
// with length less than or equal to the value.
func (a *StringAssertion) WithLengthLessThanOrEqual(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length > expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is string with length less than or equal to %d, actual is %d`,
				a.path,
				expected,
				length,
			),
			msgAndArgs...,
		)
	}

	return a
}

// That asserts that the JSON node has a string value that is satisfied by callback function.
func (a *StringAssertion) That(f func(s string) error, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if err := f(a.value); err != nil {
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

// Assert asserts that the string node has that is satisfied by the user function assertFunc.
func (a *StringAssertion) Assert(assertFunc func(t testing.TB, value string)) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	//nolint: forcetypeassert
	assertFunc(a.t.(testing.TB), a.value)

	return a
}

// matchRegexp return true if a specified regexp matches a string.
func matchRegexp(rx interface{}, s string) bool {
	var r *regexp.Regexp
	if rr, ok := rx.(*regexp.Regexp); ok {
		r = rr
	} else {
		r = regexp.MustCompile(fmt.Sprint(rx))
	}

	return r.FindStringIndex(s) != nil
}

func isOneOf(s string, ss []string) bool {
	for _, s2 := range ss {
		if s == s2 {
			return true
		}
	}

	return false
}

func formatStrings(ss []string) string {
	var b strings.Builder

	for i, s := range ss {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Quote(s))
	}

	return b.String()
}
