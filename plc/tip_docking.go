package plc

import (
	"fmt"
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
	syringeModuleDeckAndMotor = DeckNumber{Deck: d.Name, Number: K9_Syringe_Module_LHRH}
	deckAndMotor = DeckNumber{Deck: d.Name, Number: K5_Deck}

	if position, ok = consDistance["resting_position"]; !ok {
		err = fmt.Errorf("resting_position doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	distanceToTravel = positions[syringeModuleDeckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[syringeModuleDeckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[syringeModuleDeckAndMotor]["fast"], pulses, motors[syringeModuleDeckAndMotor]["ramp"], direction, syringeModuleDeckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
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
			fmt.Println("Error: ", err)
			return "", err
		}
		uniqueCartridge.WellNum = td.Position

		if cartridge, ok = cartridges[uniqueCartridge]; !ok {
			err = fmt.Errorf("cartridge doesn't exist")
			fmt.Println("Error: ", err)
			return "", err
		}
		wellPosition, ok = cartridge["distance"]
		if !ok {
			err = fmt.Errorf(" Cartridge well doesn't exist for tip docking")
			fmt.Println("Error: ", err)
			return "", err
		}

		distanceToTravel = positions[deckAndMotor] - (cartridgePosition + wellPosition)

		modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

		pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

		// 3.2 Then travel to that position

		response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
		}

		fmt.Println("deck moved to required position for docking")
		goto skipToPositionSyringeHeight
	}
	//
	// 4. Deck Flow: If it is some position on the deck
	//

	// 4.1 first calculate the accurate position on deck
	deckPosition = "pos_" + fmt.Sprintf("%d", td.Position)
	if position, ok = consDistance[deckPosition]; !ok {
		err = fmt.Errorf("%s doesn't exist for consumable distances", deckPosition)
		fmt.Println("Error: ", err)
		return "", err
	}

	// 4.2 Move to that position on deck
	distanceToTravel = positions[deckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module for tip docking. Error: %v", err)
	}

	//
	// 5. Then set the syringe height to the specified value for docking
	//
skipToPositionSyringeHeight:

	distanceToTravel = positions[deckAndMotor] - td.Height

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[syringeModuleDeckAndMotor]["steps"]) * distanceToTravel))

	//
	// 6. At that position move the syringe module to the specified height
	//
	response, err = d.setupMotor(motors[syringeModuleDeckAndMotor]["fast"], pulses, motors[syringeModuleDeckAndMotor]["ramp"], direction, syringeModuleDeckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module for tip docking. Error: %v", err)
	}

	// 7. Return success
	fmt.Println("tip docked successfully")
	return "Success", nil

}
