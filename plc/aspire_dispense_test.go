package plc

import (
	"mylab/cpagent/db"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AspireDispenseTestSuite struct {
	suite.Suite
	dbMock     *db.DBMockStore
	driverMock *MockCompact32Driver
}

var testTTObj = db.TipsTubes{
	ID:               1,
	Name:             "testTip",
	Type:             "tip",
	AllowedPositions: []int64{1, 2},
	Volume:           2,
	Height:           1.1,
}

var testCartridgeObj = []db.Cartridge{
	db.Cartridge{
		ID:          1,
		Type:        db.Cartridge1,
		Description: "test 1"},
	db.Cartridge{
		ID:          2,
		Type:        db.Cartridge2,
		Description: "test 2"},
}

var testCartridgeWellsObj = []db.CartridgeWells{
	db.CartridgeWells{
		ID:       1,
		WellNum:  4,
		Distance: 3.4,
		Height:   1.1,
		Volume:   2,
	},
	db.CartridgeWells{
		ID:       1,
		WellNum:  8,
		Distance: 3.4,
		Height:   1.1,
		Volume:   2,
	},
}

var testConsDistanceObj = []db.ConsumableDistance{
	db.ConsumableDistance{
		ID:          1001,
		Name:        string(db.Cartridge1) + "_start",
		Distance:    2,
		Description: "testing 1",
	},
	db.ConsumableDistance{
		ID:          1051,
		Name:        string(db.Cartridge2) + "_start",
		Distance:    1,
		Description: "testing 2",
	},
	db.ConsumableDistance{
		ID:          1002,
		Name:        "pickup_tip_up",
		Distance:    1,
		Description: "testing 2",
	},
}

var testMotorObj = []db.Motor{
	db.Motor{
		ID:     1,
		Name:   "test motor 1",
		Deck:   DeckA,
		Number: 10,
		Ramp:   100,
		Steps:  200,
		Slow:   1000,
		Fast:   5500,
	},
	db.Motor{
		ID:     2,
		Name:   "test motor 2",
		Deck:   DeckA,
		Number: 9,
		Ramp:   200,
		Steps:  400,
		Slow:   2000,
		Fast:   7500,
	},
}

func (suite *AspireDispenseTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	suite.dbMock.On("ListTipsTubes", mock.Anything).Return([]db.TipsTubes{testTTObj}, nil)
	suite.dbMock.On("ListCartridges", mock.Anything).Return(testCartridgeObj, nil)
	suite.dbMock.On("ListCartridgeWells").Return(testCartridgeWellsObj, nil)
	suite.dbMock.On("ListMotors").Return(testMotorObj, nil)
	suite.dbMock.On("ListConsDistances").Return(testConsDistanceObj, nil)

	LoadAllPLCFuncs(suite.dbMock)
	suite.driverMock = &MockCompact32Driver{}
}

func TestAspireDispenseTestSuite(t *testing.T) {
	suite.Run(t, new(AspireDispenseTestSuite))
}

var testUUID = uuid.New()
var testProcessUUID = uuid.New()
var recipeUUID = uuid.New()

var testAspireDispenseRecord = db.AspireDispense{
	ID:                   testUUID,
	Category:             db.WW,
	CartridgeType:        db.Cartridge1,
	SourcePosition:       4,
	AspireHeight:         2,
	AspireMixingVolume:   3,
	AspireNoOfCycles:     4,
	AspireVolume:         5,
	AspireAirVolume:      6,
	DispenseHeight:       7,
	DispenseMixingVolume: 8,
	DispenseNoOfCycles:   9,
	DestinationPosition:  8,
	ProcessID:            testProcessUUID,
}
var testdeck = Compact32Deck{
	name:    DeckA,
	WsMsgCh: make(chan string),
	WsErrCh: make(chan error),
}

func (suite *AspireDispenseTestSuite) TestAspireDispenseSuccess() {

	testdeck.DeckDriver = suite.driverMock
	suite.driverMock.On("WriteSingleCoil", mock.Anything, mock.Anything).Return(nil)

	for i := 1; i >= 0; i-- {
		suite.driverMock.On("WriteSingleRegister", mock.Anything, mock.Anything).Return([]uint8{uint8(i)}, nil)
		suite.driverMock.On("ReadCoils", mock.Anything, mock.Anything).Return([]uint8{uint8(i)}, nil)
	}

	res, err := testdeck.AspireDispense(testAspireDispenseRecord, 1)

	assert.Equal(suite.T(), "ASPIRE and DISPENSE was successful", res)
	assert.Nil(suite.T(), err)
	suite.driverMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseTestSuite) TestAspireDispenseTipTubeNotExists() {

	testdeck.DeckDriver = suite.driverMock

	res, err := testdeck.AspireDispense(testAspireDispenseRecord, 1)

	assert.Equal(suite.T(), "", res)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "testTip1 tip doesn't exist for tipstubes", err.Error())

}

func (suite *AspireDispenseTestSuite) TestAspireDispenseCartridgeNotExists() {

	testdeck.DeckDriver = suite.driverMock

	res, err := testdeck.AspireDispense(testAspireDispenseRecord, 3)

	assert.Equal(suite.T(), "", res)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "sourceCartridge doesn't exist", err.Error())

}

func (suite *AspireDispenseTestSuite) TestAspireDispenseWrongCategory() {

	testdeck.DeckDriver = suite.driverMock
	var wrongCatergory db.Category = "wrongCatergory"
	testAspireDispenseRecord.Category = wrongCatergory

	res, err := testdeck.AspireDispense(testAspireDispenseRecord, 1)

	assert.Equal(suite.T(), "", res)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "category is invalid for aspire_dispense opeartion", err.Error())
	suite.driverMock.AssertExpectations(suite.T())
	testAspireDispenseRecord.Category = db.WW
}
func (suite *AspireDispenseTestSuite) TestAspireDispenseWrongSourcePosition() {

	testdeck.DeckDriver = suite.driverMock

	testAspireDispenseRecord.SourcePosition = 10

	res, err := testdeck.AspireDispense(testAspireDispenseRecord, 1)

	assert.Equal(suite.T(), "", res)
	assert.NotNil(suite.T(), err)
	suite.driverMock.AssertExpectations(suite.T())
	assert.Equal(suite.T(), "sourceCartridge doesn't exist", err.Error())
	testAspireDispenseRecord.SourcePosition = 4
}
func (suite *AspireDispenseTestSuite) TestAspireDispenseWrongDestinationPosition() {

	testdeck.DeckDriver = suite.driverMock
	testdeck.name = "C"

	testAspireDispenseRecord.DestinationPosition = 10

	res, err := testdeck.AspireDispense(testAspireDispenseRecord, 1)

	assert.Equal(suite.T(), "", res)
	assert.NotNil(suite.T(), err)
	suite.driverMock.AssertExpectations(suite.T())
	assert.Equal(suite.T(), "destinationCartridge doesn't exist", err.Error())
	testAspireDispenseRecord.DestinationPosition = 8
}

func (suite *AspireDispenseTestSuite) TestAspireDispenseWrongDeck() {

	testdeck.DeckDriver = suite.driverMock
	testdeck.name = "C"
	aborted.Store("C", false)

	res, err := testdeck.AspireDispense(testAspireDispenseRecord, 1)

	assert.Equal(suite.T(), "", res)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed to load syringe module for deck: C", err.Error())
	suite.driverMock.AssertExpectations(suite.T())
}
