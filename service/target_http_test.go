package service

import (
	"context"
	"errors"
	"fmt"
	"mylab/cpagent/db"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TargetHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *TargetHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestTargetTestSuite(t *testing.T) {
	suite.Run(t, new(TargetHandlerTestSuite))
}

func (suite *TargetHandlerTestSuite) TestListTargetsSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ListTargets", mock.Anything).Return(
		[]db.Target{
			db.Target{Name: "test-target", ID: testUUID},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/targets",
		"/targets",
		"",
		listTargetHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"test-target","dye_id":"00000000-0000-0000-0000-000000000000"}]`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TargetHandlerTestSuite) TestListTargetsWhenDBFailure() {
	suite.dbMock.On("ListTargets", mock.Anything).Return(
		[]db.Target{},
		errors.New("error fetching targets"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/targets",
		"/targets",
		"",
		listTargetHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func makeHTTPCall(method, path, requestURL, body string, handlerFunc http.HandlerFunc) (recorder *httptest.ResponseRecorder) {
	// create a http request using the given parameters
	req, _ := http.NewRequest(method, requestURL, strings.NewReader(body))

	// test recorder created for capturing api responses
	recorder = httptest.NewRecorder()

	// create a router to serve the handler in test with the prepared request
	router := mux.NewRouter()
	router.HandleFunc(path, handlerFunc).Methods(method)

	// serve the request and write the response to recorder
	router.ServeHTTP(recorder, req)
	return
}

func makeHTTPCallWithHeader(method, path, requestURL, body string, headers map[string]string, handlerFunc http.HandlerFunc) (recorder *httptest.ResponseRecorder) {
	// create a http request using the given parameters
	req, _ := http.NewRequest(method, requestURL, strings.NewReader(body))

	for i, v := range headers {
		req.Header.Set(i, v)
	}
	// test recorder created for capturing api responses
	recorder = httptest.NewRecorder()

	// create a router to serve the handler in test with the prepared request
	router := mux.NewRouter()
	router.HandleFunc(path, handlerFunc).Methods(method)

	// serve the request and write the response to recorder
	router.ServeHTTP(recorder, req)
	return
}

func makeHTTPCallWithContext(ctxMap map[string]interface{}, method, path, requestURL, body string, handlerFunc http.HandlerFunc) (recorder *httptest.ResponseRecorder) {
	// create a http request using the given parameters
	req, _ := http.NewRequest(method, requestURL, strings.NewReader(body))

	ctx := req.Context()
	for key, value := range ctxMap {
		ctx = context.WithValue(ctx, key, value)
	}
	// test recorder created for capturing api responses
	recorder = httptest.NewRecorder()

	// create a router to serve the handler in test with the prepared request
	router := mux.NewRouter()
	router.HandleFunc(path, handlerFunc).Methods(method)

	// serve the request and write the response to recorder
	router.ServeHTTP(recorder, req.WithContext(ctx))
	return

}
