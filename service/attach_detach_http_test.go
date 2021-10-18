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
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func TestAttachDetachTestSuite(t *testing.T) {
	suite.Run(t, new(AttachDetachHandlerTestSuite))
}

var testAttachDetachRecord = db.AttachDetach{
	ID:        testUUID,
	Operation: "attach",
	Height:    60,
	ProcessID: testProcessUUID,
}

func (suite *AttachDetachHandlerTestSuite) TestCreateAttachDetachHandler() {
	t := suite.T()
	t.Run("when create attach detach record is successful", func(t *testing.T) {
		suite.dbMock.On("CreateAttachDetach", mock.Anything, mock.Anything, recipeUUID).Return(testAttachDetachRecord, nil).Once()
		body, _ := json.Marshal(testAttachDetachRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{recipe_id}",
			"/attach-detach/"+recipeUUID.String(),
			string(body),
			createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
		assert.Equal(suite.T(), string(body), recorder.Body.String())

	})

	t.Run("when create attach detach record failed", func(t *testing.T) {
		suite.dbMock.On("CreateAttachDetach", mock.Anything, mock.Anything, recipeUUID).Return(db.AttachDetach{}, responses.AttachDetachCreateError).Once()

		body, _ := json.Marshal(testAttachDetachRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{recipe_id}",
			"/attach-detach/"+recipeUUID.String(),
			string(body),
			createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.AttachDetachCreateError.Error()})

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})
	t.Run("when create attach detach recieves object with missing attributes", func(t *testing.T) {
		testAttachDetachMissingRecord := db.AttachDetach{
			ID:        testUUID,
			ProcessID: testProcessUUID,
		}
		body, _ := json.Marshal(testAttachDetachMissingRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{recipe_id}",
			"/attach-detach/"+recipeUUID.String(),
			string(body),
			createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(map[string]interface{}{"AttachDetach.Operation": "required", "error": "invalid value for field Operation"})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	t.Run("when create attach detach recieves invalid recipe id", func(t *testing.T) {

		body, _ := json.Marshal(testAttachDetachRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{recipe_id}",
			"/attach-detach/"+invalidUUID,
			string(body),
			createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.RecipeIDInvalidError.Error()})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	t.Run("when create attach detach recieves invalid object for decode", func(t *testing.T) {
		body := "{\"id\":\"38afc9ef-250e-477e-90cf-ef6448c0eb90\",\"category\":\"well_to_well\",\"cartridge_type\":\"cartridge_1\",\"source_position\":?,\"aspire_height\":2,\"aspire_mixing_volume\":3,\"aspire_no_of_cycles\":4,\"aspire_volume\":5,\"aspire_air_volume\":6,\"dispense_height\":1,\"dispense_mixing_volume\":8,\"dispense_no_of_cycles\":9,\"destination_position\":8,\"process_id\":\"46e8b814-3eb3-4cd4-9fd3-3b19a08ea456\",\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"}"

		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{recipe_id}",
			"/attach-detach/"+recipeUUID.String(),
			string(body),
			createAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.AttachDetachDecodeError.Error()})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestShowAttachDetachHandler() {
	t := suite.T()

	t.Run("when show aspire dispense record is successful", func(t *testing.T) {
		suite.dbMock.On("ShowAttachDetach", mock.Anything, recipeUUID).Return(testAttachDetachRecord, nil).Once()

		output, _ := json.Marshal(testAttachDetachRecord)
		recorder := makeHTTPCall(http.MethodGet,
			"/attach-detach/{id}",
			"/attach-detach/"+recipeUUID.String(),
			"",
			showAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})
	t.Run("when show aspire dispense record recieves invalid uuid", func(t *testing.T) {

		output, _ := json.Marshal(ErrObj{Err: responses.UUIDParseError.Error()})
		recorder := makeHTTPCall(http.MethodGet,
			"/attach-detach/{id}",
			"/attach-detach/"+invalidUUID,
			"",
			showAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})
	t.Run("when show aspire dispense record fails", func(t *testing.T) {
		suite.dbMock.On("ShowAttachDetach", mock.Anything, recipeUUID).Return(db.AttachDetach{}, responses.AttachDetachFetchError)

		output, _ := json.Marshal(ErrObj{Err: responses.AttachDetachFetchError.Error()})
		recorder := makeHTTPCall(http.MethodGet,
			"/attach-detach/{id}",
			"/attach-detach/"+recipeUUID.String(),
			"",
			showAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AttachDetachHandlerTestSuite) TestUpdateAttachDetachHandler() {
	t := suite.T()
	t.Run("when update attach detach record is successful", func(t *testing.T) {
		suite.dbMock.On("UpdateAttachDetach", mock.Anything, testAttachDetachRecord).Return(nil).Once()
		body, _ := json.Marshal(testAttachDetachRecord)
		recorder := makeHTTPCall(http.MethodPut,
			"/attach-detach/{id}",
			"/attach-detach/"+testProcessUUID.String(),
			string(body),
			updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(MsgObj{Msg: responses.AttachDetachUpdateSuccess})
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	t.Run("when update attach detach record failed", func(t *testing.T) {
		suite.dbMock.On("UpdateAttachDetach", mock.Anything, testAttachDetachRecord).Return(responses.AttachDetachCreateError).Once()

		body, _ := json.Marshal(testAttachDetachRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{id}",
			"/attach-detach/"+testProcessUUID.String(),
			string(body),
			updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.AttachDetachUpdateError.Error()})

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	t.Run("when update attach detach recieves object with missing attributes", func(t *testing.T) {
		testAttachDetachMissingRecord := db.AttachDetach{
			ID:        testUUID,
			ProcessID: testProcessUUID,
		}
		body, _ := json.Marshal(testAttachDetachMissingRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{id}",
			"/attach-detach/"+recipeUUID.String(),
			string(body),
			updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(map[string]interface{}{"AttachDetach.Operation": "required", "error": "invalid value for field Operation"})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	t.Run("when update attach detach recieves invalid recipe id", func(t *testing.T) {

		body, _ := json.Marshal(testAttachDetachRecord)
		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{recipe_id}",
			"/attach-detach/"+invalidUUID,
			string(body),
			updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.UUIDParseError.Error()})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	t.Run("when update attach detach recieves invalid object for decode", func(t *testing.T) {
		body := "{\"id\":\"38afc9ef-250e-477e-90cf-ef6448c0eb90\",\"category\":\"well_to_well\",\"cartridge_type\":\"cartridge_1\",\"source_position\":?,\"aspire_height\":2,\"aspire_mixing_volume\":3,\"aspire_no_of_cycles\":4,\"aspire_volume\":5,\"aspire_air_volume\":6,\"dispense_height\":1,\"dispense_mixing_volume\":8,\"dispense_no_of_cycles\":9,\"destination_position\":8,\"process_id\":\"46e8b814-3eb3-4cd4-9fd3-3b19a08ea456\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\"}"

		recorder := makeHTTPCall(http.MethodPost,
			"/attach-detach/{id}",
			"/attach-detach/"+recipeUUID.String(),
			string(body),
			updateAttachDetachHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.AttachDetachDecodeError.Error()})

		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())

	})

	suite.dbMock.AssertExpectations(suite.T())
}
