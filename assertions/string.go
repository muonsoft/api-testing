package assertions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

// StringAssertion is used to build a chain of assertions for the string node.
type StringAssertion struct {
	t             TestingT
	messagePrefix string
	value         string
}

func NewStringAssertion(t TestingT, messagePrefix string, value string) *StringAssertion {
	return &StringAssertion{t: t, messagePrefix: messagePrefix, value: value}
}

// IsEmpty asserts that string value equals to empty string.
func (a *StringAssertion) IsEmpty(msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value != "" {
		a.fail(
			fmt.Sprintf(`is empty string, actual is "%s"`, a.value),
			msgAndArgs...,
		)
	}

	return a
}

// IsNotEmpty asserts that string value not equals to empty string.
func (a *StringAssertion) IsNotEmpty(msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value == "" {
		a.fail(`is not empty string`, msgAndArgs...)
	}

	return a
}

// EqualTo asserts that string value equals to the given value.
func (a *StringAssertion) EqualTo(expectedValue string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value != expectedValue {
		a.fail(
			fmt.Sprintf(`equal to "%s", actual is "%s"`, expectedValue, a.value),
			msgAndArgs...,
		)
	}

	return a
}

// NotEqualTo asserts that string value not equals to the given value.
func (a *StringAssertion) NotEqualTo(expectedValue string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if a.value == expectedValue {
		a.fail(
			fmt.Sprintf(`not equal to "%s", actual is "%s"`, expectedValue, a.value),
			msgAndArgs...,
		)
	}

	return a
}

// EqualToOneOf asserts that string value equals to one of the given values.
func (a *StringAssertion) EqualToOneOf(expectedValues ...string) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if !isOneOf(a.value, expectedValues) {
		a.fail(
			fmt.Sprintf(`equal to one of values (%s), actual is "%s"`, formatStrings(expectedValues), a.value),
		)
	}

	return a
}

// Matches asserts that string value that matches the regular expression.
func (a *StringAssertion) Matches(regexp interface{}, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if !matchRegexp(regexp, a.value) {
		a.fail(
			fmt.Sprintf(`matches "%v", actual is "%s"`, regexp, a.value),
			msgAndArgs...,
		)
	}

	return a
}

// NotMatches asserts that string value that does not match the regular expression.
func (a *StringAssertion) NotMatches(regexp interface{}, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if matchRegexp(regexp, a.value) {
		a.fail(
			fmt.Sprintf(`not matches "%v", actual is "%s"`, regexp, a.value),
			msgAndArgs...,
		)
	}

	return a
}

// Contains asserts that string value that contains a string.
func (a *StringAssertion) Contains(value string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if !strings.Contains(a.value, value) {
		a.fail(
			fmt.Sprintf(`contains "%s", actual is "%s"`, value, a.value),
			msgAndArgs...,
		)
	}

	return a
}

// NotContains asserts that string value that does not contain a string.
func (a *StringAssertion) NotContains(value string, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()
	if strings.Contains(a.value, value) {
		a.fail(
			fmt.Sprintf(`not contains "%s", actual is "%s"`, value, a.value),
			msgAndArgs...,
		)
	}

	return a
}

// WithLength asserts that string value with length equal to the given value.
func (a *StringAssertion) WithLength(length int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	actual := utf8.RuneCountInString(a.value)
	if actual != length {
		a.fail(
			fmt.Sprintf(`is string with length is %d, actual is %d`, length, actual),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThan asserts that string value with length greater than the value.
func (a *StringAssertion) WithLengthGreaterThan(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length <= expected {
		a.fail(
			fmt.Sprintf(`is string with length greater than %d, actual is %d`, expected, length),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThanOrEqual asserts that string value
// with length greater than or equal to the value.
func (a *StringAssertion) WithLengthGreaterThanOrEqual(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length < expected {
		a.fail(
			fmt.Sprintf(`is string with length greater than or equal to %d, actual is %d`, expected, length),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThan asserts that string value with length less than the value.
func (a *StringAssertion) WithLengthLessThan(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length >= expected {
		a.fail(
			fmt.Sprintf(`is string with length less than %d, actual is %d`, expected, length),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThanOrEqual asserts that string value
// with length less than or equal to the value.
func (a *StringAssertion) WithLengthLessThanOrEqual(expected int, msgAndArgs ...interface{}) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	length := utf8.RuneCountInString(a.value)
	if length > expected {
		a.fail(
			fmt.Sprintf(`is string with length less than or equal to %d, actual is %d`, expected, length),
			msgAndArgs...,
		)
	}

	return a
}

func (a *StringAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	assert.Fail(a.t, a.messagePrefix+message, msgAndArgs...)
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
