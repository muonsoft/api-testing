package assertjson

import "github.com/stretchr/testify/assert"

// Asserts that the JSON node exists.
func (node *AssertNode) Exists(msgAndArgs ...interface{}) {
	if node.err != nil {
		assert.Failf(node.t, "failed asserting that json node '%s' exists", node.path, msgAndArgs...)
	}
}

// Asserts that the JSON node does not exist.
func (node *AssertNode) DoesNotExist(msgAndArgs ...interface{}) {
	if node.err == nil {
		assert.Failf(node.t, "failed asserting that json node '%s' does not exist", node.path, msgAndArgs...)
	}
}

// Asserts that the JSON node equals to the boolean with `true` value.
func (node *AssertNode) IsTrue(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.True(node.t, node.value.(bool), msgAndArgs...)
	}
}

// Asserts that the JSON node equals to the boolean with `false` value.
func (node *AssertNode) IsFalse(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.False(node.t, node.value.(bool), msgAndArgs...)
	}
}

// Asserts that the JSON node equals to `null` value.
func (node *AssertNode) IsNull(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.Nil(node.t, node.value, msgAndArgs...)
	}
}

// Asserts that the JSON node not equals to `null` value.
func (node *AssertNode) IsNotNull(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.NotNil(node.t, node.value, msgAndArgs...)
	}
}
