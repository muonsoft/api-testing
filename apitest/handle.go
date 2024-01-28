package apitest

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// RequestOption can be used to tune up http.Request.
type RequestOption func(r *http.Request)

// WithHeader option adds specific header to the request.
func WithHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Add(key, value)
	}
}

// WithContentType option adds Content-Type header to the request.
func WithContentType(contentType string) RequestOption {
	return func(r *http.Request) {
		r.Header.Add("Content-Type", contentType)
	}
}

// WithJSONContentType option adds Content-Type header to the request
// with "application/json" content type.
func WithJSONContentType() RequestOption {
	return func(r *http.Request) {
		r.Header.Add("Content-Type", "application/json; charset=utf-8")
	}
}

// WithCookie option adds a cookie to the request.
func WithCookie(cookie *http.Cookie) RequestOption {
	return func(r *http.Request) {
		r.AddCookie(cookie)
	}
}

// HandleRequest is used to test http.Handler by passing httptest.ResponseRecorder to it.
// This function returns ResponseAssertion struct as a helper to build assertions on the response.
func HandleRequest(t TestingT, handler http.Handler, request *http.Request) *ResponseAssertion {
	t.Helper()
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	return &ResponseAssertion{t: t, recorder: recorder}
}

// HandleGET is an alias for HandleRequest that builds the GET request from url and options.
func HandleGET(t TestingT, handler http.Handler, url string, options ...RequestOption) *ResponseAssertion {
	t.Helper()
	return handleRequest(t, handler, http.MethodGet, url, nil, options...)
}

// HandlePOST is an alias for HandleRequest that builds the POST request from url, body and options.
func HandlePOST(t TestingT, handler http.Handler, url string, body io.Reader, options ...RequestOption) *ResponseAssertion {
	t.Helper()
	return handleRequest(t, handler, http.MethodPost, url, body, options...)
}

// HandlePUT is an alias for HandleRequest that builds the PUT request from url, body and options.
func HandlePUT(t TestingT, handler http.Handler, url string, body io.Reader, options ...RequestOption) *ResponseAssertion {
	t.Helper()
	return handleRequest(t, handler, http.MethodPut, url, body, options...)
}

// HandlePATCH is an alias for HandleRequest that builds the PATCH request from url, body and options.
func HandlePATCH(t TestingT, handler http.Handler, url string, body io.Reader, options ...RequestOption) *ResponseAssertion {
	t.Helper()
	return handleRequest(t, handler, http.MethodPatch, url, body, options...)
}

// HandleDELETE is an alias for HandleRequest that builds the DELETE request from url and options.
func HandleDELETE(t TestingT, handler http.Handler, url string, options ...RequestOption) *ResponseAssertion {
	t.Helper()
	return handleRequest(t, handler, http.MethodDelete, url, nil, options...)
}

func handleRequest(t TestingT, handler http.Handler, method, url string, body io.Reader, options ...RequestOption) *ResponseAssertion {
	t.Helper()
	request := httptest.NewRequest(method, url, body)
	for _, setUpRequest := range options {
		setUpRequest(request)
	}
	return HandleRequest(t, handler, request)
}
