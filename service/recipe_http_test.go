package service

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"net/http"
	"testing"
	"mylab/cpagent/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type RecipeHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *RecipeHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

}

func TestRecipeTestSuite(t *testing.T) {
	suite.Run(t, new(RecipeHandlerTestSuite))
}

var testRecipeRecord = db.Recipe{
	ID:                 recipeUUID,
	Name:               "testRecipeName",
	Description:        "testDescription",
	Position1:          1,
	Position2:          2,
	Position3:          2,
	Position4:          3,
	Position5:          3,
	Cartridge1Position: 1,
	Position7:          2,
	Cartridge2Position: 2,
	Position9:          4,
	ProcessCount:       10,
	IsPublished:        false,
}
var testListRecipeRecord = []db.Recipe{
	testRecipeRecord,
}

func (suite *RecipeHandlerTestSuite) TestCreateRecipeSuccess() {

	suite.dbMock.On("CreateRecipe", mock.Anything, mock.Anything).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testRecipeRecord)

	recorder := makeHTTPCall(http.MethodPost,
		"/recipes",
		"/recipes",
		string(body),
		createRecipeHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestCreateRecipeFailure() {

	suite.dbMock.On("CreateRecipe", mock.Anything, mock.Anything).Return(db.Recipe{}, fmt.Errorf("Error creating recipe"))

	body, _ := json.Marshal(testRecipeRecord)

	recorder := makeHTTPCall(http.MethodPost,
		"/recipes",
		"/recipes",
		string(body),
		createRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Errorf("Error creating recipe")

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestListRecipesSuccess() {

	suite.dbMock.On("ListRecipes", mock.Anything, mock.Anything).Return(testListRecipeRecord, nil)
	body, _ := json.Marshal(testListRecipeRecord)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/recipes",
		"/recipes",
		"",
		listRecipesHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestListRecipesFailure() {
	suite.dbMock.On("ListRecipes", mock.Anything, mock.Anything).Return(
		[]db.Recipe{}, fmt.Errorf("Error fetching recipe"))

	recorder := makeHTTPCall(
		http.MethodGet,
		"/recipes",
		"/recipes",
		"",
		listRecipesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ""
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestShowRecipeSuccess() {

	suite.dbMock.On("ShowRecipe", mock.Anything, mock.Anything).Return(testRecipeRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/recipes/{id}",
		"/recipes/"+recipeUUID.String(),
		"",
		showRecipeHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testRecipeRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestShowRecipeFailure() {

	suite.dbMock.On("ShowRecipe", mock.Anything, mock.Anything).Return(db.Recipe{}, fmt.Errorf("Error showing recipe"))

	recorder := makeHTTPCall(http.MethodGet,
		"/recipes/{id}",
		"/recipes/"+recipeUUID.String(),
		"",
		showRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ""
	assert.Equal(suite.T(), http.StatusNotFound, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestUpdateRecipeSuccess() {

	suite.dbMock.On("UpdateRecipe", mock.Anything, mock.Anything).Return(testRecipeRecord, nil)

	body, _ := json.Marshal(testRecipeRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/recipes/{id}",
		"/recipes/"+recipeUUID.String(),
		string(body),
		updateRecipeHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(MsgObj{Msg: responses.RecipeUpdateSuccess})

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestUpdateRecipeFailure() {

	suite.dbMock.On("UpdateRecipe", mock.Anything, mock.Anything).Return(db.Recipe{}, fmt.Errorf("Error creating recipe"))

	body, _ := json.Marshal(testRecipeRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/recipes/{id}",
		"/recipes/"+recipeUUID.String(),
		string(body),
		updateRecipeHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestDeleteRecipeSuccess() {

	suite.dbMock.On("DeleteRecipe", mock.Anything, mock.Anything).Return(
		recipeUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/recipes/{id}",
		"/recipes/"+recipeUUID.String(),
		"",
		deleteRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"recipe deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestDeleteRecipeFailure() {

	suite.dbMock.On("DeleteRecipe", mock.Anything, mock.Anything).Return("", fmt.Errorf("Error deleting recipe"))

	recorder := makeHTTPCall(http.MethodDelete,
		"/recipes/{id}",
		"/recipes/"+recipeUUID.String(),
		"",
		deleteRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
