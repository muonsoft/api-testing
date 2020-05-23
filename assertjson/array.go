package assertjson

import "github.com/stretchr/testify/assert"

func (node *AssertNode) IsArrayWithElementsCount(count int, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, []interface{}{}, node.value, msgAndArgs...)
		assert.Len(node.t, node.value, count, msgAndArgs...)
	}
}
