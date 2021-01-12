package compact32

import (
	"mylab/cpagent/db"
)

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

type DeckNumber struct {
	Deck   string
	Number uint16
}

var motors = make(map[DeckNumber]map[string]uint16)
var consDistance = make(map[string]float64)
var labwares = make(map[int]string)
var tipstubes = make(map[string]map[string]float64)
var cartridges = make(map[int]map[string]float64)

func SelectAllMotors(store db.Storer) (err error) {
	allMotors, err := store.GetAllMotors()
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
	allConsDistances, err := store.GetAllConsDistances()
	if err != nil {
		return
	}

	for _, distance := range allConsDistances {
		consDistance[distance.Name] = distance.Distance
	}
	return
}

func SelectAllLabwares(store db.Storer) (err error) {
	allLabwares, err := store.GetAllLabwares()
	if err != nil {
		return
	}

	for _, labware := range allLabwares {
		labwares[labware.ID] = labware.Name
	}
	return
}

func SelectAllTipsTubes(store db.Storer) (err error) {
	allTipsTubes, err := store.GetAllTipsTubes()
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
	allCartridges, err := store.GetAllCartridges()
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
