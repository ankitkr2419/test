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
type AttachDetachHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *AttachDetachHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestAttachDetachTestSuite(t *testing.T) {
	suite.Run(t, new(AttachDetachHandlerTestSuite))
}

var testAttachDetachRecord = db.AttachDetach{
	ID:            testUUID,
	Operation:     "attach",
	OperationType: "wash",
	ProcessID:     testProcessUUID,
}

func (suite *AttachDetachHandlerTestSuite) TestCreateAttachDetachSuccess() {

	suite.dbMock.On("CreateAttachDetach", mock.Anything, mock.Anything, recipeUUID).Return(testAttachDetachRecord, nil)

	body, _ := json.Marshal(testAttachDetachRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/attach-detach/{recipe_id}",
		"/attach-detach/"+recipeUUID.String(),
		string(body),
		createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestCreateAttachDetachFailure() {

	suite.dbMock.On("CreateAttachDetach", mock.Anything, mock.Anything, recipeUUID).Return(db.AttachDetach{}, responses.AttachDetachCreateError)

	body, _ := json.Marshal(testAttachDetachRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/attach-detach/{recipe_id}",
		"/attach-detach/"+recipeUUID.String(),
		string(body),
		createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.AttachDetachCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestCreateAttachDetachInvalidUUID() {

	body, _ := json.Marshal(testAttachDetachRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/attach-detach/{recipe_id}",
		"/attach-detach/"+invalidUUID,
		string(body),
		createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.RecipeIDInvalidError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestShowAttachDetachSuccess() {

	suite.dbMock.On("ShowAttachDetach", mock.Anything, testProcessUUID).Return(testAttachDetachRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/attach-detach/{id}",
		"/attach-detach/"+testProcessUUID.String(),
		"",
		showAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testAttachDetachRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestShowAttachDetachFailure() {

	suite.dbMock.On("ShowAttachDetach", mock.Anything, testProcessUUID).Return(db.AttachDetach{}, responses.AttachDetachFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/attach-detach/{id}",
		"/attach-detach/"+testProcessUUID.String(),
		"",
		showAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.AttachDetachFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestShowAttachDetachInvalidUUID() {

	recorder := makeHTTPCall(http.MethodGet,
		"/attach-detach/{recipe_id}",
		"/attach-detach/"+invalidUUID,
		"",
		showAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestUpdateAttachDetachSuccess() {

	suite.dbMock.On("UpdateAttachDetach", mock.Anything, testAttachDetachRecord).Return(nil)

	body, _ := json.Marshal(testAttachDetachRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/attach-detach/{id}",
		"/attach-detach/"+testProcessUUID.String(),
		string(body),
		updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.AttachDetachUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestUpdateAttachDetachFailure() {

	suite.dbMock.On("UpdateAttachDetach", mock.Anything, testAttachDetachRecord).Return(responses.AttachDetachUpdateError)

	body, _ := json.Marshal(testAttachDetachRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/attach-detach/{id}",
		"/attach-detach/"+testProcessUUID.String(),
		string(body),
		updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.AttachDetachUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestUpdateAttachDetachInvalidUUID() {

	body, _ := json.Marshal(testAttachDetachRecord)
	recorder := makeHTTPCall(http.MethodPut,
		"/attach-detach/{recipe_id}",
		"/attach-detach/"+invalidUUID,
		string(body),
		updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
