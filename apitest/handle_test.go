package apitest_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/muonsoft/api-testing/apitest"
	"github.com/muonsoft/api-testing/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testURL = "/users"

func TestHandleRequest(t *testing.T) {
	tests := []struct {
		name          string
		handle        func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion
		assertRequest func(t *testing.T, request *http.Request)
	}{
		{
			name: "HandleGET",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandleGET(t, handler, testURL)
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodGet, request.Method)
				assert.Equal(t, testURL, request.URL.String())
			},
		},
		{
			name: "HandleGET - with header",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandleGET(t, handler, testURL, apitest.WithHeader("X-Test", "value"))
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodGet, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				assert.Equal(t, "value", request.Header.Get("X-Test"))
			},
		},
		{
			name: "HandleGET - with header x2",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandleGET(t,
					handler, testURL,
					apitest.WithHeader("X-Test", "foo"),
					apitest.WithHeader("X-Test", "bar"),
				)
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodGet, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				assert.Equal(t, []string{"foo", "bar"}, request.Header["X-Test"])
			},
		},
		{
			name: "HandleGET - with content type",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandleGET(t, handler, testURL, apitest.WithContentType("text/html"))
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodGet, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				assert.Equal(t, "text/html", request.Header.Get("Content-Type"))
			},
		},
		{
			name: "HandleGET - with JSON content type",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandleGET(t, handler, testURL, apitest.WithJSONContentType())
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodGet, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				assert.Equal(t, "application/json; charset=utf-8", request.Header.Get("Content-Type"))
			},
		},
		{
			name: "HandleGET - with cookie",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandleGET(t, handler, testURL, apitest.WithCookie(&http.Cookie{
					Name:  "testCookie",
					Value: "testValue",
				}))
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodGet, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				cookie, err := request.Cookie("testCookie")
				require.NoError(t, err)
				assert.Equal(t, "testValue", cookie.Value)
			},
		},
		{
			name: "HandlePOST",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandlePOST(t, handler, testURL, strings.NewReader("testBody"))
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodPost, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				body, _ := io.ReadAll(request.Body)
				assert.Equal(t, "testBody", string(body))
			},
		},
		{
			name: "HandlePUT",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandlePUT(t, handler, testURL, strings.NewReader("testBody"))
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodPut, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				body, _ := io.ReadAll(request.Body)
				assert.Equal(t, "testBody", string(body))
			},
		},
		{
			name: "HandlePATCH",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandlePATCH(t, handler, testURL, strings.NewReader("testBody"))
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodPatch, request.Method)
				assert.Equal(t, testURL, request.URL.String())
				body, _ := io.ReadAll(request.Body)
				assert.Equal(t, "testBody", string(body))
			},
		},
		{
			name: "HandleDELETE",
			handle: func(t *mock.Tester, handler http.Handler) *apitest.ResponseAssertion {
				return apitest.HandleDELETE(t, handler, testURL)
			},
			assertRequest: func(t *testing.T, request *http.Request) {
				t.Helper()
				assert.Equal(t, http.MethodDelete, request.Method)
				assert.Equal(t, testURL, request.URL.String())
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tester := &mock.Tester{}
			handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				test.assertRequest(t, request)
				writer.WriteHeader(http.StatusOK)
			})
			test.handle(tester, handler).IsOK()
		})
	}
}
