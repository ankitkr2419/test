package compact32

import (
	"mylab/cpagent/db"
)

type DeckNumber struct {
	Deck   string
	Number uint16
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
	finalSensorCutPulses                = uint16(2999)
)

// Special Speeds
const (
	homingFastSpeed = uint16(2000)
	homingSlowSpeed = uint16(500)
)

var wrotePulses = map[string]uint16{
	"A": 0,
	"B": 0,
}
var executedPulses = map[string]uint16{
	"A": 0,
	"B": 0,
}
var sensorHasCut = map[string]bool{
	"A": false,
	"B": false,
}
var aborted = map[string]bool{
	"A": false,
	"B": false,
}
var paused = map[string]bool{
	"A": false,
	"B": false,
}
var runInProgress = map[string]bool{
	"A": false,
	"B": false,
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
var labwares = make(map[int]string)
var tipstubes = make(map[string]map[string]float64)
var cartridges = make(map[int]map[string]float64)
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
			deckAndNumber.Deck = "A"
			deckAndNumber.Number = uint16(cd.ID - 1050)
			calibs[deckAndNumber] = cd.Distance
		}
	}
	return
}

//***NOTE***
// ids from 1001 - 1100 are reserved for Calibration variables.
// 1001- 1050 for deck A
// 1051- 1100 for deck B

func SelectAllLabwares(store db.Storer) (err error) {
	allLabwares, err := store.ListLabwares()
	if err != nil {
		return
	}

	for _, labware := range allLabwares {
		labwares[labware.ID] = labware.Name
	}
	return
}

func SelectAllTipsTubes(store db.Storer) (err error) {
	allTipsTubes, err := store.ListTipsTubes()
	if err != nil {
		return
	}

	for _, tiptube := range allTipsTubes {
		tipstubes[tiptube.Name] = make(map[string]float64)
		tipstubes[tiptube.Name]["labware_id"] = float64(tiptube.LabwareID)
		tipstubes[tiptube.Name]["consumable_distance_id"] = float64(tiptube.ConsumabledistanceID)
		tipstubes[tiptube.Name]["volume"] = tiptube.Volume
		tipstubes[tiptube.Name]["height"] = tiptube.Height
	}
	return
}

func SelectAllCartridge(store db.Storer) (err error) {
	allCartridges, err := store.ListCartridges()
	if err != nil {
		return
	}

	for _, cartridge := range allCartridges {
		cartridges[cartridge.ID] = make(map[string]float64)
		cartridges[cartridge.ID]["labware_id"] = float64(cartridge.LabwareID)
		cartridges[cartridge.ID]["wells_num"] = float64(cartridge.WellNum)
		cartridges[cartridge.ID]["distance"] = cartridge.Distance
		cartridges[cartridge.ID]["height"] = cartridge.Height
		cartridges[cartridge.ID]["volume"] = cartridge.Volume
	}
	return
}
