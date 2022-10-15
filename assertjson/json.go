package assertjson

import "fmt"

// WithJSON asserts that the JSON node has a string value with JSON.
func (a *StringAssertion) WithJSON(jsonAssert JSONAssertFunc) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	body := &AssertJSON{
		t:       a.t,
		message: fmt.Sprintf(`failed asserting that JSON node "%s" is string with JSON`, a.path),
	}
	body.assert([]byte(a.value), jsonAssert)

	return nil
}
