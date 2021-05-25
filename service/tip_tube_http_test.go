package service

import (
	"encoding/json"
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
type TipTubeHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *TipTubeHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestTipTubeTestSuite(t *testing.T) {
	suite.Run(t, new(TipTubeHandlerTestSuite))
}

var testTTObj = db.TipsTubes{
	ID:               1,
	Name:             "testTip",
	Type:             "test",
	AllowedPositions: []int64{1, 2},
	Volume:           2,
	Height:           1.1,
}

func (suite *TipTubeHandlerTestSuite) TestCreateTipTubeSuccess() {
	suite.dbMock.On("InsertTipsTubes", mock.Anything, mock.Anything).Return(nil)

	body, _ := json.Marshal(testTTObj)

	recorder := makeHTTPCall(http.MethodPost,
		"/tiptube",
		"/tiptube",
		string(body),
		createTipTubeHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipTubeHandlerTestSuite) TestCreateTipTubeFailure() {
	suite.dbMock.On("InsertTipsTubes", mock.Anything, mock.Anything).Return(errors.New("failed to insert tip tube"))

	body, _ := json.Marshal(testTTObj)

	recorder := makeHTTPCall(http.MethodPost,
		"/tiptube",
		"/tiptube",
		string(body),
		createTipTubeHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.NotEqual(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipTubeHandlerTestSuite) TestListTipTubeSuccess() {
	suite.dbMock.On("ListTipsTubes", mock.Anything, mock.Anything).Return([]db.TipsTubes{testTTObj}, nil)

	recorder := makeHTTPCall(http.MethodPost,
		"/tiptube/{tiptube:[a-z]*}",
		"/tiptube/tip",
		"",
		listTipsTubesHandler(Dependencies{Store: suite.dbMock}),
	)
	body, _ := json.Marshal([]db.TipsTubes{testTTObj})
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipTubeHandlerTestSuite) TestListTipTubeFailure() {
	suite.dbMock.On("ListTipsTubes", mock.Anything, mock.Anything).Return([]db.TipsTubes{}, errors.New("failed to fetch tips tubes"))

	recorder := makeHTTPCall(http.MethodPost,
		"/tiptube/{tiptube:[a-z]*}",
		"/tiptube/tip",
		"",
		listTipsTubesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"err":"error fetching tip tube record"}`)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipTubeHandlerTestSuite) TestListTipTubeInvalidInput() {

	recorder := makeHTTPCall(http.MethodPost,
		"/tiptube/{tiptube:[a-z]*}",
		"/tiptube/tp",
		"",
		listTipsTubesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"err":"error invalid tip tube arguments"}`)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
