package assertjson_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/muonsoft/api-testing/assertjson"
	"github.com/muonsoft/api-testing/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestFileHas(t *testing.T) {
	assertjson.FileHas(t, "./../test/testdata/object.json", func(json *assertjson.AssertJSON) {
		// common assertions
		json.Node("nullNode").Exists()
		json.Node("notExistingNode").DoesNotExist()
		json.Node("nullNode").IsNull()
		json.Node("stringNode").IsNotNull()
		json.Node("trueBooleanNode").IsTrue()
		json.Node("falseBooleanNode").IsFalse()
		json.Node("objectNode").EqualJSON(`{"objectKey": "objectValue"}`)

		// string assertions
		json.Node("stringNode").IsString()
		json.Node("stringNode").EqualToTheString("stringValue")
		json.Node("stringNode").AssertString(func(tb testing.TB, value string) {
			tb.Helper()
			assert.Equal(tb, "stringValue", value)
		})
		json.Node("stringNode").Matches("^string.*$")
		json.Node("stringNode").DoesNotMatch("^notMatch$")
		json.Node("stringNode").Contains("string")
		json.Node("stringNode").DoesNotContain("notContain")
		json.Node("stringNode").IsStringWithLength(11)
		json.Node("stringNode").IsStringWithLengthInRange(11, 11)

		// fluent string assertions
		json.Node("stringNode").IsString()
		json.Node("stringNode").IsString().EqualTo("stringValue")
		json.Node("stringNode").IsString().EqualToOneOf("stringValue", "nextValue")
		json.Node("stringNode").IsString().NotEqualTo("invalid")
		json.Node("stringNode").IsString().Matches("^string.*$")
		json.Node("stringNode").IsString().NotMatches("^notMatch$")
		json.Node("stringNode").IsString().Contains("string")
		json.Node("stringNode").IsString().NotContains("notContain")
		json.Node("stringNode").IsString().WithLength(11)
		json.Node("stringNode").IsString().WithLengthGreaterThan(10)
		json.Node("stringNode").IsString().WithLengthGreaterThanOrEqual(11)
		json.Node("stringNode").IsString().WithLengthLessThan(12)
		json.Node("stringNode").IsString().WithLengthLessThanOrEqual(11)
		json.Node("stringNode").IsString().That(func(s string) error {
			if s != "stringValue" {
				return fmt.Errorf("invalid")
			}
			return nil
		})
		json.Node("stringNode").IsString().Assert(func(tb testing.TB, value string) {
			tb.Helper()
			assert.Equal(tb, "stringValue", value)
		})

		// numeric assertions
		json.Node("integerNode").IsInteger()
		json.Node("integerNode").IsInteger().EqualTo(123)
		json.Node("integerNode").IsInteger().NotEqualTo(321)
		json.Node("integerNode").IsInteger().GreaterThan(122)
		json.Node("integerNode").IsInteger().GreaterThanOrEqual(123)
		json.Node("integerNode").IsInteger().LessThan(124)
		json.Node("integerNode").IsInteger().LessThanOrEqual(123)
		json.Node("integerNode").EqualToTheInteger(123)
		json.Node("integerNode").IsNumberInRange(122, 124)
		json.Node("integerNode").IsNumberGreaterThan(122)
		json.Node("integerNode").IsNumberGreaterThanOrEqual(123)
		json.Node("integerNode").IsNumberLessThan(124)
		json.Node("integerNode").IsNumberLessThanOrEqual(123)
		json.Node("floatNode").IsFloat()
		json.Node("floatNode").IsNumber()
		json.Node("floatNode").IsNumber().EqualTo(123.123)
		json.Node("floatNode").IsNumber().NotEqualTo(321.123)
		json.Node("floatNode").IsNumber().EqualToWithDelta(123.123, 0.1)
		json.Node("floatNode").IsNumber().GreaterThan(122)
		json.Node("floatNode").IsNumber().GreaterThanOrEqual(123.123)
		json.Node("floatNode").IsNumber().LessThan(124)
		json.Node("floatNode").IsNumber().LessThanOrEqual(123.123)
		json.Node("floatNode").IsNumber().GreaterThanOrEqual(122).LessThanOrEqual(124)
		json.Node("floatNode").EqualToTheFloat(123.123)
		json.Node("floatNode").IsNumberInRange(122, 124)
		json.Node("floatNode").IsNumberGreaterThan(122)
		json.Node("floatNode").IsNumberGreaterThanOrEqual(123.123)
		json.Node("floatNode").IsNumberLessThan(124)
		json.Node("floatNode").IsNumberLessThanOrEqual(123.123)

		// string values assertions
		json.Node("uuid").IsString().WithUUID()
		json.Node("uuid").IsUUID().IsNotNil().OfVersion(4).OfVariant(1)
		json.Node("uuid").IsUUID().EqualTo(uuid.FromStringOrNil("23e98a0c-26c8-410f-978f-d1d67228af87"))
		json.Node("uuid").IsUUID().NotEqualTo(uuid.FromStringOrNil("a54cbd42-b30c-4619-b89a-47375734d49c"))
		json.Node("nilUUID").IsUUID().IsNil()
		json.Node("email").IsEmail()
		json.Node("email").IsHTML5Email()
		json.Node("url").IsURL().WithSchemas("https").WithHosts("example.com")

		// time assertions
		json.Node("time").IsTime().EqualTo(time.Date(2022, time.October, 16, 12, 14, 32, 0, time.UTC))
		json.Node("time").IsTime().NotEqualTo(time.Date(2021, time.October, 16, 12, 14, 32, 0, time.UTC))
		json.Node("time").IsTime().AfterOrEqualTo(time.Date(2022, time.October, 16, 12, 14, 32, 0, time.UTC))
		json.Node("time").IsTime().After(time.Date(2021, time.October, 16, 12, 14, 32, 0, time.UTC))
		json.Node("time").IsTime().Before(time.Date(2023, time.October, 16, 12, 14, 32, 0, time.UTC))
		json.Node("time").IsTime().BeforeOrEqualTo(time.Date(2022, time.October, 16, 12, 14, 32, 0, time.UTC))
		json.Node("time").IsTime().AtDate(2022, time.October, 16)
		json.Node("date").IsDate().EqualToDate(2022, time.October, 16)
		json.Node("date").IsDate().NotEqualToDate(2021, time.October, 16)
		json.Node("date").IsDate().AfterDate(2021, time.October, 16)
		json.Node("date").IsDate().AfterOrEqualToDate(2022, time.October, 16)
		json.Node("date").IsDate().BeforeDate(2023, time.October, 16)
		json.Node("date").IsDate().BeforeOrEqualToDate(2022, time.October, 16)

		// JSON Web Token (JWT) assertion
		isJWT := json.Node("jwt").IsJWT(func(token *jwt.Token) (interface{}, error) {
			return []byte("your-256-bit-secret"), nil
		})
		isJWT.
			WithAlgorithm("HS256").
			// standard claims assertions
			WithID("abc12345").
			WithIssuer("https://issuer.example.com").
			WithSubject("https://subject.example.com").
			WithAudience([]string{"https://audience1.example.com", "https://audience2.example.com"}).
			// json assertion of header part
			WithHeader(func(json *assertjson.AssertJSON) {
				json.Node("alg").IsString().EqualTo("HS256")
				json.Node("typ").IsString().EqualTo("JWT")
			}).
			// json assertion of payload part
			WithPayload(func(json *assertjson.AssertJSON) {
				json.Node("name").IsString().EqualTo("John Doe")
			})
		// time assertions for standard claims
		isJWT.WithExpiresAt().AfterDate(2022, time.October, 26)
		isJWT.WithNotBefore().BeforeDate(2022, time.October, 27)
		isJWT.WithIssuedAt().BeforeDate(2022, time.October, 27)

		// array assertions
		json.Node("arrayNode").IsArrayWithElementsCount(1)
		json.Node("arrayNode").IsArray()
		json.Node("arrayNode").IsArray().WithLength(1)
		json.Node("arrayNode").IsArray().WithLengthGreaterThan(0)
		json.Node("arrayNode").IsArray().WithLengthGreaterThanOrEqual(1)
		json.Node("arrayNode").IsArray().WithLengthLessThan(2)
		json.Node("arrayNode").IsArray().WithLengthLessThanOrEqual(1)
		json.Node("arrayNode").IsArray().WithUniqueElements()
		json.Node("arrayNode").ForEach(func(node *assertjson.AssertNode) {
			node.IsString().EqualTo("arrayValue")
		})

		// object assertions
		json.Node("objectNode").IsObjectWithPropertiesCount(1)
		json.Node("objectNode").IsObject()
		json.Node("objectNode").IsObject().WithPropertiesCount(1)
		json.Node("objectNode").IsObject().WithPropertiesCountGreaterThan(0)
		json.Node("objectNode").IsObject().WithPropertiesCountGreaterThanOrEqual(1)
		json.Node("objectNode").IsObject().WithPropertiesCountLessThan(2)
		json.Node("objectNode").IsObject().WithPropertiesCountLessThanOrEqual(1)
		json.Node("objectNode").IsObject().WithUniqueElements()
		json.Node("objectNode").ForEach(func(node *assertjson.AssertNode) {
			node.IsString().EqualTo("objectValue")
		})

		// seek node by path elements
		json.Node("bookstore", "books", 1, "name").IsString().EqualTo("Green book")

		// use fmt.Stringer in node path
		id := uuid.FromStringOrNil("9b1100ea-986b-446b-ae7e-0c8ce7196c26")
		json.Node("hashmap", id, "key").IsString().EqualTo("value")

		// complex keys
		json.Node("@id").IsString().EqualTo("json-ld-id")
		json.Node("hydra:members").IsString().EqualTo("hydraMembers")

		// reusable assertions
		isGreenBook := func(json *assertjson.AssertJSON) {
			json.Node("id").IsInteger().EqualTo(123)
			json.Node("name").IsString().EqualTo("Green book")
		}
		json.Node("bookstore", "books", 1).Assert(isGreenBook)
		json.Node("bookstore", "bestBook").Assert(isGreenBook)
		isGreenBook(json.At("bookstore", "books", 1))
		isGreenBook(json.At("bookstore", "bestBook"))

		// get node values
		assert.Equal(t, "stringValue", json.Node("stringNode").Value())
		assert.Equal(t, "stringValue", json.Node("stringNode").String())
		assert.Equal(t, "123", json.Node("integerNode").String())
		assert.Equal(t, "123.123000", json.Node("floatNode").String())
		assert.Equal(t, 123.0, json.Node("integerNode").Float())
		assert.Equal(t, 123.123, json.Node("floatNode").Float())
		assert.Equal(t, 123, json.Node("integerNode").Integer())
		assert.Equal(t, 1, json.Node("arrayNode").IsArray().Length())
		assert.Equal(t, 1, json.Node("arrayNode").ArrayLength())
		assert.Equal(t, 1, json.Node("objectNode").IsObject().PropertiesCount())
		assert.Equal(t, 1, json.Node("objectNode").ObjectPropertiesCount())
		assert.JSONEq(t, `{"objectKey": "objectValue"}`, string(json.Node("objectNode").JSON()))
		assert.Equal(t, "2022-10-16T15:14:32+03:00", json.Node("time").Time().Format(time.RFC3339))
		assert.Equal(t, "23e98a0c-26c8-410f-978f-d1d67228af87", json.Node("uuid").IsUUID().Value().String())
		assert.Equal(t, "23e98a0c-26c8-410f-978f-d1d67228af87", json.Node("uuid").UUID().String())
		assert.Equal(t,
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiaHR0cHM6Ly9hdWRpZW5jZTEuZXhhbXBsZS5jb20iLCJodHRwczovL2F1ZGllbmNlMi5leGFtcGxlLmNvbSJdLCJleHAiOjQ4MjAzNjAxMzEsImlhdCI6MTY2Njc1NjUzMSwiaXNzIjoiaHR0cHM6Ly9pc3N1ZXIuZXhhbXBsZS5jb20iLCJqdGkiOiJhYmMxMjM0NSIsIm5hbWUiOiJKb2huIERvZSIsIm5iZiI6MTY2Njc1NjUzMSwic3ViIjoiaHR0cHM6Ly9zdWJqZWN0LmV4YW1wbGUuY29tIn0.fGUvIn-BV8bPKkZdrxUneew3_qBe-knptL9a_TkNA4M",
			json.Node("jwt").
				IsJWT(func(token *jwt.Token) (interface{}, error) {
					return []byte("your-256-bit-secret"), nil
				}).
				Value().
				Raw,
		)
		assert.Equal(t,
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiaHR0cHM6Ly9hdWRpZW5jZTEuZXhhbXBsZS5jb20iLCJodHRwczovL2F1ZGllbmNlMi5leGFtcGxlLmNvbSJdLCJleHAiOjQ4MjAzNjAxMzEsImlhdCI6MTY2Njc1NjUzMSwiaXNzIjoiaHR0cHM6Ly9pc3N1ZXIuZXhhbXBsZS5jb20iLCJqdGkiOiJhYmMxMjM0NSIsIm5hbWUiOiJKb2huIERvZSIsIm5iZiI6MTY2Njc1NjUzMSwic3ViIjoiaHR0cHM6Ly9zdWJqZWN0LmV4YW1wbGUuY29tIn0.fGUvIn-BV8bPKkZdrxUneew3_qBe-knptL9a_TkNA4M",
			json.Node("jwt").
				JWT(func(token *jwt.Token) (interface{}, error) {
					return []byte("your-256-bit-secret"), nil
				}).
				Raw,
		)
	})
}

func TestHas(t *testing.T) {
	time.Local = time.UTC

	tests := []struct {
		name         string
		json         string
		assert       assertjson.JSONAssertFunc
		wantMessages []string
	}{
		{
			name: "invalid JSON",
			json: `{`,
			wantMessages: []string{
				"data has invalid JSON: unexpected end of JSON input",
			},
		},
		{
			name: "JSON node not found",
			json: `{}`,
			assert: func(json *assertjson.AssertJSON) {
				json.At("key")
			},
			wantMessages: []string{
				`failed to find JSON node "key": [key] not found`,
			},
		},
		{
			name: "JSON each array node equal to string",
			json: `{"array": ["value", "value", "value"]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("array").ForEach(func(node *assertjson.AssertNode) {
					node.IsString().EqualTo("value")
				})
			},
		},
		{
			name: "JSON each array node equal to string fails",
			json: `{"array": ["value", "v", "value"]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("array").ForEach(func(node *assertjson.AssertNode) {
					node.IsString().EqualTo("value")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "array[1]": equal to "value", actual is "v"`,
			},
		},
		{
			name: "JSON each object node equal to string",
			json: `{"object": {"a": "value", "b": "value", "c": "value"}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("object").ForEach(func(node *assertjson.AssertNode) {
					node.IsString().EqualTo("value")
				})
			},
		},
		{
			name: "JSON each object node equal to string fails",
			json: `{"object": {"a": "value", "b": "v", "c": "value"}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("object").ForEach(func(node *assertjson.AssertNode) {
					node.IsString().EqualTo("value")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "object.b": equal to "value", actual is "v"`,
			},
		},
		{
			name: "JSON each node on not iterable",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").ForEach(func(node *assertjson.AssertNode) {})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is iterable (array or object)`,
			},
		},
		{
			name: "JSON callable assertion",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").Assert(func(json *assertjson.AssertJSON) {
					json.Node().IsString().EqualTo("value")
				})
			},
		},
		{
			name: "JSON callable assertion fails",
			json: `{"key": "unexpected"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").Assert(func(json *assertjson.AssertJSON) {
					json.Node().IsString().EqualTo("value")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": equal to "value", actual is "unexpected"`,
			},
		},
		{
			name: "JSON node exists",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").Exists()
			},
		},
		{
			name: "JSON node exists fails",
			json: `{}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").Exists()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" exists`,
			},
		},
		{
			name: "JSON node does not exist",
			json: `{}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").DoesNotExist()
			},
		},
		{
			name: "JSON node does not exist fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").DoesNotExist()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" does not exist`,
			},
		},
		{
			name: "JSON node is null",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNull()
			},
		},
		{
			name: "JSON node is null fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNull()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is null`,
			},
		},
		{
			name: "JSON node is not null",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNotNull()
			},
		},
		{
			name: "JSON node is not null fails",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNotNull()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is not null`,
			},
		},
		{
			name: "JSON node is true",
			json: `{"key": true}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTrue()
			},
		},
		{
			name: "JSON node is true fails",
			json: `{"key": false}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTrue()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is true`,
			},
		},
		{
			name: "JSON node is true on not boolean",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTrue()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is boolean`,
			},
		},
		{
			name: "JSON node is false",
			json: `{"key": false}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsFalse()
			},
		},
		{
			name: "JSON node is false fails",
			json: `{"key": true}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsFalse()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is false`,
			},
		},
		{
			name: "JSON node is false on not boolean",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsFalse()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is boolean`,
			},
		},
		{
			name: "JSON node equal to JSON",
			json: `{"key": {"key": "value"}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").EqualJSON(`{"key": "value"}`)
			},
		},
		{
			name: "JSON node equal to JSON fails",
			json: `{"key": {"k": "v"}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").EqualJSON(`{"key": "value"}`)
			},
			wantMessages: []string{
				"Not equal",
				`failed at JSON node "key"`,
			},
		},
		{
			name: "JSON node is number",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber()
			},
		},
		{
			name: "JSON node is number fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber()
			},
			wantMessages: []string{
				`value at path "key" is not a number`,
			},
		},
		{
			name: "JSON node is number equal to",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().EqualTo(123.123)
			},
		},
		{
			name: "JSON node is number equal to fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().EqualTo(321.123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": equal to 321.123000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number not equal to",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().NotEqualTo(321.123)
			},
		},
		{
			name: "JSON node is number not equal to fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().NotEqualTo(123.123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": not equal to 123.123000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number equal to with delta",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().EqualToWithDelta(123.1, 1)
			},
		},
		{
			name: "JSON node is number equal to with delta fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().EqualToWithDelta(321.123, 1)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": equal to 321.123000 with delta 1.000000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number greater than",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().GreaterThan(123.122)
			},
		},
		{
			name: "JSON node is number greater than fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().GreaterThan(123.123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": greater than 123.123000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number greater than or equal",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().GreaterThanOrEqual(123.123)
			},
		},
		{
			name: "JSON node is number greater than or equal fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().GreaterThanOrEqual(123.124)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": greater than or equal 123.124000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number less than",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().LessThan(123.124)
			},
		},
		{
			name: "JSON node is number less than fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().LessThan(123.123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": less than 123.123000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number less than or equal",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().LessThanOrEqual(123.123)
			},
		},
		{
			name: "JSON node is number less than or equal fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsNumber().LessThanOrEqual(123.122)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": less than or equal 123.122000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsNumber().
					EqualTo(0).
					EqualToWithDelta(0, 0).
					NotEqualTo(0).
					GreaterThan(0).
					GreaterThanOrEqual(0).
					LessThan(0).
					LessThanOrEqual(0)
			},
			wantMessages: []string{
				`value at path "key" is not a number`,
			},
		},
		{
			name: "JSON node is integer",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger()
			},
		},
		{
			name: "JSON node is integer fails on string",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger()
			},
			wantMessages: []string{
				`value at path "key" is not numeric`,
			},
		},
		{
			name: "JSON node is integer fails on float",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger()
			},
			wantMessages: []string{
				`value at path "key" is float, not integer`,
			},
		},
		{
			name: "JSON node is integer equal to",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().EqualTo(123)
			},
		},
		{
			name: "JSON node is integer equal to fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().EqualTo(321)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": equal to 321, actual is 123`,
			},
		},
		{
			name: "JSON node is integer not equal to",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().NotEqualTo(321)
			},
		},
		{
			name: "JSON node is integer not equal to fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().NotEqualTo(123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": not equal to 123, actual is 123`,
			},
		},
		{
			name: "JSON node is integer greater than",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().GreaterThan(122)
			},
		},
		{
			name: "JSON node is integer greater than fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().GreaterThan(123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": greater than 123, actual is 123`,
			},
		},
		{
			name: "JSON node is integer greater than or equal",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().GreaterThanOrEqual(123)
			},
		},
		{
			name: "JSON node is integer greater than or equal fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().GreaterThanOrEqual(124)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": greater than or equal 124, actual is 123`,
			},
		},
		{
			name: "JSON node is integer less than",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().LessThan(124)
			},
		},
		{
			name: "JSON node is integer less than fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().LessThan(123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": less than 123, actual is 123`,
			},
		},
		{
			name: "JSON node is integer less than or equal",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().LessThanOrEqual(123)
			},
		},
		{
			name: "JSON node is integer less than or equal fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsInteger().LessThanOrEqual(122)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": less than or equal 122, actual is 123`,
			},
		},
		{
			name: "JSON node is integer fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsInteger().
					EqualTo(0).
					NotEqualTo(0).
					GreaterThan(0).
					GreaterThanOrEqual(0).
					LessThan(0).
					LessThanOrEqual(0)
			},
			wantMessages: []string{
				`value at path "key" is not numeric`,
			},
		},
		{
			name: "JSON node is string",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString()
			},
		},
		{
			name: "JSON node is string fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is string`,
			},
		},
		{
			name: "JSON node is string equal to",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().EqualTo("value")
			},
		},
		{
			name: "JSON node is string equal to fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().EqualTo("string")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": equal to "string", actual is "value"`,
			},
		},
		{
			name: "JSON node is string not equal to",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotEqualTo("string")
			},
		},
		{
			name: "JSON node is string not equal to fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotEqualTo("value")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": not equal to "value", actual is "value"`,
			},
		},
		{
			name: "JSON node is string equal to one of",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().EqualToOneOf("value")
			},
		},
		{
			name: "JSON node is string equal to one of fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().EqualToOneOf("foo", "bar")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": equal to one of values ("foo", "bar"), actual is "value"`,
			},
		},
		{
			name: "JSON node is string matches",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().Matches(".*")
			},
		},
		{
			name: "JSON node is string matches fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().Matches("\\d+")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": matches "\d+", actual is "value"`,
			},
		},
		{
			name: "JSON node is string matches compiled regexp",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().Matches(regexp.MustCompile(".*"))
			},
		},
		{
			name: "JSON node is string matches compiled regexp fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().Matches(regexp.MustCompile(`\d+`))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": matches "\d+", actual is "value"`,
			},
		},
		{
			name: "JSON node is string not matches",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotMatches("\\d+")
			},
		},
		{
			name: "JSON node is string not matches fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotMatches(".*")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": not matches ".*", actual is "value"`,
			},
		},
		{
			name: "JSON node is string not matches compiled regexp",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotMatches(regexp.MustCompile(`\d+`))
			},
		},
		{
			name: "JSON node is string not matches compiled regexp fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotMatches(regexp.MustCompile(".*"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": not matches ".*", actual is "value"`,
			},
		},
		{
			name: "JSON node is string contains",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().Contains("alu")
			},
		},
		{
			name: "JSON node is string contains fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().Contains("string")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": contains "string", actual is "value"`,
			},
		},
		{
			name: "JSON node is string not contains",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotContains("string")
			},
		},
		{
			name: "JSON node is string not contains fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().NotContains("alu")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": not contains "alu", actual is "value"`,
			},
		},
		{
			name: "JSON node is string with length",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLength(5)
			},
		},
		{
			name: "JSON node is string with length fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLength(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with length is 4, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length greater than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthGreaterThan(4)
			},
		},
		{
			name: "JSON node is string with length greater than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthGreaterThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with length greater than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length greater than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthGreaterThanOrEqual(5)
			},
		},
		{
			name: "JSON node is string with length greater than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthGreaterThanOrEqual(6)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with length greater than or equal to 6, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length less than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthLessThan(6)
			},
		},
		{
			name: "JSON node is string with length less than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthLessThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with length less than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length less than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthLessThanOrEqual(5)
			},
		},
		{
			name: "JSON node is string with length less than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithLengthLessThanOrEqual(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with length less than or equal to 4, actual is 5`,
			},
		},
		{
			name: "JSON node is string checked by custom function",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().That(func(s string) error {
					return nil
				})
			},
		},
		{
			name: "JSON node is string checked by custom function fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().That(func(s string) error {
					return fmt.Errorf("error")
				})
			},
			wantMessages: []string{
				`failed asserting JSON node "key": error`,
			},
		},
		{
			name: "JSON node is string fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsString().
					EqualTo("").
					NotEqualTo("").
					EqualToOneOf("").
					Matches(".*").
					NotMatches(".*").
					Contains("").
					NotContains("").
					WithLength(0).
					WithLengthGreaterThan(0).
					WithLengthGreaterThanOrEqual(0).
					WithLengthLessThan(0).
					WithLengthLessThanOrEqual(0).
					WithEmail().
					WithHTML5Email().
					That(func(s string) error { return nil }).
					Assert(func(tb testing.TB, value string) { tb.Helper() })
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is string`,
			},
		},
		{
			name: "JSON node is array",
			json: `{"key": [1, 2, 3]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray()
			},
		},
		{
			name: "JSON node is array fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is array`,
			},
		},
		{
			name: "JSON node is array with length",
			json: `{"key": [1, 2, 3]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLength(3)
			},
		},
		{
			name: "JSON node is array with length fails",
			json: `{"key": [1, 2, 3]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLength(2)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is array with length is 2, actual is 3`,
			},
		},
		{
			name: "JSON node is array with length greater than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthGreaterThan(4)
			},
		},
		{
			name: "JSON node is array with length greater than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthGreaterThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is array with length greater than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is array with length greater than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthGreaterThanOrEqual(5)
			},
		},
		{
			name: "JSON node is array with length greater than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthGreaterThanOrEqual(6)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is array with length greater than or equal to 6, actual is 5`,
			},
		},
		{
			name: "JSON node is array with length less than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthLessThan(6)
			},
		},
		{
			name: "JSON node is array with length less than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthLessThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is array with length less than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is array with length less than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthLessThanOrEqual(5)
			},
		},
		{
			name: "JSON node is array with length less than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithLengthLessThanOrEqual(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is array with length less than or equal to 4, actual is 5`,
			},
		},
		{
			name: "JSON node is array with unique elements",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithUniqueElements()
			},
		},
		{
			name: "JSON node is array with unique elements fails",
			json: `{"key": [3, 2, 3, 4, 2]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"key\" is array with unique elements, duplicated elements",
			},
		},
		{
			name: "JSON node is array with unique elements fails on objects",
			json: `{
				"key": [
					{"a": "a"},
					{"b": "a"},
					{"a": "a"},
					{"a": "b"}
				]
			}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsArray().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"key\" is array with unique elements, duplicated elements",
			},
		},
		{
			name: "JSON node is array fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsArray().
					WithLength(0).
					WithLengthGreaterThan(0).
					WithLengthGreaterThanOrEqual(0).
					WithLengthLessThan(0).
					WithLengthLessThanOrEqual(0).
					WithUniqueElements().
					Length()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is array`,
			},
		},
		{
			name: "JSON node is object",
			json: `{"key": {"a": 1}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject()
			},
		},
		{
			name: "JSON node is object fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is object`,
			},
		},
		{
			name: "JSON node is object with properties count",
			json: `{"key": {"a": 1, "b": 2, "c": 3}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCount(3)
			},
		},
		{
			name: "JSON node is object with properties count fails",
			json: `{"key": {"a": 1, "b": 2, "c": 3}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCount(2)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is object with properties count is 2, actual is 3`,
			},
		},
		{
			name: "JSON node is object with properties count greater than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountGreaterThan(4)
			},
		},
		{
			name: "JSON node is object with properties count greater than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountGreaterThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is object with properties count greater than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is object with properties count greater than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountGreaterThanOrEqual(5)
			},
		},
		{
			name: "JSON node is object with properties count greater than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountGreaterThanOrEqual(6)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is object with properties count greater than or equal to 6, actual is 5`,
			},
		},
		{
			name: "JSON node is object with properties count less than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountLessThan(6)
			},
		},
		{
			name: "JSON node is object with properties count less than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountLessThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is object with properties count less than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is object with properties count less than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountLessThanOrEqual(5)
			},
		},
		{
			name: "JSON node is object with properties count less than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithPropertiesCountLessThanOrEqual(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is object with properties count less than or equal to 4, actual is 5`,
			},
		},
		{
			name: "JSON node is object with unique elements",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithUniqueElements()
			},
		},
		{
			name: "JSON node is object with unique elements fails",
			json: `{"key": {"a": 3, "b": 2, "c": 3, "d": 4, "e": 2}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"key\": is object with unique elements, duplicated elements",
			},
		},
		{
			name: "JSON node is object with unique elements fails on objects",
			json: `{
				"key": {
					"A": {"a": "a"},
					"B": {"b": "a"},
					"C": {"a": "a"},
					"D": {"a": "b"}
				}
			}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsObject().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"key\": is object with unique elements, duplicated elements",
			},
		},
		{
			name: "JSON node is object fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsObject().
					WithPropertiesCount(0).
					WithPropertiesCountGreaterThan(0).
					WithPropertiesCountGreaterThanOrEqual(0).
					WithPropertiesCountLessThan(0).
					WithPropertiesCountLessThanOrEqual(0).
					WithUniqueElements().
					PropertiesCount()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is object`,
			},
		},
		{
			name: "JSON node is UUID",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID()
			},
		},
		{
			name: "JSON node is UUID fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is UUID, actual is "value"`,
			},
		},
		{
			name: "JSON node is nil UUID",
			json: `{"key": "00000000-0000-0000-0000-000000000000"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().Nil()
			},
		},
		{
			name: "JSON node is nil UUID fails",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().Nil()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is nil UUID, actual is "bf0d10a1-d74c-436a-9db1-77c23b5e464f"`,
			},
		},
		{
			name: "JSON node is not nil UUID",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().NotNil()
			},
		},
		{
			name: "JSON node is not nil UUID fails",
			json: `{"key": "00000000-0000-0000-0000-000000000000"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().NotNil()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is not nil UUID, actual is "00000000-0000-0000-0000-000000000000"`,
			},
		},
		{
			name: "JSON node is UUID v4",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().Version(4)
			},
		},
		{
			name: "JSON node is UUID v4 fails",
			json: `{"key": "00000000-0000-0000-0000-000000000000"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().Version(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is UUID of version 4, actual is 0`,
			},
		},
		{
			name: "JSON node is UUID variant 1",
			json: `{"key": "a67e4bfc-1039-11ed-861d-0242ac120002"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().Variant(1)
			},
		},
		{
			name: "JSON node is UUID variant 1 fails",
			json: `{"key": "00000000-0000-0000-0000-000000000000"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().Variant(1)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is UUID of variant 1, actual is 0`,
			},
		},
		{
			name: "JSON node is UUID equal to",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().EqualTo(uuid.FromStringOrNil("bf0d10a1-d74c-436a-9db1-77c23b5e464f"))
			},
		},
		{
			name: "JSON node is UUID equal to fails",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().EqualTo(uuid.FromStringOrNil("01fb115c-0fdc-4072-b5ae-c517689d670c"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is UUID equal to "01fb115c-0fdc-4072-b5ae-c517689d670c", actual is "bf0d10a1-d74c-436a-9db1-77c23b5e464f"`,
			},
		},
		{
			name: "JSON node is UUID not equal to",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().NotEqualTo(uuid.FromStringOrNil("01fb115c-0fdc-4072-b5ae-c517689d670c"))
			},
		},
		{
			name: "JSON node is UUID not equal to fails",
			json: `{"key": "bf0d10a1-d74c-436a-9db1-77c23b5e464f"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsUUID().NotEqualTo(uuid.FromStringOrNil("bf0d10a1-d74c-436a-9db1-77c23b5e464f"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is UUID not equal to "bf0d10a1-d74c-436a-9db1-77c23b5e464f", actual is "bf0d10a1-d74c-436a-9db1-77c23b5e464f"`,
			},
		},
		{
			name: "JSON node is UUID fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsUUID().
					Nil().
					NotNil().
					Version(0).
					Variant(0).
					EqualTo(uuid.Nil).
					NotEqualTo(uuid.FromStringOrNil("bf0d10a1-d74c-436a-9db1-77c23b5e464f")).
					Value()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is string`,
			},
		},
		{
			name: "JSON node is email",
			json: `{"key": "user@example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsEmail()
			},
		},
		{
			name: "JSON node is email fails",
			json: `{"key": "user @ example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsEmail()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is email, actual is "user @ example.com"`,
			},
		},
		{
			name: "JSON node is HTML5 email",
			json: `{"key": "user@example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsHTML5Email()
			},
		},
		{
			name: "JSON node is HTML5 email fails",
			json: `{"key": "user @ example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsHTML5Email()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is email (HTML5 format), actual is "user @ example.com"`,
			},
		},
		{
			name: "JSON node is URL",
			json: `{"key": "https://example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL()
			},
		},
		{
			name: "JSON node is URL fails",
			json: `{"key": "invalid\\:"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is URL, actual is "invalid\:"`,
			},
		},
		{
			name: "JSON node is URL validation fails",
			json: `{"key": "http://example.com/exploit.html?not_a%hex"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is URL, actual is "http://example.com/exploit.html?not_a%hex"`,
			},
		},
		{
			name: "JSON node is URL with schemas",
			json: `{"key": "https://example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL().WithSchemas("http", "https")
			},
		},
		{
			name: "JSON node is URL with schemas fails",
			json: `{"key": "ftp://example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL().WithSchemas("http", "https")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is URL with schemas "http", "https", actual is "ftp"`,
			},
		},
		{
			name: "JSON node is URL with hosts",
			json: `{"key": "https://example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL().WithHosts("example.com", "example.net")
			},
		},
		{
			name: "JSON node is URL with hosts fails",
			json: `{"key": "https://example.dev"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL().WithHosts("example.com", "example.net")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is URL with hosts "example.com", "example.net", actual is "example.dev"`,
			},
		},
		{
			name: "JSON node is URL checked by custom function",
			json: `{"key": "https://example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL().That(func(u *url.URL) error {
					return nil
				})
			},
		},
		{
			name: "JSON node is URL checked by custom function fails",
			json: `{"key": "https://example.com"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsURL().That(func(u *url.URL) error {
					return fmt.Errorf("error")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is URL: error`,
			},
		},
		{
			name: "JSON node is URL fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsURL().
					WithHosts().
					WithSchemas().
					That(func(u *url.URL) error { return nil })
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is string`,
			},
		},
		{
			name: "JSON node is time",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime()
			},
		},
		{
			name: "JSON node is time fails",
			json: `{"key": "invalid"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is time: parsing time "invalid" as "2006-01-02T15:04:05Z07:00": cannot parse "invalid" as "2006"`,
			},
		},
		{
			name: "JSON node is time with layout",
			json: `{"key": "16 Oct 22 15:20 MSK"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTimeWithLayout(time.RFC822)
			},
		},
		{
			name: "JSON node is time with layout fails",
			json: `{"key": "invalid"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTimeWithLayout(time.RFC822)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is time: parsing time "invalid" as "02 Jan 06 15:04 MST": cannot parse "invalid" as "02"`,
			},
		},
		{
			name: "JSON node is time equal",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().EqualTo(parseTime("2022-10-16T15:14:32+03:00"))
			},
		},
		{
			name: "JSON node is time equal fails",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().EqualTo(parseTime("2022-11-17T16:15:43+03:00"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time equal to "2022-11-17T16:15:43+03:00", actual is "2022-10-16T15:14:32+03:00"`,
			},
		},
		{
			name: "JSON node is time not equal",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().NotEqualTo(parseTime("2022-11-17T16:15:43+03:00"))
			},
		},
		{
			name: "JSON node is time not equal fails",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().NotEqualTo(parseTime("2022-10-16T15:14:32+03:00"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time not equal to "2022-10-16T15:14:32+03:00", actual is "2022-10-16T15:14:32+03:00"`,
			},
		},
		{
			name: "JSON node is time after",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().After(parseTime("2022-10-16T00:00:00+03:00"))
			},
		},
		{
			name: "JSON node is time after fails",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().After(parseTime("2022-10-17T00:00:00+03:00"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time after "2022-10-17T00:00:00+03:00", actual is "2022-10-16T15:14:32+03:00"`,
			},
		},
		{
			name: "JSON node is time after or equal",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().AfterOrEqualTo(parseTime("2022-10-16T15:14:32+03:00"))
			},
		},
		{
			name: "JSON node is time after or equal fails",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().AfterOrEqualTo(parseTime("2022-10-16T15:14:33+03:00"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time after or equal to "2022-10-16T15:14:33+03:00", actual is "2022-10-16T15:14:32+03:00"`,
			},
		},
		{
			name: "JSON node is time before",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().Before(parseTime("2022-10-17T00:00:00+03:00"))
			},
		},
		{
			name: "JSON node is time before fails",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().Before(parseTime("2022-10-16T00:00:00+03:00"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time before "2022-10-16T00:00:00+03:00", actual is "2022-10-16T15:14:32+03:00"`,
			},
		},
		{
			name: "JSON node is time before or equal",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().BeforeOrEqualTo(parseTime("2022-10-16T15:14:32+03:00"))
			},
		},
		{
			name: "JSON node is time before or equal fails",
			json: `{"key": "2022-10-16T15:14:32+03:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().BeforeOrEqualTo(parseTime("2022-10-16T15:14:31+03:00"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time before or equal to "2022-10-16T15:14:31+03:00", actual is "2022-10-16T15:14:32+03:00"`,
			},
		},
		{
			name: "JSON node is time at date start",
			json: `{"key": "2022-10-16T00:00:00+00:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().AtDate(2022, time.October, 16)
			},
		},
		{
			name: "JSON node is time at date middle",
			json: `{"key": "2022-10-16T12:00:00+00:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().AtDate(2022, time.October, 16)
			},
		},
		{
			name: "JSON node is time at date end",
			json: `{"key": "2022-10-16T23:59:59+00:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().AtDate(2022, time.October, 16)
			},
		},
		{
			name: "JSON node is time at date fails",
			json: `{"key": "2022-10-17T00:00:00+00:00"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsTime().AtDate(2022, time.October, 16)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time at date "2022-10-16", actual is "2022-10-17T00:00:00Z"`,
			},
		},
		{
			name: "JSON node is date",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate()
			},
		},
		{
			name: "JSON node is date fails",
			json: `{"key": "invalid"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is time: parsing time "invalid" as "2006-01-02": cannot parse "invalid" as "2006"`,
			},
		},
		{
			name: "JSON node is date equal",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().EqualToDate(2022, time.October, 16)
			},
		},
		{
			name: "JSON node is date equal fails",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().EqualToDate(2022, time.October, 15)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time equal to "2022-10-15", actual is "2022-10-16"`,
			},
		},
		{
			name: "JSON node is date not equal",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().NotEqualToDate(2022, time.October, 15)
			},
		},
		{
			name: "JSON node is date not equal fails",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().NotEqualToDate(2022, time.October, 16)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time not equal to "2022-10-16", actual is "2022-10-16"`,
			},
		},
		{
			name: "JSON node is date after",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().AfterDate(2022, time.October, 15)
			},
		},
		{
			name: "JSON node is date after fails",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().AfterDate(2022, time.October, 16)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time after "2022-10-16", actual is "2022-10-16"`,
			},
		},
		{
			name: "JSON node is date after or equal",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().AfterOrEqualToDate(2022, time.October, 16)
			},
		},
		{
			name: "JSON node is date after or equal fails",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().AfterOrEqualToDate(2022, time.October, 17)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time after or equal to "2022-10-17", actual is "2022-10-16"`,
			},
		},
		{
			name: "JSON node is date before",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().BeforeDate(2022, time.October, 17)
			},
		},
		{
			name: "JSON node is date before fails",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().BeforeDate(2022, time.October, 16)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time before "2022-10-16", actual is "2022-10-16"`,
			},
		},
		{
			name: "JSON node is date before or equal",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().BeforeOrEqualToDate(2022, time.October, 16)
			},
		},
		{
			name: "JSON node is date before or equal fails",
			json: `{"key": "2022-10-16"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsDate().BeforeOrEqualToDate(2022, time.October, 15)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is time before or equal to "2022-10-15", actual is "2022-10-16"`,
			},
		},
		{
			name: "JSON node is time fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").
					IsTime().
					EqualTo(time.Now()).
					NotEqualTo(time.Now()).
					After(time.Now()).
					AfterOrEqualTo(time.Now()).
					Before(time.Now()).
					BeforeOrEqualTo(time.Now()).
					EqualToDate(0, 0, 0).
					NotEqualToDate(0, 0, 0).
					AfterDate(0, 0, 0).
					AfterOrEqualToDate(0, 0, 0).
					BeforeDate(0, 0, 0).
					BeforeOrEqualToDate(0, 0, 0).
					Value()
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is string`,
			},
		},
		{
			name: "JSON node is JSON",
			json: `{"key": "{\"key\": \"value\"}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsString().EqualTo("value")
				})
			},
		},
		{
			name: "JSON node is JSON fails",
			json: `{"key": "{key}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsString().EqualTo("value")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: data has invalid JSON: invalid character 'k' looking for beginning of object key string`,
			},
		},
		{
			name: "JSON node is JSON: missing node",
			json: `{"key": "{}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.At("key")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed to find JSON node "key": [key] not found`,
			},
		},
		{
			name: "JSON node is JSON: string assertion failed",
			json: `{"key": "{\"key\": 123}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsString()
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key" is string`,
			},
		},
		{
			name: "JSON node is JSON: string equal assertion failed",
			json: `{"key": "{\"key\": \"value\"}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsString().EqualTo("expected")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": equal to "expected", actual is "value"`,
			},
		},
		{
			name: "JSON node is JSON: URL assertion failed",
			json: `{"key": "{\"key\": \"http://example.com\"}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsURL().WithHosts("example.net")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": is URL with hosts "example.net", actual is "example.com"`,
			},
		},
		{
			name: "JSON node is JSON: time assertion failed",
			json: `{"key": "{\"key\": \"2022-10-16\"}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsDate().EqualToDate(2022, time.October, 15)
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": is time equal to "2022-10-15", actual is "2022-10-16"`,
			},
		},
		{
			name: "JSON node is JSON: UUID assertion failed",
			json: `{"key": "{\"key\": \"00000000-0000-0000-0000-000000000000\"}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsUUID().NotNil()
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": is not nil UUID, actual is "00000000-0000-0000-0000-000000000000"`,
			},
		},
		{
			name: "JSON node is JSON: JWT assertion failed",
			json: `{"key": "{\"key\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c\"}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsJWT(getJWTSecret).WithAlgorithm("HS512")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": is JWT with algorithm "HS512", actual is "HS256"`,
			},
		},
		{
			name: "JSON node is JSON: number assertion failed",
			json: `{"key": "{\"key\": 123}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsNumber().EqualTo(321)
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": equal to 321.000000, actual is 123.000000`,
			},
		},
		{
			name: "JSON node is JSON: integer assertion failed",
			json: `{"key": "{\"key\": 123}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsInteger().EqualTo(321)
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": equal to 321, actual is 123`,
			},
		},
		{
			name: "JSON node is JSON: array assertion failed",
			json: `{"key": "{\"key\": []}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsArray().WithLength(1)
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": is array with length is 1, actual is 0`,
			},
		},
		{
			name: "JSON node is JSON: object assertion failed",
			json: `{"key": "{\"key\": {}}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").IsObject().WithPropertiesCount(1)
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": is object with properties count is 1, actual is 0`,
			},
		},
		{
			name: "JSON node is JSON: for each array assertion failed",
			json: `{"key": "{\"key\": [\"value\"]}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").ForEach(func(node *assertjson.AssertNode) {
						node.IsString().EqualTo("expected")
					})
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key[0]": equal to "expected", actual is "value"`,
			},
		},
		{
			name: "JSON node is JSON: for each object assertion failed",
			json: `{"key": "{\"key\": {\"key\": \"value\"}}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").ForEach(func(node *assertjson.AssertNode) {
						node.IsString().EqualTo("expected")
					})
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key.key": equal to "expected", actual is "value"`,
			},
		},
		{
			name: "JSON node is JSON: callable assertion",
			json: `{"key": "{\"key\": {}}"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsString().WithJSON(func(json *assertjson.AssertJSON) {
					json.Node("key").Assert(func(json *assertjson.AssertJSON) {
						json.Node().IsObject().WithPropertiesCount(1)
					})
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is string with JSON: failed asserting that JSON node "key": is object with properties count is 1, actual is 0`,
			},
		},
		{
			name: "JSON node is JWT",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret)
			},
		},
		{
			name: "JSON node is JWT fails",
			json: `{"key": "invalid"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is JWT: token contains an invalid number of segments`,
			},
		},
		{
			name: "JSON node is JWT fails on invalid signature",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.u5ClmY6KVIUdReH0H2qpG2oyqrb8VTfJ8NzaLVxylEI"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret)
			},
			wantMessages: []string{
				`failed asserting that JSON node "key" is JWT: signature is invalid`,
			},
		},
		{
			name: "JSON node is JWT with algorithm",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret).WithAlgorithm("HS256")
			},
		},
		{
			name: "JSON node is JWT with algorithm fails",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret).WithAlgorithm("HS512")
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is JWT with algorithm "HS512", actual is "HS256"`,
			},
		},
		{
			name: "JSON node is JWT with id",
			json: jsonWithJWT(jwt.MapClaims{"jti": "12345"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithID("12345")
			},
		},
		{
			name: "JSON node is JWT with id no field",
			json: jsonWithJWT(jwt.MapClaims{}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithID("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with id ("jti") "unexpected": field does not exist`,
			},
		},
		{
			name: "JSON node is JWT with id invalid type",
			json: jsonWithJWT(jwt.MapClaims{"jti": 12345}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithID("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with id ("jti") "unexpected": string is expected`,
			},
		},
		{
			name: "JSON node is JWT with id not equal",
			json: jsonWithJWT(jwt.MapClaims{"jti": "12345"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithID("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with id ("jti") "unexpected", actual is "12345"`,
			},
		},
		{
			name: "JSON node is JWT with issuer",
			json: jsonWithJWT(jwt.MapClaims{"iss": "expected"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuer("expected")
			},
		},
		{
			name: "JSON node is JWT with issuer no field",
			json: jsonWithJWT(jwt.MapClaims{}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuer("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with issuer ("iss") "unexpected": field does not exist`,
			},
		},
		{
			name: "JSON node is JWT with issuer invalid type",
			json: jsonWithJWT(jwt.MapClaims{"iss": 12345}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuer("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with issuer ("iss") "unexpected": string is expected`,
			},
		},
		{
			name: "JSON node is JWT with issuer not equal",
			json: jsonWithJWT(jwt.MapClaims{"iss": "expected"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuer("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with issuer ("iss") "unexpected", actual is "expected"`,
			},
		},
		{
			name: "JSON node is JWT with subject",
			json: jsonWithJWT(jwt.MapClaims{"sub": "expected"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithSubject("expected")
			},
		},
		{
			name: "JSON node is JWT with subject no field",
			json: jsonWithJWT(jwt.MapClaims{}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithSubject("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with subject ("sub") "unexpected": field does not exist`,
			},
		},
		{
			name: "JSON node is JWT with subject invalid type",
			json: jsonWithJWT(jwt.MapClaims{"sub": 12345}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithSubject("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with subject ("sub") "unexpected": string is expected`,
			},
		},
		{
			name: "JSON node is JWT with subject not equal",
			json: jsonWithJWT(jwt.MapClaims{"sub": "expected"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithSubject("unexpected")
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with subject ("sub") "unexpected", actual is "expected"`,
			},
		},
		{
			name: "JSON node is JWT with audience",
			json: jsonWithJWT(jwt.MapClaims{"aud": "expected"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithAudience([]string{"expected"})
			},
		},
		{
			name: "JSON node is JWT with multiple audience",
			json: jsonWithJWT(jwt.MapClaims{"aud": []string{"one", "two"}}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithAudience([]string{"one", "two"})
			},
		},
		{
			name: "JSON node is JWT with audience no field",
			json: jsonWithJWT(jwt.MapClaims{}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithAudience([]string{"one", "two"})
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with audience ("aud") ["one", "two"]: field does not exist`,
			},
		},
		{
			name: "JSON node is JWT with audience invalid type",
			json: jsonWithJWT(jwt.MapClaims{"aud": 12345}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithAudience([]string{"unexpected"})
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with audience ("aud") ["unexpected"]: string or array of strings expected`,
			},
		},
		{
			name: "JSON node is JWT with audience not equal",
			json: jsonWithJWT(jwt.MapClaims{"aud": "expected"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithAudience([]string{"unexpected"})
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with audience ("aud") ["unexpected"], actual is ["expected"]`,
			},
		},
		{
			name: "JSON node is JWT with expires at",
			json: jsonWithJWT(jwt.MapClaims{"exp": time.Now().Add(time.Hour).Unix()}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithExpiresAt()
			},
		},
		{
			name: "JSON node is JWT with expires at no field",
			json: jsonWithJWT(jwt.MapClaims{}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithExpiresAt()
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with expires at ("exp"): field does not exist`,
			},
		},
		{
			name: "JSON node is JWT with expires at invalid type",
			json: jsonWithJWT(jwt.MapClaims{"exp": "string"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithExpiresAt()
			},
			wantMessages: []string{
				`failed asserting that JSON node "" is JWT: Token is expired`,
			},
		},
		{
			name: "JSON node is JWT with expires at failed",
			json: jsonWithJWT(jwt.MapClaims{"exp": parseTime("2100-01-01T00:00:00Z").Unix()}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithExpiresAt().AfterDate(2200, time.January, 1)
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with expires at ("exp"): is time after "2200-01-01T00:00:00Z", actual is "2100-01-01T00:00:00Z"`,
			},
		},
		{
			name: "JSON node is JWT with not before",
			json: jsonWithJWT(jwt.MapClaims{"nbf": time.Now().Add(-time.Hour).Unix()}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithNotBefore()
			},
		},
		{
			name: "JSON node is JWT with not before no field",
			json: jsonWithJWT(jwt.MapClaims{}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithNotBefore()
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with not before ("nbf"): field does not exist`,
			},
		},
		{
			name: "JSON node is JWT with not before invalid type",
			json: jsonWithJWT(jwt.MapClaims{"nbf": "string"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithNotBefore()
			},
			wantMessages: []string{
				`failed asserting that JSON node "" is JWT: Token is not valid yet`,
			},
		},
		{
			name: "JSON node is JWT with not before failed",
			json: jsonWithJWT(jwt.MapClaims{"nbf": parseTime("2000-01-01T00:00:00Z").Unix()}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithNotBefore().AfterDate(2001, time.January, 1)
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with not before ("nbf"): is time after "2001-01-01T00:00:00Z", actual is "2000-01-01T00:00:00Z"`,
			},
		},
		{
			name: "JSON node is JWT with issued at",
			json: jsonWithJWT(jwt.MapClaims{"iat": time.Now().Add(-time.Hour).Unix()}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuedAt()
			},
		},
		{
			name: "JSON node is JWT with issued at no field",
			json: jsonWithJWT(jwt.MapClaims{}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuedAt()
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with issued at ("iat"): field does not exist`,
			},
		},
		{
			name: "JSON node is JWT with issued at invalid type",
			json: jsonWithJWT(jwt.MapClaims{"iat": "string"}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuedAt()
			},
			wantMessages: []string{
				`failed asserting that JSON node "" is JWT: Token used before issued`,
			},
		},
		{
			name: "JSON node is JWT with issued at failed",
			json: jsonWithJWT(jwt.MapClaims{"iat": parseTime("2000-01-01T00:00:00Z").Unix()}),
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsJWT(getJWTSecret).WithIssuedAt().AfterDate(2001, time.January, 1)
			},
			wantMessages: []string{
				`failed asserting that JSON node "": is JWT with issued at ("iat"): is time after "2001-01-01T00:00:00Z", actual is "2000-01-01T00:00:00Z"`,
			},
		},
		{
			name: "JSON node is JWT with header",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret).WithHeader(func(json *assertjson.AssertJSON) {
					json.Node("alg").IsString().EqualTo("HS256")
				})
			},
		},
		{
			name: "JSON node is JWT with header fails",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret).WithHeader(func(json *assertjson.AssertJSON) {
					json.Node("alg").IsString().EqualTo("HS512")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is JWT with header: failed asserting that JSON node "alg": equal to "HS512", actual is "HS256"`,
			},
		},
		{
			name: "JSON node is JWT with payload",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret).WithPayload(func(json *assertjson.AssertJSON) {
					json.Node("name").IsString().EqualTo("John Doe")
				})
			},
		},
		{
			name: "JSON node is JWT with payload fails",
			json: `{"key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("key").IsJWT(getJWTSecret).WithPayload(func(json *assertjson.AssertJSON) {
					json.Node("name").IsString().EqualTo("John Smith")
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "key": is JWT with payload: failed asserting that JSON node "name": equal to "John Smith", actual is "John Doe"`,
			},
		},
		// seek by json iterator
		{
			name: "seek by json iterator",
			json: `"value"`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node().IsString().EqualTo("value")
			},
		},
		{
			name: "seek by json iterator",
			json: `{"a": {"b": {"c": "value"}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("a", "b", "c").IsString().EqualTo("value")
			},
		},
		{
			name: "seek by json iterator",
			json: `{"a": {"b": {"c": "value"}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("a", "b", String("c")).IsString().EqualTo("value")
			},
		},
		{
			name: "seek by json iterator",
			json: `{"a": {"b": {"c": ["value"]}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("a", "b", "c", 0).IsString().EqualTo("value")
			},
		},
		{
			name: "seek by json iterator",
			json: `{"a": {"b": {"c": ["value"]}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("a", "b", "c", 1).IsString()
			},
			wantMessages: []string{
				`failed to find JSON node "a.b.c[1]": [1] not found`,
			},
		},
		{
			name: "seek by json iterator",
			json: `{"a": {"b": {"c": ["value"]}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.At("a", "b", "c", 1).Node().IsString()
			},
			wantMessages: []string{
				`failed to find JSON node "a.b.c[1]": [1] not found`,
				`failed asserting that JSON node "a.b.c[1]" is string`,
			},
		},
		// deprecated behaviour: seek by json pointer path
		{
			name: "deprecated: json pointer path",
			json: `"value"`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("").IsString().EqualTo("value")
			},
		},
		{
			name: "deprecated: json pointer path",
			json: `{"a": {"b": {"c": "value"}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/a/b/c").IsString().EqualTo("value")
			},
		},
		{
			name: "deprecated: json pointer path",
			json: `{"a": {"b": {"c": ["value"]}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/a/b/c/0").IsString().EqualTo("value")
			},
		},
		{
			name: "deprecated: json pointer path",
			json: `{"a": {"/b": {"~c": ["value"]}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/a/~1b/~0c/0").IsString().EqualTo("value")
			},
		},
		{
			name: "deprecated: json pointer path",
			json: `{"a": {"b": {"c": ["value"]}}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/a/b/c/1").IsString()
			},
			wantMessages: []string{
				`failed to find JSON node "a.b.c[1]": [1] not found`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tester := &mock.Tester{}

			assertjson.Has(tester, []byte(test.json), test.assert)

			tester.AssertContains(t, test.wantMessages)
		})
	}
}

func TestAssertNode_Exists(t *testing.T) {
	tests := []struct {
		json string
		want bool
	}{
		{json: `{"key": "value"}`, want: true},
		{json: `{}`, want: false},
	}
	for _, test := range tests {
		t.Run(test.json, func(t *testing.T) {
			tester := &mock.Tester{}

			var got bool
			assertjson.Has(tester, []byte(test.json), func(json *assertjson.AssertJSON) {
				got = json.Node("key").Exists()
			})

			assert.Equal(t, test.want, got)
		})
	}
}

const tokenSecret = "your-256-bit-secret"

func getJWTSecret(token *jwt.Token) (interface{}, error) {
	return []byte(tokenSecret), nil
}

func jsonWithJWT(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return string(j)
}

func parseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

type String string

func (s String) String() string {
	return string(s)
}
