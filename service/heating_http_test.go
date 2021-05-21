package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type HeatingHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *HeatingHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestHeatingTestSuite(t *testing.T) {
	suite.Run(t, new(HeatingHandlerTestSuite))
}

var testHeatingRecord = db.Heating{
	ID:          testUUID,
	Temperature: 80,
	Duration:    60,
	FollowTemp:  true,
	ProcessID:   testProcessUUID,
}

func (suite *HeatingHandlerTestSuite) TestCreateHeatingSuccess() {

	suite.dbMock.On("CreateHeating", mock.Anything, mock.Anything, recipeUUID).Return(testHeatingRecord, nil)

	body, _ := json.Marshal(testHeatingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/heating/{recipe_id}",
		"/heating/"+recipeUUID.String(),
		string(body),
		createHeatingHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestCreateHeatingFailure() {

	suite.dbMock.On("CreateHeating", mock.Anything, mock.Anything, recipeUUID).Return(db.Heating{}, responses.HeatingCreateError)

	body, _ := json.Marshal(testHeatingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/heating/{recipe_id}",
		"/heating/"+recipeUUID.String(),
		string(body),
		createHeatingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.HeatingCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestShowHeatingSuccess() {

	suite.dbMock.On("ShowHeating", mock.Anything, testProcessUUID).Return(testHeatingRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/heating/{id}",
		"/heating/"+testProcessUUID.String(),
		"",
		showHeatingHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testHeatingRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestShowHeatingFailure() {

	suite.dbMock.On("ShowHeating", mock.Anything, testProcessUUID).Return(db.Heating{}, responses.HeatingFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/heating/{id}",
		"/heating/"+testProcessUUID.String(),
		"",
		showHeatingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.HeatingFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestUpdateHeatingSuccess() {

	suite.dbMock.On("UpdateHeating", mock.Anything, testHeatingRecord).Return(nil)

	body, _ := json.Marshal(testHeatingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/heating/{id}",
		"/heating/"+testProcessUUID.String(),
		string(body),
		updateHeatingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.HeatingUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestUpdateHeatingFailure() {

	suite.dbMock.On("UpdateHeating", mock.Anything, testHeatingRecord).Return(responses.HeatingUpdateError)

	body, _ := json.Marshal(testHeatingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/heating/{id}",
		"/heating/"+testProcessUUID.String(),
		string(body),
		updateHeatingHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.HeatingUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestCreateHeatingInvalidUUID() {

	body, _ := json.Marshal(testHeatingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/heating/{recipe_id}",
		"/heating/"+invalidUUID,
		string(body),
		createHeatingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.RecipeIDInvalidError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestShowHeatingInvalidUUID() {

	recorder := makeHTTPCall(http.MethodGet,
		"/heating/{recipe_id}",
		"/heating/"+invalidUUID,
		"",
		showHeatingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HeatingHandlerTestSuite) TestUpdateHeatingInvalidUUID() {

	body, _ := json.Marshal(testHeatingRecord)
	recorder := makeHTTPCall(http.MethodPut,
		"/heating/{recipe_id}",
		"/heating/"+invalidUUID,
		string(body),
		updateHeatingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
