package assertjson_test

import (
	"testing"

	"github.com/muonsoft/api-testing/assertjson"
	"github.com/muonsoft/api-testing/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestLinesHas_ValidLines(t *testing.T) {
	data := []byte(`{"key": "a", "id": 1}
{"other": "value", "id": 2}
`)

	ok := assertjson.LinesHas(t, data, func(lines *assertjson.AssertJSONLines) {
		lines.At(0).Node("key").Exists()
		lines.At(0).Node("key").IsString().EqualTo("a")
		lines.At(0).Node("id").IsInteger().EqualTo(1)
		lines.At(1).Node("other").IsString().EqualTo("value")
		lines.At(1).Node("id").IsInteger().EqualTo(2)
		lines.WithLength(2)
	})
	assert.True(t, ok)
}

func TestLinesHas_EmptyLinesSkipped(t *testing.T) {
	data := []byte(`{"key": "first"}

{"key": "second"}
`)

	ok := assertjson.LinesHas(t, data, func(lines *assertjson.AssertJSONLines) {
		lines.At(0).Node("key").IsString().EqualTo("first")
		lines.At(1).Node("key").IsString().EqualTo("second")
		lines.WithLength(2)
	})
	assert.True(t, ok)
}

func TestLinesHas_WithLengthMismatch(t *testing.T) {
	data := []byte(`{"a": 1}
{"b": 2}
`)

	tester := &mock.Tester{}
	assertjson.LinesHas(tester, data, func(lines *assertjson.AssertJSONLines) {
		lines.WithLength(3)
	})
	assert.True(t, tester.Failed())
}

func TestLinesHas_AtOutOfRangeNegative(t *testing.T) {
	data := []byte(`{"a": 1}
`)

	tester := &mock.Tester{}
	assertjson.LinesHas(tester, data, func(lines *assertjson.AssertJSONLines) {
		lines.At(-1).Node("a").Exists()
	})
	assert.True(t, tester.Failed())
}

func TestLinesHas_AtOutOfRangeBeyondLen(t *testing.T) {
	data := []byte(`{"a": 1}
`)

	tester := &mock.Tester{}
	assertjson.LinesHas(tester, data, func(lines *assertjson.AssertJSONLines) {
		lines.At(99).Node("a").Exists()
	})
	assert.True(t, tester.Failed())
}

func TestLinesHas_InvalidLineInMiddle(t *testing.T) {
	data := []byte(`{"a": 1}
not json
{"b": 2}
`)

	tester := &mock.Tester{}
	ok := assertjson.LinesHas(tester, data, func(lines *assertjson.AssertJSONLines) {
		lines.WithLength(3)
	})
	assert.False(t, ok)
	assert.True(t, tester.Failed())
}

func TestLines_MethodHas(t *testing.T) {
	data := []byte(`{"event": "created"}
{"event": "updated"}
`)

	lines := assertjson.Lines(t, data)
	ok := lines.Has(func(l *assertjson.AssertJSONLines) {
		l.At(0).Node("event").IsString().EqualTo("created")
		l.At(1).Node("event").IsString().EqualTo("updated")
		l.WithLength(2)
	})
	assert.True(t, ok)
}

func TestLines_AtAndLenDirectly(t *testing.T) {
	data := []byte(`{"id": 1}
{"id": 2}
`)

	lines := assertjson.Lines(t, data)
	assert.Equal(t, 2, lines.Len())
	lines.At(0).Node("id").IsInteger().EqualTo(1)
	lines.At(1).Node("id").IsInteger().EqualTo(2)
}

func TestFileLinesHas_FileNotFound(t *testing.T) {
	tester := &mock.Tester{}
	ok := assertjson.FileLinesHas(tester, "./nonexistent.ndjson", func(lines *assertjson.AssertJSONLines) {
		lines.WithLength(0)
	})
	assert.False(t, ok)
	assert.True(t, tester.Failed())
}
