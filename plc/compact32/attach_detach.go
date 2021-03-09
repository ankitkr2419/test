package compact32

import (
	"fmt"
	"math"
	"mylab/cpagent/db"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) AttachDetach(ad db.AttachDetach) (response string, err error) {

	// Operation attach and detach
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

/*Detach :
 ****** ALGORITHM *******
1. First move to 12.5 mm backward for the magnet to detach
2. Then if it is semi-detach then stay there
3. If it is full-detach then move up to 0.5 mm from the 0 position of the magnet.
4. Then move the magnet 20 mm back to avoid any chances of possible collision with the tips.
*/
func (d *Compact32Deck) Detach(operationType string) (response string, err error) {

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

	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravelBack))

	// set up motor for attach step 2 Forward Motion
	response, err = d.SetupMotor(motors[deckMagnetFwdRev]["fast"], pulses, motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}
	logger.Printf("magnet move forward 2")
skipMagnetBackToSourcePosition:

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
		goto skipMagnetbackSecPosition
	}

	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravelUp))

	// set up motor for attach step 2 Downward Motion
	response, err = d.SetupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}
	logger.Printf("magnet move upward 1")
	logger.Printf("%v", response)

skipMagnetbackSecPosition:
	// step 3 to go with the normal flow of full detach.
	// to calculate the relative position of the magnet from its current
	// position to move the magnet backward for step 2.
	if magnetUpPosition, ok = consDistance["magnet_back_step_2"]; !ok {
		err = fmt.Errorf("magnet_back_step_1 doesn't exist for consuamble distances")
		fmt.Println("Error: ", err)
		return "", err
	}
	// distance and direction setup for magnet for backward step 2
	distanceToTravelBack = positions[deckMagnetFwdRev] - magnetUpPosition

	switch {
	// distToTravel > 0 means go towards the Sensor or FWD
	case distanceToTravelUp > 0.1:
		direction = 1
	case distanceToTravelUp < -0.1:
		distanceToTravelUp *= -1
		direction = 0
	default:
		// Skip the setUpMotor Step
		goto skipMagnetToSuccessPosition
	}

	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravelUp))

	// set up motor for attach step 2 Downward Motion
	response, err = d.SetupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}

skipMagnetToSuccessPosition:
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
func (d *Compact32Deck) Attach(operationType string) (response string, err error) {

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

	pulses = uint16(math.Round(float64(motors[deckAndNumber]["steps"]) * distToTravel))

	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], pulses, motors[deckAndNumber]["ramp"], direction, deckAndNumber.Number)
	if err != nil {
		fmt.Printf("error in moving deck to required position")
		return
	}
	logger.Printf("%s \n", response)

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

	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravelDown))

	// set up motor for attach step 1 Downward Motion
	response, err = d.SetupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}
	logger.Printf("magnet move downward 1")
	logger.Printf("%v \n", response)

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
	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravelFwd))

	// set up motor for attach step 1 forward Motion
	response, err = d.SetupMotor(motors[deckMagnetFwdRev]["fast"], pulses, motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}
	logger.Printf("magnet move forward 1")
	logger.Printf("%v \n", response)

skipMagnetFwdToSourcePosition:
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
	pulses = uint16(math.Round(float64(motors[deckMagnetUpDown]["steps"]) * distanceToTravelDown))

	// set up motor for attach step 2 Downward Motion
	response, err = d.SetupMotor(motors[deckMagnetUpDown]["fast"], pulses, motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}
	logger.Printf("magnet move downward 2")
	logger.Printf("%v \n", response)

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

	pulses = uint16(math.Round(float64(motors[deckMagnetFwdRev]["steps"]) * distanceToTravelFwd))

	// set up motor for attach step 2 Forward Motion
	response, err = d.SetupMotor(motors[deckMagnetFwdRev]["fast"], uint16(2000), motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}
	logger.Printf("magnet move forward 2")
	logger.Printf("%v \n", response)

skipMagnetFwdSecToSourcePosition:
	return "Success", nil

}
