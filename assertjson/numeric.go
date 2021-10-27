package assertjson

import (
	"math"

	"github.com/stretchr/testify/assert"
)

// IsInteger asserts that the JSON node has an integer value.
func (node *AssertNode) IsInteger(msgAndArgs ...interface{}) {
	node.t.Helper()
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

// IsFloat asserts that the JSON node has a float value.
func (node *AssertNode) IsFloat(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
	}
}

// EqualToTheInteger asserts that the JSON node has an integer value equals to the given value.
func (node *AssertNode) EqualToTheInteger(expectedValue int, msgAndArgs ...interface{}) {
	node.t.Helper()
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

// EqualToTheFloat asserts that the JSON node has a float value equals to the given value.
func (node *AssertNode) EqualToTheFloat(expectedValue float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.Equal(node.t, expectedValue, node.value, msgAndArgs...)
	}
}

// IsNumberGreaterThan asserts that the JSON node has a number greater than the given value.
func (node *AssertNode) IsNumberGreaterThan(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.Greater(node.t, node.value, value, msgAndArgs...)
	}
}

// IsNumberGreaterThanOrEqual asserts that the JSON node has a number greater than or equal to the given value.
func (node *AssertNode) IsNumberGreaterThanOrEqual(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.GreaterOrEqual(node.t, node.value, value, msgAndArgs...)
	}
}

// IsNumberLessThan asserts that the JSON node has a number less than the given value.
func (node *AssertNode) IsNumberLessThan(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.Less(node.t, node.value, value, msgAndArgs...)
	}
}

// IsNumberLessThanOrEqual asserts that the JSON node has a number less than or equal to the given value.
func (node *AssertNode) IsNumberLessThanOrEqual(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.LessOrEqual(node.t, node.value, value, msgAndArgs...)
	}
}

// IsNumberInRange asserts that the JSON node has a number with value in the given range.
func (node *AssertNode) IsNumberInRange(min float64, max float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		assert.IsType(node.t, 0.0, node.value, msgAndArgs...)
		assert.GreaterOrEqual(node.t, node.value, min, msgAndArgs...)
		assert.LessOrEqual(node.t, node.value, max, msgAndArgs...)
	}
}
