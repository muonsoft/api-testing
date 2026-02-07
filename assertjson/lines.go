package assertjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/muonsoft/api-testing/internal/js"
	"github.com/stretchr/testify/assert"
)

// AssertJSONLines holds parsed JSON Lines for assertion.
type AssertJSONLines struct {
	t       TestingT
	message string
	lines   []interface{}
}

// LinesAssertFunc is a callback function used for asserting JSON Lines.
type LinesAssertFunc func(lines *AssertJSONLines)

// LinesHas parses data as JSON Lines and runs the callback for asserting lines.
// Returns false if t has already failed.
func LinesHas(t TestingT, data []byte, linesAssert LinesAssertFunc) bool {
	t.Helper()
	lines := parseLines(t, data)
	if lines == nil {
		return false
	}
	body := &AssertJSONLines{t: t, lines: lines}
	linesAssert(body)
	return !t.Failed()
}

// FileLinesHas reads a file as JSON Lines and runs the callback for asserting lines.
func FileLinesHas(t TestingT, filename string, linesAssert LinesAssertFunc) bool {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		assert.Fail(t, fmt.Sprintf(`failed to read file "%s": %s`, filename, err.Error()))
		return false
	}

	return LinesHas(t, data, linesAssert)
}

// Lines parses data as JSON Lines and returns AssertJSONLines for further use.
// Use Has() to run assertions in a callback, or call At(), Len(), WithLength() directly.
func Lines(t TestingT, data []byte) *AssertJSONLines {
	t.Helper()
	lines := parseLines(t, data)
	if lines == nil {
		return &AssertJSONLines{t: t, lines: []interface{}{}}
	}
	return &AssertJSONLines{t: t, lines: lines}
}

// FileLines reads a file as JSON Lines and returns AssertJSONLines.
func FileLines(t TestingT, filename string) *AssertJSONLines {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		assert.Fail(t, fmt.Sprintf(`failed to read file "%s": %s`, filename, err.Error()))
		return &AssertJSONLines{t: t, lines: []interface{}{}}
	}

	return Lines(t, data)
}

// Has runs the callback with this AssertJSONLines. Returns false if t has already failed.
func (l *AssertJSONLines) Has(linesAssert LinesAssertFunc) bool {
	l.t.Helper()
	linesAssert(l)
	return !l.t.Failed()
}

// At returns AssertJSON for the line at the given index (0-based).
// Fails the test if index is out of range.
func (l *AssertJSONLines) At(index int) *AssertJSON {
	l.t.Helper()
	if index < 0 || index >= len(l.lines) {
		l.fail(fmt.Sprintf(
			`JSON Lines index %d is out of range (lines count: %d)`,
			index,
			len(l.lines),
		))
		return &AssertJSON{t: l.t, path: js.NewPath(js.ArrayIndex(index)), data: nil}
	}
	return &AssertJSON{
		t:       l.t,
		message: l.message,
		path:    js.NewPath(js.ArrayIndex(index)),
		data:    l.lines[index],
	}
}

// Len returns the number of parsed lines.
func (l *AssertJSONLines) Len() int {
	l.t.Helper()
	return len(l.lines)
}

// WithLength asserts that the number of lines equals expected. Returns l for chaining.
func (l *AssertJSONLines) WithLength(expected int) *AssertJSONLines {
	l.t.Helper()
	if len(l.lines) != expected {
		l.fail(fmt.Sprintf(
			`JSON Lines count is %d, actual is %d`,
			expected,
			len(l.lines),
		))
	}
	return l
}

func (l *AssertJSONLines) fail(message string, msgAndArgs ...interface{}) {
	l.t.Helper()
	assert.Fail(l.t, l.message+message, msgAndArgs...)
}

func parseLines(t TestingT, data []byte) []interface{} {
	t.Helper()
	rawLines := bytes.Split(data, []byte("\n"))
	lines := make([]interface{}, 0, len(rawLines))
	lineNum := 0
	for _, line := range rawLines {
		lineNum++
		line = bytes.TrimSuffix(bytes.TrimSpace(line), []byte("\r"))
		if len(line) == 0 {
			continue
		}
		var value interface{}
		if err := json.Unmarshal(line, &value); err != nil {
			assert.Fail(t, fmt.Sprintf("JSON Lines line %d: %s", lineNum, err.Error()))
			return nil
		}
		lines = append(lines, value)
	}
	return lines
}
