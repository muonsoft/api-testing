package assertjson

import "github.com/stretchr/testify/assert"

// IsObjectWithPropertiesCount asserts that the JSON node is an object with given properties count.
func (node *AssertNode) IsObjectWithPropertiesCount(count int, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, map[string]interface{}{}, node.value, msgAndArgs...)
		assert.Len(node.t, node.value, count, msgAndArgs...)
	}
}
