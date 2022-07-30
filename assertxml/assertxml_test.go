package assertxml_test

import (
	"testing"

	"github.com/muonsoft/api-testing/assertxml"
)

func TestFileHas(t *testing.T) {
	assertxml.FileHas(t, "./../test/testdata/object.xml", func(xml *assertxml.AssertXML) {
		// common assertions
		xml.Node("/root/stringNode").Exists()
		xml.Node("/root/notExistingNode").DoesNotExist()

		// string assertions
		xml.Node("/root/stringNode").EqualToTheString("stringValue")
		xml.Nodef("/root/%s", "stringNode").EqualToTheString("stringValue")
	})
}
