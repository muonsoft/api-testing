package apitest

import (
	"io"
	"net/http"

	"github.com/stretchr/testify/suite"
)

// APISuite is used to build testing suite of an API.
type APISuite struct {
	Handler http.Handler
	suite.Suite
}

// SendRequest is used to test http.Handler by passing httptest.ResponseRecorder to it.
// This function returns AssertResponse struct as a helper to build assertions on the response.
func (suite *APISuite) SendRequest(request *http.Request) *AssertResponse {
	suite.T().Helper()
	suite.checkHandler()
	return HandleRequest(suite.T(), suite.Handler, request)
}

// SendGET is an alias for HandleRequest that builds the GET request from url and options.
func (suite *APISuite) SendGET(url string, options ...RequestOption) *AssertResponse {
	suite.T().Helper()
	suite.checkHandler()
	return HandleGET(suite.T(), suite.Handler, url, options...)
}

// SendPOST is an alias for HandleRequest that builds the POST request from url, body and options.
func (suite *APISuite) SendPOST(url string, body io.Reader, options ...RequestOption) *AssertResponse {
	suite.T().Helper()
	suite.checkHandler()
	return HandlePOST(suite.T(), suite.Handler, url, body, options...)
}

// SendPUT is an alias for HandleRequest that builds the PUT request from url, body and options.
func (suite *APISuite) SendPUT(url string, body io.Reader, options ...RequestOption) *AssertResponse {
	suite.T().Helper()
	suite.checkHandler()
	return HandlePUT(suite.T(), suite.Handler, url, body, options...)
}

// SendPATCH is an alias for HandleRequest that builds the PATCH request from url, body and options.
func (suite *APISuite) SendPATCH(url string, body io.Reader, options ...RequestOption) *AssertResponse {
	suite.T().Helper()
	suite.checkHandler()
	return HandlePATCH(suite.T(), suite.Handler, url, body, options...)
}

// SendDELETE is an alias for HandleRequest that builds the DELETE request from url and options.
func (suite *APISuite) SendDELETE(url string, options ...RequestOption) *AssertResponse {
	suite.T().Helper()
	suite.checkHandler()
	return HandleDELETE(suite.T(), suite.Handler, url, options...)
}

func (suite *APISuite) checkHandler() {
	if suite.Handler == nil {
		suite.T().Fatal("APISuite.Handler is nil, please use set up method to set handler")
	}
}
