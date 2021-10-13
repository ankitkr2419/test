package service

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"reflect"
	"testing"

	"github.com/dgrijalva/jwt-go"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthenticateTestSuite struct {
	suite.Suite
	dbMock *db.DBMockStore
}

func TestAuthenticateTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticateTestSuite))
}

func (suite *AuthenticateTestSuite) SetupTest() {
	config.SetSecretKey("SECRET_KEY")
	suite.dbMock = &db.DBMockStore{}
	loadUtils()

}
func (suite *AuthenticateTestSuite) TestEncodeToken() {
	t := suite.T()
	t.Run("when encode token with deck", func(t *testing.T) {
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		tokenType := reflect.TypeOf(token).Kind()
		assert.Equal(suite.T(), tokenType, reflect.String)

	})
	t.Run("when encode token without deck", func(t *testing.T) {
		token, _ := EncodeToken("test", testUUID, "tester", "", Application, map[string]string{})
		tokenType := reflect.TypeOf(token).Kind()
		assert.Equal(suite.T(), tokenType, reflect.String)
	})

}

func (suite *AuthenticateTestSuite) TestDecodeToken() {
	t := suite.T()
	t.Run("when decode token with deck", func(t *testing.T) {
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		flag, _ := decodeToken(token)

		deck, ok := flag["deck"].(string)

		assert.NotEqual(suite.T(), flag, nil)
		assert.Equal(suite.T(), ok, true)
		assert.Equal(suite.T(), deck, plc.DeckA)

	})
	t.Run("when decode token without deck", func(t *testing.T) {
		token, _ := EncodeToken("test", testUUID, "tester", "", Application, map[string]string{})
		flag, _ := decodeToken(token)

		deck, ok := flag["deck"].(string)

		assert.NotEqual(suite.T(), flag, nil)
		assert.Equal(suite.T(), ok, true)
		assert.Equal(suite.T(), deck, "")
	})
	t.Run("when decode token with incorrect access key", func(t *testing.T) {
		token, _ := EncodeToken("test", testUUID, "tester", "", Application, map[string]string{})
		config.SetSecretKey("correct_key")
		flag, err := decodeToken(token)

		deck, ok := flag["deck"].(string)

		assert.Equal(suite.T(), flag, jwt.MapClaims(jwt.MapClaims(nil)))
		assert.NotEqual(suite.T(), err, nil)
		assert.Equal(suite.T(), ok, false)
		assert.Equal(suite.T(), deck, "")
	})
}

func (suite *AuthenticateTestSuite) TestAuthenticate() {
	t := suite.T()
	t.Run("when the user is authenticated successfully", func(t *testing.T) {
		suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil).Once()
		deps := Dependencies{Store: suite.dbMock}
		Application = RTPCR
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate",
			"/test/authenticate",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	})
	t.Run("when the user authentication fails due to user auth not found", func(t *testing.T) {
		suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, responses.UserAuthNotFoundError).Once()
		deps := Dependencies{Store: suite.dbMock}
		Application = RTPCR
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate",
			"/test/authenticate",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application type is incorrect", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		Application = ""
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate",
			"/test/authenticate",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application role is not allowed", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		Application = ""
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate",
			"/test/authenticate",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application, admin),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has cross deck login", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		Application = ""
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/B",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has empty deck login", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		Application = ""
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has empty token", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + ""},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has decode token error", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		token := "invalidtoken"
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has role missing", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzQxMDc1NTMsInN1YiI6InRlc3QiLCJkZWNrIjoiQSIsImF1dGhfaWQiOiJkNGY2OTc0MS05YmNhLTQ2OTQtOTNiMy05MTkxN2I3NTIxYTciLCJhcHBfdHlwZSI6IiJ9.QBybgfXIe_sbzrEieBaL3uk4neTKAwlQhyYzEeltPpM"
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has deck is missing", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzQxMDc1NTMsInN1YiI6InRlc3QiLCJyb2xlIjoidGVzdGVyIiwiYXV0aF9pZCI6ImQ0ZjY5NzQxLTliY2EtNDY5NC05M2IzLTkxOTE3Yjc1MjFhNyIsImFwcF90eXBlIjoiIn0.Bpr7DP7RUuqaCjT5zRYISyLXZCQdtXz-QneBVdGGr_c"
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has incorect auth id", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		Application = RTPCR
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzQxMDc1NTMsInN1YiI6InRlc3QiLCJyb2xlIjoidGVzdGVyIiwiZGVjayI6IkEiLCJhdXRoX2lkIjoiaW52YWxpZC11dWlkIiwiYXBwX3R5cGUiOiJydHBjciJ9.9Gs9ekvhmaAen54PdhscfGjg9FcSBQOTsKCwjoisl90"
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, RTPCR),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has app_type is missing", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzQxMDc1NTMsInN1YiI6InRlc3QiLCJyb2xlIjoidGVzdGVyIiwiZGVjayI6IkEiLCJhdXRoX2lkIjoiaW52YWxpZC11dWlkIn0.1xNBCIPDw9sOYUT_j55zdQw7HED2dGKeUSV0esv-HwE"
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application type is none", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2MzQxMDc1NTMsInN1YiI6InRlc3QiLCJyb2xlIjoidGVzdGVyIiwiZGVjayI6IkEiLCJhdXRoX2lkIjoiaW52YWxpZC11dWlkIiwiYXBwX3R5cGUiOiIifQ.MbekRLmlBYD8_YLSkSTtweOjtjyGMtmooLdEJOnuP3o"
		Application = None
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, None),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	t.Run("when application has cross application login", func(t *testing.T) {
		deps := Dependencies{Store: suite.dbMock}
		Application = RTPCR
		token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Extraction, map[string]string{})
		recorder := makeHTTPCallWithHeader(
			http.MethodPost,
			"/test/authenticate/{deck:[A-B]?}",
			"/test/authenticate/A",
			"",
			map[string]string{"Authorization": "Bearer " + token},
			authenticate(testHandlerFunc(deps), deps, Application),
		)
		assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	})
	suite.dbMock.AssertExpectations(suite.T())

}

func (suite *AuthenticateTestSuite) TestAuthenticateWithRoleSuccess() {
	Application = Combined
	suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil)
	deps := Dependencies{Store: suite.dbMock}
	token, _ := EncodeToken("test", testUUID, "admin", plc.DeckA, Combined, map[string]string{})
	recorder := makeHTTPCallWithHeader(
		http.MethodPost,
		"/test/authenticate",
		"/test/authenticate",
		"",
		map[string]string{"Authorization": "Bearer " + token},
		authenticate(testHandlerFunc(deps), deps, Combined, admin),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

}

func (suite *AuthenticateTestSuite) TestAuthenticateWithDeckSuccess() {

	deps := Dependencies{Store: suite.dbMock}
	token, _ := EncodeToken("test", testUUID, "admin", plc.DeckA, Combined, map[string]string{})
	recorder := makeHTTPCallWithHeader(
		http.MethodPost,
		"/test/authenticate/{deck:[A-B]?}",
		"/test/authenticate/B",
		"",
		map[string]string{"Authorization": "Bearer " + token},
		authenticate(testHandlerFunc(deps), deps, admin),
	)

	output := fmt.Sprintf(`{"err":"error wrong token for deck"}`)
	assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())

}

func (suite *AuthenticateTestSuite) TestAuthenticateWithRoleFailed() {

	deps := Dependencies{Store: suite.dbMock}
	token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, Application, map[string]string{})
	recorder := makeHTTPCallWithHeader(
		http.MethodPost,
		"/test/authenticate",
		"/test/authenticate",
		"",
		map[string]string{"Authorization": "Bearer " + token},
		authenticate(testHandlerFunc(deps), deps, admin),
	)
	output := ErrObj{Err: responses.UserTokenAppNotExistError.Error()}
	outputBytes, _ := json.Marshal(output)
	assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	assert.Equal(suite.T(), string(outputBytes), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())

}

func testHandlerFunc(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		logger.Println("test handler function")
		rw.WriteHeader(http.StatusOK)
	})
}
