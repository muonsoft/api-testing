package assertjson

import (
	"fmt"
	"math"

	"github.com/stretchr/testify/assert"
)

// IsInteger asserts that the JSON node has an integer value.
// It returns IntegerAssertion to execute a chain of assertions for the node value.
func (node *AssertNode) IsInteger(msgAndArgs ...interface{}) *IntegerAssertion {
	node.t.Helper()
	if node.exists() {
		float, ok := node.value.(float64)
		if !ok {
			assert.Fail(
				node.t,
				fmt.Sprintf(`value at path "%s" is not numeric`, node.Path()),
				msgAndArgs...,
			)
			return nil
		}
		_, fractional := math.Modf(float)
		if fractional != 0 {
			assert.Fail(
				node.t,
				fmt.Sprintf(`value at path "%s" is float, not integer`, node.Path()),
				msgAndArgs...,
			)
			return nil
		}
		return &IntegerAssertion{t: node.t, path: node.Path(), value: int(float)}
	}

	return nil
}

// IsFloat asserts that the JSON node has a float value.
// It returns NumberAssertion to execute a chain of assertions for the node value.
func (node *AssertNode) IsFloat(msgAndArgs ...interface{}) *NumberAssertion {
	node.t.Helper()
	return node.IsNumber(msgAndArgs...)
}

// IsNumber asserts that the JSON node has a float value.
// It returns NumberAssertion to execute a chain of assertions for the node value.
func (node *AssertNode) IsNumber(msgAndArgs ...interface{}) *NumberAssertion {
	node.t.Helper()
	if node.exists() {
		if f, ok := node.value.(float64); ok {
			return &NumberAssertion{t: node.t, path: node.Path(), value: f}
		}
		assert.Fail(
			node.t,
			fmt.Sprintf(`value at path "%s" is not a number`, node.Path()),
			msgAndArgs...,
		)
	}

	return nil
}

// EqualToTheInteger asserts that the JSON node has an integer value equals to the given value.
// Deprecated: use IsInteger() instead.
func (node *AssertNode) EqualToTheInteger(expectedValue int, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsInteger().EqualTo(expectedValue, msgAndArgs...)
}

// EqualToTheFloat asserts that the JSON node has a float value equals to the given value.
// Deprecated: use IsNumber() instead.
func (node *AssertNode) EqualToTheFloat(expectedValue float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsNumber().EqualTo(expectedValue, msgAndArgs...)
}

// IsNumberGreaterThan asserts that the JSON node has a number greater than the given value.
// Deprecated: use IsNumber().GreaterThan() instead.
func (node *AssertNode) IsNumberGreaterThan(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsNumber().GreaterThan(value, msgAndArgs...)
}

// IsNumberGreaterThanOrEqual asserts that the JSON node has a number greater than or equal to the given value.
// Deprecated: use IsNumber().GreaterThanOrEqual() instead.
func (node *AssertNode) IsNumberGreaterThanOrEqual(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsNumber().GreaterThanOrEqual(value, msgAndArgs...)
}

// IsNumberLessThan asserts that the JSON node has a number less than the given value.
// Deprecated: use IsNumber().LessThan() instead.
func (node *AssertNode) IsNumberLessThan(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsNumber().LessThan(value, msgAndArgs...)
}

// IsNumberLessThanOrEqual asserts that the JSON node has a number less than or equal to the given value.
// Deprecated: use IsNumber().LessThanOrEqual() instead.
func (node *AssertNode) IsNumberLessThanOrEqual(value float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsNumber().LessThanOrEqual(value, msgAndArgs...)
}

// IsNumberInRange asserts that the JSON node has a number with value in the given range.
// Deprecated: use IsNumber().GreaterThanOrEqual().LessThanOrEqual() instead.
func (node *AssertNode) IsNumberInRange(min, max float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsNumber().GreaterThanOrEqual(min, msgAndArgs...)
	node.IsNumber().LessThanOrEqual(max, msgAndArgs...)
}

// NumberAssertion is used to build a chain of assertions for the numeric node.
type NumberAssertion struct {
	t     TestingT
	path  string
	value float64
}

// EqualTo asserts that the JSON node has a numeric value equals to the given value.
func (a *NumberAssertion) EqualTo(expected float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value != expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" equal to %f, actual is %f`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// EqualToWithDelta asserts that the JSON node has a numeric value equals to the given value with delta.
func (a *NumberAssertion) EqualToWithDelta(expected, delta float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	dt := a.value - expected
	if dt < -delta || dt > delta {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" equal to %f with delta %f, actual is %f`,
				a.path,
				expected,
				delta,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// GreaterThan asserts that the JSON node has a numeric value greater than the given value.
func (a *NumberAssertion) GreaterThan(expected float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value <= expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" greater than %f, actual is %f`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// GreaterThanOrEqual asserts that the JSON node has a numeric value greater than or equal to the given value.
func (a *NumberAssertion) GreaterThanOrEqual(expected float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value < expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" greater than or equal %f, actual is %f`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// LessThan asserts that the JSON node has a numeric value less than the given value.
func (a *NumberAssertion) LessThan(expected float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value >= expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" less than %f, actual is %f`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// LessThanOrEqual asserts that the JSON node has a numeric value less than or equal to the given value.
func (a *NumberAssertion) LessThanOrEqual(expected float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value > expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" less than or equal %f, actual is %f`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// IntegerAssertion is used to build a chain of assertions for the integer node.
type IntegerAssertion struct {
	t     TestingT
	path  string
	value int
}

// EqualTo asserts that the JSON node has an integer value equals to the given value.
func (a *IntegerAssertion) EqualTo(expected int, msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value != expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" equal to %d, actual is %d`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// GreaterThan asserts that the JSON node has an integer value greater than the given value.
func (a *IntegerAssertion) GreaterThan(expected int, msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value <= expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" greater than %d, actual is %d`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// GreaterThanOrEqual asserts that the JSON node has an integer value greater than or equal to the given value.
func (a *IntegerAssertion) GreaterThanOrEqual(expected int, msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value < expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" greater than or equal %d, actual is %d`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// LessThan asserts that the JSON node has an integer value less than the given value.
func (a *IntegerAssertion) LessThan(expected int, msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value >= expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" less than %d, actual is %d`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}

// LessThanOrEqual asserts that the JSON node has an integer value less than or equal to the given value.
func (a *IntegerAssertion) LessThanOrEqual(expected int, msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value > expected {
		assert.Fail(
			a.t,
			fmt.Sprintf(
				`failed asserting that JSON node "%s" less than or equal %d, actual is %d`,
				a.path,
				expected,
				a.value,
			),
			msgAndArgs...,
		)
	}

	return nil
}
