package apitest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muonsoft/api-testing/apitest"
	"github.com/muonsoft/api-testing/assertjson"
	"github.com/muonsoft/api-testing/internal/mock"
)

func TestAssertResponse(t *testing.T) {
	tests := []struct {
		name          string
		writeResponse func(w http.ResponseWriter)
		assert        func(*apitest.ResponseAssertion)
		wantMessages  []string
	}{
		{
			name: "HasCode passed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasCode(http.StatusOK)
			},
		},
		{
			name: "HasCode failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasCode(http.StatusBadRequest)
			},
			wantMessages: []string{
				"expected status code: 400 (Bad Request), actual is: 200 (OK)",
				"HTTP/1.1 200 OK",
			},
		},
		{
			name: "IsOK failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsOK()
			},
			wantMessages: []string{
				"expected status code: 200 (OK), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsCreated failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsCreated()
			},
			wantMessages: []string{
				"expected status code: 201 (Created), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsAccepted failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsAccepted()
			},
			wantMessages: []string{
				"expected status code: 202 (Accepted), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "HasNoContent failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasNoContent()
			},
			wantMessages: []string{
				"expected status code: 204 (No Content), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsBadRequest failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsBadRequest()
			},
			wantMessages: []string{
				"expected status code: 400 (Bad Request), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsUnauthorized failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsUnauthorized()
			},
			wantMessages: []string{
				"expected status code: 401 (Unauthorized), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsForbidden failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsForbidden()
			},
			wantMessages: []string{
				"expected status code: 403 (Forbidden), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsNotFound failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsNotFound()
			},
			wantMessages: []string{
				"expected status code: 404 (Not Found), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsMethodNotAllowed failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsMethodNotAllowed()
			},
			wantMessages: []string{
				"expected status code: 405 (Method Not Allowed), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsConflict failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsConflict()
			},
			wantMessages: []string{
				"expected status code: 409 (Conflict), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsUnsupportedMediaType failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsUnsupportedMediaType()
			},
			wantMessages: []string{
				"expected status code: 415 (Unsupported Media Type), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsUnprocessableEntity failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsUnprocessableEntity()
			},
			wantMessages: []string{
				"expected status code: 422 (Unprocessable Entity), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "IsInternalServerError failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsInternalServerError()
			},
			wantMessages: []string{
				"expected status code: 500 (Internal Server Error), actual is: 200 (OK)",
				"HTTP/1.1 200 OK",
			},
		},
		{
			name: "IsBadGateway failed",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.IsBadGateway()
			},
			wantMessages: []string{
				"expected status code: 502 (Bad Gateway), actual is: 500 (Internal Server Error)",
				"HTTP/1.1 500 Internal Server Error",
			},
		},
		{
			name: "HasHeader passed",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("X-Test", "value")
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasHeader("X-Test", "value")
			},
		},
		{
			name: "HasHeader - missing header",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasHeader("X-Test", "value")
			},
			wantMessages: []string{
				"response does not contain header \"X-Test\"",
			},
		},
		{
			name: "HasHeader - invalid header",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("X-Test", "invalid")
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasHeader("X-Test", "value")
			},
			wantMessages: []string{
				`response header "X-Test" is expected to be "value", actual is "invalid"`,
			},
		},
		{
			name: "HasContentType passed",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasContentType("application/json")
			},
		},
		{
			name: "HasContentType failed",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasContentType("application/json")
			},
			wantMessages: []string{
				`response header "Content-Type" is expected to be "application/json", actual is "text/html"`,
			},
		},
		{
			name: "HasJSON passed",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"ok":true}`))
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasJSON(func(json *assertjson.AssertJSON) {
					json.Node("ok").IsTrue()
				})
			},
		},
		{
			name: "HasJSON failed",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"ok":true}`))
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasJSON(func(json *assertjson.AssertJSON) {
					json.Node("ok").IsFalse()
				})
			},
			wantMessages: []string{
				`failed asserting that JSON node "ok" is false`,
			},
		},
		{
			name: "HasCookie passed",
			writeResponse: func(w http.ResponseWriter) {
				cookie := &http.Cookie{
					Name:  "testCookie",
					Value: "testValue",
				}
				http.SetCookie(w, cookie)
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasCookie("testCookie").WithValue().EqualTo("testValue")
			},
		},
		{
			name: "HasCookie missing cookie",
			writeResponse: func(w http.ResponseWriter) {
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasCookie("testCookie")
			},
			wantMessages: []string{
				`response does not contain cookie "testCookie"`,
			},
		},
		{
			name: "HasCookie.WithJWT passed",
			writeResponse: func(w http.ResponseWriter) {
				cookie := &http.Cookie{
					Name:  "testCookie",
					Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				}
				http.SetCookie(w, cookie)
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasCookie("testCookie").
					WithValue().
					WithJWT(getJWTSecret).
					WithPayload(func(json *assertjson.AssertJSON) {
						json.Node("name").IsString().EqualTo("John Doe")
					})
			},
		},
		{
			name: "HasCookie failed",
			writeResponse: func(w http.ResponseWriter) {
				cookie := &http.Cookie{
					Name:  "testCookie",
					Value: "invalid",
				}
				http.SetCookie(w, cookie)
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.HasCookie("testCookie").WithValue().EqualTo("testValue")
			},
			wantMessages: []string{
				`failed asserting that cookie "testCookie" value equal to "testValue", actual is "invalid"`,
			},
		},
		{
			name: "Print",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(`content body`))
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.Print()
			},
			wantMessages: []string{
				"HTTP/1.1 200 OK\nContent-Type: text/html\ncontent body",
			},
		},
		{
			name: "PrintJSON",
			writeResponse: func(w http.ResponseWriter) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"ok":true}`))
				w.WriteHeader(http.StatusOK)
			},
			assert: func(response *apitest.ResponseAssertion) {
				response.PrintJSON()
			},
			wantMessages: []string{
				"HTTP/1.1 200 OK\nContent-Type: application/json\n{\n\t\"ok\": true\n}",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tester := &mock.Tester{}
			handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				test.writeResponse(writer)
			})
			response := apitest.HandleRequest(tester, handler, httptest.NewRequest(http.MethodGet, "/", nil))

			test.assert(response)

			tester.AssertContains(t, test.wantMessages)
		})
	}
}

const tokenSecret = "your-256-bit-secret"

func getJWTSecret(_ *jwt.Token) (interface{}, error) {
	return []byte(tokenSecret), nil
}
