package assertxml

import "github.com/stretchr/testify/assert"

// Asserts that the XML node has a string value equals to the given value.
func (node *AssertNode) EqualToTheString(expectedValue string, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.Equal(node.t, expectedValue, node.value, msgAndArgs...)
	}
}
