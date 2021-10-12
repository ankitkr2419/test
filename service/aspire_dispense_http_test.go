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
type AspireDispenseHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *AspireDispenseHandlerTestSuite) SetupTest() {

	suite.dbMock = &db.DBMockStore{}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
	suite.dbMock.On("ListTipsTubes", mock.Anything).Return([]db.TipsTubes{plc.TestTTObj}, nil)
	suite.dbMock.On("ListCartridges", mock.Anything).Return(plc.TestCartridgeObj, nil)
	suite.dbMock.On("ListCartridgeWells").Return(plc.TestCartridgeWellsObj, nil)
	suite.dbMock.On("ListMotors", mock.Anything).Return(plc.TestMotorObj, nil)
	suite.dbMock.On("ListConsDistances").Return(plc.TestConsDistanceObj, nil)
	plc.LoadAllPLCFuncs(suite.dbMock)
}

func TestAspireDispenseTestSuite(t *testing.T) {
	suite.Run(t, new(AspireDispenseHandlerTestSuite))
}

var invalidUUID = "not-a-uuid"

var testAspireDispenseRecord = db.AspireDispense{
	ID:                   testUUID,
	Category:             db.WW,
	CartridgeType:        db.Cartridge1,
	SourcePosition:       4,
	AspireHeight:         2,
	AspireMixingVolume:   3,
	AspireNoOfCycles:     4,
	AspireVolume:         5,
	AspireAirVolume:      6,
	DispenseHeight:       1,
	DispenseMixingVolume: 8,
	DispenseNoOfCycles:   9,
	DestinationPosition:  8,
	ProcessID:            testProcessUUID,
}

var testProcessADRecord = db.Process{
	ID:             testProcessUUID,
	Name:           testName,
	Type:           testType,
	SequenceNumber: sequenceNumber,
	RecipeID:       recipeUUID,
}

func (suite *AspireDispenseHandlerTestSuite) TestCreateAspireDispenseHandler() {
	t := suite.T()
	t.Run("when create aspire dispense record is successful", func(t *testing.T) {
		suite.dbMock.On("CreateAspireDispense", mock.Anything, mock.Anything, recipeUUID).Return(testAspireDispenseRecord, nil).Once()
		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil).Once()

		body, _ := json.Marshal(testAspireDispenseRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/aspire-dispense/{recipe_id}",
			"/aspire-dispense/"+recipeUUID.String(),
			string(body),
			createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
		assert.Equal(suite.T(), string(body), recorder.Body.String())

	})

	t.Run("when create aspire dispense return recipe not found error", func(t *testing.T) {
		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, responses.RecipeFetchError).Once()

		body, _ := json.Marshal(testAspireDispenseRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/aspire-dispense/{recipe_id}",
			"/aspire-dispense/"+recipeUUID.String(),
			string(body),
			createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.RecipeFetchError.Error()})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	t.Run("when create aspire dispense record failed", func(t *testing.T) {
		suite.dbMock.On("CreateAspireDispense", mock.Anything, mock.Anything, recipeUUID).Return(db.AspireDispense{}, responses.AspireDispenseCreateError).Once()
		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil).Once()

		body, _ := json.Marshal(testAspireDispenseRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/aspire-dispense/{recipe_id}",
			"/aspire-dispense/"+recipeUUID.String(),
			string(body),
			createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.AspireDispenseCreateError.Error()})

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})
	t.Run("when create aspire dispense recieves object with missing attributes", func(t *testing.T) {
		testAspireDispenseMissingRecord := db.AspireDispense{
			ID:                   testUUID,
			Category:             db.WW,
			CartridgeType:        db.Cartridge1,
			SourcePosition:       4,
			AspireMixingVolume:   3,
			AspireNoOfCycles:     4,
			AspireVolume:         5,
			AspireAirVolume:      6,
			DispenseHeight:       1,
			DispenseMixingVolume: 8,
			DispenseNoOfCycles:   9,
			DestinationPosition:  8,
			ProcessID:            testProcessUUID,
		}
		body, _ := json.Marshal(testAspireDispenseMissingRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/aspire-dispense/{recipe_id}",
			"/aspire-dispense/"+recipeUUID.String(),
			string(body),
			createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(map[string]interface{}{"AspireDispense.AspireHeight": "required", "error": "invalid value for field AspireHeight"})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
		testAspireDispenseRecord.DestinationPosition = 8

	})

	t.Run("when create aspire dispense recieves invalid object", func(t *testing.T) {

		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil).Once()
		testAspireDispenseRecord.DestinationPosition = 3
		body, _ := json.Marshal(testAspireDispenseRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/aspire-dispense/{recipe_id}",
			"/aspire-dispense/"+recipeUUID.String(),
			string(body),
			createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.InvalidDispenseWell.Error()})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
		testAspireDispenseRecord.DestinationPosition = 8

	})
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestShowAspireDispenseHandler() {
	t := suite.T()

	t.Run("when show aspire dispense record is successful", func(t *testing.T) {
		suite.dbMock.On("ShowAspireDispense", mock.Anything, recipeUUID).Return(testAspireDispenseRecord, nil).Once()

		output, _ := json.Marshal(testAspireDispenseRecord)
		recorder := makeHTTPCall(http.MethodGet,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+recipeUUID.String(),
			"",
			showAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})
	t.Run("when show aspire dispense record recieves invalid uuid", func(t *testing.T) {

		output, _ := json.Marshal(ErrObj{Err: responses.UUIDParseError.Error()})
		recorder := makeHTTPCall(http.MethodGet,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+invalidUUID,
			"",
			showAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})
	t.Run("when show aspire dispense record fails", func(t *testing.T) {
		suite.dbMock.On("ShowAspireDispense", mock.Anything, recipeUUID).Return(db.AspireDispense{}, responses.AspireDispenseFetchError)

		output, _ := json.Marshal(ErrObj{Err: responses.AspireDispenseFetchError.Error()})
		recorder := makeHTTPCall(http.MethodGet,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+recipeUUID.String(),
			"",
			showAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestUpdateAspireDispenseHandler() {
	t := suite.T()
	t.Run("when Update aspire dispense record is successful", func(t *testing.T) {
		suite.dbMock.On("UpdateAspireDispense", mock.Anything, mock.Anything).Return(nil).Once()
		suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessADRecord, nil).Once()
		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil).Once()

		body, _ := json.Marshal(testAspireDispenseRecord)

		recorder := makeHTTPCall(http.MethodPut,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+testProcessUUID.String(),
			string(body),
			updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output := MsgObj{Msg: responses.AspireDispenseUpdateSuccess}
		outputBytes, _ := json.Marshal(output)

		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	})

	t.Run("when Update aspire dispense return recipe not found error", func(t *testing.T) {
		suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessADRecord, nil).Once()
		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, responses.RecipeFetchError).Once()

		body, _ := json.Marshal(testAspireDispenseRecord)

		recorder := makeHTTPCall(http.MethodPut,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+testProcessUUID.String(),
			string(body),
			updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output := ErrObj{Err: responses.RecipeFetchError.Error()}
		outputBytes, _ := json.Marshal(output)

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	})

	t.Run("when Update aspire dispense return process not found error", func(t *testing.T) {
		suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessADRecord, responses.ProcessFetchError).Once()

		body, _ := json.Marshal(testAspireDispenseRecord)

		recorder := makeHTTPCall(http.MethodPut,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+testProcessUUID.String(),
			string(body),
			updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output := ErrObj{Err: responses.ProcessFetchError.Error()}
		outputBytes, _ := json.Marshal(output)

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	})

	t.Run("when Update aspire dispense record failes", func(t *testing.T) {
		suite.dbMock.On("UpdateAspireDispense", mock.Anything, mock.Anything).Return(responses.AspireDispenseUpdateError).Once()
		suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessADRecord, nil).Once()
		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil).Once()

		body, _ := json.Marshal(testAspireDispenseRecord)

		recorder := makeHTTPCall(http.MethodPut,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+testProcessUUID.String(),
			string(body),
			updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output := ErrObj{Err: responses.AspireDispenseUpdateError.Error()}
		outputBytes, _ := json.Marshal(output)

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	})

	t.Run("when Update aspire dispense record recieves invalid uuid", func(t *testing.T) {

		recorder := makeHTTPCall(http.MethodPut,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+invalidUUID,
			"",
			updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output := ErrObj{Err: responses.UUIDParseError.Error()}
		outputBytes, _ := json.Marshal(output)

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	})
	t.Run("when Update aspire dispense recieves invalid request object", func(t *testing.T) {
		suite.dbMock.On("ShowProcess", mock.Anything, testProcessUUID).Return(testProcessADRecord, nil).Once()
		suite.dbMock.On("ShowRecipe", mock.Anything, recipeUUID).Return(testRecipeRecord, nil).Once()
		testAspireDispenseRecord.DestinationPosition = 3

		body, _ := json.Marshal(testAspireDispenseRecord)

		recorder := makeHTTPCall(http.MethodPut,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+testProcessUUID.String(),
			string(body),
			updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output := ErrObj{Err: responses.InvalidDispenseWell.Error()}
		outputBytes, _ := json.Marshal(output)

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	})
	t.Run("when Update aspire dispense recieves object with missing attributes", func(t *testing.T) {
		testAspireDispenseMissingRecord := db.AspireDispense{
			ID:                   testUUID,
			Category:             db.WW,
			CartridgeType:        db.Cartridge1,
			SourcePosition:       4,
			AspireMixingVolume:   3,
			AspireNoOfCycles:     4,
			AspireVolume:         5,
			AspireAirVolume:      6,
			DispenseHeight:       1,
			DispenseMixingVolume: 8,
			DispenseNoOfCycles:   9,
			DestinationPosition:  8,
			ProcessID:            testProcessUUID,
		}

		body, _ := json.Marshal(testAspireDispenseMissingRecord)

		recorder := makeHTTPCall(http.MethodPut,
			"/aspire-dispense/{id}",
			"/aspire-dispense/"+testProcessUUID.String(),
			string(body),
			updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(map[string]interface{}{"AspireDispense.AspireHeight": "required", "error": "invalid value for field AspireHeight"})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), output, recorder.Body.Bytes())

	})
	suite.dbMock.AssertExpectations(suite.T())
}
