package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type CleanupHandlerTestSuite struct {
	suite.Suite
	dbMock  *db.DBMockStore
	plcDeck map[string]plc.Extraction
}

var testDeps Dependencies

func (suite *CleanupHandlerTestSuite) SetupTest() {
	var wsMsg = make(chan string)
	var wsErr = make(chan error)
	suite.dbMock = &db.DBMockStore{}
	driverA := &plc.PLCMockStore{}
	driverB := &plc.PLCMockStore{}
	suite.plcDeck = map[string]plc.Extraction{
		plc.DeckA: driverA,
		plc.DeckB: driverB,
	}

	testDeps = Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck, WsErrCh: wsErr, WsMsgCh: wsMsg}

	initiateWebSocket()
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func TestCleanupTestSuite(t *testing.T) {
	suite.Run(t, new(CleanupHandlerTestSuite))
}

func (suite *CleanupHandlerTestSuite) TestDiscardBoxCleanupSuccess() {

	deck := plc.DeckB
	//TODO: On Mock return success
	suite.plcDeck[deck].(*plc.PLCMockStore).On("DiscardBoxCleanup").Return("Discard Box Cleanup Success", nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/discard-box/cleanup/{deck:[A-B]}",
		"/discard-box/cleanup/"+deck,
		"",
		discardBoxCleanupHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
	)

	msg := MsgObj{Msg: responses.DiscardBoxMovedSuccess, Deck: deck}
	body, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestDiscardBoxCleanupFailure() {

	deck := plc.DeckB
	//TODO: On Mock return error
	suite.plcDeck[deck].(*plc.PLCMockStore).On("DiscardBoxCleanup").Return("Discard Box Cleanup Failure", responses.DiscardBoxMoveError)

	recorder := makeHTTPCall(http.MethodGet,
		"/discard-box/cleanup/{deck:[A-B]}",
		"/discard-box/cleanup/"+deck,
		"",
		discardBoxCleanupHandler(testDeps),
	)

	err := ErrObj{Err: responses.DiscardBoxMoveError.Error(), Deck: deck}
	body, _ := json.Marshal(err)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestRestoreDeckSuccess() {

	deck := plc.DeckB
	//TODO: On Mock return success
	suite.plcDeck[deck].(*plc.PLCMockStore).On("RestoreDeck").Return("Restore Deck Success", nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/restore-deck/{deck:[A-B]}",
		"/restore-deck/"+deck,
		"",
		restoreDeckHandler(testDeps),
	)

	msg := MsgObj{Msg: responses.RestoreDeckSuccess, Deck: deck}
	body, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestRestoreDeckFailure() {

	deck := plc.DeckB
	//TODO: On Mock return error

	suite.plcDeck[deck].(*plc.PLCMockStore).On("RestoreDeck").Return("", responses.RestoreDeckError)

	recorder := makeHTTPCall(http.MethodGet,
		"/restore-deck/{deck:[A-B]}",
		"/restore-deck/"+deck,
		"",
		restoreDeckHandler(testDeps),
	)

	err := ErrObj{Err: responses.RestoreDeckError.Error(), Deck: deck}
	body, _ := json.Marshal(err)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CleanupHandlerTestSuite) TestUVCleanupSuccess() {

	deck := plc.DeckB
	time := "01:20:00"
	//TODO: On Mock return success

	suite.plcDeck[deck].(*plc.PLCMockStore).On("UVLight", time).Return("UV Success", nil).Maybe()

	recorder := makeHTTPCall(http.MethodGet,
		"/uv/{time}/{deck:[A-B]}",
		"/uv/"+time+"/"+deck,
		"",
		uvLightHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
	)

	msg := MsgObj{Msg: responses.UVCleanupProgress, Deck: deck}
	body, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func testWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		select {
		case msg := <-testDeps.WsMsgCh:
			logger.Infoln(msg)
		case err := <-testDeps.WsErrCh:
			logger.Errorln(err)
		}
	}
}

func initiateWebSocket() {

	s := httptest.NewServer(http.HandlerFunc(testWebSocket))
	defer s.Close()
	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	ws, _, _ := websocket.DefaultDialer.Dial(u, nil)

	defer ws.Close()
}
