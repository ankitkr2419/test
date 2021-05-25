package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"
	"testing"

	"github.com/google/uuid"

	"mylab/cpagent/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type RunRecipeHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *RunRecipeHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func TestRunRecipeTestSuite(t *testing.T) {
	suite.Run(t, new(RunRecipeHandlerTestSuite))
}

var invalidDeck = "I"
var runStepWise = false
var recipeID = uuid.New()

// Run Recipe Continuously Test cases
func (suite *RunRecipeHandlerTestSuite) TestRunRecipeSuccess() {

	recorder := makeHTTPCall(http.MethodGet,
		"/run/{id}/{deck:[A-B]}",
		"/run/"+recipeUUID.String()+"/"+deckB,
		"",
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	msg := MsgObj{Msg: responses.RecipeRunInProgress, Deck: deckB}

	output, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RunRecipeHandlerTestSuite) TestRunRecipeUUIDParseFailure() {

	recorder := makeHTTPCall(http.MethodGet,
		"/run/{id}/{deck:[A-B]}",
		"/run/"+invalidUUID+"/"+deckB,
		"",
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	_, err := uuid.Parse(invalidUUID)

	errObj := ErrObj{Err: err.Error(), Deck: deckB}

	output, _ := json.Marshal(errObj)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

//
// NOTE: This test case can be used in login/logout/homing handler
//

/*
func (suite *ProcessHandlerTestSuite) TestRunRecipeInvalidDeckFailure() {

	suite.dbMock.On("runRecipe", mock.Anything, mock.Anything, deckB, runStepWise, recipeID).Return("Success", nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/run/{id}/{deck:[A-B]}",
		"/run/"+ recipeUUID.String()+ "/" + invalidDeck,
		"",
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	errObj := ErrObj{Err: responses.DeckNameInvalid.Error(), Deck: invalidDeck}

	output, _ := json.Marshal(errObj)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
*/

// Step Run Test cases
func (suite *RunRecipeHandlerTestSuite) TestStepRunRecipeSuccess() {

	recorder := makeHTTPCall(http.MethodGet,
		"/step-run/{id}/{deck:[A-B]}",
		"/step-run/"+recipeUUID.String()+"/"+deckB,
		"",
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	msg := MsgObj{Msg: responses.RecipeRunInProgress, Deck: deckB}

	output, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RunRecipeHandlerTestSuite) TestStepRunRecipeUUIDParseFailure() {

	recorder := makeHTTPCall(http.MethodGet,
		"/step-run/{id}/{deck:[A-B]}",
		"/step-run/"+invalidUUID+"/"+deckB,
		"",
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	_, err := uuid.Parse(invalidUUID)

	errObj := ErrObj{Err: err.Error(), Deck: deckB}

	output, _ := json.Marshal(errObj)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RunRecipeHandlerTestSuite) TestRunNextStepSuccess() {

	runNext[deckB] = false

	recorder := makeHTTPCall(http.MethodGet,
		"/run-next-step/{deck:[A-B]}",
		"/run-next-step/"+deckB,
		"",
		runNextStepHandler(Dependencies{Store: suite.dbMock}),
	)

	go func() {
		// drain nextStep channel
		<-nextStep[deckB]
	}()

	msg := MsgObj{Msg: responses.NextStepRunInProgress, Deck: deckB}

	output, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertNotCalled(suite.T(), "populateNextStepChan", mock.Anything)
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RunRecipeHandlerTestSuite) TestRunNextStepFailure() {

	runNext[deckB] = true

	recorder := makeHTTPCall(http.MethodGet,
		"/run-next-step/{deck:[A-B]}",
		"/run-next-step/"+deckB,
		"",
		runNextStepHandler(Dependencies{Store: suite.dbMock}),
	)

	errObj := ErrObj{Err: responses.StepRunNotInProgressError.Error(), Deck: deckB}

	output, _ := json.Marshal(errObj)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
