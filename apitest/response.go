package apitest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/muonsoft/api-testing/assertjson"
	"github.com/muonsoft/api-testing/assertxml"
)

// AssertResponse is used to build assertions around httptest.ResponseRecorder.
type AssertResponse struct {
	t        testing.TB
	recorder *httptest.ResponseRecorder
}

// Recorder returns underlying httptest.ResponseRecorder.
func (r *AssertResponse) Recorder() *httptest.ResponseRecorder {
	r.t.Helper()
	return r.recorder
}

// Code returns HTTP status code of the response.
func (r *AssertResponse) Code() int {
	r.t.Helper()
	return r.recorder.Code
}

// HasCode asserts that the response has specific HTTP status code.
func (r *AssertResponse) HasCode(code int) {
	r.t.Helper()
	if r.recorder.Code != code {
		r.t.Errorf(
			"expected status code: %d (%s), actual is: %d (%s)",
			code,
			http.StatusText(code),
			r.recorder.Code,
			http.StatusText(r.recorder.Code),
		)
	}
}

// IsOK asserts that the response has an 200 Ok HTTP status code.
func (r *AssertResponse) IsOK() {
	r.t.Helper()
	r.HasCode(http.StatusOK)
}

// IsCreated asserts that the response has an 201 Created HTTP status code.
func (r *AssertResponse) IsCreated() {
	r.t.Helper()
	r.HasCode(http.StatusCreated)
}

// IsAccepted asserts that the response has an 202 Accepted HTTP status code.
func (r *AssertResponse) IsAccepted() {
	r.t.Helper()
	r.HasCode(http.StatusAccepted)
}

// HasNoContent asserts that the response has an 204 No Content HTTP status code and also checks that body is empty.
func (r *AssertResponse) HasNoContent() {
	r.t.Helper()
	r.HasCode(http.StatusNoContent)
	if r.recorder.Body.Len() > 0 {
		r.t.Errorf("response with no content unexpectedly has body with length %d", r.recorder.Body.Len())
	}
}

// IsBadRequest asserts that the response has an 400 Bad Request HTTP status code.
func (r *AssertResponse) IsBadRequest() {
	r.t.Helper()
	r.HasCode(http.StatusBadRequest)
}

// IsUnauthorized asserts that the response has an 401 Unauthorized HTTP status code.
func (r *AssertResponse) IsUnauthorized() {
	r.t.Helper()
	r.HasCode(http.StatusUnauthorized)
}

// IsForbidden asserts that the response has an 403 Forbidden HTTP status code.
func (r *AssertResponse) IsForbidden() {
	r.t.Helper()
	r.HasCode(http.StatusForbidden)
}

// IsNotFound asserts that the response has an 404 Not Found HTTP status code.
func (r *AssertResponse) IsNotFound() {
	r.t.Helper()
	r.HasCode(http.StatusNotFound)
}

// IsMethodNotAllowed asserts that the response has an 405 Method Not Allowed HTTP status code.
func (r *AssertResponse) IsMethodNotAllowed() {
	r.t.Helper()
	r.HasCode(http.StatusMethodNotAllowed)
}

// IsConflict asserts that the response has an 409 Conflict HTTP status code.
func (r *AssertResponse) IsConflict() {
	r.t.Helper()
	r.HasCode(http.StatusConflict)
}

// IsUnsupportedMediaType asserts that the response has an 415 Unsupported Media Type HTTP status code.
func (r *AssertResponse) IsUnsupportedMediaType() {
	r.t.Helper()
	r.HasCode(http.StatusUnsupportedMediaType)
}

// IsUnprocessableEntity asserts that the response has an 422 Unprocessable Entity HTTP status code.
func (r *AssertResponse) IsUnprocessableEntity() {
	r.t.Helper()
	r.HasCode(http.StatusUnprocessableEntity)
}

// IsInternalServerError asserts that the response has an 500 Internal Server Error HTTP status code.
func (r *AssertResponse) IsInternalServerError() {
	r.t.Helper()
	r.HasCode(http.StatusInternalServerError)
}

// HasHeader asserts that the response contains specific header with key and value.
func (r *AssertResponse) HasHeader(key, value string) {
	r.t.Helper()
	header := r.recorder.Header().Get(key)
	if header == "" {
		r.t.Errorf(`response does not contain header "%s"`, header)
		return
	}
	if value != header {
		r.t.Errorf(
			`response header "%s" is expected to be "%s", actual is "%s"`,
			key,
			value,
			header,
		)
	}
}

// HasContentType asserts that the response contains Content-Type header with specific value.
func (r *AssertResponse) HasContentType(contentType string) {
	r.t.Helper()
	r.HasHeader("Content-Type", contentType)
}

// HasJSON asserts that the response body contains JSON and runs JSON assertions by callback function.
func (r *AssertResponse) HasJSON(jsonAssert assertjson.JSONAssertFunc) {
	r.t.Helper()
	assertjson.Has(r.t, r.recorder.Body.Bytes(), jsonAssert)
}

// HasXML asserts that the response body contains XML and runs XML assertions by callback function.
func (r *AssertResponse) HasXML(xmlAssert assertxml.XMLAssertFunc) {
	r.t.Helper()
	assertxml.Has(r.t, r.recorder.Body.Bytes(), xmlAssert)
}

// Print prints response headers and body to console. Use it for debug purposes.
func (r *AssertResponse) Print() {
	r.t.Helper()
	headers := r.printHeaders()
	r.t.Log(headers + r.recorder.Body.String())
}

// PrintJSON prints response headers and indented JSON body to console. Use it for debug purposes.
func (r *AssertResponse) PrintJSON() {
	r.t.Helper()
	headers := r.printHeaders()
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

func (r *AssertResponse) printHeaders() string {
	r.t.Helper()
	s := &strings.Builder{}
	s.WriteString("\n")
	// nolint:bodyclose
	fmt.Fprintln(s, r.recorder.Result().Proto, r.recorder.Result().Status)
	for name, values := range r.recorder.Header() {
		fmt.Fprintf(s, "%s: %s\n", name, strings.Join(values, "; "))
	}

	return s.String()
}
