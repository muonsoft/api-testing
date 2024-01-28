package apitest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/muonsoft/api-testing/assertjson"
	"github.com/muonsoft/api-testing/assertxml"
	"github.com/stretchr/testify/assert"
)

// ResponseAssertion is used to build assertions around httptest.ResponseRecorder.
type ResponseAssertion struct {
	t        TestingT
	recorder *httptest.ResponseRecorder
}

// Recorder returns underlying httptest.ResponseRecorder.
func (r *ResponseAssertion) Recorder() *httptest.ResponseRecorder {
	r.t.Helper()
	return r.recorder
}

// Code returns HTTP status code of the response.
func (r *ResponseAssertion) Code() int {
	r.t.Helper()
	return r.recorder.Code
}

// Header returns HTTP headers of the response.
func (r *ResponseAssertion) Header() http.Header {
	r.t.Helper()
	return r.recorder.Header()
}

// Cookies returns HTTP cookies of the response.
func (r *ResponseAssertion) Cookies() []*http.Cookie {
	r.t.Helper()
	response := http.Response{Header: r.recorder.Header()}
	return response.Cookies()
}

// HasCode asserts that the response has specific HTTP status code.
func (r *ResponseAssertion) HasCode(code int) {
	r.t.Helper()
	if r.recorder.Code != code {
		assert.Fail(r.t, fmt.Sprintf(
			"expected status code: %d (%s), actual is: %d (%s)",
			code,
			http.StatusText(code),
			r.recorder.Code,
			http.StatusText(r.recorder.Code),
		))
		r.logResponse()
	}
}

// IsOK asserts that the response has an 200 Ok HTTP status code.
func (r *ResponseAssertion) IsOK() {
	r.t.Helper()
	r.HasCode(http.StatusOK)
}

// IsCreated asserts that the response has an 201 Created HTTP status code.
func (r *ResponseAssertion) IsCreated() {
	r.t.Helper()
	r.HasCode(http.StatusCreated)
}

// IsAccepted asserts that the response has an 202 Accepted HTTP status code.
func (r *ResponseAssertion) IsAccepted() {
	r.t.Helper()
	r.HasCode(http.StatusAccepted)
}

// HasNoContent asserts that the response has an 204 No Content HTTP status code and also checks that body is empty.
func (r *ResponseAssertion) HasNoContent() {
	r.t.Helper()
	r.HasCode(http.StatusNoContent)
	if r.recorder.Body.Len() > 0 {
		assert.Fail(r.t, fmt.Sprintf(
			"response with no content unexpectedly has body with length %d",
			r.recorder.Body.Len()),
		)
	}
}

// IsBadRequest asserts that the response has an 400 Bad Request HTTP status code.
func (r *ResponseAssertion) IsBadRequest() {
	r.t.Helper()
	r.HasCode(http.StatusBadRequest)
}

// IsUnauthorized asserts that the response has an 401 Unauthorized HTTP status code.
func (r *ResponseAssertion) IsUnauthorized() {
	r.t.Helper()
	r.HasCode(http.StatusUnauthorized)
}

// IsForbidden asserts that the response has an 403 Forbidden HTTP status code.
func (r *ResponseAssertion) IsForbidden() {
	r.t.Helper()
	r.HasCode(http.StatusForbidden)
}

// IsNotFound asserts that the response has an 404 Not Found HTTP status code.
func (r *ResponseAssertion) IsNotFound() {
	r.t.Helper()
	r.HasCode(http.StatusNotFound)
}

// IsMethodNotAllowed asserts that the response has an 405 Method Not Allowed HTTP status code.
func (r *ResponseAssertion) IsMethodNotAllowed() {
	r.t.Helper()
	r.HasCode(http.StatusMethodNotAllowed)
}

// IsConflict asserts that the response has an 409 Conflict HTTP status code.
func (r *ResponseAssertion) IsConflict() {
	r.t.Helper()
	r.HasCode(http.StatusConflict)
}

// IsUnsupportedMediaType asserts that the response has an 415 Unsupported Media Type HTTP status code.
func (r *ResponseAssertion) IsUnsupportedMediaType() {
	r.t.Helper()
	r.HasCode(http.StatusUnsupportedMediaType)
}

// IsUnprocessableEntity asserts that the response has an 422 Unprocessable Entity HTTP status code.
func (r *ResponseAssertion) IsUnprocessableEntity() {
	r.t.Helper()
	r.HasCode(http.StatusUnprocessableEntity)
}

// IsInternalServerError asserts that the response has an 500 Internal Server Error HTTP status code.
func (r *ResponseAssertion) IsInternalServerError() {
	r.t.Helper()
	r.HasCode(http.StatusInternalServerError)
}

// IsBadGateway asserts that the response has an 502 Bad Gateway HTTP status code.
func (r *ResponseAssertion) IsBadGateway() {
	r.t.Helper()
	r.HasCode(http.StatusBadGateway)
}

// HasHeader asserts that the response contains specific header with key and value.
func (r *ResponseAssertion) HasHeader(key, value string) {
	r.t.Helper()
	header := r.recorder.Header().Get(key)
	if header == "" {
		assert.Fail(r.t, fmt.Sprintf(`response does not contain header "%s"`, key))
		return
	}
	if value != header {
		assert.Fail(r.t, fmt.Sprintf(
			`response header "%s" is expected to be "%s", actual is "%s"`,
			key,
			value,
			header,
		))
	}
}

// HasContentType asserts that the response contains Content-Type header with specific value.
func (r *ResponseAssertion) HasContentType(contentType string) {
	r.t.Helper()
	r.HasHeader("Content-Type", contentType)
}

// HasJSON asserts that the response body contains JSON and runs JSON assertions by callback function.
func (r *ResponseAssertion) HasJSON(jsonAssert assertjson.JSONAssertFunc) {
	r.t.Helper()
	assertjson.Has(r.t, r.recorder.Body.Bytes(), jsonAssert)
}

// HasXML asserts that the response body contains XML and runs XML assertions by callback function.
func (r *ResponseAssertion) HasXML(xmlAssert assertxml.XMLAssertFunc) {
	r.t.Helper()
	assertxml.Has(r.t, r.recorder.Body.Bytes(), xmlAssert)
}

// HasCookie asserts that the response contains specific cookie in Set-Cookie header.
func (r *ResponseAssertion) HasCookie(name string) *CookieAssertion {
	r.t.Helper()

	for _, cookie := range r.Cookies() {
		if cookie.Name == name {
			return &CookieAssertion{
				t:      r.t,
				cookie: cookie,
			}
		}
	}

	assert.Fail(r.t, fmt.Sprintf(`response does not contain cookie "%s"`, name))

	return nil
}

// Print prints response headers and body to console. Use it for debug purposes.
func (r *ResponseAssertion) Print() {
	r.t.Helper()
	headers := r.formatHeaders()
	r.t.Log(headers + r.recorder.Body.String())
}

// PrintJSON prints response headers and indented JSON body to console. Use it for debug purposes.
func (r *ResponseAssertion) PrintJSON() {
	r.t.Helper()
	headers := r.formatHeaders()
	var body interface{}
	err := json.Unmarshal(r.recorder.Body.Bytes(), &body)
	if err != nil {
		r.t.Log(headers)
		r.t.Error("invalid JSON:", err)
		return
	}
	printableJSON, _ := json.MarshalIndent(body, "", "\t")
	r.t.Log(headers + string(printableJSON))
}

func (r *ResponseAssertion) logResponse() {
	r.t.Helper()
	headers := r.formatHeaders()
	var body interface{}
	err := json.Unmarshal(r.recorder.Body.Bytes(), &body)
	if err != nil {
		r.t.Log(headers + r.recorder.Body.String())
		return
	}
	printableJSON, _ := json.MarshalIndent(body, "", "\t")
	r.t.Log(headers + string(printableJSON))
}

func (r *ResponseAssertion) formatHeaders() string {
	r.t.Helper()
	s := &strings.Builder{}
	s.WriteString("\n")
	//nolint:bodyclose
	fmt.Fprintln(s, r.recorder.Result().Proto, r.recorder.Result().Status)
	for name, values := range r.recorder.Header() {
		fmt.Fprintf(s, "%s: %s\n", name, strings.Join(values, "; "))
	}

	return s.String()
}
