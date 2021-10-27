package apitest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/muonsoft/api-testing/assertjson"
	"github.com/muonsoft/api-testing/assertxml"
)

type AssertResponse struct {
	t        testing.TB
	recorder *httptest.ResponseRecorder
}

func (r *AssertResponse) Recorder() *httptest.ResponseRecorder {
	r.t.Helper()
	return r.recorder
}

func (r *AssertResponse) Code() int {
	r.t.Helper()
	return r.recorder.Code
}

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

func (r *AssertResponse) IsOK() {
	r.t.Helper()
	r.HasCode(http.StatusOK)
}

func (r *AssertResponse) IsCreated() {
	r.t.Helper()
	r.HasCode(http.StatusCreated)
}

func (r *AssertResponse) IsAccepted() {
	r.t.Helper()
	r.HasCode(http.StatusCreated)
}

func (r *AssertResponse) HasNoContent() {
	r.t.Helper()
	r.HasCode(http.StatusNoContent)
	if r.recorder.Body.Len() > 0 {
		r.t.Errorf("response with no content unexpectedly has body with length %d", r.recorder.Body.Len())
	}
}

func (r *AssertResponse) IsBadRequest() {
	r.t.Helper()
	r.HasCode(http.StatusBadRequest)
}

func (r *AssertResponse) IsUnauthorized() {
	r.t.Helper()
	r.HasCode(http.StatusUnauthorized)
}

func (r *AssertResponse) IsForbidden() {
	r.t.Helper()
	r.HasCode(http.StatusForbidden)
}

func (r *AssertResponse) IsNotFound() {
	r.t.Helper()
	r.HasCode(http.StatusNotFound)
}

func (r *AssertResponse) IsMethodNotAllowed() {
	r.t.Helper()
	r.HasCode(http.StatusMethodNotAllowed)
}

func (r *AssertResponse) IsConflict() {
	r.t.Helper()
	r.HasCode(http.StatusConflict)
}

func (r *AssertResponse) IsUnsupportedMediaType() {
	r.t.Helper()
	r.HasCode(http.StatusUnsupportedMediaType)
}

func (r *AssertResponse) IsUnprocessableEntity() {
	r.t.Helper()
	r.HasCode(http.StatusUnprocessableEntity)
}

func (r *AssertResponse) IsInternalServerError() {
	r.t.Helper()
	r.HasCode(http.StatusInternalServerError)
}

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

func (r *AssertResponse) HasContentType(contentType string) {
	r.t.Helper()
	r.HasHeader("Content-Type", contentType)
}

func (r *AssertResponse) HasJSON(jsonAssert assertjson.JSONAssertFunc) {
	r.t.Helper()
	assertjson.Has(r.t, r.recorder.Body.Bytes(), jsonAssert)
}

func (r *AssertResponse) HasXML(xmlAssert assertxml.XMLAssertFunc) {
	r.t.Helper()
	assertxml.Has(r.t, r.recorder.Body.Bytes(), xmlAssert)
}
