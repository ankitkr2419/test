package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
	"time"
)

/****ALGORITHM******
// TODO: check for height and volumes constraint at insertion process itself
variables: category, cartridgeType string,
			cartridgeID, type, source_well, destination_well, aspire_cycles, dispense_cycles int64,
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
  13. setup the syringe module motor with aspire height
  14. pickup and drop that asp_mix_vol for number of aspire_cycles
   these cycles should be fast
  15. pickup asp_vol slow
  16. move syringe module up slow till just above base
  17. take airVolume in
  18. Move slowly to destinationPosition by calculating the difference of Positions
  19. move syringe module down at fast till base
  20. setup the syringe module motor with dispense height
  21. pickup and drop that dis_mix_vol for number of dis_cycles
  22. dispense by that dis_vol
  23. move syringe module up
  24. dispense blow out air

********/

func (d *Compact32Deck) AspireDispenseBeta(ad db.AspireDispense, cartridgeID int64, tipType string) (response string, err error) {

	var sourceCartridge, destinationCartridge map[string]float64
	var sourcePosition, destinationPosition, distToTravel, position, tipHeight float64
	var ok bool
	var direction, pulses uint16
	var deckAndMotor DeckNumber
	deckAndMotor.Deck = d.name

	//-----------------
	// Get Tip Height -
	//-----------------
	var tipHeightInter interface{}
	if tipHeightInter, ok = tipstubes[tipType]["height"]; !ok {
		err = fmt.Errorf(tipType + " tip doesn't exist for tipstubes")
		fmt.Println("Error: ", err)
		return "", err
	}

	if tipHeight, ok = tipHeightInter.(float64); !ok {
		err = fmt.Errorf(tipType + " tip has unknown type!")
		fmt.Println("Error: ", err)
		return "", err
	}

	/*** GET THE CARTRIDGES
	E.g :
	********** for well_to_well category only ***********
	Suppose
		cartridgeID = 1 && cartridgeType = "extraction" && source = 2 && destination= 4
	Then
		sourceCartridge =
		- id: 2
			cartridgeID: 1
			type: "extraction"
			description: "Extraction Cartridge"
			wellNum: 2
			distance: 24.5
			height: 2
			volume: 10

	And
		destinationCartridge =
		- id: 4
			cartridgeID: 1
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
		CartridgeID:   cartridgeID,
		CartridgeType: ad.CartridgeType,
	}

	//*************************************************
	// ALGORITHM's 1 to 8 steps are implemented below *
	//*************************************************
	//  1. Check the category of operation
	//  2. if category is well_to_well then goto 3 else if category is shaker_to_well then goto 5  else 7
	//  3. store the source_well position into sourcePosition variable
	//  4. store the destination_well position into destinationPosition variable; goto 9
	//  5. store the shaker position into sourcePosition variable
	//  6. store the destination_well position into destinationPosition variable; goto 9
	//  7. store the source_well position into sourcePosition variable
	//  8. store the shaker position into destinationPosition variable
	//

	// NOTE : below position is added to sourcePosition/destinationPosition
	// But only when they are wells
	if position, ok = consDistance[string(ad.CartridgeType)+"_start"]; !ok {
		err = fmt.Errorf(string(ad.CartridgeType) + "_start doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	//----------------------
	// Get Source Position -
	//----------------------
	switch ad.Category {
	case "well_to_well", "well_to_shaker", "well_to_deck":
		uniqueCartridge.WellNum = ad.SourcePosition
		if sourceCartridge, ok = cartridges[uniqueCartridge]; !ok {
			err = fmt.Errorf("sourceCartridge doesn't exist")
			fmt.Println("Error: ", err)
			return "", err
		}
		sourcePosition, ok = sourceCartridge["distance"]
		sourcePosition += position
		fmt.Println("sourcePosition: ", sourcePosition)
	case "shaker_to_well", "shaker_to_deck":
		sourcePosition, ok = consDistance["shaker_tube"]
	case "deck_to_well", "deck_to_deck", "deck_to_shaker":
		// TODO: Check source Positions
		fmt.Println("This is the position---> ", "pos_"+fmt.Sprintf("%d", ad.SourcePosition))
		sourcePosition, ok = consDistance["pos_"+fmt.Sprintf("%d", ad.SourcePosition)]
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

	//---------------------------
	// Get Destination Position -
	//---------------------------
	switch ad.Category {
	case "well_to_well", "shaker_to_well", "deck_to_well":
		uniqueCartridge.WellNum = ad.DestinationPosition
		if destinationCartridge, ok = cartridges[uniqueCartridge]; !ok {
			err = fmt.Errorf("destinationCartridge doesn't exist")
			fmt.Println("Error: ", err)
			return "", err
		}
		fmt.Println(destinationCartridge)
		destinationPosition, ok = destinationCartridge["distance"]
		destinationPosition += position
		fmt.Println("destinationPosition: ", destinationPosition)
	case "well_to_shaker", "deck_to_shaker":
		destinationPosition, ok = consDistance["shaker_tube"]
	case "well_to_deck", "deck_to_deck", "shaker_to_deck":
		fmt.Println("This is the position---> ", "pos_"+fmt.Sprintf("%d", ad.DestinationPosition))
		destinationPosition, ok = consDistance["pos_"+fmt.Sprintf("%d", ad.DestinationPosition)]
		// default already handled in source Position
	}
	if !ok {
		err = fmt.Errorf("destination doesn't exist for dispensing")
		fmt.Println("Error: ", err)
		return "", err
	}

	//
	// 9. setup the motor of syringe module to go up atleast 30mm above deck
	//

	// TODO: Check if its only LH /RH or both !!!
	// Go UP with extraction/pcr tip
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	if position, ok = consDistance["pickup_tip_up"]; !ok {
		err = fmt.Errorf("pickup_tip_up doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Don't forget to add tipHeight for every tip we have currently attached
	distToTravel = positions[deckAndMotor] + tipHeight - position

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distToTravel > 0.1:
		direction = 1
	case distToTravel < -0.1:
		distToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipTipUp
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

skipTipUp:

	//
	// 10. calculate the current position difference for deck;
	//      if its positive then direction is 1(towards sensor) else 0(oppose sensor)
	//

	deckAndMotor.Number = K5_Deck

	distToTravel = positions[deckAndMotor] - sourcePosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distToTravel > 0.1:
		direction = 1
	case distToTravel < -0.1:
		distToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipDeckToSourcePosition
	}
	fmt.Println(distToTravel)

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	//*************************
	// REACHING ASPIRE SOURCE *
	//*************************

	//
	// 11. move deck to match the sourcePosition with help of difference calculated
	//
	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Deck to Aspire Source. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

skipDeckToSourcePosition:

	//***********
	// ASPIRING *
	//***********

	//
	//   12. move syringe module down at fast till base
	//

	// We know the concrete direction here onwards till Deck Movement
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	if position = consDistance["deck_base"]; ok {
		distToTravel = position - (positions[deckAndMotor] + tipHeight)
	} else {
		err = fmt.Errorf("deck_base doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip to deck base. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	//
	//   13. setup the syringe module motor with aspire height
	//

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * ad.AspireHeight))

	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	//
	//   14. pickup and drop that asp_mix_vol for number of aspire_cycles
	//       these cycles should be fast
	//

	deckAndMotor.Number = K10_Syringe_LHRH
	// for volume :-> 25 pulses = 1 microLitres
	// NOTE: Store volumes in microLitres only
	oneMicroLitrePulses := 25.0
	pulses = uint16(math.Round(oneMicroLitrePulses * ad.AspireMixingVolume))

	if ad.AspireMixingVolume == 0 {
		goto skipAspireCycles
	}
	for cycleNumber := int64(1); cycleNumber <= ad.AspireNoOfCycles; cycleNumber++ {
		// Aspire
		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
		if err != nil {
			return
		}

		time.Sleep(100 * time.Millisecond)

		// Dispense
		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
		if err != nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
skipAspireCycles:

	//
	//   15. pickup asp_vol slow
	//

	pulses = uint16(math.Round(oneMicroLitrePulses * ad.AspireVolume))

	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	//
	//   16. move syringe module up slow till just above base
	//
	deckAndMotor.Number = K9_Syringe_Module_LHRH

	if position, ok = consDistance["pickup_tip_up"]; !ok {
		err = fmt.Errorf("pickup_tip_up doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Don't forget to add tipHeight for every tip we have currently attached
	distToTravel = positions[deckAndMotor] + tipHeight - position
	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)

	deckAndMotor.Number = K10_Syringe_LHRH

	//
	//  17. take airVolume in
	//
	pulses = uint16(math.Round(oneMicroLitrePulses * ad.AspireAirVolume))
	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	//********************************
	// REACHING DISPENSE DESTINATION *
	//********************************

	//
	// 18. Move slowly to destinationPosition by calculating the difference of Positions
	//

	deckAndMotor.Number = K5_Deck

	distToTravel = positions[deckAndMotor] - destinationPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distToTravel > 0.1:
		direction = 1
	case distToTravel < -0.1:
		distToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipDeckToDestinationPosition
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Deck to Dispense Destination. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

skipDeckToDestinationPosition:

	//*************
	// DISPENCING *
	//*************
	//
	//   19. move syringe module down at fast till base
	//

	// We know the concrete direction here onwards
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	if position = consDistance["deck_base"]; ok {
		distToTravel = position - (positions[deckAndMotor] + tipHeight)
	} else {
		err = fmt.Errorf("deck_base doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip to deck base. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	//
	//   20. setup the syringe module motor with dispense height
	//
	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * ad.DispenseHeight))

	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}
	time.Sleep(100 * time.Millisecond)

	//
	//   21. pickup and drop that dis_mix_vol for number of dis_cycles
	//
	deckAndMotor.Number = K10_Syringe_LHRH
	pulses = uint16(math.Round(oneMicroLitrePulses * ad.DispenseMixingVolume))

	if ad.DispenseMixingVolume == 0 {
		goto skipDispenseCycles
	}
	for cycleNumber := int64(1); cycleNumber <= ad.DispenseNoOfCycles; cycleNumber++ {
		// Dispense
		// CHECK : should these operations be fast ?
		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
		if err != nil {
			return
		}

		time.Sleep(100 * time.Millisecond)

		// Aspire
		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
		if err != nil {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
skipDispenseCycles:

	//
	//    22. dispense by that dis_vol
	//
	pulses = uint16(math.Round(oneMicroLitrePulses * ad.DispenseVolume))
	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	//
	//   23. move syringe module up
	//
	deckAndMotor.Number = K9_Syringe_Module_LHRH

	if position, ok = consDistance["pickup_tip_up"]; !ok {
		err = fmt.Errorf("pickup_tip_up doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Don't forget to add tipHeight for every tip we have currently attached
	distToTravel = positions[deckAndMotor] + tipHeight - position
	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
	if err != nil {
		return
	}

	//
	//   24. dispense blow out air
	//
	deckAndMotor.Number = K10_Syringe_LHRH

	pulses = uint16(math.Round(oneMicroLitrePulses * ad.DispenseBlowVolume))
	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	return "ASPIRE and DISPENSE was successful", nil
}
