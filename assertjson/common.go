package assertjson

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stretchr/testify/assert"
)

// Exists asserts that the JSON node exists. Returns true if node exists.
func (node *AssertNode) Exists(msgAndArgs ...interface{}) bool {
	node.t.Helper()
	if node.err != nil {
		node.fail(
			fmt.Sprintf(`failed asserting that JSON node "%s" exists`, node.path.String()),
			msgAndArgs...,
		)
		return false
	}

	return true
}

// DoesNotExist asserts that the JSON node does not exist.
func (node *AssertNode) DoesNotExist(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.err == nil {
		node.fail(
			fmt.Sprintf(`failed asserting that JSON node "%s" does not exist`, node.path.String()),
			msgAndArgs...,
		)
	}
}

// IsNull asserts that the JSON node equals to `null` value.
func (node *AssertNode) IsNull(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		if !isNil(node.value) {
			node.fail(
				fmt.Sprintf(`failed asserting that JSON node "%s" is null`, node.path.String()),
				msgAndArgs...,
			)
		}
	}
}

// IsNotNull asserts that the JSON node not equals to `null` value.
func (node *AssertNode) IsNotNull(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		if isNil(node.value) {
			node.fail(
				fmt.Sprintf(`failed asserting that JSON node "%s" is not null`, node.path.String()),
				msgAndArgs...,
			)
		}
	}
}

// IsTrue asserts that the JSON node equals to the boolean with `true` value.
func (node *AssertNode) IsTrue(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		if b, ok := node.value.(bool); ok {
			if !b {
				node.fail(
					fmt.Sprintf(`failed asserting that JSON node "%s" is true`, node.path.String()),
					msgAndArgs...,
				)
			}
			return
		}
		node.fail(
			fmt.Sprintf(`failed asserting that JSON node "%s" is boolean`, node.path.String()),
			msgAndArgs...,
		)
	}
}

// IsFalse asserts that the JSON node equals to the boolean with `false` value.
func (node *AssertNode) IsFalse(msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		if b, ok := node.value.(bool); ok {
			if b {
				node.fail(
					fmt.Sprintf(`failed asserting that JSON node "%s" is false`, node.path.String()),
					msgAndArgs...,
				)
			}
			return
		}
		node.fail(
			fmt.Sprintf(`failed asserting that JSON node "%s" is boolean`, node.path.String()),
			msgAndArgs...,
		)
	}
}

// EqualJSON asserts that node is equal to JSON string.
func (node *AssertNode) EqualJSON(expected string, msgAndArgs ...interface{}) {
	node.t.Helper()
	if node.exists() {
		data, _ := json.Marshal(node.value)
		if !assert.JSONEq(node.t, expected, string(data), msgAndArgs...) {
			node.fail(fmt.Sprintf(`failed at JSON node "%s"`, node.path.String()))
		}
	}
}

// isNil checks if a specified object is nil or not, without failing.
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice,
		},
		kind,
	)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

// containsKind checks if a specified kind in the slice of kinds.
func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}
