# API testing tools for Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/muonsoft/api-testing.svg)](https://pkg.go.dev/github.com/muonsoft/api-testing)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/muonsoft/api-testing)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/muonsoft/api-testing)
![GitHub](https://img.shields.io/github/license/muonsoft/api-testing)
[![Go Report Card](https://goreportcard.com/badge/github.com/muonsoft/api-testing)](https://goreportcard.com/report/github.com/muonsoft/api-testing)
![CI](https://github.com/muonsoft/api-testing/workflows/CI/badge.svg?branch=master)

## Contents

- [Installation](#installation)
- [apitest package](#apitest-package) — HTTP handler testing
- [assertjson package](#assertjson-package) — JSON assertions
- [assertxml package](#assertxml-package) — XML assertions
- [Examples](EXAMPLES.md) — подробные примеры по темам

## Installation

```bash
go get github.com/muonsoft/api-testing
```

---

## `apitest` package

The `apitest` package provides methods for testing client-server communication.
It can be used to test `http.Handler` and build assertions on HTTP responses (status, headers, body).

### Example

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
        json.Node("ok").IsTrue()
    })
    response.Print() // prints response with headers and body
    response.PrintJSON() // prints response with headers and indented JSON body
}
```

---

## `assertjson` package

The `assertjson` package provides fluent assertions for JSON values. Nodes are selected via [JSON Pointer](https://tools.ietf.org/html/rfc6901) or path elements. Supports strings, numbers, arrays, objects, UUID, email, URL, time, and JWT.

### Example

```go
package yours

import (
    "net/http/httptest"
    "testing"

    "github.com/muonsoft/api-testing/assertjson"
)

func TestYourAPI(t *testing.T) {
    recorder := httptest.NewRecorder()
    handler := newHTTPHandler()

    request := httptest.NewRequest("GET", "/content", nil)
    handler.ServeHTTP(recorder, request)

    assertjson.Has(t, recorder.Body.Bytes(), func(json *assertjson.AssertJSON) {
        json.Node("nullNode").Exists()
        json.Node("notExistingNode").DoesNotExist()
        json.Node("stringNode").IsString().EqualTo("stringValue")
        json.Node("integerNode").IsInteger().EqualTo(123)
        json.Node("objectNode").EqualJSON(`{"objectKey": "objectValue"}`)
        json.Node("bookstore", "books", 1, "name").IsString().EqualTo("Green book")
        json.Node("bookstore", "books", 1).Print() // debug helper
    })
}
```

**JSON Lines (NDJSON):** use `LinesHas` or `Lines(t, data).Has(...)` to assert over multiple JSON objects, one per line:

```go
assertjson.LinesHas(t, body, func(lines *assertjson.AssertJSONLines) {
    lines.At(0).Node("id").EqualToTheInteger(1)
    lines.At(0).Node("name").EqualToTheString("Alice")
    lines.At(1).Node("role").EqualToTheString("admin")
    lines.WithLength(2)
})
```

Подробные примеры по темам (строки, числа, массивы, объекты, UUID, время, JWT и др.) — в [EXAMPLES.md](EXAMPLES.md). Полный API — [pkg.go.dev/github.com/muonsoft/api-testing/assertjson](https://pkg.go.dev/github.com/muonsoft/api-testing/assertjson).

---

## `assertxml` package

The `assertxml` package provides methods for testing XML values. Nodes are selected via XML Path syntax.

### Example

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

    assertxml.Has(t, recorder.Body.Bytes(), func(xml *assertxml.AssertXML) {
        // common assertions
        xml.Node("/root/stringNode").Exists()
        xml.Node("/root/notExistingNode").DoesNotExist()
  
        // string assertions
        xml.Node("/root/stringNode").EqualToTheString("stringValue")
    })
}
```
