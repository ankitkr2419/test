package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ProcessHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *ProcessHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestProcessTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessHandlerTestSuite))
}

var testUUID = uuid.New()
var testProcessUUID = uuid.New()
var recipeUUID = uuid.New()

const testName = "testName"
const testType = db.ProcessType("testType")
const sequenceNumber int64 = 1

var testProcessRecord = db.Process{
	ID:             testUUID,
	Name:           testName,
	Type:           testType,
	SequenceNumber: sequenceNumber,
	RecipeID:       recipeUUID,
}

var listProcesses = []db.Process{
	testProcessRecord,
}

func (suite *ProcessHandlerTestSuite) TestCreateProcessSuccess() {

	suite.dbMock.On("CreateProcess", mock.Anything, testProcessRecord).Return(testProcessRecord, nil)
	body, _ := json.Marshal(testProcessRecord)

	recorder := makeHTTPCall(http.MethodPost,
		"/processes",
		"/processes",
		string(body),
		createProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestCreateProcessFailure() {

	suite.dbMock.On("CreateProcess", mock.Anything, testProcessRecord).Return(db.Process{}, responses.ProcessCreateError)

	body, _ := json.Marshal(testProcessRecord)

	recorder := makeHTTPCall(http.MethodPost,
		"/processes",
		"/processes",
		string(body),
		createProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	
	err := ErrObj{Err: responses.ProcessCreateError.Error()}

	output, _ := json.Marshal(err)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestListProcessSuccess() {

	suite.dbMock.On("ListProcesses", mock.Anything, testUUID).Return(
		listProcesses, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/recipe/{id}/processes",
		"/recipe/"+testUUID.String()+"/processes",
		"",
		listProcessesHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(listProcesses)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestListProcessUUIDParseError() {

	recorder := makeHTTPCall(http.MethodGet,
		"/recipe/{id}/processes",
		"/recipe/"+invalidUUID+"/processes",
		"",
		listProcessesHandler(Dependencies{Store: suite.dbMock}),
	)
	
	err := ErrObj{Err: responses.UUIDParseError.Error()}

	output, _ := json.Marshal(err)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestListProcessFailure() {
	suite.dbMock.On("ListProcesses", mock.Anything, testUUID).Return(
		[]db.Process{}, responses.ProcessFetchError)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/recipe/{id}/processes",
		"/recipe/"+testUUID.String()+"/processes",
		"",
		listProcessesHandler(Dependencies{Store: suite.dbMock}),
	)

	err := ErrObj{Err: responses.ProcessFetchError.Error()}

	output, _ := json.Marshal(err)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestShowProcessSuccess() {

	suite.dbMock.On("ShowProcess", mock.Anything, testUUID).Return(testProcessRecord, nil)
	body, _ := json.Marshal(testProcessRecord)

	recorder := makeHTTPCall(http.MethodGet,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		showProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestShowProcessUUIDParseError() {

	recorder := makeHTTPCall(http.MethodGet,
		"/processes/{id}",
		"/processes/"+invalidUUID,
		"",
		showProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal(ErrObj{Err:responses.UUIDParseError.Error()})

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestShowProcessFailure() {

	suite.dbMock.On("ShowProcess", mock.Anything, mock.Anything).Return(db.Process{}, responses.ProcessFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		showProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal(ErrObj{Err:responses.ProcessFetchError.Error()})

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestUpdateProcessSuccess() {

	suite.dbMock.On("UpdateProcess", mock.Anything, testProcessRecord).Return(nil)

	body, _ := json.Marshal(testProcessRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		string(body),
		updateProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(MsgObj{Msg:responses.ProcessUpdateSuccess})

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestUpdateProcessUUIDParseError() {

	body, _ := json.Marshal(testProcessRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/processes/{id}",
		"/processes/"+invalidUUID,
		string(body),
		updateProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(ErrObj{Err:responses.UUIDParseError.Error()})

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestUpdateProcessFailure() {

	suite.dbMock.On("UpdateProcess", mock.Anything, mock.Anything).Return(responses.ProcessUpdateError)

	body, _ := json.Marshal(testProcessRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		string(body),
		updateProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(ErrObj{Err:responses.ProcessUpdateError.Error()})

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestDeleteProcessSuccess() {

	suite.dbMock.On("DeleteProcess", mock.Anything, mock.Anything).Return(nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		deleteProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(MsgObj{Msg:responses.ProcessDeleteSuccess})

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output) , recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestDeleteProcessFailure() {

	suite.dbMock.On("DeleteProcess", mock.Anything, mock.Anything).Return(responses.ProcessDeleteError)

	recorder := makeHTTPCall(http.MethodDelete,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		deleteProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(ErrObj{Err:responses.ProcessDeleteError.Error()})

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestDeleteProcessUUIDParseError() {

	recorder := makeHTTPCall(http.MethodDelete,
		"/processes/{id}",
		"/processes/"+invalidUUID,
		"",
		deleteProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(ErrObj{Err:responses.UUIDParseError.Error()})

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

