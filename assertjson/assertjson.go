// Package assertjson provides methods for testing JSON values.
// Selecting JSON values provided by JSON Pointer Syntax (https://tools.ietf.org/html/rfc6901).
//
// Example usage
//
//   import (
//      "net/http"
//      "net/http/httptest"
//      "testing"
//      "github.com/muonsoft/api-testing/assertjson"
//   )
//
//   func TestYourAPI(t *testing.T) {
//      recorder := httptest.NewRecorder()
//      handler := createHTTPHandler()
//
//      request, _ := http.NewRequest("GET", "/content", nil)
//      handler.ServeHTTP(recorder, request)
//
//      assertjson.Has(t, recorder.Body.Bytes(), func(json *assertjson.AssertJSON) {
//          // common assertions
//          json.Node("/nullNode").Exists()
//          json.Node("/notExistingNode").DoesNotExist()
//          json.Node("/nullNode").IsNull()
//          json.Node("/stringNode").IsNotNull()
//          json.Node("/trueBooleanNode").IsTrue()
//          json.Node("/falseBooleanNode").IsFalse()
//
//          // string assertions
//          json.Node("/stringNode").IsString()
//          json.Node("/stringNode").EqualToTheString("stringValue")
//          json.Node("/stringNode").Matches("^string.*$")
//          json.Node("/stringNode").DoesNotMatch("^notMatch$")
//          json.Node("/stringNode").Contains("string")
//          json.Node("/stringNode").DoesNotContain("notContain")
//          json.Node("/stringNode").IsStringWithLength(11)
//          json.Node("/stringNode").IsStringWithLengthInRange(11, 11)
//
//          // numeric assertions
//          json.Node("/integerNode").IsInteger()
//          json.Node("/integerNode").EqualToTheInteger(123)
//          json.Node("/integerNode").IsNumberInRange(122, 124)
//          json.Node("/integerNode").IsNumberGreaterThan(122)
//          json.Node("/integerNode").IsNumberGreaterThanOrEqual(123)
//          json.Node("/integerNode").IsNumberLessThan(124)
//          json.Node("/integerNode").IsNumberLessThanOrEqual(123)
//          json.Node("/floatNode").IsFloat()
//          json.Node("/floatNode").EqualToTheFloat(123.123)
//          json.Node("/floatNode").IsNumberInRange(122, 124)
//          json.Node("/floatNode").IsNumberGreaterThan(122)
//          json.Node("/floatNode").IsNumberGreaterThanOrEqual(123.123)
//          json.Node("/floatNode").IsNumberLessThan(124)
//          json.Node("/floatNode").IsNumberLessThanOrEqual(123.123)
//
//          // array assertions
//          json.Node("/arrayNode").IsArrayWithElementsCount(1)
//          // object assertions
//          json.Node("/objectNode").IsObjectWithPropertiesCount(1)
//
//          // json pointer expression
//          json.Node("/complexNode/items/1/key").EqualToTheString("value")
//      })
//   }
package assertjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/xeipuuv/gojsonpointer"
)

// AssertJSON - main structure that holds parsed JSON.
type AssertJSON struct {
	t    testing.TB
	path string
	data interface{}
}

// AssertNode - structure for asserting JSON node.
type AssertNode struct {
	t          testing.TB
	err        error
	pathPrefix string
	path       string
	value      interface{}
}

// JSONAssertFunc - callback function used for asserting JSON nodes.
type JSONAssertFunc func(json *AssertJSON)

// FileHas loads JSON from file and runs user callback for testing its nodes.
func FileHas(t testing.TB, filename string, jsonAssert JSONAssertFunc) {
	t.Helper()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("failed to read file '%s': %s", filename, err.Error())
	} else {
		Has(t, data, jsonAssert)
	}
}

// Has - loads JSON from byte slice and runs user callback for testing its nodes.
func Has(t testing.TB, data []byte, jsonAssert JSONAssertFunc) {
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
		j.t.Errorf(`failed to find json node "%s": %v`, path, err)
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

func (node *AssertNode) exists() bool {
	node.t.Helper()
	if node.err != nil {
		node.t.Errorf(`failed to find json node "%s": %v`, node.pathPrefix+node.path, node.err)
	}

	return node.err == nil
}
