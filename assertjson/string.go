package assertjson

import "github.com/stretchr/testify/assert"

func (node *AssertNode) EqualToTheString(expectedValue string, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.Equal(node.t, expectedValue, node.value, msgAndArgs...)
	}
}

func (node *AssertNode) Matches(regexp string, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.Regexp(node.t, regexp, node.value, msgAndArgs...)
	}
}

func (node *AssertNode) DoesNotMatch(regexp string, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.NotRegexp(node.t, regexp, node.value, msgAndArgs...)
	}
}

func (node *AssertNode) Contains(contain string, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.Contains(node.t, node.value, contain, msgAndArgs...)
	}
}

func (node *AssertNode) DoesNotContain(contain string, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.NotContains(node.t, node.value, contain, msgAndArgs...)
	}
}

func (node *AssertNode) IsStringWithLength(length int, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.Equal(node.t, len(node.value.(string)), length, msgAndArgs...)
	}
}

func (node *AssertNode) IsStringWithLengthInRange(min int, max int, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, "", node.value, msgAndArgs...)
		assert.GreaterOrEqual(node.t, len(node.value.(string)), min, msgAndArgs...)
		assert.LessOrEqual(node.t, len(node.value.(string)), max, msgAndArgs...)
	}
}
