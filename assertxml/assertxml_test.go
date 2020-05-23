package assertxml

import "testing"

func TestFileHas(t *testing.T) {
	FileHas(t, "./../test/fixtures/object.xml", func(xml *AssertXML) {
		// common assertions
		xml.Node("/root/stringNode").Exists()
		xml.Node("/root/notExistingNode").DoesNotExist()

		// string assertions
		xml.Node("/root/stringNode").EqualToTheString("stringValue")
	})
}
