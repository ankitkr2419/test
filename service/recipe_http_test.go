package service

import (
	"encoding/json"
	"fmt"
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

var(
	pos1 int64 = 1
	pos2 int64 = 2
	pos3 int64 = 3
	pos4 int64 = 4
	car1 int64 = 1
	car2 int64 = 2
)

var testRecipeRecord = db.Recipe{
	ID:                 recipeUUID,
	Name:               "testRecipeName",
	Description:        "testDescription",
	Position1:          &pos1,
	Position2:          &pos1,
	Position3:          &pos1,
	Position4:          &pos2,
	Position5:          &pos2,
	Cartridge1Position: &car1,
	Position7:          &pos3,
	Cartridge2Position: &car2,
	Position9:          &pos4,
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
	output := `{"err":"error fetching Recipe list"}`
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
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), `{"err":"error fetching Recipe record"}`, recorder.Body.String())

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
	assert.Equal(suite.T(), `{"err":"error updating Recipe record"}`, recorder.Body.String())

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
	assert.Equal(suite.T(), `{"msg":"Recipe record deleted successfully"}`, recorder.Body.String())

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
	assert.Equal(suite.T(), `{"err":"error deleting Recipe record"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
