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
type ShakingHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *ShakingHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestShakingTestSuite(t *testing.T) {
	suite.Run(t, new(ShakingHandlerTestSuite))
}

var testShakingRecord = db.Shaker{
	ID:          testUUID,
	Temperature: 80,
	FollowTemp:  true,
	RPM1:        6500,
	Time1:       80,
	ProcessID:   testProcessUUID,
}

func (suite *ShakingHandlerTestSuite) TestCreateShakingSuccess() {

	suite.dbMock.On("CreateShaking", mock.Anything, mock.Anything, mock.Anything).Return(testShakingRecord, nil)

	body, _ := json.Marshal(testShakingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/shaking/{recipe_id}",
		"/shaking/"+recipeUUID.String(),
		string(body),
		createShakingHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ShakingHandlerTestSuite) TestCreateShakingFailure() {

	suite.dbMock.On("CreateShaking", mock.Anything, mock.Anything, recipeUUID).Return(db.Shaker{}, responses.ShakingCreateError)

	body, _ := json.Marshal(testShakingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/shaking/{recipe_id}",
		"/shaking/"+recipeUUID.String(),
		string(body),
		createShakingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.ShakingCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ShakingHandlerTestSuite) TestShowShakingSuccess() {

	suite.dbMock.On("ShowShaking", mock.Anything, testProcessUUID).Return(testShakingRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/shaking/{id}",
		"/shaking/"+testProcessUUID.String(),
		"",
		showShakingHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testShakingRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ShakingHandlerTestSuite) TestShowShakingFailure() {

	suite.dbMock.On("ShowShaking", mock.Anything, testProcessUUID).Return(testShakingRecord, responses.ShakingFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/shaking/{id}",
		"/shaking/"+testProcessUUID.String(),
		"",
		showShakingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.ShakingFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ShakingHandlerTestSuite) TestUpdateShakingSuccess() {

	suite.dbMock.On("UpdateShaking", mock.Anything, mock.Anything).Return(nil)

	body, _ := json.Marshal(testShakingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/shaking/{id}",
		"/shaking/"+testProcessUUID.String(),
		string(body),
		updateShakingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.ShakingUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ShakingHandlerTestSuite) TestUpdateShakingFailure() {

	suite.dbMock.On("UpdateShaking", mock.Anything, mock.Anything).Return(responses.ShakingUpdateError)

	body, _ := json.Marshal(testShakingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/shaking/{id}",
		"/shaking/"+testProcessUUID.String(),
		string(body),
		updateShakingHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.ShakingUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
