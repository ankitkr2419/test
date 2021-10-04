package plc

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
		response, err = d.attach(ad.Height)
		if err != nil {
			logger.Errorln("error in attach process :", err)
		}
		magnetState.Store(d.name, attached)
		return
	case "detach":
		response, err = d.detach()
		if err != nil {
			logger.Errorln("error in detach process: ", err)
		}
		return
	}
	return

}

/*Detach :
 ****** ALGORITHM *******
1. First move to 12.5 mm backward for the magnet to detach
2. Then if it is semi-detach then stay there
3. move up to 0.5 mm from the 0 position of the magnet.
4. Then move the magnet 20 mm back to avoid any chances of possible collision with the tips.
*/
func (d *Compact32Deck) detach() (response string, err error) {

	// TODO: Check if already detached, then avoid all below claculations
	var magnetBackPosition, magnetUpPosition float64
	var ok bool
	var direction, pulses uint16

	deckMagnetUpDown := DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}
	// motor desc for magnet fwd/rev motion
	deckMagnetFwdRev := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}

	// TODO: This to be taken from Sensor

	// step 1 to calculate the relative position of the magnet from its current
	// position to move the magnet backward for step 1.
	if magnetBackPosition, ok = consDistance["magnet_backward_step_1"]; !ok {
		err = fmt.Errorf("magnet_backward_step_1 doesn't exist for consuamble distances")
		logger.Errorln(err)
		return "", err
	}

	// distance and direction setup for magnet for forward step 1
	distanceToTravel := Positions[deckMagnetFwdRev] - magnetBackPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for detach step 1 Backward Motion
	response, err = d.setupMotor(Motors[deckMagnetFwdRev]["fast"], pulses, Motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	// step 2 to go with the normal flow of full detach.
	// to calculate the relative position of the magnet from its current
	// position to move the magnet Upward for step 2.
	if magnetUpPosition, ok = consDistance["magnet_up_step_1"]; !ok {
		err = fmt.Errorf("magnet_up_step_1 doesn't exist for consuamble distances")
		logger.Errorln(err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravel = Positions[deckMagnetUpDown] - magnetUpPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckMagnetUpDown]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Downward Motion
	response, err = d.setupMotor(Motors[deckMagnetUpDown]["fast"], pulses, Motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}

	// step 3 to go with the normal flow of full detach.
	// to calculate the relative position of the magnet from its current
	// position to move the magnet backward for step 2.
	if magnetBackPosition, ok = consDistance["magnet_back_step_2"]; !ok {
		err = fmt.Errorf("magnet_back_step_2 doesn't exist for consuamble distances")
		logger.Errorln(err)
		return "", err
	}
	// distance and direction setup for magnet for backward step 2
	distanceToTravel = Positions[deckMagnetFwdRev] - magnetBackPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Downward Motion
	response, err = d.setupMotor(Motors[deckMagnetFwdRev]["fast"], pulses, Motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	magnetState.Store(d.name, detached)

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
func (d *Compact32Deck) attach(height int64) (response string, err error) {

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
		logger.Errorln(err)
		return "", err
	}

	distanceToTravel := Positions[deckAndNumber] - deckPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckAndNumber]["steps"]) * distanceToTravel))

	response, err = d.setupMotor(Motors[deckAndNumber]["fast"], pulses, Motors[deckAndNumber]["ramp"], direction, deckAndNumber.Number)
	if err != nil {
		logger.Errorln("error in moving deck to required position")
		return
	}

	// step 1 to calculate the relative position of the magnet from its current
	// position to move the magnet downward for step 1.
	if magnetDownFirstPosition, ok = consDistance["magnet_down_step_1"]; !ok {
		err = fmt.Errorf("magnet_down_step_1 doesn't exist for consuamble distances")
		logger.Errorln(err)
		return "", err
	}
	// distance and direction setup for magnet down step 1
	distanceToTravel = Positions[deckMagnetUpDown] - magnetDownFirstPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckMagnetUpDown]["steps"]) * distanceToTravel))

	// set up motor for attach step 1 Downward Motion
	response, err = d.setupMotor(Motors[deckMagnetUpDown]["fast"], pulses, Motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}

	// TODO: This to be taken from Sensor

	// step 2 to calculate the relative position of the magnet from its current
	// position to move the magnet forward for step 1.
	if magnetFwdFirstPosition, ok = consDistance["magnet_forward_step_1"]; !ok {
		err = fmt.Errorf("magnet_forward_step_1 doesn't exist for consuamble distances")
		logger.Errorln(err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravel = Positions[deckMagnetFwdRev] - magnetFwdFirstPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for attach step 1 forward Motion
	response, err = d.setupMotor(Motors[deckMagnetFwdRev]["fast"], pulses, Motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	// Setting magnet attched here cause there has been sigificant movement from magnet
	magnetState.Store(d.name, attached)

	// step 3 to calculate the relative position of the magnet from its current
	// position to move the magnet Downward for step 2.
	magnetDownSecPosition = float64(height)

	// distance and direction setup for magnet for forward step 1
	distanceToTravel = Positions[deckMagnetUpDown] - magnetDownSecPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckMagnetUpDown]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Downward Motion
	response, err = d.setupMotor(Motors[deckMagnetUpDown]["fast"], pulses, Motors[deckMagnetUpDown]["ramp"], direction, deckMagnetUpDown.Number)
	if err != nil {
		return
	}

	// step 4 to calculate the relative position of the magnet from its current
	// position to move the magnet Downward for step 2.
	if magnetFwdSecPosition, ok = consDistance["magnet_forward_step_2"]; !ok {
		err = fmt.Errorf("magnet_forward_step_2 doesn't exist for consuamble distances")
		logger.Errorln(err)
		return "", err
	}
	// distance and direction setup for magnet for forward step 1
	distanceToTravel = Positions[deckMagnetFwdRev] - magnetFwdSecPosition

	modifyDirectionAndDistanceToTravel(&distanceToTravel, &direction)

	pulses = uint16(math.Round(float64(Motors[deckMagnetFwdRev]["steps"]) * distanceToTravel))

	// set up motor for attach step 2 Forward Motion
	response, err = d.setupMotor(Motors[deckMagnetFwdRev]["fast"], pulses, Motors[deckMagnetFwdRev]["ramp"], direction, deckMagnetFwdRev.Number)
	if err != nil {
		return
	}

	return "Success", nil
}
