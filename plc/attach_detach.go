package plc

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
)

func (d *Compact32Deck) AttachDetach(ad db.AttachDetach) (response string, err error) {

	// TODO : Remove this State storing to attach/detach individual process
	// Operation attach and detach
	switch ad.Operation {
	case "attach":
		response, err = d.attach(ad.OperationType)
		if err != nil {
			fmt.Printf("error in attach process %v \n", err.Error())
		}
		magnetState.Store(d.name, attached)
		return
	case "detach":
		response, err = d.detach(ad.OperationType)
		if err != nil {
			fmt.Printf("error in attach process %v \n", err.Error())
		}

		// NOTE: Below string literal "semi_detach" is dependent on db schema
		// operation_type is its'd db variable
		// Make sure that db changes are reflected at here as well
		if ad.OperationType == "semi_detach" {
			magnetState.Store(d.name, semiDetached)
		} else {
			magnetState.Store(d.name, detached)
		}
		return
	}
	return

}

/*Detach :
 ****** ALGORITHM *******
1. First move to 12.5 mm backward for the magnet to detach
2. Then if it is semi-detach then stay there
3. If it is full-detach then move up to 0.5 mm from the 0 position of the magnet.
4. Then move the magnet 20 mm back to avoid any chances of possible collision with the tips.
*/
func (d *Compact32Deck) detach(operationType string) (response string, err error) {

	// TODO: Check if already detached, then avoid all below claculations
	var magnetBackPosition, magnetUpPosition float64
	var ok bool
	var direction, pulses uint16

	deckMagnetUpDown := DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}
	// motor desc for magnet fwd/rev motion
	deckMagnetFwdRev := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}
	// step 1 to calculate the relative position of the magnet from its current
	// position to move the magnet backward for step 1.
	if magnetBackPosition, ok = consDistance["magnet_backward_step_1"]; !ok {
		err = fmt.Errorf("magnet_backward_step_1 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	// distance and direction setup for magnet for forward step 1
	distanceToTravel := positions[deckMagnetFwdRev] - magnetBackPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Forward Motion
	response, err = d.setupMotor(motors[deckMagnetFwdRev]["fast"], pulses, motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	// if operation type is semi detach then return after doing the first backward step.
	if operationType == "semi_detach" {
		return "Success", nil
	}

	// step 2 to go with the normal flow of full detach.
	// to calculate the relative position of the magnet from its current
	// position to move the magnet Upward for step 2.
	if magnetUpPosition, ok = consDistance["magnet_up_step_1"]; !ok {
		err = fmt.Errorf("magnet_up_step_1 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravel = positions[deckMagnetUpDown] - magnetUpPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Downward Motion
	response, err = d.setupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}

	// step 3 to go with the normal flow of full detach.
	// to calculate the relative position of the magnet from its current
	// position to move the magnet backward for step 2.
	if magnetBackPosition, ok = consDistance["magnet_back_step_2"]; !ok {
		err = fmt.Errorf("magnet_back_step_2 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for backward step 2
	distanceToTravel = positions[deckMagnetFwdRev] - magnetBackPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Downward Motion
	response, err = d.setupMotor(motors[deckMagnetFwdRev]["fast"], pulses, motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	return "Success", nil

}

/*Attach :
****** ALGORITHM *******
1. Move the deck to the specified position (taken from the consumable_config)
2. First move the magnet 0.5 mm down
3. Then Move the magnet to at a distance of 12.5 mm forward
4. Move the magnet down to 75 mm down behind shaker for attach
5. At last move 5.5 mm forward for the magnet to attach
*/
func (d *Compact32Deck) attach(operationType string) (response string, err error) {

	// TODO: Check if already attached, then avoid all below claculations
	var deckPosition, magnetDownFirstPosition, magnetFwdFirstPosition, magnetDownSecPosition, magnetFwdSecPosition float64
	var ok bool
	var direction, pulses uint16

	//-----------deck
	// move the deck to the position where the magnet position is appropriate to shaker
	deckAndNumber := DeckNumber{Deck: d.name, Number: K5_Deck}
	// motor desc for magnet up/down motion
	deckMagnetUpDown := DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}
	// motor desc for magnet fwd/rev motion
	deckMagnetFwdRev := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}
	// get the consumable deck position

	// deck will be at this position from homing for attach or detach operation
	// We are using shaker_tube to maintain consistency
	if deckPosition, ok = consDistance["shaker_tube"]; !ok {
		err = fmt.Errorf("shaker_tube doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distanceToTravel := positions[deckAndNumber] - deckPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckAndNumber]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(motors[deckAndNumber]["fast"], pulses, motors[deckAndNumber]["ramp"], direction, deckAndNumber.Number)
	if err != nil {
		fmt.Printf("error in moving deck to required position")
		return
	}

	// step 1 to calculate the relative position of the magnet from its current
	// position to move the magnet downward for step 1.
	if magnetDownFirstPosition, ok = consDistance["magnet_down_step_1"]; !ok {
		err = fmt.Errorf("magnet_down_step_1 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet down step 1
	distanceToTravel = positions[deckMagnetUpDown] - magnetDownFirstPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravel))

	// set up motor for attach step 1 Downward Motion
	response, err = d.setupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}

	// step 2 to calculate the relative position of the magnet from its current
	// position to move the magnet forward for step 1.
	if magnetFwdFirstPosition, ok = consDistance["magnet_forward_step_1"]; !ok {
		err = fmt.Errorf("magnet_forward_step_1 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravel = positions[deckMagnetFwdRev] - magnetFwdFirstPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for attach step 1 forward Motion
	response, err = d.setupMotor(motors[deckMagnetFwdRev]["fast"], pulses, motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	// TODO: Set the height according to the operation type (wash,lysis,illusion)
	// For now it has to be kept same for all.
	// step 3 to calculate the relative position of the magnet from its current
	// position to move the magnet Downward for step 2.
	if magnetDownSecPosition, ok = consDistance["magnet_down_step_2"]; !ok {
		err = fmt.Errorf("magnet_down_step_2 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravel = positions[deckMagnetUpDown] - magnetDownSecPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Downward Motion
	response, err = d.setupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}

	// step 4 to calculate the relative position of the magnet from its current
	// position to move the magnet Downward for step 2.
	if magnetFwdSecPosition, ok = consDistance["magnet_forward_step_2"]; !ok {
		err = fmt.Errorf("magnet_forward_step_2 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravel = positions[deckMagnetFwdRev] - magnetFwdSecPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Forward Motion
	response, err = d.setupMotor(motors[deckMagnetFwdRev]["fast"], pulses, motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	return "Success", nil
}

func (d *Compact32Deck) fullDetach() (response string, err error) {
	// Calling AttachDetach below as this handles magnetState implicitly
	// WARNING: Be careful of below string literals "detach" and "full_detach",
	// any changes in db schema of magnets should be reflected in these as well.
	response, err = d.AttachDetach(db.AttachDetach{Operation: "detach", OperationType: "full_detach"})
	if err != nil {
		fmt.Printf("error in magnet detach process %v \n", err.Error())
	}
	return
}
