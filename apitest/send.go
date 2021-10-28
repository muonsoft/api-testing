package apitest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// RequestOption can be used to tune up http.Request.
type RequestOption func(r *http.Request)

// WithHeader options adds specific header to the request.
func WithHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Set(key, value)
	}
}

// HandleRequest is used to test http.Handler by passing httptest.ResponseRecorder to it.
// This function returns AssertResponse struct as a helper to build assertions on the response.
func HandleRequest(tb testing.TB, handler http.Handler, request *http.Request) *AssertResponse {
	tb.Helper()
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	return &AssertResponse{t: tb, recorder: recorder}
}

// HandleGET is an alias for HandleRequest that builds the GET request from url and options.
func HandleGET(tb testing.TB, handler http.Handler, url string, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return handleRequest(tb, handler, http.MethodGet, url, nil, options...)
}

// HandlePOST is an alias for HandleRequest that builds the POST request from url, body and options.
func HandlePOST(tb testing.TB, handler http.Handler, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return handleRequest(tb, handler, http.MethodPost, url, body, options...)
}

// HandlePUT is an alias for HandleRequest that builds the PUT request from url, body and options.
func HandlePUT(tb testing.TB, handler http.Handler, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return handleRequest(tb, handler, http.MethodPut, url, body, options...)
}

// HandlePATCH is an alias for HandleRequest that builds the PATCH request from url, body and options.
func HandlePATCH(tb testing.TB, handler http.Handler, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return handleRequest(tb, handler, http.MethodPatch, url, body, options...)
}

// HandleDELETE is an alias for HandleRequest that builds the DELETE request from url and options.
func HandleDELETE(tb testing.TB, handler http.Handler, url string, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return handleRequest(tb, handler, http.MethodDelete, url, nil, options...)
}

func handleRequest(tb testing.TB, handler http.Handler, method, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	request := httptest.NewRequest(method, url, body)
	for _, setUpRequest := range options {
		setUpRequest(request)
	}
	return HandleRequest(tb, handler, request)
}
