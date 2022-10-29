package assertjson

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/stretchr/testify/assert"
)

// IsObjectWithPropertiesCount asserts that the JSON node is an object with given properties count.
// Deprecated: use IsObject().WithPropertiesCount() instead.
func (node *AssertNode) IsObjectWithPropertiesCount(count int, msgAndArgs ...interface{}) {
	node.t.Helper()
	node.IsObject().WithPropertiesCount(count, msgAndArgs...)
}

// IsObject asserts that the JSON node is an object.
// It returns ObjectAssertion to execute a chain of assertions for the node value.
func (node *AssertNode) IsObject(msgAndArgs ...interface{}) *ObjectAssertion {
	node.t.Helper()
	if node.exists() {
		if object, ok := node.value.(map[string]interface{}); ok {
			return &ObjectAssertion{
				t:       node.t,
				message: fmt.Sprintf(`%sfailed asserting that JSON node "%s": `, node.message, node.path.String()),
				path:    node.path.String(),
				value:   object,
			}
		}
		node.fail(
			fmt.Sprintf(`failed asserting that JSON node "%s" is object`, node.path.String()),
			msgAndArgs...,
		)
	}

	return nil
}

// ObjectAssertion is used to build a chain of assertions for the object node.
type ObjectAssertion struct {
	t       TestingT
	message string
	path    string
	value   map[string]interface{}
}

// WithPropertiesCount asserts that the JSON node is an object with properties count equal to the given value.
func (a *ObjectAssertion) WithPropertiesCount(expected int, msgAndArgs ...interface{}) *ObjectAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) != expected {
		a.fail(
			fmt.Sprintf(
				`is object with properties count is %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithPropertiesCountGreaterThan asserts that the JSON node is an object with properties count greater than the value.
func (a *ObjectAssertion) WithPropertiesCountGreaterThan(expected int, msgAndArgs ...interface{}) *ObjectAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) <= expected {
		a.fail(
			fmt.Sprintf(
				`is object with properties count greater than %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithPropertiesCountGreaterThanOrEqual asserts that the JSON node is an object
// with propertiesCount greater than or equal to the value.
func (a *ObjectAssertion) WithPropertiesCountGreaterThanOrEqual(expected int, msgAndArgs ...interface{}) *ObjectAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) < expected {
		a.fail(
			fmt.Sprintf(
				`is object with properties count greater than or equal to %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithPropertiesCountLessThan asserts that the JSON node is an object with properties count less than the value.
func (a *ObjectAssertion) WithPropertiesCountLessThan(expected int, msgAndArgs ...interface{}) *ObjectAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) >= expected {
		a.fail(
			fmt.Sprintf(
				`is object with properties count less than %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithPropertiesCountLessThanOrEqual asserts that the JSON node is an object
// with properties count less than or equal to the value.
func (a *ObjectAssertion) WithPropertiesCountLessThanOrEqual(expected int, msgAndArgs ...interface{}) *ObjectAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	if len(a.value) > expected {
		a.fail(
			fmt.Sprintf(
				`is object with properties count less than or equal to %d, actual is %d`,
				expected,
				len(a.value),
			),
			msgAndArgs...,
		)
	}

	return a
}

// WithUniqueElements asserts that the JSON node is an object with unique elements.
func (a *ObjectAssertion) WithUniqueElements(msgAndArgs ...interface{}) *ObjectAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	uniques := make(map[string][]string, len(a.value))
	keys := make([]string, 0, len(a.value))

	for k, value := range a.value {
		raw, _ := json.Marshal(value)
		key := string(raw)
		if _, exist := uniques[key]; !exist {
			keys = append(keys, key)
		}
		uniques[key] = append(uniques[key], k)
	}

	duplicates := make([]string, 0)
	for _, key := range keys {
		if len(uniques[key]) > 1 {
			duplicates = append(duplicates, fmt.Sprintf(
				"value %s is duplicated at %s",
				key,
				strings.Join(quoteAll(uniques[key]), ", "),
			))
		}
	}

	if len(duplicates) > 0 {
		a.fail(
			fmt.Sprintf(
				"is object with unique elements, duplicated elements:\n%s",
				strings.Join(duplicates, ";\n"),
			),
			msgAndArgs...,
		)
	}

	return a
}

// PropertiesCount returns array underlying object properties count.
func (a *ObjectAssertion) PropertiesCount() int {
	if a == nil {
		return 0
	}
	a.t.Helper()

	return len(a.value)
}

// ObjectPropertiesCount asserts that JSON node is an object and return its properties count.
// It is an alias for IsObject().PropertiesCount().
func (node *AssertNode) ObjectPropertiesCount() int {
	return node.IsObject().PropertiesCount()
}

func (a *ObjectAssertion) fail(message string, msgAndArgs ...interface{}) {
	a.t.Helper()
	assert.Fail(a.t, a.message+message, msgAndArgs...)
}

func quoteAll(s []string) []string {
	ss := make([]string, len(s))
	for i, v := range s {
		ss[i] = strconv.Quote(v)
	}
	return ss
}
