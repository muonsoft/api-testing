package apitest_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/muonsoft/api-testing/apitest"
	"github.com/muonsoft/api-testing/assertjson"
)

func ExampleSendGET() {
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("request method:", request.Method)
		fmt.Println("request url:", request.URL.String())

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(`{"ok":true}`))
	})

	response := apitest.SendGET(&testing.T{}, handler, "/example")

	response.IsOK()
	response.HasContentType("application/json")
	response.HasJSON(func(json *assertjson.AssertJSON) {
		json.Node("ok").IsTrue()
	})
	// Output:
	// request method: GET
	// request url: /example
}
