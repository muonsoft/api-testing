# Examples

Подробные примеры использования пакетов по темам.

## Содержание

- [apitest](#apitest)
  - [GET-запрос и проверка JSON](#apitest---get-запрос-и-проверка-json)
- [assertjson](#assertjson)
  - [Базовые проверки](#assertjson---базовые-проверки)
  - [Строки](#assertjson---строки)
  - [Числа](#assertjson---числа)
  - [UUID, email, URL](#assertjson---uuid-email-url)
  - [Время и даты](#assertjson---время-и-даты)
  - [Массивы и объекты](#assertjson---массивы-и-объекты)
  - [Пути и переиспользуемые проверки](#assertjson---пути-и-переиспользуемые-проверки)
  - [JWT](#assertjson---jwt)
  - [JSON Lines (NDJSON)](#assertjson---json-lines-ndjson)
  - [Получение значений и отладка](#assertjson---получение-значений-и-отладка)
- [assertxml](#assertxml)
  - [Базовые проверки XML](#assertxml---базовые-проверки-xml)

---

## apitest

### apitest — GET-запрос и проверка JSON

```go
package yours

import (
    "net/http"
    "testing"

    "github.com/muonsoft/api-testing/apitest"
    "github.com/muonsoft/api-testing/assertjson"
)

func TestYourAPI(t *testing.T) {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"ok":true}`))
    })

    response := apitest.HandleGET(t, handler, "/example")

    response.IsOK()
    response.HasContentType("application/json")
    response.HasJSON(func(json *assertjson.AssertJSON) {
        json.Node("ok").IsTrue()
    })
    response.Print()
    response.PrintJSON()
}
```

---

## assertjson

Примеры предполагают, что JSON загружен из `recorder.Body.Bytes()` или файла через `assertjson.Has(t, data, ...)` / `assertjson.FileHas(t, filename, ...)`.

### assertjson — Базовые проверки

```go
assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    json.Node("nullNode").Exists()
    json.Node("notExistingNode").DoesNotExist()
    json.Node("nullNode").IsNull()
    json.Node("stringNode").IsNotNull()
    json.Node("trueBooleanNode").IsTrue()
    json.Node("falseBooleanNode").IsFalse()
    json.Node("objectNode").EqualJSON(`{"objectKey": "objectValue"}`)
})
```

### assertjson — Строки

```go
assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    json.Node("stringNode").IsString()
    json.Node("stringNode").Matches("^string.*$")
    json.Node("stringNode").DoesNotMatch("^notMatch$")
    json.Node("stringNode").Contains("string")
    json.Node("stringNode").DoesNotContain("notContain")

    // fluent-цепочки
    json.Node("emptyString").IsString().IsEmpty()
    json.Node("stringNode").IsString().IsNotEmpty()
    json.Node("stringNode").IsString().EqualTo("stringValue")
    json.Node("stringNode").IsString().EqualToOneOf("stringValue", "nextValue")
    json.Node("stringNode").IsString().NotEqualTo("invalid")
    json.Node("stringNode").IsString().WithLength(11)
    json.Node("stringNode").IsString().WithLengthGreaterThan(10)
    json.Node("stringNode").IsString().WithLengthLessThan(12)
    json.Node("stringNode").IsString().WithPrefix("string")
    json.Node("stringNode").IsString().WithSuffix("Value")
    json.Node("idNode").IsString().WithInteger().EqualTo(42)
    json.Node("amountNode").IsString().WithNumber().EqualTo(12.5)
    json.Node("stringNode").IsString().That(func(s string) error {
        if s != "stringValue" {
            return fmt.Errorf("invalid")
        }
        return nil
    })
})
```

### assertjson — Числа

```go
assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    json.Node("integerNode").IsInteger()
    json.Node("zeroInteger").IsInteger().IsZero()
    json.Node("integerNode").IsInteger().IsNotZero()
    json.Node("integerNode").IsInteger().EqualTo(123)
    json.Node("integerNode").IsInteger().GreaterThan(122)
    json.Node("integerNode").IsInteger().LessThanOrEqual(123)

    json.Node("floatNode").IsFloat()
    json.Node("floatNode").IsNumber()
    json.Node("floatNode").IsNumber().EqualTo(123.123)
    json.Node("floatNode").IsNumber().EqualToWithDelta(123.123, 0.1)
    json.Node("floatNode").IsNumber().GreaterThanOrEqual(122).LessThanOrEqual(124)
})
```

### assertjson — UUID, email, URL

```go
import "github.com/gofrs/uuid/v5"

assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    json.Node("uuid").IsString().WithUUID()
    json.Node("uuid").IsUUID().IsNotNil().OfVersion(4).OfVariant(1)
    json.Node("uuid").IsUUID().EqualTo(uuid.FromStringOrNil("23e98a0c-26c8-410f-978f-d1d67228af87"))
    json.Node("nilUUID").IsUUID().IsNil()

    json.Node("email").IsEmail()
    json.Node("email").IsHTML5Email()
    json.Node("url").IsURL().WithSchemas("https").WithHosts("example.com")
})
```

### assertjson — Время и даты

```go
import "time"

assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    json.Node("time").IsTime().EqualTo(time.Date(2022, time.October, 16, 12, 14, 32, 0, time.UTC))
    json.Node("time").IsTime().After(time.Date(2021, time.October, 16, 12, 14, 32, 0, time.UTC))
    json.Node("time").IsTime().Before(time.Date(2023, time.October, 16, 12, 14, 32, 0, time.UTC))
    json.Node("time").IsTime().AtDate(2022, time.October, 16)

    json.Node("date").IsDate().EqualToDate(2022, time.October, 16)
    json.Node("date").IsDate().AfterDate(2021, time.October, 16)
    json.Node("date").IsDate().BeforeOrEqualToDate(2022, time.October, 16)
})
```

### assertjson — Массивы и объекты

```go
assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    json.Node("arrayNode").IsArray()
    json.Node("arrayNode").IsArray().WithLength(1)
    json.Node("arrayNode").IsArray().WithLengthGreaterThan(0)
    json.Node("arrayNode").IsArray().WithUniqueElements()
    json.Node("arrayNode").ForEach(func(node *assertjson.AssertNode) {
        node.IsString().EqualTo("arrayValue")
    })

    // массив строк (IsStrings)
    json.Node("arrayNode").IsStrings()
    json.Node("arrayNode").IsStrings().EqualTo("arrayValue")
    json.Node("arrayNode").IsStrings().WithUniqueValues()
    json.Node("arrayNode").IsStrings().Contains("arrayValue")
    json.Node("arrayNode").IsStrings().WithLength(1)
    json.Node("arrayNode").IsStrings().WithLengthGreaterThan(0)
    json.Node("arrayNode").IsStrings().WithLengthLessThan(2)
    json.Node("arrayNode").IsStrings().That(func(values []string) error {
        if len(values) == 0 {
            return fmt.Errorf("expected non-empty array")
        }
        return nil
    })
    json.Node("arrayNode").IsStrings().Assert(func(tb testing.TB, values []string) {
        assert.NotEmpty(tb, values)
    })

    json.Node("objectNode").IsObject()
    json.Node("objectNode").IsObject().WithPropertiesCount(1)
    json.Node("objectNode").IsObject().WithPropertiesCountGreaterThan(0)
    json.Node("objectNode").ForEach(func(node *assertjson.AssertNode) {
        node.IsString().EqualTo("objectValue")
    })
})
```

### assertjson — Пути и переиспользуемые проверки

```go
import "github.com/gofrs/uuid/v5"

assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    // путь по элементам
    json.Node("bookstore", "books", 1, "name").IsString().EqualTo("Green book")

    // fmt.Stringer в пути
    id := uuid.FromStringOrNil("9b1100ea-986b-446b-ae7e-0c8ce7196c26")
    json.Node("hashmap", id, "key").IsString().EqualTo("value")

    // сложные ключи (JSON-LD, Hydra и т.п.)
    json.Node("@id").IsString().EqualTo("json-ld-id")
    json.Node("hydra:members").IsString().EqualTo("hydraMembers")

    // переиспользуемая проверка
    isGreenBook := func(json *assertjson.AssertJSON) {
        json.Node("id").IsInteger().EqualTo(123)
        json.Node("name").IsString().EqualTo("Green book")
    }
    json.Node("bookstore", "books", 1).Assert(isGreenBook)
    json.Node("bookstore", "bestBook").Assert(isGreenBook)
    isGreenBook(json.At("bookstore", "books", 1))
    isGreenBook(json.At("bookstore", "bestBook"))
})
```

### assertjson — JWT

```go
import (
    "time"
    "github.com/muonsoft/api-testing/jwt"
)

assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    isJWT := json.Node("jwt").IsJWT(func(token *jwt.Token) (interface{}, error) {
        return []byte("your-256-bit-secret"), nil
    })
    isJWT.
        WithAlgorithm("HS256").
        WithID("abc12345").
        WithIssuer("https://issuer.example.com").
        WithSubject("https://subject.example.com").
        WithAudience([]string{"https://audience1.example.com", "https://audience2.example.com"}).
        WithHeader(func(json *assertjson.AssertJSON) {
            json.Node("alg").IsString().EqualTo("HS256")
            json.Node("typ").IsString().EqualTo("JWT")
        }).
        WithPayload(func(json *assertjson.AssertJSON) {
            json.Node("name").IsString().EqualTo("John Doe")
        })
    isJWT.WithExpiresAt().AfterDate(2022, time.October, 26)
    isJWT.WithNotBefore().BeforeDate(2022, time.October, 27)
    isJWT.WithIssuedAt().BeforeDate(2022, time.October, 27)
})

// Проверка JWT из строки (без контекста JSON)
assertjson.IsJWT(t, rawJWTString, keyFunc).WithPayload(func(json *assertjson.AssertJSON) {
    json.Node("name").IsString().EqualTo("John Doe")
})
```

### assertjson — JSON Lines (NDJSON)

Данные в формате JSON Lines (одна строка JSON на строку текста) проверяются через `LinesHas` или `Lines(t, data).Has(...)`. Пустые строки пропускаются.

```go
// Пакетная функция (аналог Has для одного JSON)
assertjson.LinesHas(t, body, func(lines *assertjson.AssertJSONLines) {
    lines.At(0).Node("id").EqualToTheInteger(1)
    lines.At(0).Node("name").EqualToTheString("Alice")
    lines.At(1).Node("id").EqualToTheInteger(2)
    lines.At(1).Node("role").EqualToTheString("admin")
    lines.WithLength(2)
})

// Метод на результате Lines
assertjson.Lines(t, body).Has(func(lines *assertjson.AssertJSONLines) {
    lines.At(0).Node("event").EqualToTheString("created")
    lines.At(1).Node("event").EqualToTheString("updated")
})

// Файл
assertjson.FileLinesHas(t, "events.ndjson", func(lines *assertjson.AssertJSONLines) {
    lines.At(0).Node("event").EqualToTheString("created")
})
```

### assertjson — Получение значений и отладка

```go
import (
    "time"
    "github.com/stretchr/testify/assert"
)

assertjson.Has(t, data, func(json *assertjson.AssertJSON) {
    _ = json.Node("stringNode").Value()
    _ = json.Node("stringNode").String()
    _ = json.Node("integerNode").Integer()
    _ = json.Node("integerNode").Float()
    _ = json.Node("floatNode").Float()
    _ = json.Node("arrayNode").IsArray().Length()
    _ = json.Node("objectNode").IsObject().PropertiesCount()
    _ = json.Node("objectNode").JSON()
    _ = json.Node("time").Time().Format(time.RFC3339)
    _ = json.Node("uuid").IsUUID().Value().String()

    assert.Equal(t, "stringValue", json.Node("stringNode").String())
    assert.Equal(t, 123, json.Node("integerNode").Integer())

    // вывод узла в тестовый лог (отладка)
    json.Node("bookstore", "books", 1).Print()
})
```

---

## assertxml

### assertxml — Базовые проверки XML

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
        xml.Node("/root/stringNode").Exists()
        xml.Node("/root/notExistingNode").DoesNotExist()
        xml.Node("/root/stringNode").EqualToTheString("stringValue")
    })
}
```
