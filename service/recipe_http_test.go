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
type RecipeHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *RecipeHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestRecipeTestSuite(t *testing.T) {
	suite.Run(t, new(RecipeHandlerTestSuite))
}

var testRecipeUUID = uuid.New()
var testRecipeName = "testRecipeName"
var testDescription = "testDescription"
var position1 int64 = 1
var position2 int64 = 2
var position3 int64 = 3
var position4 int64 = 4
var position5 int64 = 5
var cartridge1pos int64 = 1
var position7 int64 = 6
var cartridge2pos int64 = 2
var position9 int64 = 7

func (suite *RecipeHandlerTestSuite) TestCreateRecipeSuccess() {

	suite.dbMock.On("CreateRecipe", mock.Anything, mock.Anything).Return(db.Recipe{
		ID:                 testRecipeUUID,
		Name:               testRecipeName,
		Description:        testDescription,
		Position1:          position1,
		Position2:          position2,
		Position3:          position3,
		Position4:          position4,
		Position5:          position5,
		Cartridge1Position: cartridge1pos,
		Position7:          position7,
		Cartridge2Position: cartridge2pos,
		Position9:          position9,
	}, nil)

	body := fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","pos_1":%d,"pos_2":%d,"pos_3":%d,"pos_4":%d,"pos_5":%d,"pos_cartridge_1":%d,"pos_7":%d,"pos_cartridge_2":%d,"pos_9":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testRecipeUUID, testRecipeName, testDescription, position1, position2, position3, position4, position5, cartridge1pos, position7, cartridge2pos, position9)
	recorder := makeHTTPCall(http.MethodPost,
		"/recipe",
		"/recipe",
		body,
		createRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","pos_1":%d,"pos_2":%d,"pos_3":%d,"pos_4":%d,"pos_5":%d,"pos_cartridge_1":%d,"pos_7":%d,"pos_cartridge_2":%d,"pos_9":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testRecipeUUID, testRecipeName, testDescription, position1, position2, position3, position4, position5, cartridge1pos, position7, cartridge2pos, position9)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestCreateRecipeFailure() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("CreateRecipe", mock.Anything, mock.Anything).Return(db.Recipe{}, fmt.Errorf("Error creating recipe"))

	body := fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","pos_1":%d,"pos_2":%d,"pos_3":%d,"pos_4":%d,"pos_5":%d,"pos_cartridge_1":%d,"pos_7":%d,"pos_cartridge_1":%d,"pos_9":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testRecipeUUID, testRecipeName, testDescription, position1, position2, position3, position4, position5, cartridge1pos, position7, cartridge2pos, position9)

	recorder := makeHTTPCall(http.MethodPost,
		"/recipe",
		"/recipe",
		body,
		createRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Errorf("Error creating recipe")

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestListRecipesSuccess() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("ListRecipes", mock.Anything, mock.Anything).Return(
		[]db.Recipe{
			db.Recipe{
				ID:                 testRecipeUUID,
				Name:               testRecipeName,
				Description:        testDescription,
				Position1:          position1,
				Position2:          position2,
				Position3:          position3,
				Position4:          position4,
				Position5:          position5,
				Cartridge1Position: cartridge1pos,
				Position7:          position7,
				Cartridge2Position: cartridge2pos,
				Position9:          position9,
			},
		}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/recipe",
		"/recipe",
		"",
		listRecipesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"%s","description":"%s","pos_1":%d,"pos_2":%d,"pos_3":%d,"pos_4":%d,"pos_5":%d,"pos_cartridge_1":%d,"pos_7":%d,"pos_cartridge_2":%d,"pos_9":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`, testRecipeUUID, testRecipeName, testDescription, position1, position2, position3, position4, position5, cartridge1pos, position7, cartridge2pos, position9)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestListRecipesFailure() {
	suite.dbMock.On("ListRecipes", mock.Anything, mock.Anything).Return(
		[]db.Recipe{}, fmt.Errorf("Error fetching recipe"))

	recorder := makeHTTPCall(
		http.MethodGet,
		"/recipe",
		"/recipe",
		"",
		listRecipesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ""
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestShowRecipeSuccess() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("ShowRecipe", mock.Anything, mock.Anything).Return(db.Recipe{
		ID:                 testRecipeUUID,
		Name:               testRecipeName,
		Description:        testDescription,
		Position1:          position1,
		Position2:          position2,
		Position3:          position3,
		Position4:          position4,
		Position5:          position5,
		Cartridge1Position: cartridge1pos,
		Position7:          position7,
		Cartridge2Position: cartridge2pos,
		Position9:          position9,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/recipe/{id}",
		"/recipe/"+testRecipeUUID.String(),
		"",
		showRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","pos_1":%d,"pos_2":%d,"pos_3":%d,"pos_4":%d,"pos_5":%d,"pos_cartridge_1":%d,"pos_7":%d,"pos_cartridge_2":%d,"pos_9":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testRecipeUUID, testRecipeName, testDescription, position1, position2, position3, position4, position5, cartridge1pos, position7, cartridge2pos, position9)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestShowRecipeFailure() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("ShowRecipe", mock.Anything, mock.Anything).Return(db.Recipe{}, fmt.Errorf("Error showing recipe"))

	recorder := makeHTTPCall(http.MethodGet,
		"/recipe/{id}",
		"/recipe/"+testRecipeUUID.String(),
		"",
		showRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ""
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestUpdateRecipeSuccess() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("UpdateRecipe", mock.Anything, mock.Anything).Return(db.Recipe{
		ID:                 testRecipeUUID,
		Name:               testRecipeName,
		Description:        testDescription,
		Position1:          position1,
		Position2:          position2,
		Position3:          position3,
		Position4:          position4,
		Position5:          position5,
		Cartridge1Position: cartridge1pos,
		Position7:          position7,
		Cartridge2Position: cartridge2pos,
		Position9:          position9,
	}, nil)

	body := fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","pos_1":%d,"pos_2":%d,"pos_3":%d,"pos_4":%d,"pos_5":%d,"pos_cartridge_1":%d,"pos_7":%d,"pos_cartridge_2":%d,"pos_9":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testRecipeUUID, testRecipeName, testDescription, position1, position2, position3, position4, position5, cartridge1pos, position7, cartridge2pos, position9)

	recorder := makeHTTPCall(http.MethodPut,
		"/recipe/{id}",
		"/recipe/"+testRecipeUUID.String(),
		body,
		updateRecipeHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"recipe updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestUpdateRecipeFailure() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("UpdateRecipe", mock.Anything, mock.Anything).Return(db.Recipe{}, fmt.Errorf("Error creating recipe"))

	body := fmt.Sprintf(`{"id":"%s","name":"%s","description":"%s","pos_1":%d,"pos_2":%d,"pos_3":%d,"pos_4":%d,"pos_5":%d,"pos_cartridge_1":%d,"pos_7":%d,"pos_cartridge_2":%d,"pos_9":%d,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testRecipeUUID, testRecipeName, testDescription, position1, position2, position3, position4, position5, cartridge1pos, position7, cartridge2pos, position9)

	recorder := makeHTTPCall(http.MethodPut,
		"/recipe/{id}",
		"/recipe/"+testRecipeUUID.String(),
		body,
		updateRecipeHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestDeleteRecipeSuccess() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("DeleteRecipe", mock.Anything, mock.Anything).Return(
		testRecipeUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/recipe/{id}",
		"/recipe/"+testRecipeUUID.String(),
		"",
		deleteRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"recipe deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *RecipeHandlerTestSuite) TestDeleteRecipeFailure() {
	testRecipeUUID := uuid.New()
	suite.dbMock.On("DeleteRecipe", mock.Anything, mock.Anything).Return("", fmt.Errorf("Error deleting recipe"))

	recorder := makeHTTPCall(http.MethodDelete,
		"/recipe/{id}",
		"/recipe/"+testRecipeUUID.String(),
		"",
		deleteRecipeHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
