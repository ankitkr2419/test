package compact32

import (
	"fmt"
	"math"
	"time"
)

/****ALGORITHM******
// TODO: check for height and volumes constraint at insertion process itself
variables: category, cartridgeType string,
			labwareID, type, source_well, destination_well, aspire_cycles, dispense_cycles int64,
			asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol, dis_vol, dis_blow float64

  1. Check the category of operation
  2. if category is well_to_well then goto 3 else if category is shaker_to_well then goto 5  else 7
  3. store the source_well position into sourcePosition variable
  4. store the destination_well position into destinationPosition variable; goto 9
  5. store the shaker position into sourcePosition variable
  6. store the destination_well position into destinationPosition variable; goto 9
  7. store the source_well position into sourcePosition variable
  8. store the shaker position into destinationPosition variable
  9. setup the motor of syringe module to go up atleast 30mm above deck
  10. calculate the current position difference for deck;
   		if its positive then direction is 1(towards sensor) else 0(oppose sensor)
  11. move deck to match the sourcePosition with help of difference calculated
  12. move syringe module down at fast till base
  13. setup the syringe motor with aspire height
  14. pickup and drop that asp_mix_vol for number of aspire_cycles
  15. blow the air out
  16. pickup asp_vol + air
  17. move syringe module up
  18. Move slowly to destinationPosition by calculating the difference of Positions
  19. move syringe module down at fast till base
  20. setup the syringe motor with dispense height
  21. pickup and drop that dis_mix_vol for number of dis_cycles
  22. blow the air out
  23. move syringe module up

********/

func (d *Compact32Deck) AspireDispense(category, cartridgeType string, labwareID, source, destination, aspire_cycles, dispense_cycles int64, asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol, dis_vol, dis_blow float64) (response string, err error) {

	if runInProgress[d.name] {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}
	sensorHasCut[d.name] = false
	aborted[d.name] = false
	runInProgress[d.name] = true
	defer d.ResetRunInProgress()

	var sourceCartridge, destinationCartridge map[string]float64
	var sourcePosition, destinationPosition, distToTravel, position float64
	var ok bool
	var direction, pulses uint16
	var deckAndMotor DeckNumber

	deckAndMotor.Deck = d.name

	//
	// ALGORITHM's 1 to 8 steps are implemented below
	//

	/*** GET THE CARTRIDGES
	E.g :
	********** for well_to_well category only ***********
	Suppose
		labwareID = 1 && cartridgeType = "extraction" && source = 2 && destination= 4
	Then
		sourceCartridge =
		- id: 2
			labwareID: 1
			type: "extraction"
			description: "Extraction Cartridge"
			wellNum: 2
			distance: 24.5
			height: 2
			volume: 10

	And
		destinationCartridge =
		- id: 4
			labwareID: 1
			type: "extraction"
			description: "Extraction Cartridge"
			wellNum: 4
			distance: 41.20
			height: 2
			volume: 10

	// For category= well_to_shaker ignore the destinationCartridge
	 	And for category= shaker_to_well ignore the sourceCartridge
	*/

	uniqueCartridge := UniqueCartridge{
		LabwareID:     labwareID,
		CartridgeType: cartridgeType,
	}

	// Get Source Position
	switch category {
	case "well_to_well", "well_to_shaker":
		uniqueCartridge.WellNum = source
		// What if we ignore the below ok ?
		if sourceCartridge, ok = cartridges[uniqueCartridge]; ok {
			err = fmt.Errorf("sourceCartridge doesn't exist")
			fmt.Println("Error: ", err)
			return "", err
		}
		sourcePosition, ok = sourceCartridge["distance"]
	case "shaker_to_well":
		sourcePosition, ok = consDistance["shaker_tube"]
	default:
		err = fmt.Errorf("category is invalid for aspire_dispense opeartion")
		fmt.Println("Error: ", err)
		return "", err
	}
	if !ok {
		err = fmt.Errorf("source doesn't exist for aspiring")
		fmt.Println("Error: ", err)
		return "", err
	}

	// NOTE : below position is added to sourcePosition as well as destinationPosition
	if position, ok = consDistance[cartridgeType+"_cartridge_start"]; ok {
		sourcePosition += position
		fmt.Println("sourcePosition: ", sourcePosition)
	} else {
		err = fmt.Errorf(cartridgeType + "_cartridge_start doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	// Get Destination Position
	switch category {
	case "well_to_well", "shaker_to_well":
		uniqueCartridge.WellNum = destination
		// What if we ignore the below ok ?
		if destinationCartridge, ok = cartridges[uniqueCartridge]; ok {
			err = fmt.Errorf("destinationCartridge doesn't exist")
			fmt.Println("Error: ", err)
			return "", err
		}
		destinationPosition, ok = destinationCartridge["distance"]
	case "well_to_shaker":
		destinationPosition, ok = consDistance["shaker_tube"]
	}
	if !ok {
		err = fmt.Errorf("destination doesn't exist for dispensing")
		fmt.Println("Error: ", err)
		return "", err
	}
	destinationPosition += position
	fmt.Println("destinationPosition: ", destinationPosition)

	/*
		Id: 303
		Name : pickup_piercing_tip_up
		Distance: 59
		Description: syringe module goes this much up at fast after picking up the piercing tip

		Just like above invent this
		Id: 311
		Name : pickup_extraction_tip_up
		Distance: 118
		Descrption: syringe module goes this much up at fast after picking up the extraction tip

		And this
		Id: 312
		Name : pickup_pcr_tip_up
		Distance: 118
		Descrption: syringe module goes this much up at fast after picking up the pcr tip

	*/

	//
	// ALGORITHM's 9th step is implemented below
	//

	// TODO: Check if its only LH /RH or both !!!
	// Go UP with extraction/pcr tip
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	if position, ok = consDistance["pickup_"+cartridgeType+"_tip_up"]; ok {
		distToTravel = positions[deckAndMotor] - position
	} else {
		err = fmt.Errorf("pickup_" + cartridgeType + "_tip_up doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distToTravel > 0:
		direction = 1
	case distToTravel < 0:
		distToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipExtractionTipUp
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

skipExtractionTipUp:

	//
	// ALGORITHM's 10th step is implemented below
	//

	deckAndMotor.Number = K5_Deck

	distToTravel = positions[deckAndMotor] - sourcePosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distToTravel > 0:
		direction = 1
	case distToTravel < 0:
		distToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipDeckToSourcePosition
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	//
	// ALGORITHM's 11th step is implemented below
	//
	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Deck to Aspire Source. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

skipDeckToSourcePosition:

	//
	// ALGORITHM's 12th step is implemented below
	//

	// We know the concrete direction here onwards
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	// TODO Add deck_base in consDistance
	if position = consDistance["deck_base"]; ok {
		distToTravel = position - positions[deckAndMotor]
	} else {
		err = fmt.Errorf("deck_base doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	// IMPLEMENT the for loop for aspiring mix and stuff

	return
}
