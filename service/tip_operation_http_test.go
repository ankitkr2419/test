package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TipOperationHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *TipOperationHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
	suite.dbMock.On("ListTipsTubes", mock.Anything).Return([]db.TipsTubes{plc.TestTTObj}, nil)
	suite.dbMock.On("ListCartridges", mock.Anything).Return(plc.TestCartridgeObj, nil)
	suite.dbMock.On("ListCartridgeWells").Return(plc.TestCartridgeWellsObj, nil)
	suite.dbMock.On("ListMotors", mock.Anything).Return(plc.TestMotorObj, nil)
	suite.dbMock.On("ListConsDistances").Return(plc.TestConsDistanceObj, nil)

	plc.LoadAllPLCFuncsExceptUtils(suite.dbMock)

}

func TestTipOperationTestSuite(t *testing.T) {
	suite.Run(t, new(TipOperationHandlerTestSuite))
}

var testTipOperationRecord = db.TipOperation{
	ID:        testUUID,
	Type:      "pickup",
	Position:  3,
	ProcessID: testProcessUUID,
}

func (suite *TipOperationHandlerTestSuite) TestCreateTipOperationSuccess() {

	suite.dbMock.On("CreateTipOperation", mock.Anything, mock.Anything, recipeUUID).Return(testTipOperationRecord, nil)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testTipOperationRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/tip-operation/{recipe_id}",
		"/tip-operation/"+recipeUUID.String(),
		string(body),
		createTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestCreateTipOperationFailure() {

	suite.dbMock.On("CreateTipOperation", mock.Anything, mock.Anything, recipeUUID).Return(testTipOperationRecord, responses.TipOperationCreateError)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testTipOperationRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/tip-operation/{recipe_id}",
		"/tip-operation/"+recipeUUID.String(),
		string(body),
		createTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.TipOperationCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestShowTipOperationSuccess() {

	suite.dbMock.On("ShowTipOperation", mock.Anything, testProcessUUID).Return(testTipOperationRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/tip-operation/{id}",
		"/tip-operation/"+testProcessUUID.String(),
		"",
		showTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testTipOperationRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestShowTipOperationFailure() {

	suite.dbMock.On("ShowTipOperation", mock.Anything, testProcessUUID).Return(testTipOperationRecord, responses.TipOperationFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/tip-operation/{id}",
		"/tip-operation/"+testProcessUUID.String(),
		"",
		showTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.TipOperationFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestUpdateTipOperationSuccess() {

	suite.dbMock.On("UpdateTipOperation", mock.Anything, testTipOperationRecord).Return(nil)
	suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessRecord, nil)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testTipOperationRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/tip-operation/{id}",
		"/tip-operation/"+testProcessUUID.String(),
		string(body),
		updateTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.TipOperationUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestUpdateTipOperationFailure() {

	suite.dbMock.On("UpdateTipOperation", mock.Anything, testTipOperationRecord).Return(responses.TipOperationUpdateError)
	suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessRecord, nil)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)
	body, _ := json.Marshal(testTipOperationRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/tip-operation/{id}",
		"/tip-operation/"+testProcessUUID.String(),
		string(body),
		updateTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.TipOperationUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestCreateTipOperationInvalidUUID() {

	body, _ := json.Marshal(testTipOperationRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/tip-operation/{recipe_id}",
		"/tip-operation/"+invalidUUID,
		string(body),
		createTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.RecipeIDInvalidError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestShowTipOperationInvalidUUID() {

	recorder := makeHTTPCall(http.MethodGet,
		"/tip-operation/{recipe_id}",
		"/tip-operation/"+invalidUUID,
		"",
		showTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipOperationHandlerTestSuite) TestUpdateTipOperationInvalidUUID() {

	body, _ := json.Marshal(testTipOperationRecord)
	recorder := makeHTTPCall(http.MethodPut,
		"/tip-operation/{recipe_id}",
		"/tip-operation/"+invalidUUID,
		string(body),
		updateTipOperationHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
