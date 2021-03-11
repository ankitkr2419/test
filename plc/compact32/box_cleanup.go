package compact32

import (
	"fmt"
	"math"
)


func (d *Compact32Deck) DiscardBoxCleanup() (response string, err error) {

	var position, distToTravel float64
	var ok bool
	var pulses uint16
	deckAndMotor := DeckNumber{Deck: d.name, Number: K5_Deck}

	aborted[d.name] = false
	if runInProgress[d.name] {
		err = fmt.Errorf("previous run is already in progress... wait or abort it")
		return
	}

	runInProgress[d.name] = true
	defer d.ResetRunInProgress()


	fmt.Println("Deck is moving to discard_box_open_position")


	if position, ok = consDistance["discard_box_open_position"]; !ok {
		err = fmt.Errorf("discard_box_open_position doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distToTravel =  position - positions[deckAndMotor]

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	// We know concrete direction here, its REV
	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], REV, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was an issue moving deck REV to discard_box_open_position. Error: %v", err)
	}

	fmt.Println("Moved Deck to Cleanup Discard Box Successfully")

	return "DISCARD BOX CLEANUP SUCCESS", nil
}


func (d *Compact32Deck) RestoreDeck() (response string, err error) {

	var position, distToTravel float64
	var ok bool
	var pulses uint16
	deckAndMotor := DeckNumber{Deck: d.name, Number: K5_Deck}

	aborted[d.name] = false
	if runInProgress[d.name] {
		err = fmt.Errorf("previous run is already in progress... wait or abort it")
		return
	}

	runInProgress[d.name] = true
	defer d.ResetRunInProgress()

	fmt.Println("Deck is moving to deck_start")

	if position, ok = consDistance["deck_start"]; !ok {
		err = fmt.Errorf("deck_start doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distToTravel =  positions[deckAndMotor] - position

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	// We know concrete direction here, its FWD
	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], FWD, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was an issue moving deck FWD to deck_start. Error: %v", err)
	}

	fmt.Println("Moved Deck back to homing position")

	return "DECK RESTORED SUCCESS", nil
}