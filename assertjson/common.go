package assertjson

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

// Exists asserts that the JSON node exists.
func (node *AssertNode) Exists(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.err != nil {
		assert.Fail(
			node.t,
			fmt.Sprintf(`failed asserting that json node "%s" exists`, node.pathPrefix+node.path),
			msgAndArgs...,
		)
	}
}

// DoesNotExist asserts that the JSON node does not exist.
func (node *AssertNode) DoesNotExist(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.err == nil {
		assert.Fail(
			node.t,
			fmt.Sprintf(`failed asserting that json node "%s" does not exist`, node.pathPrefix+node.path),
			msgAndArgs...,
		)
	}
}

// IsTrue asserts that the JSON node equals to the boolean with `true` value.
func (node *AssertNode) IsTrue(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.True(node.t, node.value.(bool), msgAndArgs...)
	}
}

// IsFalse asserts that the JSON node equals to the boolean with `false` value.
func (node *AssertNode) IsFalse(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.False(node.t, node.value.(bool), msgAndArgs...)
	}
}

// IsNull asserts that the JSON node equals to `null` value.
func (node *AssertNode) IsNull(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.Nil(node.t, node.value, msgAndArgs...)
	}
}

// IsNotNull asserts that the JSON node not equals to `null` value.
func (node *AssertNode) IsNotNull(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.NotNil(node.t, node.value, msgAndArgs...)
	}
}
