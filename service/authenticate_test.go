package service

import (
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"net/http"
	"reflect"
	"testing"

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
func (suite *AuthenticateTestSuite) TestEncodeTokenWithDeck() {

	token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, map[string]string{})
	tokenType := reflect.TypeOf(token).Kind()
	assert.Equal(suite.T(), tokenType, reflect.String)

}

func (suite *AuthenticateTestSuite) TestEncodeTokenWithoutDeck() {
	token, _ := EncodeToken("test", testUUID, "tester", "", map[string]string{})
	tokenType := reflect.TypeOf(token).Kind()
	assert.Equal(suite.T(), tokenType, reflect.String)

}

func (suite *AuthenticateTestSuite) TestDecodeTokenWithDeck() {

	token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, map[string]string{})
	flag, _ := decodeToken(token)

	deck, ok := flag["deck"].(string)

	assert.NotEqual(suite.T(), flag, nil)
	assert.Equal(suite.T(), ok, true)
	assert.Equal(suite.T(), deck, plc.DeckA)

}

func (suite *AuthenticateTestSuite) TestDecodeTokenWithoutDeck() {
	token, _ := EncodeToken("test", testUUID, "tester", "", map[string]string{})
	flag, _ := decodeToken(token)

	deck, ok := flag["deck"].(string)

	assert.NotEqual(suite.T(), flag, nil)
	assert.Equal(suite.T(), ok, true)
	assert.Equal(suite.T(), deck, "")
}

func (suite *AuthenticateTestSuite) TestAuthenticateSuccess() {
	suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil)
	deps := Dependencies{Store: suite.dbMock}
	token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, map[string]string{})
	recorder := makeHTTPCallWithHeader(
		http.MethodPost,
		"/test/authenticate",
		"/test/authenticate",
		"",
		map[string]string{"Authorization": "Bearer " + token},
		authenticate(testHandlerFunc(deps), deps),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	suite.dbMock.AssertExpectations(suite.T())

}

func (suite *AuthenticateTestSuite) TestAuthenticateWithRoleSuccess() {
	suite.dbMock.On("ShowUserAuth", mock.Anything, testUserObj.Username, mock.Anything).Return(testUserAuthObj, nil)
	deps := Dependencies{Store: suite.dbMock}
	token, _ := EncodeToken("test", testUUID, "admin", plc.DeckA, map[string]string{})
	recorder := makeHTTPCallWithHeader(
		http.MethodPost,
		"/test/authenticate",
		"/test/authenticate",
		"",
		map[string]string{"Authorization": "Bearer " + token},
		authenticate(testHandlerFunc(deps), deps, admin),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

}

func (suite *AuthenticateTestSuite) TestAuthenticateWithDeckSuccess() {

	deps := Dependencies{Store: suite.dbMock}
	token, _ := EncodeToken("test", testUUID, "admin", plc.DeckA, map[string]string{})
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
	token, _ := EncodeToken("test", testUUID, "tester", plc.DeckA, map[string]string{})
	recorder := makeHTTPCallWithHeader(
		http.MethodPost,
		"/test/authenticate",
		"/test/authenticate",
		"",
		map[string]string{"Authorization": "Bearer " + token},
		authenticate(testHandlerFunc(deps), deps, admin),
	)
	output := fmt.Sprintf(`{"err":"error invalid role"}`)
	assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())

}

func testHandlerFunc(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		logger.Println("test handler function")
		rw.WriteHeader(http.StatusOK)
	})
}
