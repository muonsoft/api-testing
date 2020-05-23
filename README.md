# API testing tools for Golang

![CI](https://github.com/muonsoft/api-testing/workflows/CI/badge.svg?branch=master)

## `assertjson` package

The `assertjson` package provides methods for testing JSON values. Selecting JSON values provided by JSON Path Syntax.

Example

```go
package yours

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/muonsoft/api-testing/assertjson"
)

func TestYourAPI(t *testing.T) {
    recorder := httptest.NewRecorder()
    handler := createHTTPHandler()

    request, _ := http.NewRequest("GET", "/content", nil)
    handler.ServeHTTP(recorder, request)

    assertjson.Has(t, recorder.Body.Bytes(), func(json *assertjson.AssertJSON) {
        // common assertions
        json.Node("$.nullNode").Exists()
        json.Node("$.notExistingNode").DoesNotExist()
        json.Node("$.nullNode").IsNull()
        json.Node("$.stringNode").IsNotNull()
        json.Node("$.trueBooleanNode").IsTrue()
        json.Node("$.falseBooleanNode").IsFalse()

        // string assertions
        json.Node("$.stringNode").IsString()
        json.Node("$.stringNode").EqualToTheString("stringValue")
        json.Node("$.stringNode").Matches("^string.*$")
        json.Node("$.stringNode").DoesNotMatch("^notMatch$")
        json.Node("$.stringNode").Contains("string")
        json.Node("$.stringNode").DoesNotContain("notContain")
        json.Node("$.stringNode").IsStringWithLength(11)
        json.Node("$.stringNode").IsStringWithLengthInRange(11, 11)

        // numeric assertions
        json.Node("$.integerNode").IsInteger()
        json.Node("$.integerNode").EqualToTheInteger(123)
        json.Node("$.integerNode").IsNumberInRange(122, 124)
        json.Node("$.integerNode").IsNumberGreaterThan(122)
        json.Node("$.integerNode").IsNumberGreaterThanOrEqual(123)
        json.Node("$.integerNode").IsNumberLessThan(124)
        json.Node("$.integerNode").IsNumberLessThanOrEqual(123)
        json.Node("$.floatNode").IsFloat()
        json.Node("$.floatNode").EqualToTheFloat(123.123)
        json.Node("$.floatNode").IsNumberInRange(122, 124)
        json.Node("$.floatNode").IsNumberGreaterThan(122)
        json.Node("$.floatNode").IsNumberGreaterThanOrEqual(123.123)
        json.Node("$.floatNode").IsNumberLessThan(124)
        json.Node("$.floatNode").IsNumberLessThanOrEqual(123.123)

        // array assertions
        json.Node("$.arrayNode").IsArrayWithElementsCount(1)

        // json path expression
        json.Node("$.complexNode.items[1].key").EqualToTheString("value")
    })
}
```

## `assertxml` package

The `assertjson` package provides methods for testing XML values. Selecting XML values provided by XML Path Syntax.

Example

```go
package yours

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/muonsoft/api-testing/assertxml"
)

func TestYourAPI(t *testing.T) {
    recorder := httptest.NewRecorder()
    handler := createHTTPHandler()

    request, _ := http.NewRequest("GET", "/content", nil)
    handler.ServeHTTP(recorder, request)

    assertxml.Has(t, recorder.Body.Bytes(), func(xml *AssertXML) {
        // common assertions
        xml.Node("/root/stringNode").Exists()
        xml.Node("/root/notExistingNode").DoesNotExist()
  
        // string assertions
        xml.Node("/root/stringNode").EqualToTheString("stringValue")
    })
}
```
