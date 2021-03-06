package service

import (
	"encoding/json"
	"errors"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	logger "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.

type UserHandlerTestSuite struct {
	suite.Suite
	dbMock  *db.DBMockStore
	plcDeck map[string]plc.Extraction
}

func (suite *UserHandlerTestSuite) SetupTest() {
	config.SetSecretKey("SECRET_KEY")
	loadUtils()
	suite.dbMock = &db.DBMockStore{}
	driverA := &plc.PLCMockStore{}
	driverB := &plc.PLCMockStore{}
	suite.plcDeck = map[string]plc.Extraction{
		plc.DeckA: driverA,
		plc.DeckB: driverB,
	}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func (suite *UserHandlerTestSuite) TearDownTest() {
	t := suite.T()
	suite.dbMock.AssertExpectations(t)
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

func (suite *UserHandlerTestSuite) TestValidateUserHandler() {
	t := suite.T()
	t.Run("when deck user login is successful", func(t *testing.T) {
		suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(testUserObj, nil).Once()
		suite.dbMock.On("InsertUserAuths", mock.Anything, testUserObj.Username).Return(testUUID, nil).Once()

		body, _ := json.Marshal(testUser)

		recorder := makeHTTPCall(
			http.MethodPost,
			"/login/A",
			"/login/A",
			string(body),
			validateUserHandler(Dependencies{Store: suite.dbMock}),
		)

		// TODO: Unmarshal rcorder.Body into map[string]string and compare for msg and Role
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	})

	t.Run("when non-deck user login is successful", func(t *testing.T) {
		suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(testUserObj, nil).Once()
		suite.dbMock.On("InsertUserAuths", mock.Anything, testUserObj.Username).Return(testUUID, nil).Once()

		body, _ := json.Marshal(testUser)

		recorder := makeHTTPCall(
			http.MethodPost,
			"/login",
			"/login",
			string(body),
			validateUserHandler(Dependencies{Store: suite.dbMock}),
		)

		// TODO: Validate with JSON Token
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	})

	t.Run("when user not found in records", func(t *testing.T) {
		suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(testUserObj, errors.New("Record Not Found")).Once()

		body, _ := json.Marshal(testUser)

		recorder := makeHTTPCall(
			http.MethodPost,
			"/login",
			"/login",
			string(body),
			validateUserHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.UserInvalidError.Error()})
		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when user authentication info couldn't be inserted into db", func(t *testing.T) {
		suite.dbMock.On("ValidateUser", mock.Anything, testUserObj).Return(testUserObj, nil).Once()
		suite.dbMock.On("InsertUserAuths", mock.Anything, testUserObj.Username).Return(testUUID, errors.New("failed to insert user auth")).Once()

		body, _ := json.Marshal(testUser)
		logger.Infoln(string(body))
		recorder := makeHTTPCall(
			http.MethodPost,
			"/login",
			"/login",
			string(body),
			validateUserHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.UserAuthError.Error()})

		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})
}

func (suite *UserHandlerTestSuite) TestCreateUserHandler() {
	t := suite.T()
	t.Run("when new user created successfully", func(t *testing.T) {
		suite.dbMock.On("InsertUser", mock.Anything, testUserObj).Return(nil).Once()

		body, _ := json.Marshal(testUser)
		recorder := makeHTTPCall(http.MethodPost,
			"/users",
			"/users",
			string(body),
			createUserHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(MsgObj{Msg: responses.UserCreateSuccess})
		assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when new user couldn't be created", func(t *testing.T) {
		suite.dbMock.On("InsertUser", mock.Anything, testUserObj).Return(errors.New("cannot insert new user")).Once()

		body, _ := json.Marshal(testUser)
		recorder := makeHTTPCall(http.MethodPost,
			"/users",
			"/users",
			string(body),
			createUserHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.UserInsertError.Error()})
		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when creating user with invalid role", func(t *testing.T) {
		testUser := testUser
		testUser.Role = ""

		testUserObj := testUserObj
		testUserObj.Role = ""

		suite.dbMock.On("InsertUser", mock.Anything, testUserObj).Return(errors.New("invalid role for new user")).Once()

		body, _ := json.Marshal(testUser)
		recorder := makeHTTPCall(http.MethodPost,
			"/users",
			"/users",
			string(body),
			createUserHandler(Dependencies{Store: suite.dbMock}),
		)

		output, _ := json.Marshal(ErrObj{Err: responses.UserInsertError.Error()})
		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})
}

func (suite *UserHandlerTestSuite) TestLogoutUserHandler() {
	t := suite.T()
	t.Run("when RTPCR application user WITH deck logouts successfully", func(t *testing.T) {
		Application = RTPCR

		suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil).Once()
		suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(nil).Once()

		//first need to login to test logout
		deckUserLogin.Store(plc.DeckA, testUserAuthObj.Username)
		defer func() {
			deckUserLogin.Store(plc.DeckA, blank)
		}()
		token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, plc.DeckA, Application, map[string]string{})
		testTokenA := "Bearer " + token
		body, _ := json.Marshal(testUser)

		recorder := makeHTTPCallWithHeader(
			http.MethodDelete,
			"/logout/{deck:[A-B]?}",
			"/logout/"+plc.DeckA,
			string(body),
			map[string]string{"Authorization": testTokenA},
			logoutUserHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
		)

		output, _ := json.Marshal(MsgObj{Msg: responses.UserLogoutSuccess})
		assert.Equal(suite.T(), string(output), recorder.Body.String())
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	})

	t.Run("when Extraction application user WITH deck logouts successfully", func(t *testing.T) {
		Application = Extraction

		suite.plcDeck[plc.DeckA].(*plc.PLCMockStore).On("SetEngineerOrAdminLogged", false).Return().Once()

		suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil).Once()
		suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(nil).Once()

		//first need to login to test logout
		deckUserLogin.Store(plc.DeckA, testUserAuthObj.Username)
		defer func() {
			deckUserLogin.Store(plc.DeckA, blank)
		}()
		token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, plc.DeckA, Application, map[string]string{})
		testTokenA := "Bearer " + token
		body, _ := json.Marshal(testUser)

		recorder := makeHTTPCallWithHeader(
			http.MethodDelete,
			"/logout/{deck:[A-B]?}",
			"/logout/"+plc.DeckA,
			string(body),
			map[string]string{"Authorization": testTokenA},
			logoutUserHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
		)

		output, _ := json.Marshal(MsgObj{Msg: responses.UserLogoutSuccess})
		assert.Equal(suite.T(), string(output), recorder.Body.String())
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	})

	t.Run("when RTPCR application user WITHOUT deck logouts successfully", func(t *testing.T) {
		Application = RTPCR

		suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil).Once()
		suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(nil).Once()

		token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, blank, Application, map[string]string{})
		testTokenA := "Bearer " + token
		body, _ := json.Marshal(testUser)

		recorder := makeHTTPCallWithHeader(
			http.MethodDelete,
			"/logout/{deck:[A-B]?}",
			"/logout/",
			string(body),
			map[string]string{"Authorization": testTokenA},
			logoutUserHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
		)

		output, _ := json.Marshal(MsgObj{Msg: responses.UserLogoutSuccess})
		assert.Equal(suite.T(), string(output), recorder.Body.String())
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	})

	t.Run("when RTPCR application user WITH deck logout is unsuccessful", func(t *testing.T) {
		suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil).Maybe()
		suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(responses.UserAuthDataDeleteError).Maybe()

		//first need to login to test logout
		deckUserLogin.Store(plc.DeckA, testUserAuthObj.Username)
		defer func() {
			deckUserLogin.Store(plc.DeckA, blank)
		}()

		token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, plc.DeckA, Application, map[string]string{})
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

		output, _ := json.Marshal(ErrObj{Err: responses.UserAuthDataDeleteError.Error()})
		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when RTPCR application user WITHOUT deck logout is unsuccessful", func(t *testing.T) {
		suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil).Maybe()
		suite.dbMock.On("DeleteUserAuth", mock.Anything, testUserAuthObj).Return(responses.UserAuthDataDeleteError).Maybe()

		token, _ := EncodeToken(testUserAuthObj.Username, testUserAuthObj.AuthID, testUserObj.Role, blank, Application, map[string]string{})
		testToken := "Bearer " + token
		body, _ := json.Marshal(testUser)

		recorder := makeHTTPCallWithHeader(
			http.MethodDelete,
			"/logout/{deck:[A-B]?}",
			"/logout/",
			string(body),
			map[string]string{"Authorization": testToken},
			logoutUserHandler(Dependencies{Store: suite.dbMock}),
		)

		output, _ := json.Marshal(ErrObj{Err: responses.UserAuthDataDeleteError.Error()})
		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

}

func (suite *UserHandlerTestSuite) TestUpdateUserHandler() {
	t := suite.T()
	t.Run("when user info is updated successfully", func(t *testing.T) {
		newUserObj := testUserObj
		newUserObj.Username = "new"

		suite.dbMock.On("UpdateUser", mock.Anything, newUserObj, testUserObj.Username).Return(nil).Once()

		// Password without MD5
		newUserObj.Password = testUser.Password

		body, _ := json.Marshal(newUserObj)
		recorder := makeHTTPCall(http.MethodPut,
			"/users/{old_username}",
			"/users/"+testUserObj.Username,
			string(body),
			updateUserHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(MsgObj{Msg: responses.UserUpdateSuccess})
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when invalid data is inserted or user update", func(t *testing.T) {
		newUserObj := ""

		body, _ := json.Marshal(newUserObj)
		recorder := makeHTTPCall(http.MethodPut,
			"/users/{old_username}",
			"/users/a",
			string(body),
			updateUserHandler(Dependencies{Store: suite.dbMock}),
		)
		output, _ := json.Marshal(ErrObj{Err: responses.UserDecodeError.Error()})
		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when user updation failed", func(t *testing.T) {
		newUserObj := testUserObj
		newUserObj.Username = "new"

		suite.dbMock.On("UpdateUser", mock.Anything, newUserObj, testUserObj.Username).Return(responses.UserUpdateError).Once()

		// Password without MD5
		newUserObj.Password = testUser.Password
		body, _ := json.Marshal(newUserObj)
		recorder := makeHTTPCall(http.MethodPut,
			"/users/{old_username}",
			"/users/"+testUserObj.Username,
			string(body),
			updateUserHandler(Dependencies{Store: suite.dbMock}),
		)

		output, _ := json.Marshal(ErrObj{Err: responses.UserUpdateError.Error()})
		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})
}

func (suite *UserHandlerTestSuite) TestDeleteUserHandler() {
	t := suite.T()
	t.Run("when user is deleted successfully", func(t *testing.T) {
		testUserObj := testUserObj

		ctx := map[string]interface{}{
			db.ContextKeyUsername: "main",
		}

		suite.dbMock.On("DeleteUser", mock.Anything, testUserObj.Username).Return(nil).Once()

		recorder := makeHTTPCallWithContext(ctx,
			http.MethodDelete,
			"/users/{username}",
			"/users/"+testUserObj.Username,
			"",
			deleteUserHandler(Dependencies{Store: suite.dbMock}),
		)

		output, _ := json.Marshal(MsgObj{Msg: responses.UserDeleteSuccess})
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	// ASK: @Paramita is the Deck independent?
	t.Run("when deleting the logged in user", func(t *testing.T) {
		testUserObj := testUserObj

		ctx := map[string]interface{}{
			db.ContextKeyUsername: testUserObj.Username,
		}

		recorder := makeHTTPCallWithContext(ctx,
			http.MethodDelete,
			"/users/{username}",
			"/users/"+testUserObj.Username,
			"",
			deleteUserHandler(Dependencies{Store: suite.dbMock}),
		)

		output, _ := json.Marshal(ErrObj{Err: responses.SameUserDeleteError.Error()})
		assert.Equal(suite.T(), http.StatusForbidden, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when non existing user is being deleted", func(t *testing.T) {
		testUserObj := testUserObj

		ctx := map[string]interface{}{
			db.ContextKeyUsername: "main",
		}

		suite.dbMock.On("DeleteUser", mock.Anything, testUserObj.Username).Return(responses.ZeroRowsAffectedError).Once()

		recorder := makeHTTPCallWithContext(ctx,
			http.MethodDelete,
			"/users/{username}",
			"/users/"+testUserObj.Username,
			"",
			deleteUserHandler(Dependencies{Store: suite.dbMock}),
		)

		output, _ := json.Marshal(ErrObj{Err: responses.UserNotFoundError.Error()})
		assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

	t.Run("when user deletion is unsuccessful", func(t *testing.T) {
		testUserObj := testUserObj

		ctx := map[string]interface{}{
			db.ContextKeyUsername: "main",
		}

		suite.dbMock.On("DeleteUser", mock.Anything, testUserObj.Username).Return(responses.UserDeleteError).Once()

		recorder := makeHTTPCallWithContext(ctx,
			http.MethodDelete,
			"/users/{username}",
			"/users/"+testUserObj.Username,
			"",
			deleteUserHandler(Dependencies{Store: suite.dbMock}),
		)

		output, _ := json.Marshal(ErrObj{Err: responses.UserDeleteError.Error()})
		assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
		assert.Equal(suite.T(), string(output), recorder.Body.String())
	})

}
