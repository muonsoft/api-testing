package apitest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RequestOption func(r *http.Request)

func WithHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Set(key, value)
	}
}

func SendRequest(tb testing.TB, handler http.Handler, request *http.Request) *AssertResponse {
	tb.Helper()
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	return &AssertResponse{t: tb, recorder: recorder}
}

func SendGET(tb testing.TB, handler http.Handler, url string, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return sendRequest(tb, handler, http.MethodGet, url, nil, options...)
}

func SendPOST(tb testing.TB, handler http.Handler, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return sendRequest(tb, handler, http.MethodPost, url, body, options...)
}

func SendPUT(tb testing.TB, handler http.Handler, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return sendRequest(tb, handler, http.MethodPut, url, body, options...)
}

func SendPATCH(tb testing.TB, handler http.Handler, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return sendRequest(tb, handler, http.MethodPatch, url, body, options...)
}

func SendDELETE(tb testing.TB, handler http.Handler, url string, options ...RequestOption) *AssertResponse {
	tb.Helper()
	return sendRequest(tb, handler, http.MethodPatch, url, nil, options...)
}

func sendRequest(tb testing.TB, handler http.Handler, method, url string, body io.Reader, options ...RequestOption) *AssertResponse {
	tb.Helper()
	request := httptest.NewRequest(method, url, body)
	for _, setUpRequest := range options {
		setUpRequest(request)
	}
	return SendRequest(tb, handler, request)
}
