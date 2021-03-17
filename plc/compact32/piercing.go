package compact32

import (
	"fmt"
	"math"
	"sort"
	"mylab/cpagent/db"
)

/****ALGORITHM******

1. get cartridge start distance
2. Pierce well after well
	2.1 Move deck to the well position
	2.2 Pierce and come back up
	2.3 Repeat step 2.1 and 2.2 till well exists
********/

func (d *Compact32Deck) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {

	var deckAndMotor DeckNumber
	var position, cartridgeStart, piercingHeight, distToTravel float64
	var ok bool
	var direction, pulses, piercingPulses uint16
	// []int has direct method to get slice sorted
	var wellsToBePierced []int
	deckAndMotor.Deck = d.name
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	uniqueCartridge := UniqueCartridge{
		CartridgeID:   cartridgeID,
		CartridgeType: pi.Type,
	}

	// 1. get cartridge start distance

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

	// 2. Pierce well after well

	//*************************
	// Pierce Well after Well *
	//*************************

	// Calculation below considers syringe module as glued with tip
	// And we go to piercingHeight
	distToTravel = piercingHeight - positions[deckAndMotor]
	// We know concrete direction here
	// piercingHeight will be less

	piercingPulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))


	for _, well := range pi.CartridgeWells {
		wellsToBePierced = append(wellsToBePierced, int(well))
	}

	// sort wells in Ascending Order
	sort.Ints(wellsToBePierced)
	
	for _, wellNumber := range wellsToBePierced {
		//
		// 2.1 Move deck to the well position
		//
		deckAndMotor.Number = K5_Deck
		uniqueCartridge.WellNum = int64(wellNumber)
		if position, ok = cartridges[uniqueCartridge]["distance"]; !ok {
			err = fmt.Errorf("distance doesn't exist for well number %d", wellNumber)
			fmt.Println("Error: ", err)
			return "", err
		}

		// here disToTravel moves our deck to well position
		// position + cartridgeStart is the distance of first well on deck for wellNumber = 1
		distToTravel = positions[deckAndMotor] - (position + cartridgeStart)

		switch {
		// distToTravel > 0 means go towards the Sensor or FWD
		case distToTravel > minimumMoveDistance:
			direction = 1
		case distToTravel < (minimumMoveDistance * -1):
			distToTravel *= -1
			direction = 0
		default:
			// Skip the setUpMotor Step
			goto skipDeckMovement
		}

		pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Deck to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

	skipDeckMovement:
		fmt.Println("Completed Move Deck to reach the wellNum ", wellNumber)

		// 2.2 Pierce and come back up

		// WE know concrete direction here, its DOWN
		deckAndMotor.Number = K9_Syringe_Module_LHRH

		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], piercingPulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module DOWN to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		fmt.Println("Pierced WellNumber: ", wellNumber)

		// WE know concrete direction here, its UP
		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], piercingPulses, motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module UP to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		fmt.Println("Got Up from WellNumber: ", wellNumber)

		// 2.3 Repeat step 2.1 and  2.2 till another well exists
	}

	return "Successfully completed piercing operation", nil
}
