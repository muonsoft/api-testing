// Package assertjson provides methods for testing JSON values. Selecting JSON values provided by JSON Path Syntax.
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
//          json.Node("$.nullNode").Exists()
//          json.Node("$.notExistingNode").DoesNotExist()
//          json.Node("$.nullNode").IsNull()
//          json.Node("$.stringNode").IsNotNull()
//          json.Node("$.trueBooleanNode").IsTrue()
//          json.Node("$.falseBooleanNode").IsFalse()
//
//          // string assertions
//          json.Node("$.stringNode").IsString()
//          json.Node("$.stringNode").EqualToTheString("stringValue")
//          json.Node("$.stringNode").Matches("^string.*$")
//          json.Node("$.stringNode").DoesNotMatch("^notMatch$")
//          json.Node("$.stringNode").Contains("string")
//          json.Node("$.stringNode").DoesNotContain("notContain")
//          json.Node("$.stringNode").IsStringWithLength(11)
//          json.Node("$.stringNode").IsStringWithLengthInRange(11, 11)
//
//          // numeric assertions
//          json.Node("$.integerNode").IsInteger()
//          json.Node("$.integerNode").EqualToTheInteger(123)
//          json.Node("$.integerNode").IsNumberInRange(122, 124)
//          json.Node("$.integerNode").IsNumberGreaterThan(122)
//          json.Node("$.integerNode").IsNumberGreaterThanOrEqual(123)
//          json.Node("$.integerNode").IsNumberLessThan(124)
//          json.Node("$.integerNode").IsNumberLessThanOrEqual(123)
//          json.Node("$.floatNode").IsFloat()
//          json.Node("$.floatNode").EqualToTheFloat(123.123)
//          json.Node("$.floatNode").IsNumberInRange(122, 124)
//          json.Node("$.floatNode").IsNumberGreaterThan(122)
//          json.Node("$.floatNode").IsNumberGreaterThanOrEqual(123.123)
//          json.Node("$.floatNode").IsNumberLessThan(124)
//          json.Node("$.floatNode").IsNumberLessThanOrEqual(123.123)
//
//          // array assertions
//          json.Node("$.arrayNode").IsArrayWithElementsCount(1)
//
//          // json path expression
//          json.Node("$.complexNode.items[1].key").EqualToTheString("value")
//      })
//   }
package assertjson

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yalp/jsonpath"
	"io/ioutil"
	"testing"
)

// AssertJSON - main structure that holds parsed JSON.
type AssertJSON struct {
	t    *testing.T
	data interface{}
}

// AssertNode - structure for asserting JSON node.
type AssertNode struct {
	t     *testing.T
	err   error
	path  string
	value interface{}
}

// JSONAssertFunc - callback function used for asserting JSON nodes.
type JSONAssertFunc func(json *AssertJSON)

// FileHas loads JSON from file and runs user callback for testing its nodes.
func FileHas(t *testing.T, filename string, jsonAssert JSONAssertFunc) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		assert.Failf(t, "failed to read file '%s': %s", filename, err.Error())
	}
	Has(t, data, jsonAssert)
}

// Has - loads JSON from byte slice and runs user callback for testing its nodes.
func Has(t *testing.T, data []byte, jsonAssert JSONAssertFunc) {
	body := &AssertJSON{t: t}
	err := json.Unmarshal(data, &body.data)
	if err != nil {
		assert.Failf(t, "data has invalid JSON: %s", err.Error())
	} else {
		jsonAssert(body)
	}
}

// Node searches for JSON node by JSON Path Syntax. Returns struct for asserting the node values.
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
