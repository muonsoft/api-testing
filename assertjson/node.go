package assertjson

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/stretchr/testify/assert"
)

// AssertNode - structure for asserting JSON node.
type AssertNode struct {
	t          TestingT
	err        error
	pathPrefix string
	path       string
	value      interface{}
}

// Value returns JSON node value as an interface. If node does not exist it returns nil.
func (node *AssertNode) Value() interface{} {
	node.t.Helper()
	if node.exists() {
		return node.value
	}

	return nil
}

// String returns the JSON node value as a string. If node does not exist it returns an empty string.
// If the node value is an integer it returns a formatted integer.
// If the node value is a float it returns a formatted float.
// Otherwise, it logs the error and returns an empty string.
func (node *AssertNode) String() string {
	node.t.Helper()
	if node.exists() {
		if s, ok := node.value.(string); ok {
			return s
		}
		if f, ok := node.value.(float64); ok {
			if n, f := math.Modf(f); f == 0 {
				return strconv.Itoa(int(n))
			}
			return fmt.Sprintf("%f", f)
		}
		assert.Fail(node.t, fmt.Sprintf(`json node at "%s" cannot be converted into string`, node.Path()))
	}

	return ""
}

// Float returns the JSON node value as a 64-bit float. If node does not exist it returns a zero.
// If node value is not numeric, it logs the error and returns a zero value.
func (node *AssertNode) Float() float64 {
	node.t.Helper()
	if node.exists() {
		if f, ok := node.value.(float64); ok {
			return f
		}
		assert.Fail(node.t, fmt.Sprintf(`json node at "%s" cannot be converted into float`, node.Path()))
	}

	return 0
}

// Integer returns the JSON node value as an integer. If node does not exist it returns a zero.
// If node value is not an integer, it logs the error and returns a zero value.
func (node *AssertNode) Integer() int {
	node.t.Helper()
	if node.exists() {
		if f, ok := node.value.(float64); ok {
			if n, f := math.Modf(f); f == 0 {
				return int(n)
			}
			assert.Fail(node.t, fmt.Sprintf(`json node at "%s" is not an integer`, node.Path()))
		} else {
			assert.Fail(node.t, fmt.Sprintf(`json node at "%s" cannot be converted into float`, node.Path()))
		}
	}

	return 0
}

// JSON returns the JSON node value as a marshaled JSON.
func (node *AssertNode) JSON() []byte {
	node.t.Helper()
	if node.exists() {
		data, _ := json.Marshal(node.value)
		return data
	}

	return nil
}

// Path returns current node path as string.
func (node *AssertNode) Path() string {
	return node.pathPrefix + node.path
}

func (node *AssertNode) fail() {
	assert.Fail(node.t, fmt.Sprintf(`failed at json node "%s"`, node.Path()))
}

func (node *AssertNode) exists() bool {
	node.t.Helper()
	if node.err != nil {
		node.t.Errorf(`failed to find json node "%s": %v`, node.Path(), node.err)
	}

	return node.err == nil
}
