package service

import (
	// "fmt"
	"mylab/cpagent/db"
	"net/http"
	"testing"

	// "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mylab/cpagent/responses"
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
}

func TestRunRecipeTestSuite(t *testing.T) {
	suite.Run(t, new(RunRecipeHandlerTestSuite))
}

var deckB = "B"
var invalidDeck ="I"
var runStepWise = false
var invalidUUID = "not-a-uuid"

func (suite *ProcessHandlerTestSuite) TestRunRecipeSuccess() {

	suite.dbMock.On("runRecipe", mock.Anything, mock.Anything, deckB, runStepWise, recipeID).Return("Success", nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/run/{id}/{deck:[A-B]}",
		"/run/"+ recipeUUID+ "/" + deckB,
		nil,
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	msg := MsgObj{Msg: responses.RecipeRunInProgress, Deck: deck}

	output := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestRunRecipeUUIDParseFailure() {

	suite.dbMock.On("runRecipe", mock.Anything, mock.Anything, deckB, runStepWise, recipeID).Return("Success", nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/run/{id}/{deck:[A-B]}",
		"/run/"+ invalidUUID+ "/" + deckB,
		nil,
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	_, err := uuid.Parse(invalidUUID)

	errObj := ErrObj{Err: err.Error(), Deck: deckB}

	output := json.Marshal(errObj)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ProcessHandlerTestSuite) TestRunRecipeInvalidDeckFailure() {

	suite.dbMock.On("runRecipe", mock.Anything, mock.Anything, deckB, runStepWise, recipeID).Return("Success", nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/run/{id}/{deck:[A-B]}",
		"/run/"+ recipeUUID+ "/" + invalidDeck,
		nil,
		runRecipeHandler(Dependencies{Store: suite.dbMock}, false),
	)

	errObj := ErrObj{Err: responses.DeckNameInvalid.Error(), Deck: invalidDeck}

	output := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
