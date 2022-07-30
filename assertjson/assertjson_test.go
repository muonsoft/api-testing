package assertjson_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/muonsoft/api-testing/assertjson"
	"github.com/muonsoft/api-testing/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestFileHas(t *testing.T) {
	assertjson.FileHas(t, "./../test/testdata/object.json", func(json *assertjson.AssertJSON) {
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
		json.Node("/stringNode").EqualToTheString("stringValue")
		json.Node("/stringNode").AssertString(func(tb testing.TB, value string) {
			tb.Helper()
			assert.Equal(tb, "stringValue", value)
		})
		json.Node("/stringNode").Matches("^string.*$")
		json.Node("/stringNode").DoesNotMatch("^notMatch$")
		json.Node("/stringNode").Contains("string")
		json.Node("/stringNode").DoesNotContain("notContain")
		json.Node("/stringNode").IsStringWithLength(11)
		json.Node("/stringNode").IsStringWithLengthInRange(11, 11)

		// fluent string assertions
		json.Node("/stringNode").IsString()
		json.Node("/stringNode").IsString().EqualTo("stringValue")
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
		json.Node("/stringNode").IsString().Assert(func(tb testing.TB, value string) {
			tb.Helper()
			assert.Equal(tb, "stringValue", value)
		})

		// numeric assertions
		json.Node("/integerNode").IsInteger()
		json.Node("/integerNode").IsInteger().EqualTo(123)
		json.Node("/integerNode").IsInteger().GreaterThan(122)
		json.Node("/integerNode").IsInteger().GreaterThanOrEqual(123)
		json.Node("/integerNode").IsInteger().LessThan(124)
		json.Node("/integerNode").IsInteger().LessThanOrEqual(123)
		json.Node("/integerNode").EqualToTheInteger(123)
		json.Node("/integerNode").IsNumberInRange(122, 124)
		json.Node("/integerNode").IsNumberGreaterThan(122)
		json.Node("/integerNode").IsNumberGreaterThanOrEqual(123)
		json.Node("/integerNode").IsNumberLessThan(124)
		json.Node("/integerNode").IsNumberLessThanOrEqual(123)
		json.Node("/floatNode").IsFloat()
		json.Node("/floatNode").IsNumber()
		json.Node("/floatNode").IsNumber().EqualTo(123.123)
		json.Node("/floatNode").IsNumber().EqualToWithDelta(123.123, 0.1)
		json.Node("/floatNode").IsNumber().GreaterThan(122)
		json.Node("/floatNode").IsNumber().GreaterThanOrEqual(123.123)
		json.Node("/floatNode").IsNumber().LessThan(124)
		json.Node("/floatNode").IsNumber().LessThanOrEqual(123.123)
		json.Node("/floatNode").IsNumber().GreaterThanOrEqual(122).LessThanOrEqual(124)
		json.Node("/floatNode").EqualToTheFloat(123.123)
		json.Node("/floatNode").IsNumberInRange(122, 124)
		json.Node("/floatNode").IsNumberGreaterThan(122)
		json.Node("/floatNode").IsNumberGreaterThanOrEqual(123.123)
		json.Node("/floatNode").IsNumberLessThan(124)
		json.Node("/floatNode").IsNumberLessThanOrEqual(123.123)

		// array assertions
		json.Node("/arrayNode").IsArrayWithElementsCount(1)
		json.Node("/arrayNode").IsArray()
		json.Node("/arrayNode").IsArray().WithLength(1)
		json.Node("/arrayNode").IsArray().WithLengthGreaterThan(0)
		json.Node("/arrayNode").IsArray().WithLengthGreaterThanOrEqual(1)
		json.Node("/arrayNode").IsArray().WithLengthLessThan(2)
		json.Node("/arrayNode").IsArray().WithLengthLessThanOrEqual(1)
		json.Node("/arrayNode").IsArray().WithUniqueElements()

		// object assertions
		json.Node("/objectNode").IsObjectWithPropertiesCount(1)
		json.Node("/objectNode").IsObject()
		json.Node("/objectNode").IsObject().WithPropertiesCount(1)
		json.Node("/objectNode").IsObject().WithPropertiesCountGreaterThan(0)
		json.Node("/objectNode").IsObject().WithPropertiesCountGreaterThanOrEqual(1)
		json.Node("/objectNode").IsObject().WithPropertiesCountLessThan(2)
		json.Node("/objectNode").IsObject().WithPropertiesCountLessThanOrEqual(1)
		json.Node("/objectNode").IsObject().WithUniqueElements()

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
		assert.JSONEq(t, `{"objectKey": "objectValue"}`, string(json.Node("/objectNode").JSON()))
	})
}

func TestHas(t *testing.T) {
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
				json.At("/key")
			},
			wantMessages: []string{
				`failed to find JSON node "/key": Object has no key 'key'`,
			},
		},
		{
			name: "JSON node exists",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").Exists()
			},
		},
		{
			name: "JSON node exists fails",
			json: `{}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").Exists()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" exists`,
			},
		},
		{
			name: "JSON node does not exist",
			json: `{}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").DoesNotExist()
			},
		},
		{
			name: "JSON node does not exist fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").DoesNotExist()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" does not exist`,
			},
		},
		{
			name: "JSON node is null",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNull()
			},
		},
		{
			name: "JSON node is null fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNull()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is null`,
			},
		},
		{
			name: "JSON node is not null",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNotNull()
			},
		},
		{
			name: "JSON node is not null fails",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNotNull()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is not null`,
			},
		},
		{
			name: "JSON node is true",
			json: `{"key": true}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsTrue()
			},
		},
		{
			name: "JSON node is true fails",
			json: `{"key": false}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsTrue()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is true`,
			},
		},
		{
			name: "JSON node is true on not boolean",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsTrue()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is boolean`,
			},
		},
		{
			name: "JSON node is false",
			json: `{"key": false}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsFalse()
			},
		},
		{
			name: "JSON node is false fails",
			json: `{"key": true}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsFalse()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is false`,
			},
		},
		{
			name: "JSON node is false on not boolean",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsFalse()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is boolean`,
			},
		},
		{
			name: "JSON node equal to JSON",
			json: `{"key": {"key": "value"}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").EqualJSON(`{"key": "value"}`)
			},
		},
		{
			name: "JSON node equal to JSON fails",
			json: `{"key": {"k": "v"}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").EqualJSON(`{"key": "value"}`)
			},
			wantMessages: []string{
				"Not equal",
				`failed at json node "/key"`,
			},
		},
		{
			name: "JSON node is number",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber()
			},
		},
		{
			name: "JSON node is number fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber()
			},
			wantMessages: []string{
				`value at path "/key" is not a number`,
			},
		},
		{
			name: "JSON node is number equal to",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().EqualTo(123.123)
			},
		},
		{
			name: "JSON node is number equal to fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().EqualTo(321.123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" equal to 321.123000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number equal to with delta",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().EqualToWithDelta(123.1, 1)
			},
		},
		{
			name: "JSON node is number equal to with delta fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().EqualToWithDelta(321.123, 1)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" equal to 321.123000 with delta 1.000000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number greater than",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().GreaterThan(123.122)
			},
		},
		{
			name: "JSON node is number greater than fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().GreaterThan(123.123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" greater than 123.123000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number greater than or equal",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().GreaterThanOrEqual(123.123)
			},
		},
		{
			name: "JSON node is number greater than or equal fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().GreaterThanOrEqual(123.124)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" greater than or equal 123.124000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number less than",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().LessThan(123.124)
			},
		},
		{
			name: "JSON node is number less than fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().LessThan(123.123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" less than 123.123000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number less than or equal",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().LessThanOrEqual(123.123)
			},
		},
		{
			name: "JSON node is number less than or equal fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsNumber().LessThanOrEqual(123.122)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" less than or equal 123.122000, actual is 123.123000`,
			},
		},
		{
			name: "JSON node is number fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").
					IsNumber().
					EqualTo(0).
					EqualToWithDelta(0, 0).
					GreaterThan(0).
					GreaterThanOrEqual(0).
					LessThan(0).
					LessThanOrEqual(0)
			},
			wantMessages: []string{
				`value at path "/key" is not a number`,
			},
		},
		{
			name: "JSON node is integer",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger()
			},
		},
		{
			name: "JSON node is integer fails on string",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger()
			},
			wantMessages: []string{
				`value at path "/key" is not numeric`,
			},
		},
		{
			name: "JSON node is integer fails on float",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger()
			},
			wantMessages: []string{
				`value at path "/key" is float, not integer`,
			},
		},
		{
			name: "JSON node is integer equal to",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().EqualTo(123)
			},
		},
		{
			name: "JSON node is integer equal to fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().EqualTo(321)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" equal to 321, actual is 123`,
			},
		},
		{
			name: "JSON node is integer greater than",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().GreaterThan(122)
			},
		},
		{
			name: "JSON node is integer greater than fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().GreaterThan(123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" greater than 123, actual is 123`,
			},
		},
		{
			name: "JSON node is integer greater than or equal",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().GreaterThanOrEqual(123)
			},
		},
		{
			name: "JSON node is integer greater than or equal fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().GreaterThanOrEqual(124)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" greater than or equal 124, actual is 123`,
			},
		},
		{
			name: "JSON node is integer less than",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().LessThan(124)
			},
		},
		{
			name: "JSON node is integer less than fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().LessThan(123)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" less than 123, actual is 123`,
			},
		},
		{
			name: "JSON node is integer less than or equal",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().LessThanOrEqual(123)
			},
		},
		{
			name: "JSON node is integer less than or equal fails",
			json: `{"key": 123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsInteger().LessThanOrEqual(122)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" less than or equal 122, actual is 123`,
			},
		},
		{
			name: "JSON node is integer fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").
					IsInteger().
					EqualTo(0).
					GreaterThan(0).
					GreaterThanOrEqual(0).
					LessThan(0).
					LessThanOrEqual(0)
			},
			wantMessages: []string{
				`value at path "/key" is not numeric`,
			},
		},
		{
			name: "JSON node is string",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString()
			},
		},
		{
			name: "JSON node is string fails",
			json: `{"key": 123.123}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is string`,
			},
		},
		{
			name: "JSON node is string equal to",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().EqualTo("value")
			},
		},
		{
			name: "JSON node is string equal to fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().EqualTo("string")
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" equal to "string", actual is "value"`,
			},
		},
		{
			name: "JSON node is string matches",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().Matches(".*")
			},
		},
		{
			name: "JSON node is string matches fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().Matches("\\d+")
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" matches "\d+", actual is "value"`,
			},
		},
		{
			name: "JSON node is string matches compiled regexp",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().Matches(regexp.MustCompile(".*"))
			},
		},
		{
			name: "JSON node is string matches compiled regexp fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().Matches(regexp.MustCompile(`\d+`))
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" matches "\d+", actual is "value"`,
			},
		},
		{
			name: "JSON node is string not matches",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().NotMatches("\\d+")
			},
		},
		{
			name: "JSON node is string not matches fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().NotMatches(".*")
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" not matches ".*", actual is "value"`,
			},
		},
		{
			name: "JSON node is string not matches compiled regexp",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().NotMatches(regexp.MustCompile(`\d+`))
			},
		},
		{
			name: "JSON node is string not matches compiled regexp fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().NotMatches(regexp.MustCompile(".*"))
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" not matches ".*", actual is "value"`,
			},
		},
		{
			name: "JSON node is string contains",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().Contains("alu")
			},
		},
		{
			name: "JSON node is string contains fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().Contains("string")
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" contains "string", actual is "value"`,
			},
		},
		{
			name: "JSON node is string not contains",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().NotContains("string")
			},
		},
		{
			name: "JSON node is string not contains fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().NotContains("alu")
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" not contains "alu", actual is "value"`,
			},
		},
		{
			name: "JSON node is string with length",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLength(5)
			},
		},
		{
			name: "JSON node is string with length fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLength(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is string with length is 4, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length greater than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthGreaterThan(4)
			},
		},
		{
			name: "JSON node is string with length greater than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthGreaterThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is string with length greater than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length greater than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthGreaterThanOrEqual(5)
			},
		},
		{
			name: "JSON node is string with length greater than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthGreaterThanOrEqual(6)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is string with length greater than or equal to 6, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length less than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthLessThan(6)
			},
		},
		{
			name: "JSON node is string with length less than",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthLessThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is string with length less than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is string with length less than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthLessThanOrEqual(5)
			},
		},
		{
			name: "JSON node is string with length less than or equal",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().WithLengthLessThanOrEqual(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is string with length less than or equal to 4, actual is 5`,
			},
		},
		{
			name: "JSON node is string checked by custom function",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().That(func(s string) error {
					return nil
				})
			},
		},
		{
			name: "JSON node is string checked by custom function fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsString().That(func(s string) error {
					return fmt.Errorf("error")
				})
			},
			wantMessages: []string{
				`failed asserting JSON node "/key": error`,
			},
		},
		{
			name: "JSON node is string fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").
					IsString().
					EqualTo("").
					Matches(".*").
					NotMatches(".*").
					Contains("").
					NotContains("").
					WithLength(0).
					WithLengthGreaterThan(0).
					WithLengthGreaterThanOrEqual(0).
					WithLengthLessThan(0).
					WithLengthLessThanOrEqual(0).
					That(func(s string) error { return nil }).
					Assert(func(tb testing.TB, value string) { tb.Helper() })
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is string`,
			},
		},
		{
			name: "JSON node is array",
			json: `{"key": [1, 2, 3]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray()
			},
		},
		{
			name: "JSON node is array fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is array`,
			},
		},
		{
			name: "JSON node is array with length",
			json: `{"key": [1, 2, 3]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLength(3)
			},
		},
		{
			name: "JSON node is array with length fails",
			json: `{"key": [1, 2, 3]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLength(2)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is array with length is 2, actual is 3`,
			},
		},
		{
			name: "JSON node is array with length greater than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthGreaterThan(4)
			},
		},
		{
			name: "JSON node is array with length greater than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthGreaterThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is array with length greater than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is array with length greater than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthGreaterThanOrEqual(5)
			},
		},
		{
			name: "JSON node is array with length greater than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthGreaterThanOrEqual(6)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is array with length greater than or equal to 6, actual is 5`,
			},
		},
		{
			name: "JSON node is array with length less than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthLessThan(6)
			},
		},
		{
			name: "JSON node is array with length less than",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthLessThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is array with length less than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is array with length less than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthLessThanOrEqual(5)
			},
		},
		{
			name: "JSON node is array with length less than or equal",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithLengthLessThanOrEqual(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is array with length less than or equal to 4, actual is 5`,
			},
		},
		{
			name: "JSON node is array with unique elements",
			json: `{"key": [1, 2, 3, 4, 5]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithUniqueElements()
			},
		},
		{
			name: "JSON node is array with unique elements fails",
			json: `{"key": [3, 2, 3, 4, 2]}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsArray().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"/key\" is array with unique elements, duplicated elements",
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
				json.Node("/key").IsArray().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"/key\" is array with unique elements, duplicated elements",
			},
		},
		{
			name: "JSON node is array fails once for a chain",
			json: `{"key": null}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").
					IsArray().
					WithLength(0).
					WithLengthGreaterThan(0).
					WithLengthGreaterThanOrEqual(0).
					WithLengthLessThan(0).
					WithLengthLessThanOrEqual(0).
					WithUniqueElements()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is array`,
			},
		},
		{
			name: "JSON node is object",
			json: `{"key": {"a": 1}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject()
			},
		},
		{
			name: "JSON node is object fails",
			json: `{"key": "value"}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject()
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is object`,
			},
		},
		{
			name: "JSON node is object with properties count",
			json: `{"key": {"a": 1, "b": 2, "c": 3}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCount(3)
			},
		},
		{
			name: "JSON node is object with properties count fails",
			json: `{"key": {"a": 1, "b": 2, "c": 3}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCount(2)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is object with properties count is 2, actual is 3`,
			},
		},
		{
			name: "JSON node is object with properties count greater than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountGreaterThan(4)
			},
		},
		{
			name: "JSON node is object with properties count greater than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountGreaterThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is object with properties count greater than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is object with properties count greater than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountGreaterThanOrEqual(5)
			},
		},
		{
			name: "JSON node is object with properties count greater than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountGreaterThanOrEqual(6)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is object with properties count greater than or equal to 6, actual is 5`,
			},
		},
		{
			name: "JSON node is object with properties count less than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountLessThan(6)
			},
		},
		{
			name: "JSON node is object with properties count less than",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountLessThan(5)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is object with properties count less than 5, actual is 5`,
			},
		},
		{
			name: "JSON node is object with properties count less than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountLessThanOrEqual(5)
			},
		},
		{
			name: "JSON node is object with properties count less than or equal",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithPropertiesCountLessThanOrEqual(4)
			},
			wantMessages: []string{
				`failed asserting that JSON node "/key" is object with properties count less than or equal to 4, actual is 5`,
			},
		},
		{
			name: "JSON node is object with unique elements",
			json: `{"key": {"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithUniqueElements()
			},
		},
		{
			name: "JSON node is object with unique elements fails",
			json: `{"key": {"a": 3, "b": 2, "c": 3, "d": 4, "e": 2}}`,
			assert: func(json *assertjson.AssertJSON) {
				json.Node("/key").IsObject().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"/key\" is object with unique elements, duplicated elements",
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
				json.Node("/key").IsObject().WithUniqueElements()
			},
			wantMessages: []string{
				"failed asserting that JSON node \"/key\" is object with unique elements, duplicated elements",
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
				got = json.Node("/key").Exists()
			})

			assert.Equal(t, test.want, got)
		})
	}
}
