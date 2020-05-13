package service

import (
	"errors"
	"mylab/mylabdiscoveries/db"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TargetHandlerTestSuite))
}

func (suite *TargetHandlerTestSuite) TestListUsersSuccess() {
	suite.dbMock.On("ListTarget",mock.Anything).Return(
		[]db.Target{
			db.Target{Name: "test-target", ID: 1},
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

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `[{"id":1,"name":"test-target","dye_id":0,"well_id":0,"template_id":0,"ct":0}]`, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TargetHandlerTestSuite) TestListUsersWhenDBFailure() {
	suite.dbMock.On("ListTarget", mock.Anything).Return(
		[]db.Target{},
		errors.New("error fetching user records"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/targets",
		"/targets",
		"",
		listTargetHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(),"", recorder.Body.String())
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
