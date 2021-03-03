package compact32

import (
	"fmt"
	"math"
)

/****ALGORITHM******
1. Move Syringe Module to Resting position
2. Move Deck to the tip's position
3. Move Syringe Module down fast to tip's base
4. Move Syringe Module down slow to tip's inside
5. Move Syringe Module up with tip to Resting position.

********/

func (d *Compact32Deck) TipPickup(pos int64) (response string, err error) {

	// **************
	// Tip PickUp	*
	//***************
	var deckAndMotor DeckNumber
	var position, distToTravel, restingPos float64
	var direction, pulses uint16
	var ok bool
	deckAndMotor.Deck = d.name

	//
	// 1. Move Syringe Module to Resting position
	//

	deckAndMotor.Number = K9_Syringe_Module_LHRH

	fmt.Println("Moving Syringe Module to resting position")
	if restingPos, ok = consDistance["resting_position"]; !ok {
		err = fmt.Errorf("resting_position doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	distToTravel = positions[deckAndMotor] - restingPos

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distToTravel > 0.1:
		direction = 1
	case distToTravel < -0.1:
		distToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipSyringeModuleMove
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syringe Module to resting position. Error: %v", err)
	}

skipSyringeModuleMove:

	//
	// 2. Move Deck to the tip's position
	//

	deckAndMotor.Number = K5_Deck

	fmt.Println("Moving Deck to pos_" + fmt.Sprintf("%d", pos))
	if position, ok = consDistance["pos_"+fmt.Sprintf("%d", pos)]; !ok {
		err = fmt.Errorf("pos_" + fmt.Sprintf("%d", pos) + " doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	distToTravel = positions[deckAndMotor] - position

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distToTravel > 0.1:
		direction = 1
	case distToTravel < -0.1:
		distToTravel *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipDeckMove
	}

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Deck to tip source. Error: %v", err)
	}

skipDeckMove:

	//
	// 3. Move Syringe Module down fast to tip's base
	//

	deckAndMotor.Number = K9_Syringe_Module_LHRH

	fmt.Println("Moving Syringe to tip's base")
	if position, ok = consDistance["syringe_module_fast_down"]; !ok {
		err = fmt.Errorf("syringe_module_fast_down doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Here syringe_module_fast_down will awlays be greater
	// than resting_position
	distToTravel = position - positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to tip's base. Error: %v", err)
	}

	//
	// 4. Move Syringe Module down slow to tip's inside
	//

	fmt.Println("Moving Syringe to tip's inside")
	if position, ok = consDistance["syringe_module_slow_down"]; !ok {
		err = fmt.Errorf("syringe_module_slow_down doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Here syringe_module_slow_down will awlays be greater
	// than syringe_module_fast_down
	distToTravel = position - positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["slow"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to tip's inside. Error: %v", err)
	}

	//
	// 5. Move Syringe Module up with tip to Resting position.
	//

	fmt.Println("Moving Syringe Module to Resting Position")

	// Here resting_position will awlays be lesser
	// than whatever position earlier
	distToTravel = positions[deckAndMotor] - restingPos

	// We know Concrete Direction here, its UP

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to resting position. Error: %v", err)
	}

	return "Tip PickUp was successfull", nil

}
