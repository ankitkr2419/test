package plc

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
	"strconv"
	"strings"
	"time"
)

func (d *Compact32Deck) DiscardBoxCleanup() (response string, err error) {

	if !d.IsMachineHomed() {
		err = fmt.Errorf("Please home the machine first!")
		return
	}

	if d.IsRunInProgress() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return
	}

	var position, distanceToTravel float64
	var ok bool
	var pulses uint16
	deckAndMotor := DeckNumber{Deck: d.Name, Number: K5_Deck}

	d.SetRunInProgress()
	defer d.ResetRunInProgress()

	fmt.Println("Deck is moving to discard_box_open_position")

	if position, ok = consDistance["discard_box_open_position"]; !ok {
		err = fmt.Errorf("discard_box_open_position doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distanceToTravel = position - positions[deckAndMotor]

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	// We know concrete direction here, its REV
	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], REV, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was an issue moving deck REV to discard_box_open_position. Error: %v", err)
	}

	fmt.Println("Moved Deck to Cleanup Discard Box Successfully")

	return "DISCARD BOX CLEANUP SUCCESS", nil
}

func (d *Compact32Deck) RestoreDeck() (response string, err error) {

	if !d.IsMachineHomed() {
		err = fmt.Errorf("Please home the machine first!")
		return
	}

	if d.IsRunInProgress() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return
	}

	var position, distanceToTravel float64
	var ok bool
	var pulses uint16
	deckAndMotor := DeckNumber{Deck: d.Name, Number: K5_Deck}

	d.SetRunInProgress()
	defer d.ResetRunInProgress()

	fmt.Println("Deck is moving to deck_start")

	if position, ok = consDistance["deck_start"]; !ok {
		err = fmt.Errorf("deck_start doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distanceToTravel = positions[deckAndMotor] - position

	pulses = uint16(math.Round(float64(motors[deckAndMotor]["steps"]) * distanceToTravel))

	// We know concrete direction here, its FWD
	response, err = d.setupMotor(motors[deckAndMotor]["fast"], pulses, motors[deckAndMotor]["ramp"], FWD, deckAndMotor.Number)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("There was an issue moving deck FWD to deck_start. Error: %v", err)
	}

	fmt.Println("Moved Deck back to homing position")

	return "DECK RESTORED SUCCESS", nil
}

/*
ALGORITHM
	1. 	Calculate UV Time in Seconds
	1.  Start UV Light
	2.  Add delay
	3.  Monitor for PAUSE and abort or completion
	4.  If Paused then monitor for resumed
*/

func (d *Compact32Deck) UVLight(uvTime string) (response string, err error) {

	if !d.IsMachineHomed() {
		err = fmt.Errorf("Please home the machine first!")
		d.WsErrCh <- err

		return
	}

	if d.IsRunInProgress() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		d.WsErrCh <- err

		return
	}

	// totalTime is UVLight timer time in Seconds
	// timeElapsed is the time from start to pause

	var totalTime int64

	d.SetRunInProgress()
	defer d.ResetRunInProgress()

	//
	// 1. 	Calculate UV Time in Seconds
	//
	totalTime, err = calculateUVTimeInSeconds(uvTime)
	if err != nil {

		return "", err
	}

	//
	// 2. Start UV Light
	//
	response, err = d.switchOnUVLight()
	if err != nil {
		d.WsErrCh <- err
		return
	}
	d.setUVLightInProgress()
	defer d.resetUVLightInProgress()

	//
	// 3. Add delay
	//

	delay := db.Delay{
		DelayTime: totalTime,
	}

	response, err = d.AddDelay(delay)
	if err != nil {
		d.WsErrCh <- err
		return
	}

	d.WsMsgCh <- fmt.Sprintf("success_uvlight_UV Light Completed Successfully %v", d.Name)
	return "UV Light Completed Successfully", nil
}

func (d *Compact32Deck) waitUntilResumed(deck string) (response string, err error) {
	for {
		time.Sleep(time.Millisecond * 300)
		if !d.isMachineInPausedState() {
			// when resumed go again to timer start
			return "Resumed", nil
		}

		if d.isMachineInAbortedState() {
			err = fmt.Errorf("Operation was Aborted!")
			return "", err
		}
	}
}

func calculateUVTimeInSeconds(uvTime string) (totalTime int64, err error) {

	var hours, minutes, seconds int64
	timeArr := strings.Split(uvTime, ":")
	if len(timeArr) != 3 {
		err = fmt.Errorf("time format isn't of the form HH:MM:SS")
		return 0, err
	}

	hours, err = parseIntRange(timeArr[0], "hours", 0, 24)
	if err != nil {
		return 0, err
	}

	minutes, err = parseIntRange(timeArr[1], "minutes", 0, 59)
	if err != nil {
		return 0, err
	}

	seconds, err = parseIntRange(timeArr[2], "seconds", 0, 59)
	if err != nil {
		return 0, err
	}

	totalTime = hours*60*60 + minutes*60 + seconds

	if totalTime < minimumUVLightOnTime {
		err = fmt.Errorf("please check your time. minimum time is : %v seconds", minimumUVLightOnTime)
		return 0, err
	}

	return
}

func parseIntRange(timeString, unit string, min, max int64) (value int64, err error) {
	value, err = strconv.ParseInt(timeString, 10, 64)
	if err != nil || value > max || value < min {
		err = fmt.Errorf("please check %v format, valid range: [%d,%d]", unit, min, max)
		return 0, err
	}
	return
}
