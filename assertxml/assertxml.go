package assertxml

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/xmlpath.v2"
	"io/ioutil"
	"testing"
)

type AssertXML struct {
	t   *testing.T
	xml *xmlpath.Node
}

type AssertNode struct {
	t     *testing.T
	found bool
	path  string
	value string
}

type XMLAssertFunc func(xml *AssertXML)

func FileHas(t *testing.T, filename string, xmlAssert XMLAssertFunc) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		assert.Failf(t, "failed to read file '%s': %s", filename, err.Error())
	}
	Has(t, data, xmlAssert)
}

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
