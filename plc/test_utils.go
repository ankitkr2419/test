package plc

import "mylab/cpagent/db"

var TestTTObj = db.TipsTubes{
	ID:               1,
	Name:             "testTip",
	Type:             "tip",
	AllowedPositions: []int64{1, 2},
	Volume:           2,
	Height:           1.1,
}

var TestCartridgeObj = []db.Cartridge{
	db.Cartridge{
		ID:          1,
		Type:        db.Cartridge1,
		Description: "test 1"},
	db.Cartridge{
		ID:          2,
		Type:        db.Cartridge2,
		Description: "test 2"},
}

var TestCartridgeWellsObj = []db.CartridgeWells{
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

var TestConsDistanceObj = []db.ConsumableDistance{
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
	db.ConsumableDistance{
		ID:          11,
		Name:        "deck_base",
		Distance:    135,
		Description: "testing ",
	},
}

var TestMotorObj = []db.Motor{
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
