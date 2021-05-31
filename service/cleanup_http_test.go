package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type CleanupHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *CleanupHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestcleanupTestSuite(t *testing.T) {
	suite.Run(t, new(CleanupHandlerTestSuite))
}

func (suite *CleanupHandlerTestSuite) TestDiscardBoxCleanupSuccess() {

	deck := deckB
	//TODO: On Mock return success

	recorder := makeHTTPCall(http.MethodGet,
		"/discard-box/cleanup/{deck:[A-B]}",
		"/discard-box/cleanup/"+deck,
		"",
		discardBoxCleanupHandler(Dependencies{Store: suite.dbMock}),
	)

	msg := MsgObj{Msg: responses.DiscardBoxMovedSuccess, Deck: deck}
	body, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestDiscardBoxCleanupFailure() {

	deck := deckB
	//TODO: On Mock return error

	recorder := makeHTTPCall(http.MethodGet,
		"/discard-box/cleanup/{deck:[A-B]}",
		"/discard-box/cleanup/"+deck,
		"",
		discardBoxCleanupHandler(Dependencies{Store: suite.dbMock}),
	)

	err := ErrObj{Err: responses.DiscardBoxMoveError.Error(), Deck: deck}
	body, _ := json.Marshal(err)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestRestoreDeckSuccess() {

	deck := deckB
	//TODO: On Mock return success

	recorder := makeHTTPCall(http.MethodGet,
		"/restore-deck/{deck:[A-B]}",
		"/restore-deck/"+deck,
		"",
		restoreDeckHandler(Dependencies{Store: suite.dbMock}),
	)

	msg := MsgObj{Msg: responses.RestoreDeckSuccess, Deck: deck}
	body, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestRestoreDeckFailure() {

	deck := deckB
	//TODO: On Mock return error

	recorder := makeHTTPCall(http.MethodGet,
		"/restore-deck/{deck:[A-B]}",
		"/restore-deck/"+deck,
		"",
		restoreDeckHandler(Dependencies{Store: suite.dbMock}),
	)

	err := ErrObj{Err: responses.RestoreDeckError.Error(), Deck: deck}
	body, _ := json.Marshal(err)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestUVCleanupSuccess() {

	deck := deckB
	time := "01:20:00"
	//TODO: On Mock return success

	recorder := makeHTTPCall(http.MethodGet,
		"/uv/{time}/{deck:[A-B]}",
		"/uv/"+time+"/"+deck,
		"",
		uvLightHandler(Dependencies{Store: suite.dbMock}),
	)

	msg := MsgObj{Msg: responses.UVCleanupProgress, Deck: deck}
	body, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
