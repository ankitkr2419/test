package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mylab/cpagent/config"
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

type UserHandlerTestSuite struct {
	suite.Suite
	dbMock *db.DBMockStore
}

func (suite *UserHandlerTestSuite) SetupTest() {
	config.SetSecretKey("SECRET_KEY")
	loadUtils()
	suite.dbMock = &db.DBMockStore{}
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

var testUser = db.User{
	Username: "test",
	Password: "test",
	Role:     admin,
}
var testUserObj = db.User{
	Username: "test",
	Password: MD5Hash(testUser.Password),
	Role:     admin,
}
var testUserAuthObj = db.UserAuth{
	Username: "test",
	AuthID:   testUUID,
}

func (suite *UserHandlerTestSuite) TestValidateUsersWithDeckSuccess() {

	suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(nil)
	suite.dbMock.On("InsertUserAuths", mock.Anything, testUserObj.Username).Return(testUUID, nil)

	body, _ := json.Marshal(testUser)

	recorder := makeHTTPCall(
		http.MethodPost,
		"/login/A",
		"/login/A",
		string(body),
		validateUserHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestValidateUsersWithoutDeckSuccess() {

	suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(nil)
	suite.dbMock.On("InsertUserAuths", mock.Anything, testUserObj.Username).Return(testUUID, nil)

	body, _ := json.Marshal(testUser)

	recorder := makeHTTPCall(
		http.MethodPost,
		"/login",
		"/login",
		string(body),
		validateUserHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestValidateUsersWhenUserNotFound() {

	suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(errors.New("Record Not Found"))

	body, _ := json.Marshal(testUser)

	recorder := makeHTTPCall(
		http.MethodPost,
		"/login",
		"/login",
		string(body),
		validateUserHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"err":"error invalid user"}`)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestValidateUsersWhenUserAuthInsertFailed() {

	suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(nil)
	suite.dbMock.On("InsertUserAuths", mock.Anything, testUserObj.Username).Return(testUUID, errors.New("failed to insert user auth"))

	body, _ := json.Marshal(testUser)
	logger.Infoln(string(body))
	recorder := makeHTTPCall(
		http.MethodPost,
		"/login",
		"/login",
		string(body),
		validateUserHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"err":"error in storing user authentication data"}`)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestCreateUserSuccess() {

	suite.dbMock.On("InsertUser", mock.Anything, testUserObj).Return(testUser, nil)

	body, _ := json.Marshal(testUser)
	recorder := makeHTTPCall(http.MethodPost,
		"/user",
		"/user",
		string(body),
		createUserHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"msg":"user successfully created"}`)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestCreateUserFailed() {

	suite.dbMock.On("InsertUser", mock.Anything, testUserObj).Return(testUser, errors.New("cannot insert new user"))

	body, _ := json.Marshal(testUser)
	recorder := makeHTTPCall(http.MethodPost,
		"/user",
		"/user",
		string(body),
		createUserHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"err":"error in inserting user"}`)
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestCreateUserValidationFailed() {

	testUser.Role = ""

	body, _ := json.Marshal(testUser)
	recorder := makeHTTPCall(http.MethodPost,
		"/user",
		"/user",
		string(body),
		createUserHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	suite.dbMock.AssertExpectations(suite.T())
	testUser.Role = admin

}

func (suite *UserHandlerTestSuite) TestLogoutWithDeckSuccess() {
	suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil)
	suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(nil)

	//first need to login to test logout
	userLogin.Store(plc.DeckA, true)
	token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, plc.DeckA, map[string]string{})
	testTokenA := "Bearer " + token
	body, _ := json.Marshal(testUser)

	recorder := makeHTTPCallWithHeader(
		http.MethodDelete,
		"/logout/{deck:[A-B]?}",
		"/logout/"+plc.DeckA,
		string(body),
		map[string]string{"Authorization": testTokenA},
		logoutUserHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestLogoutWithoutDeckSuccess() {
	suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil)
	suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(nil)

	token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, "", map[string]string{})
	testTokenA := "Bearer " + token
	body, _ := json.Marshal(testUser)

	recorder := makeHTTPCallWithHeader(
		http.MethodDelete,
		"/logout/{deck:[A-B]?}",
		"/logout/",
		string(body),
		map[string]string{"Authorization": testTokenA},
		logoutUserHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestLogoutWithDeckFailure() {
	suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, errors.New("failed to fetch user auth record"))

	//first need to login to test logout
	userLogin.Store(plc.DeckA, true)

	token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, plc.DeckA, map[string]string{})
	testTokenA := "Bearer " + token
	body, _ := json.Marshal(testUser)

	recorder := makeHTTPCallWithHeader(
		http.MethodDelete,
		"/logout/{deck:[A-B]?}",
		"/logout/"+plc.DeckA,
		string(body),
		map[string]string{"Authorization": testTokenA},
		logoutUserHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"err":"error in authenticating user"}`)
	assert.Equal(suite.T(), http.StatusForbidden, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestLogoutWithDeckDeleteFailure() {
	suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil)
	suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(responses.UserAuthDataDeleteError)

	//first need to login to test logout
	userLogin.Store(plc.DeckA, true)

	token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, plc.DeckA, map[string]string{})
	testTokenA := "Bearer " + token
	body, _ := json.Marshal(testUser)

	recorder := makeHTTPCallWithHeader(
		http.MethodDelete,
		"/logout/{deck:[A-B]?}",
		"/logout/"+plc.DeckA,
		string(body),
		map[string]string{"Authorization": testTokenA},
		logoutUserHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"err":"%s"}`, responses.UserAuthDataDeleteError.Error())
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
