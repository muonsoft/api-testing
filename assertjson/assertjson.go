// Package assertjson provides methods for testing JSON values.
// Selecting JSON values provided by JSON Pointer Syntax (https://tools.ietf.org/html/rfc6901).
package assertjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/xeipuuv/gojsonpointer"
)

// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Helper()
	Errorf(format string, args ...interface{})
}

// AssertJSON - main structure that holds parsed JSON.
type AssertJSON struct {
	t    TestingT
	path string
	data interface{}
}

// JSONAssertFunc - callback function used for asserting JSON nodes.
type JSONAssertFunc func(json *AssertJSON)

// FileHas loads JSON from file and runs user callback for testing its nodes.
func FileHas(t TestingT, filename string, jsonAssert JSONAssertFunc) {
	t.Helper()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("failed to read file '%s': %s", filename, err.Error())
	} else {
		Has(t, data, jsonAssert)
	}
}

// Has - loads JSON from byte slice and runs user callback for testing its nodes.
func Has(t TestingT, data []byte, jsonAssert JSONAssertFunc) {
	t.Helper()
	body := &AssertJSON{t: t}
	err := json.Unmarshal(data, &body.data)
	if err != nil {
		t.Errorf("data has invalid JSON: %s", err.Error())
	} else {
		jsonAssert(body)
	}
}

// Node searches for JSON node by JSON Path Syntax. Returns struct for asserting the node values.
func (j *AssertJSON) Node(path string) *AssertNode {
	j.t.Helper()
	var value interface{}

	pointer, err := gojsonpointer.NewJsonPointer(path)
	if err == nil {
		value, _, err = pointer.Get(j.data)
	}

	return &AssertNode{
		t:          j.t,
		err:        err,
		pathPrefix: j.path,
		path:       path,
		value:      value,
	}
}

// Nodef searches for JSON node by JSON Path Syntax. Returns struct for asserting the node values.
// It calculates path by applying fmt.Sprintf function.
func (j *AssertJSON) Nodef(format string, a ...interface{}) *AssertNode {
	j.t.Helper()
	return j.Node(fmt.Sprintf(format, a...))
}

// At is used to test assertions on some node in a batch. It returns AssertJSON object on that node.
func (j *AssertJSON) At(path string) *AssertJSON {
	j.t.Helper()
	var value interface{}

	pointer, err := gojsonpointer.NewJsonPointer(path)
	if err == nil {
		value, _, err = pointer.Get(j.data)
	}
	if err != nil {
		j.t.Errorf(`failed to find JSON node "%s": %v`, path, err)
	}

	return &AssertJSON{
		t:    j.t,
		path: j.path + path,
		data: value,
	}
}

// Atf is used to test assertions on some node in a batch. It returns AssertJSON object on that node.
// It calculates path by applying fmt.Sprintf function.
func (j *AssertJSON) Atf(format string, a ...interface{}) *AssertJSON {
	j.t.Helper()
	return j.At(fmt.Sprintf(format, a...))
}
