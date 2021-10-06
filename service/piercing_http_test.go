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
type PiercingHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *PiercingHandlerTestSuite) SetupTest() {
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

func TestPiercingTestSuite(t *testing.T) {
	suite.Run(t, new(PiercingHandlerTestSuite))
}

var testPiercingRecord = db.Piercing{
	ID:             testUUID,
	Type:           db.Cartridge1,
	CartridgeWells: []int64{4, 8},
	ProcessID:      testProcessUUID,
	Heights:        []int64{1, 3},
}

func (suite *PiercingHandlerTestSuite) TestCreatePiercingSuccess() {

	suite.dbMock.On("CreatePiercing", mock.Anything, mock.Anything, recipeUUID).Return(testPiercingRecord, nil)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testPiercingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/piercing/{recipe_id}",
		"/piercing/"+recipeUUID.String(),
		string(body),
		createPiercingHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestCreatePiercingFailure() {

	suite.dbMock.On("CreatePiercing", mock.Anything, mock.Anything, recipeUUID).Return(db.Piercing{}, responses.PiercingCreateError)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testPiercingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/piercing/{recipe_id}",
		"/piercing/"+recipeUUID.String(),
		string(body),
		createPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.PiercingCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestShowPiercingSuccess() {

	suite.dbMock.On("ShowPiercing", mock.Anything, testProcessUUID).Return(testPiercingRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		"",
		showPiercingHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testPiercingRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestShowPiercingFailure() {

	suite.dbMock.On("ShowPiercing", mock.Anything, testProcessUUID).Return(db.Piercing{}, responses.PiercingFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		"",
		showPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.PiercingFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestUpdatePiercingSuccess() {

	suite.dbMock.On("UpdatePiercing", mock.Anything, testPiercingRecord).Return(nil)
	suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessRecord, nil)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testPiercingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		string(body),
		updatePiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.PiercingUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestUpdatePiercingFailure() {

	suite.dbMock.On("UpdatePiercing", mock.Anything, testPiercingRecord).Return(responses.PiercingUpdateError)
	suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessRecord, nil)
	suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testPiercingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		string(body),
		updatePiercingHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.PiercingUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestCreatePiercingInvalidUUID() {

	body, _ := json.Marshal(testPiercingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/piercing/{recipe_id}",
		"/piercing/"+invalidUUID,
		string(body),
		createPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.RecipeIDInvalidError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestShowPiercingInvalidUUID() {

	recorder := makeHTTPCall(http.MethodGet,
		"/piercing/{recipe_id}",
		"/piercing/"+invalidUUID,
		"",
		showPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestUpdatePiercingInvalidUUID() {

	body, _ := json.Marshal(testPiercingRecord)
	recorder := makeHTTPCall(http.MethodPut,
		"/piercing/{recipe_id}",
		"/piercing/"+invalidUUID,
		string(body),
		updatePiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.UUIDParseError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
