// Package assertjson provides methods for testing JSON values.
// Nodes are selected by path elements passed to Node and At.
package assertjson

import (
	"encoding/json"
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/muonsoft/api-testing/internal/js"
	"github.com/stretchr/testify/assert"
)

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Helper()
	Errorf(format string, args ...interface{})
	Log(args ...interface{})
	Failed() bool
}

// AssertJSON - main structure that holds parsed JSON.
type AssertJSON struct {
	t       TestingT
	message string
	path    *js.Path
	data    interface{}
}

func NewAssertJSON(t TestingT, message string, data interface{}) *AssertJSON {
	return &AssertJSON{t: t, message: message, data: data}
}

// JSONAssertFunc - callback function used for asserting JSON nodes.
type JSONAssertFunc func(json *AssertJSON)

// FileHas loads JSON from file and runs user callback for testing its nodes.
func FileHas(t TestingT, filename string, jsonAssert JSONAssertFunc) bool {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		assert.Fail(t, fmt.Sprintf(`failed to read file "%s": %s`, filename, err.Error()))

		return false
	}

	return Has(t, data, jsonAssert)
}

// Has - loads JSON from byte slice and runs user callback for testing its nodes.
func Has(t TestingT, data []byte, jsonAssert JSONAssertFunc) bool {
	t.Helper()
	body := &AssertJSON{t: t}
	body.assert(data, jsonAssert)

	return !t.Failed()
}

// Node searches for JSON node by JSON Path Syntax. Returns struct for asserting the node values.
func (j *AssertJSON) Node(path ...interface{}) *AssertNode {
	j.t.Helper()

	node := &AssertNode{t: j.t, message: j.message}

	path = preprocessPath(path)
	jspath, err := js.PathFromAny(path...)
	if err != nil {
		node.fail(fmt.Sprintf("parse path: %s", err.Error()))
	} else {
		node.path = j.path.With(jspath)
		node.value, node.err = getValueByPath(j.data, path...)
	}

	return node
}

// At is used to test assertions on some node in a batch. It returns AssertJSON object on that node.
func (j *AssertJSON) At(path ...interface{}) *AssertJSON {
	j.t.Helper()
	a := &AssertJSON{t: j.t}

	path = preprocessPath(path)
	jsPath, err := js.PathFromAny(path...)
	if err != nil {
		a.fail(fmt.Sprintf("parse path: %s", err.Error()))
	}
	a.path = j.path.With(jsPath)

	a.data, err = getValueByPath(j.data, path...)
	if err != nil {
		j.fail(fmt.Sprintf(`failed to find JSON node "%s": %v`, a.path.String(), err))
	}

	return a
}

func (j *AssertJSON) assert(data []byte, jsonAssert JSONAssertFunc) {
	j.t.Helper()
	err := json.Unmarshal(data, &j.data)
	if err != nil {
		j.fail(fmt.Sprintf("data has invalid JSON: %s", err.Error()))
	} else {
		jsonAssert(j)
	}
}

func (j *AssertJSON) fail(message string, msgAndArgs ...interface{}) {
	j.t.Helper()
	assert.Fail(j.t, j.message+message, msgAndArgs...)
}

func preprocessPath(path []interface{}) []interface{} {
	for i := range path {
		if s, ok := path[i].(fmt.Stringer); ok {
			path[i] = s.String()
		}
	}

	return path
}

func getValueByPath(data interface{}, path ...interface{}) (interface{}, error) {
	v := jsoniter.Wrap(data)
	for _, e := range path {
		if _, ok := e.(int); ok {
			if _, ok := v.GetInterface().(map[string]interface{}); ok {
				return nil, fmt.Errorf("value of type int is not assignable to type string")
			}
		}

		v = v.Get(e)
		if v.LastError() != nil {
			return nil, v.LastError()
		}
	}

	return v.GetInterface(), nil
}
