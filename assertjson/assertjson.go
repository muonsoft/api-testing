package assertjson

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yalp/jsonpath"
	"io/ioutil"
	"testing"
)

type AssertJSON struct {
	t    *testing.T
	data interface{}
}

type AssertNode struct {
	t     *testing.T
	err   error
	path  string
	value interface{}
}

type JSONAssertFunc func(json *AssertJSON)

func FileHas(t *testing.T, filename string, jsonAssert JSONAssertFunc) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		assert.Failf(t, "failed to read file '%s': %s", filename, err.Error())
	}
	Has(t, data, jsonAssert)
}

func Has(t *testing.T, data []byte, jsonAssert JSONAssertFunc) {
	body := &AssertJSON{t: t}
	err := json.Unmarshal(data, &body.data)
	if err != nil {
		assert.Failf(t, "data has invalid JSON: %s", err.Error())
	} else {
		jsonAssert(body)
	}
}

func (j *AssertJSON) Node(path string) *AssertNode {
	value, err := jsonpath.Read(j.data, path)

	return &AssertNode{
		t:     j.t,
		err:   err,
		path:  path,
		value: value,
	}
}

func (node *AssertNode) exists() bool {
	if node.err != nil {
		assert.Fail(node.t, fmt.Sprintf("failed to find json node '%s': %v", node.path, node.err))
	}

	return node.err == nil
}
