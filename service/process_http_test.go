package service

import (
	"fmt"
	"mylab/cpagent/db"
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
var testName = "testName"
var testType = "testType"
var recipeUUID = uuid.New()
var sequenceNumber int64 = 1

func (suite *ProcessHandlerTestSuite) TestCreateProcessSuccess() {

	suite.dbMock.On("CreateProcess", mock.Anything, mock.Anything).Return(db.Process{
		ID:             testUUID,
		Name:           testName,
		Type:           testType,
		RecipeID:       recipeUUID,
		SequenceNumber: sequenceNumber,
	}, nil)

	body := fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","recipe_id":"%s","sequence_num":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testName, testType, recipeUUID, sequenceNumber)
	recorder := makeHTTPCall(http.MethodPost,
		"/processes",
		"/processes",
		body,
		createProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","recipe_id":"%s","sequence_num":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testName, testType, recipeUUID, sequenceNumber)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestCreateProcessFailure() {
	testUUID := uuid.New()
	suite.dbMock.On("CreateProcess", mock.Anything, mock.Anything).Return(db.Process{}, fmt.Errorf("Error creating process"))

	body := fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","recipe_id":"%s","sequence_num":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testName, testType, recipeUUID, sequenceNumber)

	recorder := makeHTTPCall(http.MethodPost,
		"/processes",
		"/processes",
		body,
		createProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Errorf("Error creating process")

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestListProcessSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ListProcesses", mock.Anything, mock.Anything).Return(
		[]db.Process{
			db.Process{
				ID:             testUUID,
				Name:           testName,
				Type:           testType,
				RecipeID:       recipeUUID,
				SequenceNumber: sequenceNumber,
			},
		}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		listProcessesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"%s","type":"%s","recipe_id":"%s","sequence_num":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`, testUUID, testName, testType, recipeUUID, sequenceNumber)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestListProcessFailure() {
	suite.dbMock.On("ListProcesses", mock.Anything, mock.Anything).Return(
		[]db.Process{}, fmt.Errorf("Error fetching process"))

	recorder := makeHTTPCall(
		http.MethodGet,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		listProcessesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ""
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestShowProcessSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ShowProcess", mock.Anything, mock.Anything).Return(db.Process{
		ID:             testUUID,
		Name:           testName,
		Type:           testType,
		RecipeID:       recipeUUID,
		SequenceNumber: sequenceNumber,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		showProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","recipe_id":"%s","sequence_num":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testName, testType, recipeUUID, sequenceNumber)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestShowProcessFailure() {
	testUUID := uuid.New()
	suite.dbMock.On("ShowProcess", mock.Anything, mock.Anything).Return(db.Process{}, fmt.Errorf("Error showing process"))

	recorder := makeHTTPCall(http.MethodGet,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		showProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ""
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestUpdateProcessSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("UpdateProcess", mock.Anything, mock.Anything).Return(db.Process{
		ID:             testUUID,
		Name:           testName,
		Type:           testType,
		RecipeID:       recipeUUID,
		SequenceNumber: sequenceNumber,
	}, nil)

	body := fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","recipe_id":"%s","sequence_num":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testName, testType, recipeUUID, sequenceNumber)

	recorder := makeHTTPCall(http.MethodPut,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		body,
		updateProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"process updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestUpdateProcessFailure() {
	testUUID := uuid.New()
	suite.dbMock.On("UpdateProcess", mock.Anything, mock.Anything).Return(db.Process{}, fmt.Errorf("Error creating process"))

	body := fmt.Sprintf(`{"id":"%s","name":"%s","type":"%s","recipe_id":"%s","sequence_num":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testName, testType, recipeUUID, sequenceNumber)

	recorder := makeHTTPCall(http.MethodPut,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		body,
		updateProcessHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestDeleteProcessSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteProcess", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		deleteProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"process deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestDeleteProcessFailure() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteProcess", mock.Anything, mock.Anything).Return("", fmt.Errorf("Error deleting process"))

	recorder := makeHTTPCall(http.MethodDelete,
		"/processes/{id}",
		"/processes/"+testUUID.String(),
		"",
		deleteProcessHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
