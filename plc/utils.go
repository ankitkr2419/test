package plc

import (
	"context"
	"mylab/cpagent/db"

	"sync"

	logger "github.com/sirupsen/logrus"
)

var HeatingCycleComplete, CycleComplete, DataCapture, ExperimentRunning, LidPidTuningInProgress bool
var CurrentCycleTemperature, CurrentLidTemp float32
var CurrentCycle uint16

type DeckNumber struct {
	Deck   string
	Number uint16
}

// Each Cartridge can be uniquely identified by these fields
type UniqueCartridge struct {
	CartridgeID   int64
	CartridgeType db.CartridgeType
	WellNum       int64
}

// motors
const (
	K1_Syringe_Module_LH = uint16(iota + 1)
	K2_Syringe_Module_RH
	K3_Syringe_LH
	K4_Syringe_RH
	K5_Deck
	K6_Magnet_Up_Down
	K7_Magnet_Rev_Fwd
	K8_Shaker
	K9_Syringe_Module_LHRH
	K10_Syringe_LHRH
)

// All these are special + max Pulses
const (
	initialSensorCutDeckPulses          = uint16(59199)
	initialSensorCutSyringePulses       = uint16(26666)
	initialSensorCutSyringeModulePulses = uint16(29999)
	initialSensorCutMagnetPulses        = uint16(29999)
	moveOppositeSensorPulses            = uint16(19999)
	reverseAfterNonCutPulses            = uint16(2000)
	reverseAfterNonCutPulsesMagnet      = uint16(400)
	finalSensorCutPulses                = uint16(2999)
	minimumPulsesThreshold              = uint16(50)
	highestUint16                       = uint16(65535)
)

// 120 Seconds is the minimum UVLight On Time
const (
	minimumUVLightOnTime int64 = 2 * 60
)

// here we are hardcoding the shaker no in future this is to be fetched dynamically.
// 3 is the value that needs to be passed for heating both the shakers.
const (
	shaker = uint16(3)
)

// Special Speeds
const (
	homingFastSpeed     = uint16(2000)
	homingSlowSpeed     = uint16(500)
	homingDeckFastSpeed = uint16(3000)
)

// Magnet States
const (
	detached = iota
	attached
)

// Syringe module States
const (
	InDeck = iota
	OutDeck
)

// Deck Names
const (
	DeckA = "A"
	DeckB = "B"
)

var deckRecipe map[string]db.Recipe
var deckProcesses map[string][]db.Process
var wrotePulses, executedPulses, aborted, paused, homed, EngineerOrAdminLogged sync.Map
var runInProgress, magnetState, timerInProgress, heaterInProgress sync.Map
var uvLightInProgress, syringeModuleState, shakerInProgress, tipDiscardInProgress, motorOperationCompleted sync.Map
var pIDCalibrationInProgress sync.Map

// tipHeight is the Height of tip from syringe's base
var tipHeight map[string]float64

// Special variables for both deck operation
var BothDeckHomingInProgress bool
var homingPercent, currentProcess sync.Map

// variable map Registers to keep track of machine related variables.
var motorNumReg, speedReg, directionReg, rampReg, pulseReg, onReg sync.Map

func loadUtils() {

	wrotePulses.Store(DeckA, uint16(0))
	wrotePulses.Store(DeckB, uint16(0))
	executedPulses.Store(DeckA, uint16(0))
	executedPulses.Store(DeckB, uint16(0))
	aborted.Store(DeckA, false)
	aborted.Store(DeckB, false)
	paused.Store(DeckA, false)
	paused.Store(DeckB, false)
	runInProgress.Store(DeckA, false)
	runInProgress.Store(DeckB, false)
	timerInProgress.Store(DeckA, false)
	timerInProgress.Store(DeckB, false)
	heaterInProgress.Store(DeckA, false)
	heaterInProgress.Store(DeckB, false)
	shakerInProgress.Store(DeckA, false)
	shakerInProgress.Store(DeckB, false)
	tipDiscardInProgress.Store(DeckA, false)
	tipDiscardInProgress.Store(DeckB, false)
	uvLightInProgress.Store(DeckA, false)
	uvLightInProgress.Store(DeckB, false)
	magnetState.Store(DeckA, detached)
	magnetState.Store(DeckB, detached)
	syringeModuleState.Store(DeckA, OutDeck)
	syringeModuleState.Store(DeckB, OutDeck)
	homed.Store(DeckA, false)
	homed.Store(DeckB, false)
	motorOperationCompleted.Store(DeckA, false)
	motorOperationCompleted.Store(DeckB, false)
	pIDCalibrationInProgress.Store("A", false)
	pIDCalibrationInProgress.Store("B", false)
	EngineerOrAdminLogged.Store("A", false)
	EngineerOrAdminLogged.Store("B", false)

	deckRecipe = map[string]db.Recipe{
		DeckA: db.Recipe{},
		DeckB: db.Recipe{},
	}

	deckProcesses = map[string][]db.Process{
		DeckA: []db.Process{},
		DeckB: []db.Process{},
	}

	tipHeight = map[string]float64{
		DeckA: 0,
		DeckB: 0,
	}

	BothDeckHomingInProgress = false
	homingPercent.Store(DeckA, float64(0))
	homingPercent.Store(DeckB, float64(0))
	currentProcess.Store(DeckA, int64(-1))
	currentProcess.Store(DeckB, int64(-1))

	motorNumReg.Store(DeckA, uint16(0))
	motorNumReg.Store(DeckB, uint16(0))
	speedReg.Store(DeckA, uint16(0))
	speedReg.Store(DeckB, uint16(0))
	directionReg.Store(DeckA, uint16(0))
	directionReg.Store(DeckB, uint16(0))
	rampReg.Store(DeckA, uint16(0))
	rampReg.Store(DeckB, uint16(0))
	pulseReg.Store(DeckA, uint16(0))
	pulseReg.Store(DeckB, uint16(0))
	onReg.Store(DeckA, OFF)
	onReg.Store(DeckB, OFF)
}

// positions = map[deck(A or B)]map[motor number(1 to 10)]distance(only positive)
var Positions = map[DeckNumber]float64{
	// Deck A and its Motors
	DeckNumber{Deck: DeckA, Number: K1_Syringe_Module_LH}:   0,
	DeckNumber{Deck: DeckA, Number: K2_Syringe_Module_RH}:   0,
	DeckNumber{Deck: DeckA, Number: K3_Syringe_LH}:          0,
	DeckNumber{Deck: DeckA, Number: K4_Syringe_RH}:          0,
	DeckNumber{Deck: DeckA, Number: K5_Deck}:                0,
	DeckNumber{Deck: DeckA, Number: K6_Magnet_Up_Down}:      0,
	DeckNumber{Deck: DeckA, Number: K7_Magnet_Rev_Fwd}:      0,
	DeckNumber{Deck: DeckA, Number: K8_Shaker}:              0,
	DeckNumber{Deck: DeckA, Number: K9_Syringe_Module_LHRH}: 0,
	DeckNumber{Deck: DeckA, Number: K10_Syringe_LHRH}:       0,
	// Deck B and its Motors
	DeckNumber{Deck: DeckB, Number: K1_Syringe_Module_LH}:   0,
	DeckNumber{Deck: DeckB, Number: K2_Syringe_Module_RH}:   0,
	DeckNumber{Deck: DeckB, Number: K3_Syringe_LH}:          0,
	DeckNumber{Deck: DeckB, Number: K4_Syringe_RH}:          0,
	DeckNumber{Deck: DeckB, Number: K5_Deck}:                0,
	DeckNumber{Deck: DeckB, Number: K6_Magnet_Up_Down}:      0,
	DeckNumber{Deck: DeckB, Number: K7_Magnet_Rev_Fwd}:      0,
	DeckNumber{Deck: DeckB, Number: K8_Shaker}:              0,
	DeckNumber{Deck: DeckB, Number: K9_Syringe_Module_LHRH}: 0,
	DeckNumber{Deck: DeckB, Number: K10_Syringe_LHRH}:       0,
	//***WARNING
	//* Careful when dealing with K1, K2, K3 and K4
}

var Motors = make(map[DeckNumber]map[string]uint16)
var consDistance = make(map[string]float64)
var tipstubes = make(map[int64]map[string]interface{})
var labwares = make(map[int]string)
var cartridges = make(map[UniqueCartridge]map[string]float64)
var Calibs = make(map[DeckNumber]float64)

func LoadAllPLCFuncs(store db.Storer) (err error) {

	err = selectAllMotors(store)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Select All Motors failed")
		return
	}

	err = selectAllConsDistances(store)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Select All Cosumable Distances failed")
		return
	}

	err = selectAllTipsTubes(store)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Select All Tips and Tubes failed")
		return
	}

	err = selectAllCartridges(store)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Select All Cartridge failed")
		return
	}

	loadUtils()

	return nil
}

func selectAllMotors(store db.Storer) (err error) {
	allMotors, err := store.ListMotors(context.Background())
	if err != nil {
		return
	}

	for _, motor := range allMotors {
		deckNumber := DeckNumber{Deck: motor.Deck, Number: uint16(motor.Number)}
		Motors[deckNumber] = make(map[string]uint16)
		Motors[deckNumber]["ramp"] = uint16(motor.Ramp)
		Motors[deckNumber]["steps"] = uint16(motor.Steps)
		Motors[deckNumber]["slow"] = uint16(motor.Slow)
		Motors[deckNumber]["fast"] = uint16(motor.Fast)
	}
	return
}

func selectAllConsDistances(store db.Storer) (err error) {
	allConsDistances, err := store.ListConsDistances()
	if err != nil {
		return
	}

	for _, cd := range allConsDistances {
		consDistance[cd.Name] = cd.Distance
		deckAndNumber := DeckNumber{}
		switch {
		case cd.ID > 1000 && cd.ID <= 1010:
			deckAndNumber.Deck = DeckA
			deckAndNumber.Number = uint16(cd.ID - 1000)
			Calibs[deckAndNumber] = cd.Distance
		case cd.ID > 1050 && cd.ID <= 1060:
			deckAndNumber.Deck = DeckB
			deckAndNumber.Number = uint16(cd.ID - 1050)
			Calibs[deckAndNumber] = cd.Distance
		}
	}
	logger.Infoln("Calibs:--->", Calibs)
	return
}

//***NOTE***
// ids from 1001 - 1100 are reserved for Calibration variables.
// 1001- 1050 for deck A
// 1051- 1100 for deck B

func selectAllTipsTubes(store db.Storer) (err error) {
	ttype := ""
	allTipsTubes, err := store.ListTipsTubes(ttype)
	if err != nil {
		return
	}

	for _, tiptube := range allTipsTubes {
		tipstubes[tiptube.ID] = make(map[string]interface{})
		tipstubes[tiptube.ID]["name"] = tiptube.Name
		tipstubes[tiptube.ID]["id"] = tiptube.ID
		tipstubes[tiptube.ID]["type"] = tiptube.Type
		tipstubes[tiptube.ID]["allowed_positions"] = tiptube.AllowedPositions
		tipstubes[tiptube.ID]["volume"] = tiptube.Volume
		tipstubes[tiptube.ID]["height"] = tiptube.Height
		tipstubes[tiptube.ID]["tt_base"] = tiptube.TtBase
	}
	return
}

func selectAllCartridges(store db.Storer) (err error) {

	// here passing context since we need username and as when the binary runs the
	// first time there is no login information hence setting username as main.
	ctx := context.WithValue(context.Background(), db.ContextKeyUsername, "main")

	allCartridges, err := store.ListCartridges(ctx)
	if err != nil {
		return
	}

	cartridgeType := make(map[int64]db.CartridgeType)

	for _, cartridge := range allCartridges {
		cartridgeType[cartridge.ID] = cartridge.Type
	}

	allCartridgeWells, err := store.ListCartridgeWells()
	if err != nil {
		return
	}

	for _, well := range allCartridgeWells {
		uniqueCartridge := UniqueCartridge{
			CartridgeID:   well.ID,
			CartridgeType: cartridgeType[well.ID],
			WellNum:       well.WellNum,
		}
		cartridges[uniqueCartridge] = make(map[string]float64)
		cartridges[uniqueCartridge]["distance"] = well.Distance
		cartridges[uniqueCartridge]["height"] = well.Height
		cartridges[uniqueCartridge]["volume"] = well.Volume
	}
	return
}

// modifyDirectionAndDistanceToTravel will make distanceToTravel positive and the direction correct
func modifyDirectionAndDistanceToTravel(distanceToTravel *float64, direction *uint16) {
	// distanceToTravel > 0 means go towards the Sensor
	if *distanceToTravel > 0 {
		*direction = 1
	} else {
		*distanceToTravel *= -1
		*direction = 0
	}
}
