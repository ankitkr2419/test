package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
)

func checkCartridgeType(recipe db.Recipe, cT db.CartridgeType, cartridgeID *int64) error {
	switch cT {
	case db.Cartridge1:
		if recipe.Cartridge1Position == nil {
			return responses.RecipeCartridge1Missing
		}
		*cartridgeID = *recipe.Cartridge1Position

	case db.Cartridge2:
		if recipe.Cartridge2Position == nil {
			return responses.RecipeCartridge2Missing
		}
		*cartridgeID = *recipe.Cartridge2Position
	default:
		return responses.InvalidCartridgeType
	}
	return nil
}

func isDeckPositionInvalid(position int64) bool {
	if position == cartridge1Pos || position == cartridge2Pos ||
		position < minAspDisDeckPos || position > maxDeckPosition {
		return true
	}
	return false
}

func createCartridgeWell(cartridgeID int64, cT db.CartridgeType, wellNum int64) plc.UniqueCartridge {
	return plc.UniqueCartridge{
		CartridgeID:   cartridgeID,
		CartridgeType: cT,
		WellNum:       wellNum,
	}
}
