package assertjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileHas(t *testing.T) {
	FileHas(t, "./../test/testdata/object.json", func(json *AssertJSON) {
		// common assertions
		json.Node("/nullNode").Exists()
		json.Node("/notExistingNode").DoesNotExist()
		json.Node("/nullNode").IsNull()
		json.Node("/stringNode").IsNotNull()
		json.Node("/trueBooleanNode").IsTrue()
		json.Node("/falseBooleanNode").IsFalse()

		// string assertions
		json.Node("/stringNode").IsString()
		json.Node("/stringNode").EqualToTheString("stringValue")
		json.Node("/stringNode").AssertString(func(t testing.TB, value string) {
			assert.Equal(t, "stringValue", value)
		})
		json.Node("/stringNode").Matches("^string.*$")
		json.Node("/stringNode").DoesNotMatch("^notMatch$")
		json.Node("/stringNode").Contains("string")
		json.Node("/stringNode").DoesNotContain("notContain")
		json.Node("/stringNode").IsStringWithLength(11)
		json.Node("/stringNode").IsStringWithLengthInRange(11, 11)

		// numeric assertions
		json.Node("/integerNode").IsInteger()
		json.Node("/integerNode").EqualToTheInteger(123)
		json.Node("/integerNode").IsNumberInRange(122, 124)
		json.Node("/integerNode").IsNumberGreaterThan(122)
		json.Node("/integerNode").IsNumberGreaterThanOrEqual(123)
		json.Node("/integerNode").IsNumberLessThan(124)
		json.Node("/integerNode").IsNumberLessThanOrEqual(123)
		json.Node("/floatNode").IsFloat()
		json.Node("/floatNode").EqualToTheFloat(123.123)
		json.Node("/floatNode").IsNumberInRange(122, 124)
		json.Node("/floatNode").IsNumberGreaterThan(122)
		json.Node("/floatNode").IsNumberGreaterThanOrEqual(123.123)
		json.Node("/floatNode").IsNumberLessThan(124)
		json.Node("/floatNode").IsNumberLessThanOrEqual(123.123)

		// array assertions
		json.Node("/arrayNode").IsArrayWithElementsCount(1)

		// object assertions
		json.Node("/objectNode").IsObjectWithPropertiesCount(1)

		// json pointer expression
		json.Node("/complexNode/items/1/key").EqualToTheString("value")
		json.Nodef("/complexNode/items/%d/key", 1).EqualToTheString("value")

		// complex keys
		json.Node("/@id").EqualToTheString("json-ld-id")
		json.Node("/hydra:members").EqualToTheString("hydraMembers")

		// complex assertions
		json.At("/complexNode").Node("/items/1/key").EqualToTheString("value")
		json.Atf("/complexNode/%s", "items").Node("/1/key").EqualToTheString("value")
	})
}
