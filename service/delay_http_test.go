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
type DelayHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *DelayHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestDelayTestSuite(t *testing.T) {
	suite.Run(t, new(DelayHandlerTestSuite))
}

var testDelayRecord = db.Delay{
	ID:        testUUID,
	DelayTime: 60,
	ProcessID: testProcessUUID,
}

func (suite *DelayHandlerTestSuite) TestCreateDelaySuccess() {

	suite.dbMock.On("CreateDelay", mock.Anything, testDelayRecord).Return(testDelayRecord, nil)

	body, _ := json.Marshal(testDelayRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/delay",
		"/delay",
		string(body),
		createDelayHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *DelayHandlerTestSuite) TestCreateDelayFailure() {

	suite.dbMock.On("CreateDelay", mock.Anything, testDelayRecord).Return(db.Delay{}, responses.DelayCreateError)

	body, _ := json.Marshal(testDelayRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/delay",
		"/delay",
		string(body),
		createDelayHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.DelayCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *DelayHandlerTestSuite) TestShowDelaySuccess() {

	suite.dbMock.On("ShowDelay", mock.Anything, testProcessUUID).Return(testDelayRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/delay/{id}",
		"/delay/"+testProcessUUID.String(),
		"",
		showDelayHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testDelayRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *DelayHandlerTestSuite) TestShowDelayFailure() {

	suite.dbMock.On("ShowDelay", mock.Anything, testProcessUUID).Return(db.Delay{}, responses.DelayFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/delay/{id}",
		"/delay/"+testProcessUUID.String(),
		"",
		showDelayHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.DelayFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *DelayHandlerTestSuite) TestUpdateDelaySuccess() {

	suite.dbMock.On("UpdateDelay", mock.Anything, testDelayRecord).Return(nil)

	body, _ := json.Marshal(testDelayRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/delay/{id}",
		"/delay/"+testProcessUUID.String(),
		string(body),
		updateDelayHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.DelayUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *DelayHandlerTestSuite) TestUpdateDelayFailure() {

	suite.dbMock.On("UpdateDelay", mock.Anything, testDelayRecord).Return(responses.DelayUpdateError)

	body, _ := json.Marshal(testDelayRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/delay/{id}",
		"/delay/"+testProcessUUID.String(),
		string(body),
		updateDelayHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.DelayUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
