package plc

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
	"sort"
)

/****ALGORITHM******

1. get cartridge start distance
2. Pierce well after well
	2.1 Move deck to the well position
	2.2 Pierce and come back up
	2.3 Repeat step 2.1 and 2.2 till well exists
********/

func (d *Compact32Deck) Piercing(pi db.Piercing, cartridgeID int64, tip db.TipsTubes) (response string, err error) {

	var deckAndMotor DeckNumber
	var position, cartridgeStart, piercingHeight, distanceToTravel, tipHeight, deckBase float64
	var ok bool
	var direction, pulses, piercingPulses, afterPiercingRestPulses uint16
	// []int has direct method to get slice sorted
	var wellsToBePierced []int
	deckAndMotor.Deck = d.name
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	uniqueCartridge := UniqueCartridge{
		CartridgeID:   cartridgeID,
		CartridgeType: pi.Type,
	}
	//
	// 1. get cartridge start distance
	//
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

	//
	// 2. Pierce well after well
	//
	//*************************
	// Pierce Well after Well *
	//*************************

	// Calculation below considers syringe module as glued with tip
	// And we go to piercingHeight


	// Get Deck Base
	if deckBase, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumables")
		fmt.Println("Error: ", err)
		return "", err
	}

	//-----------------
	// Get Tip Height -
	//-----------------
	var tipHeightInter interface{}
	if tipHeightInter, ok = tipstubes[tip.ID]["height"]; !ok {
		err = fmt.Errorf("%v tip doesn't exist for tipstubes", tip.ID)
		fmt.Println("Error: ", err)
		return "", err
	}

	if tipHeight, ok = tipHeightInter.(float64); !ok {
		err = fmt.Errorf("%v tip has unknown ID!", tip.ID)
		fmt.Println("Error: ", err)
		return "", err
	}

	if tipHeightInter, ok = consDistance["slow_inside"]; !ok {
		err = fmt.Errorf("slow_inside doesn't exist for consumables")
		fmt.Println("Error: ", err)
		return "", err
	}

	if position, ok = tipHeightInter.(float64); !ok {
		err = fmt.Errorf("couldn't type cast slow_inside")
		fmt.Println("Error: ", err)
		return "", err
	}

	tipHeight -= position

	distanceToTravel = (deckBase + piercingHeight) - (Positions[deckAndMotor] + tipHeight)
	// We know concrete direction here
	// piercingHeight will be less

	piercingPulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))
	// after piercing is completed we need to get the tip to its resting positon
	afterPiercingRestPulses = piercingPulses

	for _, well := range pi.CartridgeWells {
		wellsToBePierced = append(wellsToBePierced, int(well))
	}

	// sort wells in Ascending Order
	sort.Ints(wellsToBePierced)

	for i, wellNumber := range wellsToBePierced {
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
		distanceToTravel = Positions[deckAndMotor] - (position + cartridgeStart)

		modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

		pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Deck to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		fmt.Println("Completed Move Deck to reach the wellNum ", wellNumber)

		// 2.2 Pierce and come back up

		// WE know concrete direction here, its DOWN
		deckAndMotor.Number = K9_Syringe_Module_LHRH

		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], piercingPulses, Motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
		// TODO: Use defer d.setIndeck as in aspire_dispense
		// Even if err has occured let's store syringeModuleState as inDeck
		syringeModuleState.Store(d.name, InDeck)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module DOWN to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		fmt.Println("Pierced WellNumber: ", wellNumber)

		// change piercingPulses just before going up after piercing the first well
		if i == 0 {
			// For wells other than first piercing height will be less
			if piercingHeight, ok = consDistance["piercing_tip_above_well_position"]; !ok {
				err = fmt.Errorf("piercing_tip_above_well_position doesn't exist for consumable distances")
				fmt.Println("Error: ", err)
				return "", err
			}

			// piercingHeight will be always less than current position
			distanceToTravel = Positions[deckAndMotor] - piercingHeight 

			piercingPulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))
		}

		// if its last well then go to resting position up
		if i == len(wellsToBePierced)-1 {
			piercingPulses = afterPiercingRestPulses
		}
		// WE know concrete direction here, its UP
		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], piercingPulses, Motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Syringe Module UP to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}
		// TODO: Use defer d.setIndeck as in aspire_dispense
		// Only after successful coming out do we say its OutDeck completely
		syringeModuleState.Store(d.name, OutDeck)

		fmt.Println("Got Up from WellNumber: ", wellNumber)

		// 2.3 Repeat step 2.1 and  2.2 till another well exists
	}

	return "Successfully completed piercing operation", nil
}
