package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
)

/* ****** ALGORITHM *******

 */
func (d *Compact32Deck) AttachDetach(ad db.AttachDetach) (response string, err error) {

	switch ad.Operation {
	case "attach":
		response, err = d.Attach(ad.OperationType)
		if err != nil {
			fmt.Printf("error in attach process %v \n", err.Error())
		}
		return
	case "detach":
		response, err = d.Detach(ad.OperationType)
		if err != nil {
			fmt.Printf("error in attach process %v \n", err.Error())
		}
		return
	}
	return

}
func (d *Compact32Deck) Detach(operationType string) (response string, err error) {

	fmt.Println("In Detach process")
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
	distanceToTravelBack := positions[deckMagnetFwdRev] - magnetBackPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravelBack > 0.1:
		direction = 1
	case distanceToTravelBack < -0.1:
		distanceToTravelBack *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipMagnetBackToSourcePosition
	}

	fmt.Printf("distance to travel fwd 2 %v \n", distanceToTravelBack)
	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravelBack))

	// set up motor for attach step 2 Forward Motion
	response, err = d.SetupMotor(motors[deckMagnetFwdRev]["fast"], uint16(2000), motors[deckMagnetFwdRev]["ramp"], direction, K7_Magnet_Rev_Fwd)
	if err != nil {
		return
	}
	fmt.Println("magnet move forward 2")
	fmt.Println(response)
skipMagnetBackToSourcePosition:

	// if operation type is semi detach then return after doing the first backward step.
	if operationType == "semi_detach" {
		return "Success", nil
	}

	// step 3 to go with the normal flow of full detach.
	// to calculate the relative position of the magnet from its current
	// position to move the magnet Downward for step 2.
	if magnetUpPosition, ok = consDistance["magnet_up_step_1"]; !ok {
		err = fmt.Errorf("magnet_down_step_2 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravelUp := positions[deckMagnetUpDown] - magnetUpPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravelUp > 0.1:
		direction = 1
	case distanceToTravelUp < -0.1:
		distanceToTravelUp *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipMagnetDownSecToSourcePosition
	}
	fmt.Printf("distance to travel up 1 %v \n", distanceToTravelUp)
	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravelUp))

	// set up motor for attach step 2 Downward Motion
	response, err = d.SetupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, K6_Magnet_Up_Down)
	if err != nil {
		return
	}
	fmt.Println("magnet move upward 1")
	fmt.Println(response)
skipMagnetDownSecToSourcePosition:
	return "Success", nil

}

func (d *Compact32Deck) Attach(operationType string) (response string, err error) {

	fmt.Println("In Attach process")

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

	if deckPosition, ok = consDistance["deck_move_for_magnet_attach"]; !ok {
		err = fmt.Errorf("deck_move_for_magnet_attach doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distToTravel := positions[deckAndNumber] - deckPosition
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
	pulses = uint16(math.Round(float64(motors[deckAndNumber]["steps"]) * distToTravel))
	fmt.Println("steps and pulses")
	fmt.Println(motors[deckAndNumber]["steps"])
	fmt.Println(pulses)

	fmt.Println("Deck is moving backward")

	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], pulses, motors[deckAndNumber]["ramp"], direction, deckAndNumber.Number)
	if err != nil {
		fmt.Printf("error in moving deck to required position")
		return
	}
	fmt.Println(response)

skipDeckToSourcePosition:

	// step 1 to calculate the relative position of the magnet from its current
	// position to move the magnet downward for step 1.
	if magnetDownFirstPosition, ok = consDistance["magnet_down_step_1"]; !ok {
		err = fmt.Errorf("magnet_down_step_1 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet down step 1
	distanceToTravelDown := positions[deckMagnetUpDown] - magnetDownFirstPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravelDown > 0.1:
		direction = 1
	case distanceToTravelDown < -0.1:
		distanceToTravelDown *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipMagnetDownToSourcePosition
	}

	fmt.Printf("distance to travel down %v \n", distanceToTravelDown)
	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravelDown))

	// set up motor for attach step 1 Downward Motion
	response, err = d.SetupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, K6_Magnet_Up_Down)
	if err != nil {
		return
	}
	fmt.Println("magnet move downward 1")
	fmt.Println(response)

skipMagnetDownToSourcePosition:

	// step 2 to calculate the relative position of the magnet from its current
	// position to move the magnet forward for step 1.
	if magnetFwdFirstPosition, ok = consDistance["magnet_forward_step_1"]; !ok {
		err = fmt.Errorf("magnet_forward_step_1 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravelFwd := positions[deckMagnetFwdRev] - magnetFwdFirstPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravelFwd > 0.1:
		direction = 1
	case distanceToTravelFwd < -0.1:
		distanceToTravelFwd *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipMagnetFwdToSourcePosition
	}
	fmt.Printf("distance to travel forward %v \n", distanceToTravelFwd)
	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravelFwd))

	// set up motor for attach step 1 forward Motion
	response, err = d.SetupMotor(motors[deckMagnetFwdRev]["fast"], pulses, motors[deckMagnetFwdRev]["ramp"], direction, K7_Magnet_Rev_Fwd)
	if err != nil {
		return
	}
	fmt.Println("magnet move forward 1")
	fmt.Println(response)

skipMagnetFwdToSourcePosition:
	// step 3 to calculate the relative position of the magnet from its current
	// position to move the magnet Downward for step 2.
	if magnetDownSecPosition, ok = consDistance["magnet_down_step_2"]; !ok {
		err = fmt.Errorf("magnet_down_step_2 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravelDown = positions[deckMagnetUpDown] - magnetDownSecPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravelDown > 0.1:
		direction = 1
	case distanceToTravelDown < -0.1:
		distanceToTravelDown *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipMagnetDownSecToSourcePosition
	}
	fmt.Printf("distance to travel down 2 %v \n", distanceToTravelDown)
	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravelDown))

	// set up motor for attach step 2 Downward Motion
	response, err = d.SetupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, K6_Magnet_Up_Down)
	if err != nil {
		return
	}
	fmt.Println("magnet move downward 2")
	fmt.Println(response)

skipMagnetDownSecToSourcePosition:
	// step 4 to calculate the relative position of the magnet from its current
	// position to move the magnet Downward for step 2.
	if magnetFwdSecPosition, ok = consDistance["magnet_forward_step_2"]; !ok {
		err = fmt.Errorf("magnet_forward_step_2 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravelFwd = positions[deckMagnetFwdRev] - magnetFwdSecPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravelFwd > 0.1:
		direction = 1
	case distanceToTravelFwd < -0.1:
		distanceToTravelFwd *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipMagnetFwdSecToSourcePosition
	}

	fmt.Printf("distance to travel fwd 2 %v \n", distanceToTravelFwd)
	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravelFwd))

	// set up motor for attach step 2 Forward Motion
	response, err = d.SetupMotor(motors[deckMagnetFwdRev]["fast"], uint16(2000), motors[deckMagnetFwdRev]["ramp"], direction, K7_Magnet_Rev_Fwd)
	if err != nil {
		return
	}
	fmt.Println("magnet move forward 2")
	fmt.Println(response)

skipMagnetFwdSecToSourcePosition:
	return "Success", nil

}
