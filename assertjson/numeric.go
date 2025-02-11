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
			node.fail(
				fmt.Sprintf(`value at path "%s" is not numeric`, node.path.String()),
				msgAndArgs...,
			)
			return nil
		}
		_, fractional := math.Modf(float)
		if fractional != 0 {
			node.fail(
				fmt.Sprintf(`value at path "%s" is float, not integer`, node.path.String()),
				msgAndArgs...,
			)
			return nil
		}
		return &IntegerAssertion{
			t:       node.t,
			message: fmt.Sprintf(`%sfailed asserting that JSON node "%s": `, node.message, node.path.String()),
			path:    node.path.String(),
			value:   int(float),
		}
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
			return &NumberAssertion{
				t:       node.t,
				message: fmt.Sprintf(`%sfailed asserting that JSON node "%s": `, node.message, node.path.String()),
				path:    node.path.String(),
				value:   f,
			}
		}
		node.fail(
			fmt.Sprintf(`value at path "%s" is not a number`, node.path.String()),
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
func (node *AssertNode) IsNumberInRange(vmin, vmax float64, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsNumber().GreaterThanOrEqual(vmin, msgAndArgs...)
	node.IsNumber().LessThanOrEqual(vmax, msgAndArgs...)
}

// NumberAssertion is used to build a chain of assertions for the numeric node.
type NumberAssertion struct {
	t       TestingT
	message string
	path    string
	value   float64
}

// IsZero asserts that the JSON node has a numeric value equals to zero.
func (a *NumberAssertion) IsZero(msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value != 0 {
		a.fail(
			fmt.Sprintf(`is zero, actual is %f`, a.value),
			msgAndArgs...,
		)
	}

	return nil
}

// IsNotZero asserts that the JSON node has a numeric value not equals to zero.
func (a *NumberAssertion) IsNotZero(msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value == 0 {
		a.fail(`is not zero`, msgAndArgs...)
	}

	return nil
}

// EqualTo asserts that the JSON node has a numeric value equals to the given value.
func (a *NumberAssertion) EqualTo(expected float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value != expected {
		a.fail(
			fmt.Sprintf(`equal to %f, actual is %f`, expected, a.value),
			msgAndArgs...,
		)
	}

	return nil
}

// NotEqualTo asserts that the JSON node has a numeric value not equals to the given value.
func (a *NumberAssertion) NotEqualTo(expected float64, msgAndArgs ...interface{}) *NumberAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value == expected {
		a.fail(
			fmt.Sprintf(`not equal to %f, actual is %f`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`equal to %f with delta %f, actual is %f`, expected, delta, a.value),
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
		a.fail(
			fmt.Sprintf(`greater than %f, actual is %f`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`greater than or equal %f, actual is %f`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`less than %f, actual is %f`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`less than or equal %f, actual is %f`, expected, a.value),
			msgAndArgs...,
		)
	}

	return nil
}

func (a *NumberAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	assert.Fail(a.t, a.message+message, msgAndArgs...)
}

// IntegerAssertion is used to build a chain of assertions for the integer node.
type IntegerAssertion struct {
	t       TestingT
	message string
	path    string
	value   int
}

// IsZero asserts that the JSON node has an integer value equals to 0.
func (a *IntegerAssertion) IsZero(msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value != 0 {
		a.fail(
			fmt.Sprintf(`is zero, actual is %d`, a.value),
			msgAndArgs...,
		)
	}

	return nil
}

// IsNotZero asserts that the JSON node has an integer value not equals to 0.
func (a *IntegerAssertion) IsNotZero(msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value == 0 {
		a.fail(`is not zero`, msgAndArgs...)
	}

	return nil
}

// EqualTo asserts that the JSON node has an integer value equals to the given value.
func (a *IntegerAssertion) EqualTo(expected int, msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value != expected {
		a.fail(
			fmt.Sprintf(`equal to %d, actual is %d`, expected, a.value),
			msgAndArgs...,
		)
	}

	return nil
}

// NotEqualTo asserts that the JSON node has an integer value not equals to the given value.
func (a *IntegerAssertion) NotEqualTo(expected int, msgAndArgs ...interface{}) *IntegerAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if a.value == expected {
		a.fail(
			fmt.Sprintf(`not equal to %d, actual is %d`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`greater than %d, actual is %d`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`greater than or equal %d, actual is %d`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`less than %d, actual is %d`, expected, a.value),
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
		a.fail(
			fmt.Sprintf(`less than or equal %d, actual is %d`, expected, a.value),
			msgAndArgs...,
		)
	}

	return nil
}

func (a *IntegerAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	assert.Fail(a.t, a.message+message, msgAndArgs...)
}
