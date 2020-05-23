package assertjson

import "github.com/stretchr/testify/assert"

func (node *AssertNode) Exists(msgAndArgs ...interface{}) {
	if node.err != nil {
		assert.Failf(node.t, "failed asserting that json node '%s' exists", node.path, msgAndArgs...)
	}
}

func (node *AssertNode) DoesNotExist(msgAndArgs ...interface{}) {
	if node.err == nil {
		assert.Failf(node.t, "failed asserting that json node '%s' does not exist", node.path, msgAndArgs...)
	}
}

func (node *AssertNode) IsTrue(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.True(node.t, node.value.(bool), msgAndArgs...)
	}
}

func (node *AssertNode) IsFalse(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.False(node.t, node.value.(bool), msgAndArgs...)
	}
}

func (node *AssertNode) IsNull(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.Nil(node.t, node.value, msgAndArgs...)
	}
}

func (node *AssertNode) IsNotNull(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.NotNil(node.t, node.value, msgAndArgs...)
	}
}
