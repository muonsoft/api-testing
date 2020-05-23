package assertxml

import "github.com/stretchr/testify/assert"

// Asserts that the XML node exists.
func (node *AssertNode) Exists(msgAndArgs ...interface{}) {
	if !node.found {
		assert.Failf(node.t, "failed asserting that xml node '%s' exists", node.path, msgAndArgs...)
	}
}

// Asserts that the XML node does not exist.
func (node *AssertNode) DoesNotExist(msgAndArgs ...interface{}) {
	if node.found {
		assert.Failf(node.t, "failed asserting that xml node '%s' does not exist", node.path, msgAndArgs...)
	}
}
