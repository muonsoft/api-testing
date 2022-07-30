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
//     func TestYourAPI(t testing.TB) {
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
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/xmlpath.v2"
)

// AssertXML - main structure that holds parsed XML.
type AssertXML struct {
	t   testing.TB
	xml *xmlpath.Node
}

// AssertNode - structure for asserting XML node.
type AssertNode struct {
	t     testing.TB
	found bool
	path  string
	value string
}

// XMLAssertFunc - callback function used for asserting XML nodes.
type XMLAssertFunc func(xml *AssertXML)

// FileHas loads XML from file and runs user callback for testing its nodes.
func FileHas(tb testing.TB, filename string, xmlAssert XMLAssertFunc) {
	tb.Helper()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		assert.Failf(tb, "failed to read file '%s': %s", filename, err.Error())
	} else {
		Has(tb, data, xmlAssert)
	}
}

// Has loads XML from byte slice and runs user callback for testing its nodes.
func Has(tb testing.TB, data []byte, xmlAssert XMLAssertFunc) {
	tb.Helper()
	xml, err := xmlpath.Parse(bytes.NewReader(data))
	body := &AssertXML{
		t:   tb,
		xml: xml,
	}
	if err != nil {
		assert.Failf(tb, "data has invalid XML: %s", err.Error())
	} else {
		xmlAssert(body)
	}
}

// Node searches for XML node by XML Path Syntax. Returns struct for asserting the node values.
func (x *AssertXML) Node(path string) *AssertNode {
	x.t.Helper()
	p := xmlpath.MustCompile(path)
	value, found := p.String(x.xml)

	return &AssertNode{
		t:     x.t,
		found: found,
		path:  path,
		value: value,
	}
}

// Nodef searches for XML node by XML Path Syntax. Returns struct for asserting the node values.
// It calculates path by applying fmt.Sprintf function.
func (x *AssertXML) Nodef(format string, a ...interface{}) *AssertNode {
	x.t.Helper()
	return x.Node(fmt.Sprintf(format, a...))
}

func (node *AssertNode) exists() bool {
	node.t.Helper()
	if !node.found {
		node.t.Errorf(`failed to find XML node "%s"`, node.path)
	}

	return node.found
}
