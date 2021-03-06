package plc

import (
	"fmt"
	"math"
	"mylab/cpagent/db"

	logger "github.com/sirupsen/logrus"
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
	var position, cartridgeStart, distanceToTravel, deckBase float64
	var ok bool
	var direction, pulses uint16
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
		logger.Errorln(err)
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
		logger.Errorln(err)
		return "", err
	}

	for i, wellNumber := range pi.CartridgeWells {
		//
		// 2.1 Move deck to the well position
		//
		deckAndMotor.Number = K5_Deck
		uniqueCartridge.WellNum = int64(wellNumber)
		if position, ok = cartridges[uniqueCartridge]["distance"]; !ok {
			err = fmt.Errorf("distance doesn't exist for well number %d", wellNumber)
			logger.Errorln(err)
			return "", err
		}

		// here disToTravel moves our deck to well position
		// position + cartridgeStart is the distance of first well on deck for wellNumber = 1
		distanceToTravel = Positions[deckAndMotor] - (position + cartridgeStart)

		modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

		pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

		_, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving Deck to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		logger.Infoln("Completed Move Deck to reach the wellNum ", wellNumber)

		// 2.2 Pierce and come back up

		// WE know concrete direction here, its DOWN
		deckAndMotor.Number = K9_Syringe_Module_LHRH

		// Go Down by well height + Base
		distanceToTravel = (deckBase + float64(pi.Heights[i])) - (Positions[deckAndMotor] + tipHeight[d.name])

		pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving Syringe Module DOWN to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		logger.Infoln("Pierced WellNumber: ", wellNumber)

		// WE know concrete direction here, its UP
		// Come Up by well height + Base
		response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving Syringe Module UP to Cartridge WellNum %d. Error: %v", wellNumber, err)
		}

		logger.Infoln("Got Up from WellNumber: ", wellNumber)
		// 2.3 Repeat step 2.1 and  2.2 till another well exists
	}

	return "Successfully completed piercing operation", nil
}
