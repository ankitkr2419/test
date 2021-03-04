package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
)

/*TipDocking : to place the tip at rest position.
------ALOGORITHM
1. First move the syringe module to the resting position
2. Check it is not deck position  type so that cartidge_flow can be proceeded.
3. Cartridge Flow : 1.1 first calculate the accurate position of the well
                    1.2 Then travel to that position
					1.3 Then set the syringe height to the specified value for docking
					1.4 Return Success
4. Deck Flow: If it is some position on the deck
5. Move to that position on the deck by calculating the deck position
6. At that position move the syring module to the specified height
7. Return success
*/
func (d *Compact32Deck) TipDocking(td db.TipDock, cartridgeID int64) (response string, err error) {

	var position, distanceToTravel, cartridgePosition, wellPosition float64
	var cartridge map[string]float64
	var direction, pulses uint16
	var deckPosition string
	var ok bool
	var deckAndMotor, syringeModuleDeckAndMotor DeckNumber

	// Step 1: move the syringe module to resting position
	// Deck and number for syringe module
	syringeModuleDeckAndMotor = DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}
	deckAndMotor = DeckNumber{Deck: d.name, Number: K5_Deck}

	if position, ok = consDistance["resting_position"]; !ok {
		err = fmt.Errorf("resting_position doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distanceToTravel = positions[syringeModuleDeckAndMotor] - position
	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravel > 0.1:
		direction = 1
	case distanceToTravel < -0.1:
		distanceToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipToRestPosition
	}
	pulses = uint16(math.Round(float64(motors[syringeModuleDeckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.SetupMotor(motors[syringeModuleDeckAndMotor]["fast"], pulses, motors[syringeModuleDeckAndMotor]["ramp"], direction, syringeModuleDeckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}

skipToRestPosition:
	if td.Type != "deck" {

		uniqueCartridge := UniqueCartridge{
			CartridgeID:   cartridgeID,
			CartridgeType: db.CartridgeType(td.Type),
		}

		// Step 2 : move the deck to the absolute position of cartridge.
		// get the cartridge well distance
		// distance to cartridge start + distance to the specified well
		if cartridgePosition, ok = consDistance[td.Type+"_start"]; !ok {
			err = fmt.Errorf(td.Type + "_start doesn't exist for consuamble distances")
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
		switch {
		// distToTravel > 0 means go towards the Sensor or FWD
		case distanceToTravel > 0.1:
			direction = 1
		case distanceToTravel < -0.1:
			distanceToTravel *= -1
			direction = 0
		default:
			// Skip the setUpMotor Step
			goto skipToPositionSyringeHeight
		}
		pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
		}

		fmt.Println("deck moved to required position for docking")
		goto skipToPositionSyringeHeight
	}
	// Step 5 : if the tip dock type is deck

	// calculate the position of the deck where the syringe module needs to dock
	deckPosition = "pos_" + fmt.Sprintf("%d", td.Position)
	if position, ok = consDistance[deckPosition]; !ok {
		err = fmt.Errorf("%s doesn't exist for consuamble distances", deckPosition)
		fmt.Println("Error: ", err)
		return "", err
	}

	//move the deck to the specified deck position
	distanceToTravel = positions[deckAndMotor] - position
	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravel > 0.1:
		direction = 1
	case distanceToTravel < -0.1:
		distanceToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipToPositionSyringeHeight
	}
	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module for tip docking. Error: %v", err)
	}

skipToPositionSyringeHeight:

	// Step 3 : Now move the syringe module to the given height
	deckAndMotor = DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}

	distanceToTravel = positions[deckAndMotor] - td.Height
	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravel > 0.1:
		direction = 1
	case distanceToTravel < -0.1:
		distanceToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipToCompletion
	}
	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module for tip docking. Error: %v", err)
	}

skipToCompletion:
	// Step 4: completed process
	fmt.Println("tip docked successfully")
	return "Success", nil

}
