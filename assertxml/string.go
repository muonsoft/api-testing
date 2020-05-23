package assertxml

import "github.com/stretchr/testify/assert"

func (node *AssertNode) EqualToTheString(expectedValue string, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.Equal(node.t, expectedValue, node.value, msgAndArgs...)
	}
}
