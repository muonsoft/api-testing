# API testing tools for Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/muonsoft/api-testing.svg)](https://pkg.go.dev/github.com/muonsoft/api-testing)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/muonsoft/api-testing)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/muonsoft/api-testing)
![GitHub](https://img.shields.io/github/license/muonsoft/api-testing)
[![Go Report Card](https://goreportcard.com/badge/github.com/muonsoft/api-testing)](https://goreportcard.com/report/github.com/muonsoft/api-testing)
![CI](https://github.com/muonsoft/api-testing/workflows/CI/badge.svg?branch=master)

## `apitest` package

The `apitest` package provides methods for testing client-server communication.
It can be used to test `http.Handler` to build complex assertions on the HTTP responses.

Example

```go
package yours

import (
    "net/http"
    "testing"
    "github.com/muonsoft/api-testing/apitest"
    "github.com/muonsoft/api-testing/assertjson"
)

func TestYourAPI(t *testing.T) {
    handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        writer.WriteHeader(http.StatusOK)
        writer.Header().Set("Content-Type", "application/json")
        writer.Write([]byte(`{"ok":true}`))
    })
    
    // HandleGET builds and sends GET request to handler
    response := apitest.HandleGET(t, handler, "/example")
    
    response.IsOK()
    response.HasContentType("application/json")
    response.HasJSON(func(json *assertjson.AssertJSON) {
        json.Node("/ok").IsTrue()
    })
    response.Print() // prints response with headers and body
    response.PrintJSON() // prints response with headers and indented JSON body
}
```

## `assertjson` package

The `assertjson` package provides methods for testing JSON values. Selecting JSON values provided by [JSON Pointer Syntax](https://tools.ietf.org/html/rfc6901).

Example

```go
package yours

import (
    "fmt"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gofrs/uuid"
    "github.com/golang-jwt/jwt/v4"
    "github.com/muonsoft/api-testing/assertjson"
    "github.com/stretchr/testify/assert"
)

func TestYourAPI(t *testing.T) {
    recorder := httptest.NewRecorder()
    handler := newHTTPHandler()

    request := httptest.NewRequest("GET", "/content", nil)
    handler.ServeHTTP(recorder, request)

    assertjson.Has(t, recorder.Body.Bytes(), func(json *assertjson.AssertJSON) {
        // common assertions
        json.Node("/nullNode").Exists()
        json.Node("/notExistingNode").DoesNotExist()
        json.Node("/nullNode").IsNull()
        json.Node("/stringNode").IsNotNull()
        json.Node("/trueBooleanNode").IsTrue()
        json.Node("/falseBooleanNode").IsFalse()
        json.Node("/objectNode").EqualJSON(`{"objectKey": "objectValue"}`)

        // string assertions
        json.Node("/stringNode").IsString()
        json.Node("/stringNode").Matches("^string.*$")
        json.Node("/stringNode").DoesNotMatch("^notMatch$")
        json.Node("/stringNode").Contains("string")
        json.Node("/stringNode").DoesNotContain("notContain")

        // fluent string assertions
        json.Node("/stringNode").IsString()
        json.Node("/stringNode").IsString().EqualTo("stringValue")
        json.Node("/stringNode").IsString().EqualToOneOf("stringValue", "nextValue")
        json.Node("/stringNode").IsString().NotEqualTo("invalid")
        json.Node("/stringNode").IsString().Matches("^string.*$")
        json.Node("/stringNode").IsString().NotMatches("^notMatch$")
        json.Node("/stringNode").IsString().Contains("string")
        json.Node("/stringNode").IsString().NotContains("notContain")
        json.Node("/stringNode").IsString().WithLength(11)
        json.Node("/stringNode").IsString().WithLengthGreaterThan(10)
        json.Node("/stringNode").IsString().WithLengthGreaterThanOrEqual(11)
        json.Node("/stringNode").IsString().WithLengthLessThan(12)
        json.Node("/stringNode").IsString().WithLengthLessThanOrEqual(11)
        json.Node("/stringNode").IsString().That(func(s string) error {
            if s != "stringValue" {
                return fmt.Errorf("invalid")
            }
            return nil
        })
        json.Node("/stringNode").IsString().Assert(func(t testing.TB, value string) {
            assert.Equal(t, "stringValue", value)
        })

        // numeric assertions
        json.Node("/integerNode").IsInteger()
        json.Node("/integerNode").IsInteger().EqualTo(123)
        json.Node("/integerNode").IsInteger().NotEqualTo(321)
        json.Node("/integerNode").IsInteger().GreaterThan(122)
        json.Node("/integerNode").IsInteger().GreaterThanOrEqual(123)
        json.Node("/integerNode").IsInteger().LessThan(124)
        json.Node("/integerNode").IsInteger().LessThanOrEqual(123)
        json.Node("/floatNode").IsFloat()
        json.Node("/floatNode").IsNumber()
        json.Node("/floatNode").IsNumber().EqualTo(123.123)
        json.Node("/floatNode").IsNumber().NotEqualTo(321.123)
        json.Node("/floatNode").IsNumber().EqualToWithDelta(123.123, 0.1)
        json.Node("/floatNode").IsNumber().GreaterThan(122)
        json.Node("/floatNode").IsNumber().GreaterThanOrEqual(123.123)
        json.Node("/floatNode").IsNumber().LessThan(124)
        json.Node("/floatNode").IsNumber().LessThanOrEqual(123.123)
        json.Node("/floatNode").IsNumber().GreaterThanOrEqual(122).LessThanOrEqual(124)

        // string values assertions
        json.Node("/uuid").IsString().WithUUID()
        json.Node("/uuid").IsUUID().NotNil().Version(4).Variant(1)
        json.Node("/uuid").IsUUID().EqualTo(uuid.FromStringOrNil("23e98a0c-26c8-410f-978f-d1d67228af87"))
        json.Node("/uuid").IsUUID().NotEqualTo(uuid.FromStringOrNil("a54cbd42-b30c-4619-b89a-47375734d49c"))
        json.Node("/nilUUID").IsUUID().Nil()
        json.Node("/email").IsEmail()
        json.Node("/email").IsHTML5Email()
        json.Node("/url").IsURL().WithSchemas("https").WithHosts("example.com")
        json.Node("/jwt").
            IsJWT(func(token *jwt.Token) (interface{}, error) {
                return []byte("your-256-bit-secret"), nil
            }).
            Algorithm("HS256").
            Header(func(json *assertjson.AssertJSON) {
                json.Node("/alg").IsString().EqualTo("HS256")
                json.Node("/typ").IsString().EqualTo("JWT")
            }).
            Payload(func(json *assertjson.AssertJSON) {
                json.Node("/name").IsString().EqualTo("John Doe")
            })

        // time assertions
        json.Node("/time").IsTime().EqualTo(time.Date(2022, time.October, 16, 12, 14, 32, 0, time.UTC))
        json.Node("/time").IsTime().NotEqualTo(time.Date(2021, time.October, 16, 12, 14, 32, 0, time.UTC))
        json.Node("/time").IsTime().AfterOrEqualTo(time.Date(2022, time.October, 16, 12, 14, 32, 0, time.UTC))
        json.Node("/time").IsTime().After(time.Date(2021, time.October, 16, 12, 14, 32, 0, time.UTC))
        json.Node("/time").IsTime().Before(time.Date(2023, time.October, 16, 12, 14, 32, 0, time.UTC))
        json.Node("/time").IsTime().BeforeOrEqualTo(time.Date(2022, time.October, 16, 12, 14, 32, 0, time.UTC))
        json.Node("/date").IsDate().EqualToDate(2022, time.October, 16)
        json.Node("/date").IsDate().NotEqualToDate(2021, time.October, 16)
        json.Node("/date").IsDate().AfterDate(2021, time.October, 16)
        json.Node("/date").IsDate().AfterOrEqualToDate(2022, time.October, 16)
        json.Node("/date").IsDate().BeforeDate(2023, time.October, 16)
        json.Node("/date").IsDate().BeforeOrEqualToDate(2022, time.October, 16)

        // array assertions
        json.Node("/arrayNode").IsArray()
        json.Node("/arrayNode").IsArray().WithLength(1)
        json.Node("/arrayNode").IsArray().WithLengthGreaterThan(0)
        json.Node("/arrayNode").IsArray().WithLengthGreaterThanOrEqual(1)
        json.Node("/arrayNode").IsArray().WithLengthLessThan(2)
        json.Node("/arrayNode").IsArray().WithLengthLessThanOrEqual(1)
        json.Node("/arrayNode").IsArray().WithUniqueElements()
        json.Node("/arrayNode").ForEach(func(node *assertjson.AssertNode) {
            node.IsString().EqualTo("arrayValue")
        })

        // object assertions
        json.Node("/objectNode").IsObject()
        json.Node("/objectNode").IsObject().WithPropertiesCount(1)
        json.Node("/objectNode").IsObject().WithPropertiesCountGreaterThan(0)
        json.Node("/objectNode").IsObject().WithPropertiesCountGreaterThanOrEqual(1)
        json.Node("/objectNode").IsObject().WithPropertiesCountLessThan(2)
        json.Node("/objectNode").IsObject().WithPropertiesCountLessThanOrEqual(1)
        json.Node("/objectNode").IsObject().WithUniqueElements()
        json.Node("/objectNode").ForEach(func(node *assertjson.AssertNode) {
            node.IsString().EqualTo("objectValue")
        })

        // json pointer expression
        json.Node("/complexNode/items/1/key").IsString().EqualTo("value")
        json.Nodef("/complexNode/items/%d/key", 1).IsString().EqualTo("value")

        // complex keys
        json.Node("/@id").IsString().EqualTo("json-ld-id")
        json.Node("/hydra:members").IsString().EqualTo("hydraMembers")

        // complex assertions
        json.At("/complexNode").Node("/items/1/key").IsString().EqualTo("value")
        json.Atf("/complexNode/%s", "items").Node("/1/key").IsString().EqualTo("value")

        // get node values
        assert.Equal(t, "stringValue", json.Node("/stringNode").Value())
        assert.Equal(t, "stringValue", json.Node("/stringNode").String())
        assert.Equal(t, "123", json.Node("/integerNode").String())
        assert.Equal(t, "123.123000", json.Node("/floatNode").String())
        assert.Equal(t, 123.0, json.Node("/integerNode").Float())
        assert.Equal(t, 123.123, json.Node("/floatNode").Float())
        assert.Equal(t, 123, json.Node("/integerNode").Integer())
        assert.Equal(t, 1, json.Node("/arrayNode").IsArray().Length())
        assert.Equal(t, 1, json.Node("/arrayNode").ArrayLength())
        assert.Equal(t, 1, json.Node("/objectNode").IsObject().PropertiesCount())
        assert.Equal(t, 1, json.Node("/objectNode").ObjectPropertiesCount())
        assert.JSONEq(t, `{"objectKey": "objectValue"}`, string(json.Node("/objectNode").JSON()))
        assert.Equal(t, "2022-10-16T15:14:32+03:00", json.Node("/time").Time().Format(time.RFC3339))
        assert.Equal(t, "23e98a0c-26c8-410f-978f-d1d67228af87", json.Node("/uuid").IsUUID().Value().String())
        assert.Equal(t, "23e98a0c-26c8-410f-978f-d1d67228af87", json.Node("/uuid").UUID().String())
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
