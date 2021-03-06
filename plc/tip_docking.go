package plc

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"math"
	"mylab/cpagent/db"
)

/*TipDocking : to place the tip at rest position.
------ALGORITHM
1. First move the syringe module to the resting position
2. Check it is not deck position type so that cartidge_flow can be proceed. Else skip to deck flow
3. Cartridge Flow :
	3.1 first calculate the accurate position of the well
    3.2 Then travel to that position
4. Deck Flow: If it is some position on the deck
	4.1 first calculate the accurate position on deck
	4.2 Move to that position on deck
5. Then set the syringe height to the specified value for docking
6. At that position move the syringe module to the specified height
7. Return success
*/
func (d *Compact32Deck) TipDocking(td db.TipDock, cartridgeID int64) (response string, err error) {

	var position, distanceToTravel, cartridgePosition, wellPosition float64
	var cartridge map[string]float64
	var direction, pulses uint16
	var deckPosition string
	var ok bool
	var deckAndMotor, syringeModuleDeckAndMotor DeckNumber

	//
	// Step 1: move the syringe module to resting position
	//
	syringeModuleDeckAndMotor = DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}
	deckAndMotor = DeckNumber{Deck: d.name, Number: K5_Deck}

	if position, ok = consDistance["resting_position"]; !ok {
		err = fmt.Errorf("resting_position doesn't exist for consumable distances")
		logger.Errorln(err)
		return "", err
	}
	distanceToTravel = Positions[syringeModuleDeckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[syringeModuleDeckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[syringeModuleDeckAndMotor]["fast"], pulses, Motors[syringeModuleDeckAndMotor]["ramp"], direction, syringeModuleDeckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}

	//
	// 2. Check it is not deck position type so that cartidge_flow can be proceed. Else skip to deck flow
	//
	if td.Type != "deck" {
		//
		// 3. Cartridge Flow :
		//

		uniqueCartridge := UniqueCartridge{
			CartridgeID:   cartridgeID,
			CartridgeType: db.CartridgeType(td.Type),
		}

		// 3.1 first calculate the accurate position of the well

		// distance to cartridge start + distance to the specified well
		if cartridgePosition, ok = consDistance[td.Type+"_start"]; !ok {
			err = fmt.Errorf(td.Type + "_start doesn't exist for consumable distances")
			logger.Errorln(err)
			return "", err
		}
		uniqueCartridge.WellNum = td.Position

		if cartridge, ok = cartridges[uniqueCartridge]; !ok {
			err = fmt.Errorf("cartridge doesn't exist")
			logger.Errorln(err)
			return "", err
		}
		wellPosition, ok = cartridge["distance"]
		if !ok {
			err = fmt.Errorf(" Cartridge well doesn't exist for tip docking")
			logger.Errorln(err)
			return "", err
		}

		distanceToTravel = Positions[deckAndMotor] - (cartridgePosition + wellPosition)

		modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

		pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

		// 3.2 Then travel to that position

		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
		}

		logger.Infoln("deck moved to required position for docking")
		goto skipToPositionSyringeHeight
	}
	//
	// 4. Deck Flow: If it is some position on the deck
	//

	// 4.1 first calculate the accurate position on deck
	deckPosition = "pos_" + fmt.Sprintf("%d", td.Position)
	if position, ok = consDistance[deckPosition]; !ok {
		err = fmt.Errorf("%s doesn't exist for consumable distances", deckPosition)
		logger.Errorln(err)
		return "", err
	}

	// 4.2 Move to that position on deck
	distanceToTravel = Positions[deckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module for tip docking. Error: %v", err)
	}

	//
	// 5. Then set the syringe height to the specified value for docking
	//
skipToPositionSyringeHeight:

	distanceToTravel = Positions[deckAndMotor] - td.Height

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[syringeModuleDeckAndMotor]["steps"]) * distanceToTravel))

	//
	// 6. At that position move the syringe module to the specified height
	//
	response, err = d.setupMotor(Motors[syringeModuleDeckAndMotor]["fast"], pulses, Motors[syringeModuleDeckAndMotor]["ramp"], direction, syringeModuleDeckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module for tip docking. Error: %v", err)
	}

	// 7. Return success
	logger.Infoln("tip docked successfully")
	return "Success", nil

}
