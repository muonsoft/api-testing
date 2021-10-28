package assertjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// IsString asserts that the JSON node has a string value.
func (node *AssertNode) IsString(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
	}
}

// EqualToTheString asserts that the JSON node has a string value equals to the given value.
func (node *AssertNode) EqualToTheString(expectedValue string, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.Equal(node.t, expectedValue, node.value, msgAndArgs...)
	}
}

// Matches asserts that the JSON node has a string value that matches the regular expression.
func (node *AssertNode) Matches(regexp string, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.Regexp(node.t, regexp, node.value, msgAndArgs...)
	}
}

// DoesNotMatch asserts that the JSON node has a string value that does not match the regular expression.
func (node *AssertNode) DoesNotMatch(regexp string, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.NotRegexp(node.t, regexp, node.value, msgAndArgs...)
	}
}

// Contains asserts that the JSON node has a string value that contains a string.
func (node *AssertNode) Contains(contain string, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.Contains(node.t, node.value, contain, msgAndArgs...)
	}
}

// DoesNotContain asserts that the JSON node has a string value that does not contain a string.
func (node *AssertNode) DoesNotContain(contain string, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.NotContains(node.t, node.value, contain, msgAndArgs...)
	}
}

// IsStringWithLength asserts that the JSON node has a string value with length equal to the given value.
func (node *AssertNode) IsStringWithLength(length int, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.Equal(node.t, len(node.value.(string)), length, msgAndArgs...)
	}
}

// IsStringWithLengthInRange asserts that the JSON node has a string value with length in the given range.
func (node *AssertNode) IsStringWithLengthInRange(min int, max int, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.GreaterOrEqual(node.t, len(node.value.(string)), min, msgAndArgs...)
		assert.LessOrEqual(node.t, len(node.value.(string)), max, msgAndArgs...)
	}
}

// AssertString asserts that the JSON node has a string value and it is satisfied by the user function assertFunc.
func (node *AssertNode) AssertString(assertFunc func(t testing.TB, value string)) {
	node.t.Helper()
	if node.exists() && assert.IsType(node.t, "", node.value) {
		assertFunc(node.t, node.value.(string))
	}
}
