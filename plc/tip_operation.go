package plc

import (
	"fmt"
	"math"
	"mylab/cpagent/config"
	"mylab/cpagent/db"

	logger "github.com/sirupsen/logrus"
)

const (
	airReserve = 300 // micro litres
)

/****ALGORITHM******
1. Check Operation type
2. If its pickup then call Tip PickUp with position to pick up from
3. Else call Tip Discard
********/

func (d *Compact32Deck) TipOperation(to db.TipOperation) (response string, err error) {

	switch to.Type {
	case db.PickupTip:
		response, err = d.tipPickup(to.Position)
	case db.DiscardTip:
		response, err = d.tipDiscard()
	}
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue doing Tip Operation. Error: %v", err)
	}

	return "Tip Operation was successfull", nil

}

/****ALGORITHM******
1. Move Deck to the tip's position
2. Move Syringe Module down fast to tip's base
3. Move Syringe Module down really slow to tip's inside
	// 3.1 Update Tip Height
4. Move Syringe Module up with tip to Resting position.

********/

func (d *Compact32Deck) tipPickup(pos int64) (response string, err error) {

	// **************
	// Tip PickUp	*
	//***************
	var deckAndMotor DeckNumber
	var position, distanceToTravel, pickedTipHeight, ttBase, slowInside, deckBase, oneMicroLitrePulses float64
	var direction, pulses uint16
	var ok, aspireAir bool
	deckAndMotor.Deck = d.name

	//
	// 1. Move Deck to the tip's position
	//

	deckAndMotor.Number = K5_Deck

	logger.Infoln("Moving Deck to pos_", pos)
	if position, ok = consDistance["pos_"+fmt.Sprintf("%d", pos)]; !ok {
		err = fmt.Errorf("pos_" + fmt.Sprintf("%d", pos) + " doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}
	distanceToTravel = Positions[deckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Deck to tip source. Error: %v", err)
	}

	//
	// 2. Move Syringe Module down fast to tip's base
	//

	deckAndMotor.Number = K9_Syringe_Module_LHRH

	// TODO: Handle this in non-harcoded way as tips add up in future
	// We can do this by being aware of the tip.
	// So, in future add a field like height of tip above the deck

	logger.Infoln("Moving Syringe to tip's base")
	if position, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}

	// How do we know which tip exists?
	// By ID of that tip/tube
	// We need recipe details to get tips tubes

	recipe := deckRecipe[d.name]
	if recipe.Name == "" {
		err = fmt.Errorf("no recipe in progress for deck %v", d.name)
		logger.Errorln("Error: ", err)
		return "", err
	}

	var id *int64
	switch pos {
	case 1:
		id = recipe.Position1
		aspireAir = true
	case 2:
		id = recipe.Position2
		aspireAir = true
	case 3:
		id = recipe.Position3
		aspireAir = true
	case 4:
		id = recipe.Position4
	case 5:
		id = recipe.Position5
	}

	if id == nil {
		err = fmt.Errorf("no tip exists for position %v", pos)
		logger.Errorln("Error: ", err)
		return "", err
	}

	// convert interface to float64
	if ttBase, ok = tipstubes[*id]["tt_base"].(float64); !ok {
		err = fmt.Errorf("tts_base doesn't exist for tip with ID %v", id)
		logger.Errorln("Error: ", err)
		return "", err
	}

	if pickedTipHeight, ok = tipstubes[*id]["height"].(float64); !ok {
		err = fmt.Errorf("tts_base doesn't exist for tip with ID %v", id)
		logger.Errorln("Error: ", err)
		return "", err
	}

	distanceToTravel = position - ttBase - Positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to tip's base. Error: %v", err)
	}

	//
	// 3. Move Syringe Module down slow to tip's inside
	//

	logger.Infoln("Moving Syringe to tip's inside")
	if slowInside, ok = consDistance["slow_inside"]; !ok {
		err = fmt.Errorf("slow_inside doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}

	if deckBase, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}

	distanceToTravel = slowInside

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	// Giving it real slow speed
	tipHeight[d.name] = pickedTipHeight - slowInside

	_, err = d.setupMotor(homingSlowSpeed, pulses, Motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)

	// 3.1 Update Tip Height
	// Even if error occurs we must be of the opinion that tip has been picked up
	// Set tipHeight to pickedTipHeight minus the slow_inside

	if err != nil {
		logger.Errorln("Error: ", err)
		return "", fmt.Errorf("There was issue moving Syinge Module to tip's inside. Error: %v", err)
	}

	//
	// 4. Move Syringe Module up with tip to a resting Position.
	//

	if position, ok = consDistance["pickup_tip_up"]; !ok {
		err = fmt.Errorf("pickup_tip_up doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}

	logger.Infoln("Moving Syringe Module to PickupTip")

	// go pickup_tip_up mm above
	distanceToTravel = Positions[deckAndMotor] + tipHeight[d.name] - (deckBase - position)

	// We know Concrete Direction here, its UP

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	_, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to %v. Error: %v", "PickupTip", err)
	}

	// Aspire some Air
	if !aspireAir {
		goto skipAspireAir
	}
	deckAndMotor.Number = K10_Syringe_LHRH
	oneMicroLitrePulses = float64(config.GetMicroLitrePulses())
	pulses = uint16(math.Round(oneMicroLitrePulses * airReserve))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], ASPIRE, deckAndMotor.Number)
	if err != nil {
		return
	}

skipAspireAir:
	return "Tip PickUp was successfull", nil

}

/****ALGORITHM******
1. Move Deck to the big hole's position
2. Move Syringe Module down fast to deck's base
3. Move Syringe Module down really slow till enough inside big hole
4. Move Deck to the small hole's position
5. Move Syringe Module up slow with tip till deck base, to drop off the tip.
	// 5.1 update tip Height
6. Move Syringe Module up fast with tip to Resting position.
*/

// TODO: Currently only discarding at Discard box so avoiding at_pickup_passing condition
func (d *Compact32Deck) tipDiscard() (response string, err error) {

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
		logger.Errorln("Error: ", err)
		return "", err
	}

	//
	//  1. Move Deck to the big hole's position
	//

	deckAndMotor.Number = K5_Deck

	logger.Infoln("Moving Deck to discard_big_hole")
	if position, ok = consDistance["discard_big_hole"]; !ok {
		err = fmt.Errorf("discard_big_hole doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}
	distanceToTravel = Positions[deckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Deck to discard_big_hole source. Error: %v", err)
	}

	//
	// 2. Move Syringe Module fast to deck base
	//

	// This discard has to start here else tip will dash along deck for discard
	d.setTipDiscardInProgress()
	defer d.resetTipDiscardInProgress()

	deckAndMotor.Number = K9_Syringe_Module_LHRH

	logger.Infoln("Moving Syringe to deck's base")
	if position, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}

	distanceToTravel = Positions[deckAndMotor] - position

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], direction, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to deck's base. Error: %v", err)
	}

	//
	// 3. Move Syringe Module down really slow till enough inside big hole
	//

	logger.Infoln("Moving Syringe to deck base inside")
	if position, ok = consDistance["syringe_module_slow_down_for_discard"]; !ok {
		err = fmt.Errorf("syringe_module_slow_down_for_discard doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}
	// Here syringe_module_slow_down_for_discard will always be greater
	// than tipFast
	distanceToTravel = position - Positions[deckAndMotor]

	// We know Concrete Direction here, its DOWN

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	// Giving it real slow speed
	response, err = d.setupMotor(homingSlowSpeed, pulses, Motors[deckAndMotor]["ramp"], DOWN, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to discard_big_hole's inside. Error: %v", err)
	}

	//
	//  4. Move Deck to the small hole's position
	//

	// Dispense complete any remaining air
	deckAndMotor.Number = K10_Syringe_LHRH
	logger.Infoln("Dispense Complete is to start")

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], initialSensorCutSyringePulses, Motors[deckAndMotor]["ramp"], DISPENSE, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return
	}

	deckAndMotor.Number = K5_Deck

	logger.Infoln("Moving Deck to discard_small_hole")
	if position, ok = consDistance["discard_small_hole"]; !ok {
		err = fmt.Errorf("discard_small_hole doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}
	distanceToTravel = Positions[deckAndMotor] - position

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	// We know concrete direction, here its towards sensor/ FWD
	response, err = d.setupMotor(homingSlowSpeed, pulses, Motors[deckAndMotor]["ramp"], FWD, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Deck to discard_big_hole source. Error: %v", err)
	}

	//
	// 5. Move Syringe Module up slow with tip till deck base, to drop off the tip.
	//
	deckAndMotor.Number = K9_Syringe_Module_LHRH

	logger.Infoln("Moving Syringe Module to drop off the tip")

	if position, ok = consDistance["deck_base"]; !ok {
		err = fmt.Errorf("deck_base doesn't exist for consumable distances")
		logger.Errorln("Error: ", err)
		return "", err
	}
	// Here deck_base will always be lesser
	// than syringe_module_slow_down_for_discard
	distanceToTravel = Positions[deckAndMotor] - position

	// We know Concrete Direction here, its UP
	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	// Giving it real slow speed
	response, err = d.setupMotor(homingSlowSpeed, pulses, Motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to deck base. Error: %v", err)
	}

	// 5.1 update tip Height
	// Set tipHeight to 0
	tipHeight[d.name] = 0

	//
	// 6. Move Syringe Module up fast with tip to Resting position.
	//

	logger.Infoln("Moving Syringe Module to Resting Position")

	// Here syringe_parking will always be lesser
	// than deck_base
	distanceToTravel = Positions[deckAndMotor] - parkingPos

	// We know Concrete Direction here, its UP

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], UP, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was issue moving Syinge Module to resting position. Error: %v", err)
	}

	return "Tip Discard was successful", nil
}
