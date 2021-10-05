package plc

import (
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	logger "github.com/sirupsen/logrus"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type DelayTestSuite struct {
	suite.Suite
	dbMock     *db.DBMockStore
	driverMock *MockCompact32Driver
}

var testProcesses = []db.Process{
	db.Process{
		ID:             testUUID,
		Name:           "one",
		Type:           "test 1",
		SequenceNumber: 1,
		RecipeID:       recipeUUID,
	},
	db.Process{
		ID:             testUUID,
		Name:           "two",
		Type:           "test 2",
		SequenceNumber: 2,
		RecipeID:       recipeUUID,
	},
	db.Process{
		ID:             testUUID,
		Name:           "three",
		Type:           "test 3",
		SequenceNumber: 3,
		RecipeID:       recipeUUID,
	},
}

func (suite *DelayTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	suite.dbMock.On("ListTipsTubes", mock.Anything).Return([]db.TipsTubes{TestTTObj}, nil)
	suite.dbMock.On("ListCartridges", mock.Anything).Return(TestCartridgeObj, nil)
	suite.dbMock.On("ListCartridgeWells").Return(TestCartridgeWellsObj, nil)
	suite.dbMock.On("ListMotors", mock.Anything).Return(TestMotorObj, nil)
	suite.dbMock.On("ListConsDistances").Return(TestConsDistanceObj, nil)

	LoadAllPLCFuncs(suite.dbMock)
	suite.driverMock = &MockCompact32Driver{}

}

var testDelayRecord = db.Delay{
	ID:        testUUID,
	DelayTime: 10,
	ProcessID: testProcessUUID,
}

func TestDelayTestSuite(t *testing.T) {
	suite.Run(t, new(DelayTestSuite))
}

var upgrader = websocket.Upgrader{}

func testWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		select {
		case msg := <-testdeck.WsMsgCh:
			logger.Infoln(msg)
		case err := <-testdeck.WsErrCh:
			logger.Errorln(err)
		}
	}
}

func (suite *DelayTestSuite) TestDelaySuccess() {
	initiateWebSocket()
	testdeck.DeckDriver = suite.driverMock

	res, err := testdeck.AddDelay(testDelayRecord, false)

	assert.Equal(suite.T(), "SUCCESS", res)
	assert.Nil(suite.T(), err)
	suite.driverMock.AssertExpectations(suite.T())
}

func (suite *DelayTestSuite) TestDelayMachineAbortedSuccess() {
	initiateWebSocket()
	testdeck.DeckDriver = suite.driverMock
	testDelayRecord.DelayTime = 20
	go func() {
		time.Sleep(200 * time.Millisecond)
		testdeck.setAborted()
	}()
	res, err := testdeck.AddDelay(testDelayRecord, false)

	assert.Equal(suite.T(), "", res)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), responses.AbortedError, err.Error())
	suite.driverMock.AssertExpectations(suite.T())
	testDelayRecord.DelayTime = 5
}

func (suite *DelayTestSuite) TestDelayUVLightInProgressSuccess() {
	initiateWebSocket()
	testdeck.setUVLightInProgress()
	testdeck.DeckDriver = suite.driverMock

	res, err := testdeck.AddDelay(testDelayRecord, false)

	assert.Equal(suite.T(), "SUCCESS", res)
	assert.Nil(suite.T(), err)
	suite.driverMock.AssertExpectations(suite.T())
}

func (suite *DelayTestSuite) TestDelayRecipeProgressSuccess() {
	initiateWebSocket()
	deckProcesses[testdeck.name] = testProcesses
	testDelayRecord.DelayTime = 2
	testdeck.ResetAborted()
	testdeck.ResetPaused()
	testdeck.SetRunInProgress()
	go func() {
		for i := 2; i >= -2; i-- {
			testdeck.SetCurrentProcessNumber(int64(i))
		}
	}()
	testdeck.DeckDriver = suite.driverMock

	res, err := testdeck.AddDelay(testDelayRecord, true)

	assert.Equal(suite.T(), "SUCCESS", res)
	assert.Nil(suite.T(), err)
	suite.driverMock.AssertExpectations(suite.T())
	testDelayRecord.DelayTime = 10
}

func initiateWebSocket() {
	s := httptest.NewServer(http.HandlerFunc(testWebSocket))
	defer s.Close()
	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	ws, _, _ := websocket.DefaultDialer.Dial(u, nil)

	defer ws.Close()
}
