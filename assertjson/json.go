package assertjson

// WithJSON asserts that the JSON node has a string value with JSON.
func (a *StringAssertion) WithJSON(jsonAssert JSONAssertFunc) *StringAssertion {
	if a == nil {
		return nil
	}
	a.t.Helper()

	body := &AssertJSON{
		t:       a.t,
		message: a.message + `is string with JSON: `,
	}
	body.assert([]byte(a.value), jsonAssert)

	return nil
}
