package service

import (
	"errors"
	"fmt"
	"mylab/cpagent/db"
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
	suite.dbMock = &db.DBMockStore{}
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) TestValidateUsersSuccess() {

	username := "test"
	password := "test"

	suite.dbMock.On("ValidateUser", mock.Anything, mock.Anything).Return(
		nil,
		nil,
	)

	body := fmt.Sprintf(`{"username": "%s","password":"%s"}`, username, password)

	recorder := makeHTTPCall(
		http.MethodPost,
		"/users/{username}/validate",
		"/users/"+username+"/validate",
		body,
		validateUserHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"msg":"Validated User Sucessfully"}`)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestValidateUsersFail() {

	username := "test"
	password := "test"

	suite.dbMock.On("ValidateUser", mock.Anything, mock.Anything).Return(
		nil,
		errors.New("Record Not Found"),
	)

	body := fmt.Sprintf(`{"username": "%s","password":"%s"}`, username, password)

	recorder := makeHTTPCall(
		http.MethodPost,
		"/users/{username}/validate",
		"/users/"+username+"/validate",
		body,
		validateUserHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"msg":"Invalid User"}`)
	assert.Equal(suite.T(), http.StatusExpectationFailed, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
