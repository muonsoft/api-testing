package assertjson

import (
	"github.com/stretchr/testify/assert"
	"math"
)

func (node *AssertNode) IsInteger(msgAndArgs ...interface{}) {
	if node.exists() {
		float, ok := node.value.(float64)
		if !ok {
			assert.Failf(node.t, "value at path '%s' is not numeric", node.path, msgAndArgs...)
		}
		_, fractional := math.Modf(float)
		if fractional != 0 {
			assert.Failf(node.t, "value at path '%s' is float, not integer", node.path, msgAndArgs...)
		}
	}
}

func (node *AssertNode) IsFloat(msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
	}
}

func (node *AssertNode) EqualToTheInteger(expectedValue int, msgAndArgs ...interface{}) {
	if node.exists() {
		float, ok := node.value.(float64)
		if !ok {
			assert.Failf(node.t, "value at path '%s' is not numeric", node.path, msgAndArgs...)
		}
		integer, fractional := math.Modf(float)
		if fractional != 0 {
			assert.Failf(node.t, "value at path '%s' is float, not integer", node.path, msgAndArgs...)
		}
		assert.Equal(node.t, expectedValue, int(integer), msgAndArgs...)
	}
}

func (node *AssertNode) EqualToTheFloat(expectedValue float64, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.Equal(node.t, expectedValue, node.value, msgAndArgs...)
	}
}

func (node *AssertNode) IsNumberGreaterThan(value float64, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.Greater(node.t, node.value, value, msgAndArgs...)
	}
}

func (node *AssertNode) IsNumberGreaterThanOrEqual(value float64, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.GreaterOrEqual(node.t, node.value, value, msgAndArgs...)
	}
}

func (node *AssertNode) IsNumberLessThan(value float64, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.Less(node.t, node.value, value, msgAndArgs...)
	}
}

func (node *AssertNode) IsNumberLessThanOrEqual(value float64, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.LessOrEqual(node.t, node.value, value, msgAndArgs...)
	}
}

func (node *AssertNode) IsNumberInRange(min float64, max float64, msgAndArgs ...interface{}) {
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.GreaterOrEqual(node.t, node.value, min, msgAndArgs...)
		assert.LessOrEqual(node.t, node.value, max, msgAndArgs...)
	}
}
