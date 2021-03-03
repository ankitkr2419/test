package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
	"time"
)

/****ALGORITHM******
1. Call Tip Pick Up at position 3
2. Move deck to extraction cartridge start
3. Pierce well after well
4. Call Tip Discard

********/

func (d *Compact32Deck) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {

	var deckAndMotor DeckNumber
	var position, extractionCartridgeStart, piercingHeight float64
	var ok bool
	deckAndMotor.Deck = d.name
	deckAndMotor.Number = K5_Deck
	uniqueCartridge := UniqueCartridge{
		CartridgeID:   cartridgeID,
		CartridgeType: pi.CartridgeType,
	}
	// 1. Call Tip Pick Up at position 3

	// 3rd position is where by default piercing tip is present
	response, err = d.TipPickup(3)
	if err != nil {
		return
	}

	// 2. Move Deck to Extraction Cartridge start

	if extractionCartridgeStart, ok = consDistance["extraction_cartridge_start"]; !ok {
		err = fmt.Errorf("extraction_cartridge_start doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	if piercingHeight, ok = consDistance["piercing_height"]; !ok {
		err = fmt.Errorf("piercing_height doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	// distToTravel = positions[deckAndMotor] - position

	// switch {
	// // distToTravel > 0 means go towards the Sensor or FWD
	// case distToTravel > 0.1:
	// 	direction = 1
	// case distToTravel < -0.1:
	// 	distToTravel *= -1
	// 	direction = 0
	// default:
	// 	// Skip the setUpMotor Step
	// 	goto skipMoveToExtractionCartridgeStart
	// }

	// pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	// response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", fmt.Errorf("There was issue moving Syringe Module with tip. Error: %v", err)
	// }

	// skipMoveToExtractionCartridgeStart:

	// fmt.Println("Completed Move Deck to Extraction Cartridge start")

	// 3. Pierce well after well

	//*************************
	// Pierce Well after Well *
	//*************************

	for _, c := range pi.CartridgeWells {

		// Pierce Well
		uniqueCartridge.WellNum = c
		if position, ok = cartridges[uniqueCartridge]["distance"]; !ok {
			err = fmt.Errorf("distance doesn't exist for well number %d", c)
			fmt.Println("Error: ", err)
			return "", err
		}

		distToTravel = positions[deckAndMotor] - (position + extractionCartridgeStart)

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

		deckAndMotor.Number = K5_Deck
		fmt.Println("Completed Move Deck to reach the wellNum ", c)

		pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

		response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue moving Deck to Cartridge WellNum %d. Error: %v", c, err)
		}

	skipDeckMovement:
		// Pierce the Well
		// Move Syringe Module Down Fwd 31mm
		response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_height"])), uint16(motors[1].Ramp), uint16(Fwd), uint16(motors[1].Number), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Down Fwd 31mm wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

		// Move Syringe Module Up Rev Fwd 31mm
		response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_height"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
		if err != nil {
			return
		}

		fmt.Println("Completed Move Syringe Module Up Rev 31mm wellNum ", c.WellNum)
		time.Sleep(1 * time.Second)

	}

	return "RUN SUCCESS", nil
}

//**************
// Tip Discard *
//**************

// Move Deck Fwd 150.54mm

// distToTravel = cd["discard_big_hole"] - at_X
// response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
// if err != nil {
// 	return
// }

// at_X = cd["discard_big_hole"]

// fmt.Println("Completed Move Deck Fwd 150.54mm")
// time.Sleep(1 * time.Second)

// // Move Syringe Module Down Fwd 83.9mm
// response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["piercing_tip_discard_height"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
// if err != nil {
// 	return
// }

// fmt.Println("Completed Move Syringe Module Down Fwd 83.9mm")
// time.Sleep(1 * time.Second)

// // Move Deck Rev 6.8mm to cut Hold of that tip
// distToTravel = at_X - cd["discard_small_hole"]
// response, err = m.setUpMotor(uint16(motors[0].Fast), uint16(math.Round(float64(motors[0].Steps)*distToTravel)), uint16(motors[0].Ramp), uint16(Fwd), uint16(motors[0].Number), On, Off)
// if err != nil {
// 	return
// }

// at_X = cd["discard_small_hole"]

// fmt.Println("Completed Move Deck Rev 6.8mm to cut Hold of that tip")
// time.Sleep(1 * time.Second)

// // Move Syringe Module Up Slow Rev 7.5mm
// response, err = m.setUpMotor(uint16(motors[1].Slow), uint16(math.Round(float64(motors[1].Steps)*cd["discard_tip_slow_up"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
// if err != nil {
// 	return
// }

// fmt.Println("Completed Move Syringe Module Up Slow Rev 7.5mm")
// time.Sleep(1 * time.Second)

// // Move Syringe Module Up Rev 137.5mm
// response, err = m.setUpMotor(uint16(motors[1].Fast), uint16(math.Round(float64(motors[1].Steps)*cd["syringe_end_max"])), uint16(motors[1].Ramp), uint16(Rev), uint16(motors[1].Number), On, Off)
// if err != nil {
// 	return
// }

// fmt.Println("Completed Move Syringe Module Up Rev 137.5mm")
// time.Sleep(1 * time.Second)

// fmt.Println("Home the machine now.")
