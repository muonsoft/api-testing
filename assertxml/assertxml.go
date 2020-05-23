// Package assertxml provides methods for testing XML values. Selecting XML values provided by XML Path Syntax.
//
// Example usage
//    import (
//        "net/http"
//        "net/http/httptest"
//        "testing"
//        "github.com/muonsoft/api-testing/assertxml"
//     )
//
//     func TestYourAPI(t *testing.T) {
//        recorder := httptest.NewRecorder()
//        handler := createHTTPHandler()
//
//        request, _ := http.NewRequest("GET", "/content", nil)
//        handler.ServeHTTP(recorder, request)
//
//        assertxml.Has(t, recorder.Body.Bytes(), func(xml *AssertXML) {
//            // common assertions
//            xml.Node("/root/stringNode").Exists()
//            xml.Node("/root/notExistingNode").DoesNotExist()
//
//            // string assertions
//            xml.Node("/root/stringNode").EqualToTheString("stringValue")
//        })
//     }
package assertxml

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/xmlpath.v2"
	"io/ioutil"
	"testing"
)

// AssertXML - main structure that holds parsed XML.
type AssertXML struct {
	t   *testing.T
	xml *xmlpath.Node
}

// AssertNode - structure for asserting XML node.
type AssertNode struct {
	t     *testing.T
	found bool
	path  string
	value string
}

// XMLAssertFunc - callback function used for asserting XML nodes.
type XMLAssertFunc func(xml *AssertXML)

// FileHas loads XML from file and runs user callback for testing its nodes.
func FileHas(t *testing.T, filename string, xmlAssert XMLAssertFunc) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		assert.Failf(t, "failed to read file '%s': %s", filename, err.Error())
	}
	Has(t, data, xmlAssert)
}

// Has loads XML from byte slice and runs user callback for testing its nodes.
func Has(t *testing.T, data []byte, xmlAssert XMLAssertFunc) {
	xml, err := xmlpath.Parse(bytes.NewReader(data))
	body := &AssertXML{
		t:   t,
		xml: xml,
	}
	if err != nil {
		assert.Failf(t, "data has invalid XML: %s", err.Error())
	} else {
		xmlAssert(body)
	}
}

// Node searches for XML node by XML Path Syntax. Returns struct for asserting the node values.
func (x *AssertXML) Node(path string) *AssertNode {
	p := xmlpath.MustCompile(path)
	value, found := p.String(x.xml)

	return &AssertNode{
		t:     x.t,
		found: found,
		path:  path,
		value: value,
	}
}

func (node *AssertNode) exists() bool {
	if !node.found {
		assert.Fail(node.t, fmt.Sprintf("failed to find XML node '%s'", node.path))
	}

	return node.found
}
