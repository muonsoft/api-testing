package assertjson

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/stretchr/testify/assert"
)

// IsArrayWithElementsCount asserts that the JSON node is an array with given elements count.
// Deprecated: use IsArray().WithLength() instead.
func (node *AssertNode) IsArrayWithElementsCount(count int, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsArray().WithLength(count, msgAndArgs...)
}

// IsArray asserts that the JSON node is an array.
// It returns ArrayAssertion to execute a chain of assertions for the node value.
func (node *AssertNode) IsArray(msgAndArgs ...interface{}) *ArrayAssertion {
	node.t.Helper()
	if node.exists() {
		if array, ok := node.value.([]interface{}); ok {
			return &ArrayAssertion{t: node.t, message: node.message, path: node.Path(), value: array}
		}
		node.fail(
			fmt.Sprintf(`failed asserting that JSON node "%s" is array`, node.Path()),
			msgAndArgs...,
		)
	}

	return nil
}

// ArrayAssertion is used to build a chain of assertions for the array node.
type ArrayAssertion struct {
	t       TestingT
	message string
	path    string
	value   []interface{}
}

// WithLength asserts that the JSON node is an array with length equal to the given value.
func (a *ArrayAssertion) WithLength(expected int, msgAndArgs ...interface{}) *ArrayAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) != expected {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is array with length is %d, actual is %d`,
				a.path,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThan asserts that the JSON node is an array with length greater than the value.
func (a *ArrayAssertion) WithLengthGreaterThan(expected int, msgAndArgs ...interface{}) *ArrayAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) <= expected {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is array with length greater than %d, actual is %d`,
				a.path,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthGreaterThanOrEqual asserts that the JSON node is an array
// with length greater than or equal to the value.
func (a *ArrayAssertion) WithLengthGreaterThanOrEqual(expected int, msgAndArgs ...interface{}) *ArrayAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) < expected {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is array with length greater than or equal to %d, actual is %d`,
				a.path,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThan asserts that the JSON node is an array with length less than the value.
func (a *ArrayAssertion) WithLengthLessThan(expected int, msgAndArgs ...interface{}) *ArrayAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) >= expected {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is array with length less than %d, actual is %d`,
				a.path,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithLengthLessThanOrEqual asserts that the JSON node is an array
// with length less than or equal to the value.
func (a *ArrayAssertion) WithLengthLessThanOrEqual(expected int, msgAndArgs ...interface{}) *ArrayAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) > expected {
		a.fail(
			fmt.Sprintf(
				`failed asserting that JSON node "%s" is array with length less than or equal to %d, actual is %d`,
				a.path,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithUniqueElements asserts that the JSON node is an array with unique elements.
func (a *ArrayAssertion) WithUniqueElements(msgAndArgs ...interface{}) *ArrayAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	uniques := make(map[string][]int, len(a.value))
	keys := make([]string, 0, len(a.value))

	for i, value := range a.value {
		raw, _ := json.Marshal(value)
		key := string(raw)
		if _, exist := uniques[key]; !exist {
			keys = append(keys, key)
		}
		uniques[key] = append(uniques[key], i)
	}

	duplicates := make([]string, 0)
	for _, key := range keys {
		if len(uniques[key]) > 1 {
			duplicates = append(duplicates, fmt.Sprintf(
				"value %s is duplicated at %s",
				key,
				strings.Join(intsToStrings(uniques[key]), ", "),
			))
		}
	}

	if len(duplicates) > 0 {
		a.fail(
			fmt.Sprintf(
				"failed asserting that JSON node \"%s\" is array with unique elements, duplicated elements:\n%s",
				a.path,
				strings.Join(duplicates, ";\n"),
			),
			msgAndArgs...,
		)
	}

	return a
}

// Length returns array underlying array length.
func (a *ArrayAssertion) Length() int {
	if a == nil {
		return 0
	}
	a.t.Helper()

	return len(a.value)
}

func (a *ArrayAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	if a.message != "" {
		message = a.message + ": " + message
	}
	assert.Fail(a.t, message, msgAndArgs...)
}

// ArrayLength asserts that JSON node is array and return its length.
// It is an alias for IsArray().Length().
func (node *AssertNode) ArrayLength() int {
	return node.IsArray().Length()
}

func intsToStrings(ints []int) []string {
	s := make([]string, len(ints))
	for i, v := range ints {
		s[i] = strconv.Itoa(v)
	}
	return s
}
