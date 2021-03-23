package compact32

import (
	"fmt"
	"mylab/cpagent/db"
	"sync"
)

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
	highestUint16 						= uint16(65535)
)

// 120 Seconds is the minimum UVLight On Time
const (
	minimumUVLightOnTime int64 = 2 * 60
)

// Special Speeds
const (
	homingFastSpeed = uint16(2000)
	homingSlowSpeed = uint16(500)
	homingDeckFastSpeed = uint16(3000)
)

// Magnet States
const (
	detached = iota
	semiDetached
	attached
)

var wrotePulses, executedPulses, aborted, paused, runInProgress, magnetState, timerInProgress, homed sync.Map

func LoadUtils() {
	wrotePulses.Store("A", uint16(0))
	wrotePulses.Store("B", uint16(0))
	executedPulses.Store("A", uint16(0))
	executedPulses.Store("B", uint16(0))
	aborted.Store("A", false)
	aborted.Store("B", false)
	paused.Store("A", false)
	paused.Store("B", false)
	runInProgress.Store("A", false)
	runInProgress.Store("B", false)
	timerInProgress.Store("A", false)
	timerInProgress.Store("B", false)
	magnetState.Store("A", detached)
	magnetState.Store("B", detached)
	homed.Store("A", false)
	homed.Store("B", false)
}

// positions = map[deck(A or B)]map[motor number(1 to 10)]distance(only positive)
var positions = map[DeckNumber]float64{
	// Deck A and its Motors
	DeckNumber{Deck: "A", Number: K1_Syringe_Module_LH}:   0,
	DeckNumber{Deck: "A", Number: K2_Syringe_Module_RH}:   0,
	DeckNumber{Deck: "A", Number: K3_Syringe_LH}:          0,
	DeckNumber{Deck: "A", Number: K4_Syringe_RH}:          0,
	DeckNumber{Deck: "A", Number: K5_Deck}:                0,
	DeckNumber{Deck: "A", Number: K6_Magnet_Up_Down}:      0,
	DeckNumber{Deck: "A", Number: K7_Magnet_Rev_Fwd}:      0,
	DeckNumber{Deck: "A", Number: K8_Shaker}:              0,
	DeckNumber{Deck: "A", Number: K9_Syringe_Module_LHRH}: 0,
	DeckNumber{Deck: "A", Number: K10_Syringe_LHRH}:       0,
	// Deck B and its Motors
	DeckNumber{Deck: "B", Number: K1_Syringe_Module_LH}:   0,
	DeckNumber{Deck: "B", Number: K2_Syringe_Module_RH}:   0,
	DeckNumber{Deck: "B", Number: K3_Syringe_LH}:          0,
	DeckNumber{Deck: "B", Number: K4_Syringe_RH}:          0,
	DeckNumber{Deck: "B", Number: K5_Deck}:                0,
	DeckNumber{Deck: "B", Number: K6_Magnet_Up_Down}:      0,
	DeckNumber{Deck: "B", Number: K7_Magnet_Rev_Fwd}:      0,
	DeckNumber{Deck: "B", Number: K8_Shaker}:              0,
	DeckNumber{Deck: "B", Number: K9_Syringe_Module_LHRH}: 0,
	DeckNumber{Deck: "B", Number: K10_Syringe_LHRH}:       0,
	//***WARNING
	//* Careful when dealing with K1, K2, K3 and K4
}

var motors = make(map[DeckNumber]map[string]uint16)
var consDistance = make(map[string]float64)
var tipstubes = make(map[string]map[string]interface{})
var labwares = make(map[int]string)
var cartridges = make(map[UniqueCartridge]map[string]float64)
var calibs = make(map[DeckNumber]float64)

func SelectAllMotors(store db.Storer) (err error) {
	allMotors, err := store.ListMotors()
	if err != nil {
		return
	}

	for _, motor := range allMotors {
		deckNumber := DeckNumber{Deck: motor.Deck, Number: uint16(motor.Number)}
		motors[deckNumber] = make(map[string]uint16)
		motors[deckNumber]["ramp"] = uint16(motor.Ramp)
		motors[deckNumber]["steps"] = uint16(motor.Steps)
		motors[deckNumber]["slow"] = uint16(motor.Slow)
		motors[deckNumber]["fast"] = uint16(motor.Fast)
	}
	return
}

func SelectAllConsDistances(store db.Storer) (err error) {
	allConsDistances, err := store.ListConsDistances()
	if err != nil {
		return
	}

	for _, cd := range allConsDistances {
		consDistance[cd.Name] = cd.Distance
		deckAndNumber := DeckNumber{}
		switch {
		case cd.ID > 1000 && cd.ID <= 1010:
			deckAndNumber.Deck = "A"
			deckAndNumber.Number = uint16(cd.ID - 1000)
			calibs[deckAndNumber] = cd.Distance
		case cd.ID > 1050 && cd.ID <= 1060:
			deckAndNumber.Deck = "B"
			deckAndNumber.Number = uint16(cd.ID - 1050)
			calibs[deckAndNumber] = cd.Distance
		}
	}
	fmt.Println("Calibs:--->", calibs)
	return
}

//***NOTE***
// ids from 1001 - 1100 are reserved for Calibration variables.
// 1001- 1050 for deck A
// 1051- 1100 for deck B

func SelectAllTipsTubes(store db.Storer) (err error) {
	allTipsTubes, err := store.ListTipsTubes()
	if err != nil {
		return
	}

	for _, tiptube := range allTipsTubes {
		tipstubes[tiptube.Name] = make(map[string]interface{})
		tipstubes[tiptube.Name]["id"] = tiptube.ID
		tipstubes[tiptube.Name]["type"] = tiptube.Type
		tipstubes[tiptube.Name]["allowed_positions"] = tiptube.AllowedPositions
		tipstubes[tiptube.Name]["volume"] = tiptube.Volume
		tipstubes[tiptube.Name]["height"] = tiptube.Height
	}
	return
}

func SelectAllCartridges(store db.Storer) (err error) {
	allCartridges, err := store.ListCartridges()
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
