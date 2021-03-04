package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
)

/****ALGORITHM******
1. Call Tip Pick Up at position 3
2. get cartridge start distance
3. Pierce well after well
	3.1 Move deck to the well position
	3.2 Pierce and come back up
4. Call Tip Discard

********/

func (d *Compact32Deck) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {

	var deckAndMotor DeckNumber
	var position, cartridgeStart, piercingHeight, distToTravel float64
	var ok bool
	var direction, pulses, piercingPulses uint16
	deckAndMotor.Deck = d.name
	deckAndMotor.Number = K5_Deck
	uniqueCartridge := UniqueCartridge{
		CartridgeID:   cartridgeID,
		CartridgeType: pi.Type,
	}

	//*************
	// Tip Pickup *
	//*************

	// 1. Call Tip Pick Up at position 3

	// 3rd position is where by default piercing tip is present
	// TODO: Think about removing hard coded position 3
	// One way is to separate out the tip pickup operation
	response, err = d.TipPickup(3)
	if err != nil {
		return
	}

	// 2. get cartridge start distance

	if cartridgeStart, ok = consDistance[string(pi.Type)+"_start"]; !ok {
		err = fmt.Errorf(string(pi.Type) + "_start doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	// piercingHeight is dependent on cartridge type
	if piercingHeight, ok = consDistance["piercing_height_"+string(pi.Type)]; !ok {
		err = fmt.Errorf("piercing_height_" + string(pi.Type) + " doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	// 3. Pierce well after well

	//*************************
	// Pierce Well after Well *
	//*************************

	deckAndMotor.Number = K9_Syringe_Module_LHRH

	// Calculation below considers syringe module as glued with tip
	// And we go to piercingHeight
	distToTravel = piercingHeight - positions[deckAndMotor]
	// We know concrete direction here
	// piercingHeight will be less

	piercingPulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	for _, wellNumber := range pi.CartridgeWells {
		//
		// 3.1 Move deck to the well position
		//
		deckAndMotor.Number = K5_Deck
		uniqueCartridge.WellNum = wellNumber
		if position, ok = cartridges[uniqueCartridge]["distance"]; !ok {
			err = fmt.Errorf("distance doesn't exist for well number %d", wellNumber)
			fmt.Println("Error: ", err)
			return "", err
		}

		// here disToTravel moves our deck to well position
		// position + cartridgeStart is the distance of first well on deck
		distToTravel = positions[deckAndMotor] - (position + cartridgeStart)

		switch {
		// distToTravel > 0 means go towards the Sensor or FWD
		case distToTravel > 0.1:
			direction = 1
		case distToTravel < -0.1:
			distToTravel *= -1
			direction = 0
		default:
			// Skip the setUpMotor Step
			goto skipDeckMovement
		}

		fmt.Println("Completed Move Deck to reach the wellNum ", wellNumber)

		pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Deck to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

	skipDeckMovement:
		fmt.Println("Completed Move Deck to reach the wellNum ", wellNumber)

		// 3.2 Pierce and come back up

		// WE know concrete direction here, its DOWN
		deckAndMotor.Number = K9_Syringe_Module_LHRH

		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], piercingPulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module DOWN to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		// WE know concrete direction here, its UP
		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], piercingPulses, motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module UP to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}
	}

	//**************
	// Tip Discard *
	//**************
	response, err = d.TipDiscard()
	if err != nil {
		err = fmt.Errorf("there was a problem while discarding the piercing tip")
		return "", err
	}

	return "Successfully completed tip operation", nil
}
