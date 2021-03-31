package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
)

/****ALGORITHM******
1. Check Operation type
2. If its pickup then call Tip PickUp with position to pick up from
3. Else call Tip Discard
********/

func (d *Compact32Deck) TipOperation(to db.TipOperation) (response string, err error) {

	switch to.Type {
	case db.PickupTip:
		response, err = d.TipPickup(to.Position)
	case db.DiscardTip:
		response, err = d.TipDiscard()
	}
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue doing Tip Operation. Error: %v", err)
	}

	return "Tip Operation was successfull", nil

}

/****ALGORITHM******
1. Move Deck to the tip's position
2. Move Syringe Module down fast to tip's base
3. Move Syringe Module down really slow to tip's inside
4. Move Syringe Module up with tip to Resting position.

********/

func (d *Compact32Deck) TipPickup(pos int64) (response string, err error) {

	// **************
	// Tip PickUp	*
	//***************
	var deckAndMotor DeckNumber
	var position, distanceToTravel, restingPos float64
	var direction, pulses uint16
	var tipFast, tipSlow, restingPositionString string
	var ok bool
	deckAndMotor.Deck = d.name


	//
	// 1. Move Deck to the tip's position
	//

	deckAndMotor.Number = K5_Deck

	fmt.Println("Moving Deck to pos_" + fmt.Sprintf("%d", pos))
	if position, ok = consDistance["pos_"+fmt.Sprintf("%d", pos)]; !ok {
		err = fmt.Errorf("pos_" + fmt.Sprintf("%d", pos) + " doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	distanceToTravel = positions[deckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Deck to tip source. Error: %v", err)
	}

	//
	// 2. Move Syringe Module down fast to tip's base
	//

	deckAndMotor.Number = K9_Syringe_Module_LHRH

	// TODO: Handle this in non-harcoded way as tips add up in future
	// We can do this by being aware of the tip.
	// So, in future add a field like height of tip above the deck
	switch pos {
	// extraction tip
	case 1, 2:
	tipFast = "syringe_module_fast_down_1000_tip"
	tipSlow = "syringe_module_slow_down_1000_tip"
	// piercing tip
	case 3:
		tipFast = "syringe_module_fast_down_piercing_tip"
		tipSlow = "syringe_module_slow_down_piercing_tip"
	}
	fmt.Println("Moving Syringe to tip's base")
	if position, ok = consDistance[tipFast]; !ok {
		err = fmt.Errorf(tipFast + " doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Here tipFast will always be greater
	// than resting_position
	distanceToTravel = position - positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to tip's base. Error: %v", err)
	}

	//
	// 3. Move Syringe Module down slow to tip's inside
	//

	fmt.Println("Moving Syringe to tip's inside")
	if position, ok = consDistance[tipSlow]; !ok {
		err = fmt.Errorf(tipSlow + " doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Here tipSlow will always be greater
	// than tipFast
	distanceToTravel = position - positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	// Giving it real slow speed
	response, err = d.setupMotor(homingSlowSpeed, pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to tip's inside. Error: %v", err)
	}

	//
	// 4. Move Syringe Module up with tip to restingPositionString.
	//

	fmt.Println("Moving Syringe Module to ", restingPositionString)

	// Here resting_position will always be lesser
	// than whatever position earlier
	distanceToTravel = positions[deckAndMotor] - restingPos

	// We know Concrete Direction here, its UP

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to %v. Error: %v", restingPositionString, err)
	}

	return "Tip PickUp was successfull", nil

}

/****ALGORITHM******
1. Move Deck to the big hole's position
2. Move Syringe Module down fast to deck's base
3. Move Syringe Module down really slow till enough inside big hole
4. Move Deck to the small hole's position
5. Move Syringe Module up slow with tip till deck base, to drop off the tip.
6. Move Syringe Module up fast with tip to Resting position.
*/

// TODO: Currently only discarding at Discard box so avoiding at_pickup_passing condition
func (d *Compact32Deck) TipDiscard() (response string, err error) {

	/*
	 ************* *
	 * Tip Discard *
	 ***************
	 */

	var deckAndMotor DeckNumber
	var position, distanceToTravel, parkingPos float64
	var direction, pulses uint16
	var ok bool
	deckAndMotor.Deck = d.name

	if parkingPos, ok = consDistance["syringe_parking"]; !ok {
		err = fmt.Errorf("syringe_parking doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	//
	//  1. Move Deck to the big hole's position
	//

	deckAndMotor.Number = K5_Deck

	fmt.Println("Moving Deck to discard_big_hole")
	if position, ok = consDistance["discard_big_hole"]; !ok {
		err = fmt.Errorf("discard_big_hole doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	distanceToTravel = positions[deckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Deck to discard_big_hole source. Error: %v", err)
	}

	//
	// 2. Move Syringe Module down fast to deck base
	//

	deckAndMotor.Number = K9_Syringe_Module_LHRH

	fmt.Println("Moving Syringe to tip's base")
	if position, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Here deck_base will always be greater
	// than syringe_parking
	distanceToTravel = position - positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to deck's base. Error: %v", err)
	}

	//
	// 3. Move Syringe Module down really slow till enough inside big hole
	//

	fmt.Println("Moving Syringe to deck base inside")
	if position, ok = consDistance["syringe_module_slow_down_for_discard"]; !ok {
		err = fmt.Errorf("syringe_module_slow_down_for_discard doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Here syringe_module_slow_down_for_discard will always be greater
	// than tipFast
	distanceToTravel = position - positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	// Giving it real slow speed
	response, err = d.setupMotor(homingSlowSpeed, pulses, motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to discard_big_hole's inside. Error: %v", err)
	}

	//
	//  4. Move Deck to the small hole's position
	//

	deckAndMotor.Number = K5_Deck

	fmt.Println("Moving Deck to discard_small_hole")
	if position, ok = consDistance["discard_small_hole"]; !ok {
		err = fmt.Errorf("discard_small_hole doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	distanceToTravel = positions[deckAndMotor] - position

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	// We know concrete direction, here its towards sensor/ FWD
	response, err = d.setupMotor(homingSlowSpeed, pulses, motors[deckAndMotor]["ramp"], FWD, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Deck to discard_big_hole source. Error: %v", err)
	}

	//
	// 5. Move Syringe Module up slow with tip till deck base, to drop off the tip.
	//
	deckAndMotor.Number = K9_Syringe_Module_LHRH

	fmt.Println("Moving Syringe Module to drop off the tip")

	if position, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// Here deck_base will always be lesser
	// than syringe_module_slow_down_for_discard
	distanceToTravel = positions[deckAndMotor] - position

	// We know Concrete Direction here, its UP
	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	// Giving it real slow speed
	response, err = d.setupMotor(homingSlowSpeed, pulses, motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to deck base. Error: %v", err)
	}

	//
	// 6. Move Syringe Module up fast with tip to Resting position.
	//

	fmt.Println("Moving Syringe Module to Resting Position")

	// Here syringe_parking will always be lesser
	// than deck_base
	distanceToTravel = positions[deckAndMotor] - parkingPos

	// We know Concrete Direction here, its UP

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to resting position. Error: %v", err)
	}

	return "Tip Discard was successful", nil
}
