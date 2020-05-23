package assertjson

import "github.com/stretchr/testify/assert"

// Asserts that the JSON node is an array with given elements count.
func (node *AssertNode) IsArrayWithElementsCount(count int, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, []interface{}{}, node.value, msgAndArgs...)
		assert.Len(node.t, node.value, count, msgAndArgs...)
	}
}
