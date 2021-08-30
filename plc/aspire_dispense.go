package plc

import (
	"fmt"
	"math"
	"mylab/cpagent/config"
	"mylab/cpagent/db"

	logger "github.com/sirupsen/logrus"
)

const (
	extraDispense = 5 // microlitres
)

/****ALGORITHM******
// TODO: check for height and volumes constraint at insertion process itself
variables: category, cartridgeType string,
			cartridgeID, type, source_well, destination_well, aspire_cycles, dispense_cycles int64,
			asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol float64

  1. Check the category of operation
  2. if category is well_to_well then goto 3 else if category is shaker_to_well then goto 5  else 7
  3. store the source_well position into sourcePosition variable
  4. store the destination_well position into destinationPosition variable; goto 9
  5. store the shaker position into sourcePosition variable
  6. store the destination_well position into destinationPosition variable; goto 9
  7. store the source_well position into sourcePosition variable
  8. store the shaker position into destinationPosition variable
  9. setup the motor of syringe module to go up atleast 30mm above deck(not required)
  10. calculate the current position difference for deck;
   		if its positive then direction is 1(towards sensor) else 0(oppose sensor)
  11. move deck to match the sourcePosition with help of difference calculated
  12. move syringe module down at fast till almost base
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
  22. Dispense completely + extraDispense

********/

func (d *Compact32Deck) AspireDispense(ad db.AspireDispense, cartridgeID int64) (response string, err error) {

	var sourceCartridge, destinationCartridge map[string]float64
	var sourcePosition, destinationPosition, distanceToTravel, position, deckBase, aboveDeck, pickUpTip float64
	var ok, dispenseComplete bool
	var direction, pulses uint16
	var deckAndMotor DeckNumber
	deckAndMotor.Deck = d.name

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
		err = fmt.Errorf(string(ad.CartridgeType) + "_start doesn't exist for consumable distances")
		logger.Errorln(err)
		return "", err
	}

	//----------------------
	// Get Source Position -
	//----------------------
	switch ad.Category {
	case db.WW, db.WS, db.WD:
		uniqueCartridge.WellNum = ad.SourcePosition
		if sourceCartridge, ok = cartridges[uniqueCartridge]; !ok {
			err = fmt.Errorf("sourceCartridge doesn't exist")
			logger.Errorln(err)
			return "", err
		}
		sourcePosition, ok = sourceCartridge["distance"]
		sourcePosition += position
		logger.Infoln("sourcePosition: ", sourcePosition)
	case db.SW, db.SD:
		sourcePosition, ok = consDistance["shaker_tube"]
	case db.DW, db.DD, db.DS:
		// TODO: Check source Positions
		logger.Infoln("This is the position---> ", "pos_"+fmt.Sprintf("%d", ad.SourcePosition))
		sourcePosition, ok = consDistance["pos_"+fmt.Sprintf("%d", ad.SourcePosition)]
	default:
		err = fmt.Errorf("category is invalid for aspire_dispense opeartion")
		logger.Errorln(err)
		return "", err
	}
	if !ok {
		err = fmt.Errorf("source doesn't exist for aspiring")
		logger.Errorln(err)
		return "", err
	}

	//---------------------------
	// Get Destination Position -
	//---------------------------
	switch ad.Category {
	case db.WW, db.SW, db.DW:
		uniqueCartridge.WellNum = ad.DestinationPosition
		if destinationCartridge, ok = cartridges[uniqueCartridge]; !ok {
			err = fmt.Errorf("destinationCartridge doesn't exist")
			logger.Errorln(err)
			return "", err
		}
		logger.Infoln("destinationCartridge :", destinationCartridge)
		destinationPosition, ok = destinationCartridge["distance"]
		destinationPosition += position
		logger.Infoln("destinationPosition: ", destinationPosition)
	case db.WS, db.DS:
		destinationPosition, ok = consDistance["shaker_tube"]
	case db.WD, db.DD, db.SD:
		logger.Infoln("This is the position---> ", "pos_"+fmt.Sprintf("%d", ad.DestinationPosition))
		destinationPosition, ok = consDistance["pos_"+fmt.Sprintf("%d", ad.DestinationPosition)]
		// Completely dispense in PCR Tube/Extraction Tube
		if ad.DestinationPosition == 11 || ad.DestinationPosition == 9 {
			dispenseComplete = true
		}
		// default already handled in source Position
	}
	if !ok {
		err = fmt.Errorf("destination doesn't exist for dispensing")
		logger.Errorln(err)
		return "", err
	}

	//
	// 9. setup the motor of syringe module to go up atleast 30mm above deck
	//
	// this step is not required because if the dispencing and aspiring wells are same
	// the syringe does not need to come above the deck.
	// and if it is aspiring from some other well then it is handled in the below code from step 10.

	//
	// 10. calculate the current position difference for deck;
	//      if its positive then direction is 1(towards sensor) else 0(oppose sensor)
	//

	deckAndMotor.Number = K5_Deck

	distanceToTravel = Positions[deckAndMotor] - sourcePosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	//*************************
	// REACHING ASPIRE SOURCE *
	//*************************

	//
	// 11. move deck to match the sourcePosition with help of difference calculated
	//
	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Deck to Aspire Source. Error: %v", err)
	}

	//***********
	// ASPIRING *
	//***********

	//
	//   12. move syringe module down at fast till almost base
	//

	// We know the concrete direction here onwards till Deck Movement
	deckAndMotor.Number = K9_Syringe_Module_LHRH

	if deckBase, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}

	if aboveDeck, ok = consDistance["above_deck"]; !ok {
		err = fmt.Errorf("above_deck doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}

	if val, ok := syringeModuleState.Load(d.name); !ok {
		err = fmt.Errorf("failed to load syringe module for deck: %s", d.name)
		return "", err
	} else {
		if val == InDeck {
			goto skipToAspireInside
		}
	}

	distanceToTravel = (Positions[deckAndMotor] + tipHeight[d.name] + aboveDeck) - deckBase

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip to deck base. Error: %v", err)
	}

skipToAspireInside:
	//
	//   13. setup the syringe module motor with aspire height
	//
	distanceToTravel = (Positions[deckAndMotor] + tipHeight[d.name]) - (ad.AspireHeight + deckBase)

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["slow"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}

	//
	//   14. pickup and drop that asp_mix_vol for number of aspire_cycles
	//       these cycles should be fast
	//

	deckAndMotor.Number = K10_Syringe_LHRH
	// for volume :-> 25 pulses = 1 microLitres
	// NOTE: Store volumes in microLitres only
	oneMicroLitrePulses := float64(config.GetMicroLitrePulses())
	pulses = uint16(math.Round(oneMicroLitrePulses * ad.AspireMixingVolume))

	if ad.AspireNoOfCycles == 0 {
		goto skipAspireCycles
	}
	for cycleNumber := int64(1); cycleNumber <= ad.AspireNoOfCycles; cycleNumber++ {
		// Aspire
		response, err = d.setupMotor(Motors[deckAndMotor]["slow"], pulses, Motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
		if err != nil {
			return
		}

		// Dispense
		// TODO: Call a separate function for this kind of setup, as it only DISPENCING
		response, err = d.setupMotor(Motors[deckAndMotor]["slow"], pulses, Motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
		if err != nil {
			return
		}
	}
skipAspireCycles:

	//
	//   15. pickup asp_vol slow
	//

	pulses = uint16(math.Round(oneMicroLitrePulses * ad.AspireVolume))

	response, err = d.setupMotor(Motors[deckAndMotor]["slow"], pulses, Motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
	if err != nil {
		return
	}

	//
	//   16. 1 move syringe module up slow till just above base
	//
	deckAndMotor.Number = K9_Syringe_Module_LHRH

	distanceToTravel = (Positions[deckAndMotor] + tipHeight[d.name] + aboveDeck) - deckBase

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["slow"], pulses, Motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)

	//
	//   16. 2 move syringe module above pickup_tip_up(27 mm) from the deck
	//
	if pickUpTip, ok = consDistance["pickup_tip_up"]; !ok {
		err = fmt.Errorf("pickup_tip_up doesn't exist for consumable distances")
		logger.Errorln(err)
		return "", err
	}
	// Don't forget to add tipHeight for every tip we have currently attached
	distanceToTravel = pickUpTip

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)
	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)

	deckAndMotor.Number = K10_Syringe_LHRH

	//
	//  17. take airVolume in
	//
	pulses = uint16(math.Round(oneMicroLitrePulses * ad.AspireAirVolume))
	response, err = d.setupMotor(Motors[deckAndMotor]["slow"], pulses, Motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
	if err != nil {
		return
	}

	//********************************
	// REACHING DISPENSE DESTINATION *
	//********************************

	//
	// 18. Move slowly to destinationPosition by calculating the difference of Positions
	//

	deckAndMotor.Number = K5_Deck

	distanceToTravel = Positions[deckAndMotor] - destinationPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Deck to Dispense Destination. Error: %v", err)
	}

	//*************
	// DISPENCING *
	//*************
	//
	//   19. move syringe module down at fast till base
	//

	// We know the concrete direction here onwards
	deckAndMotor.Number = K9_Syringe_Module_LHRH
	distanceToTravel = deckBase - (Positions[deckAndMotor] + tipHeight[d.name] + aboveDeck)

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip to deck base. Error: %v", err)
	}

	//
	//   20. setup the syringe module motor with dispense height
	//
	dispensePos := deckBase + ad.DispenseHeight
	distanceToTravel = Positions[deckAndMotor] + tipHeight[d.name] - dispensePos
	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))
	logger.Warnln("going to dispensing height", ad.DispenseHeight)

	response, err = d.setupMotor(Motors[deckAndMotor]["slow"], pulses, Motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}

	//
	//   21. pickup and drop that dis_mix_vol for number of dis_cycles
	//
	deckAndMotor.Number = K10_Syringe_LHRH
	pulses = uint16(math.Round(oneMicroLitrePulses * ad.DispenseMixingVolume))

	if ad.DispenseNoOfCycles == 0 {
		goto skipDispenseCycles
	}
	for cycleNumber := int64(1); cycleNumber <= ad.DispenseNoOfCycles; cycleNumber++ {
		// Dispense
		// CHECK : should these operations be fast ?
		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
		if err != nil {
			return
		}

		// Aspire
		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
		if err != nil {
			return
		}
	}
skipDispenseCycles:

	//
	// 22. Dispense completely + extraDispense
	//

	if dispenseComplete {
		logger.Infoln("Dispense Complete is to start")

		response, err = d.setupMotor(homingSlowSpeed, initialSensorCutSyringePulses, Motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
		if err != nil {
			return
		}
	} else {
		pulses = uint16(math.Round(oneMicroLitrePulses * (ad.AspireAirVolume + ad.AspireVolume + extraDispense)))

		logger.Infoln("Syringe is moving down and dispensingalong with extraDispense")
		// TODO:  Note down Bio team's required speed
		response, err = d.setupMotor(homingSlowSpeed, pulses, Motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
		if err != nil {
			return
		}
	}

	logger.Info("Aspire Dispense success")

	return "ASPIRE and DISPENSE was successful", nil
}

func (d *Compact32Deck) SyringeRestPosition() (response string, err error) {
	var distanceToTravel, position float64
	var direction uint16
	var ok bool

	syringeModuleDeckAndMotor := DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}

	logger.Infoln("syringe module is inDeck, moving it to rest position")

	if position, ok = consDistance["resting_position"]; !ok {
		err = fmt.Errorf("resting_position doesn't exist for consumable distances")
		logger.Errorln(err)
		return "", err
	}
	distanceToTravel = Positions[syringeModuleDeckAndMotor] + tipHeight[d.name] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses := uint16(math.Round(float64(Motors[syringeModuleDeckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[syringeModuleDeckAndMotor]["fast"], pulses, Motors[syringeModuleDeckAndMotor]["ramp"], direction, syringeModuleDeckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	}

	return
}

func (d *Compact32Deck) setSyringeState() (err error) {
	var deckBase, indeckSafe float64
	var ok bool

	if deckBase, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return
	}

	if indeckSafe, ok = consDistance["indeck_safe"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return
	}

	deckAndMotor := DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}
	// In case the operation was aborted after syringe module went inside the deck!
	// If syringe + tipHeight is below certain level then that is inDeck
	if (Positions[deckAndMotor] + tipHeight[d.name]) >= (deckBase - indeckSafe) {
		syringeModuleState.Store(d.name, InDeck)
	} else {
		syringeModuleState.Store(d.name, OutDeck)
	}
	return
}
