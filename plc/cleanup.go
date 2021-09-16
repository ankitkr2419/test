package plc

import (
	"fmt"
	"math"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) DiscardBoxCleanup() (response string, err error) {

	if !d.IsMachineHomed() {
		err = responses.PleaseHomeMachineError
		return
	}

	if d.IsRunInProgress() {
		err = responses.PreviousRunInProgressError
		return
	}

	var position, distanceToTravel float64
	var ok bool
	var pulses uint16
	deckAndMotor := DeckNumber{Deck: d.name, Number: K5_Deck}

	d.SetRunInProgress()
	defer d.ResetRunInProgress()

	logger.Infoln("Deck is moving to discard_box_open_position")

	if position, ok = consDistance["discard_box_open_position"]; !ok {
		err = fmt.Errorf("discard_box_open_position doesn't exist for consumable distances")
		logger.Errorln(err)
		return "", err
	}

	distanceToTravel = position - Positions[deckAndMotor]

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	// We know concrete direction here, its REV
	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], REV, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was an issue moving deck REV to discard_box_open_position. Error: %v", err)
	}

	logger.Infoln("Moved Deck to Cleanup Discard Box Successfully")

	return "DISCARD BOX CLEANUP SUCCESS", nil
}

func (d *Compact32Deck) RestoreDeck() (response string, err error) {

	if !d.IsMachineHomed() {
		err = responses.PleaseHomeMachineError
		return
	}

	if d.IsRunInProgress() {
		err = responses.PreviousRunInProgressError
		return
	}

	var position, distanceToTravel float64
	var ok bool
	var pulses uint16
	deckAndMotor := DeckNumber{Deck: d.name, Number: K5_Deck}

	d.SetRunInProgress()
	defer d.ResetRunInProgress()

	logger.Infoln("Deck is moving to deck_start")

	if position, ok = consDistance["deck_start"]; !ok {
		err = fmt.Errorf("deck_start doesn't exist for consumable distances")
		logger.Errorln(err)
		return "", err
	}

	distanceToTravel = Positions[deckAndMotor] - position

	pulses = uint16(math.Round(float64(Motors[deckAndMotor]["steps"]) * distanceToTravel))

	// We know concrete direction here, its FWD
	response, err = d.setupMotor(Motors[deckAndMotor]["fast"], pulses, Motors[deckAndMotor]["ramp"], FWD, deckAndMotor.Number)
	if err != nil {
		logger.Errorln(err)
		return "", fmt.Errorf("There was an issue moving deck FWD to deck_start. Error: %v", err)
	}

	logger.Infoln("Moved Deck back to homing position")

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
	defer func() {
		if err != nil {
			logger.Errorln(err.Error())
			d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
		}
	}()

	if !d.IsMachineHomed() {
		err = responses.PleaseHomeMachineError
		return
	}

	if d.IsRunInProgress() {
		err = responses.PreviousRunInProgressError
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
	if totalTime < minimumUVLightOnTime {
		err = fmt.Errorf("please check your time. minimum time is : %v seconds", minimumUVLightOnTime)
		return "", err
	}

	//
	// 2. Start UV Light
	//
	response, err = d.switchOnUVLight()
	if err != nil {
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

	response, err = d.AddDelay(delay, false)
	if err != nil {
		return
	}

	return "UV Light Completed Successfully", nil
}

func (d *Compact32Deck) waitUntilResumed(deck string) (response string, err error) {
	for {
		time.Sleep(time.Millisecond * 300)

		if d.isMachineInAbortedState() {
			return "", responses.AbortedError
		}

		if !d.isMachineInPausedState() {
			// when resumed go again to timer start
			return "Resumed", nil
		}

	}
}

func calculateUVTimeInSeconds(uvTime string) (totalTime int64, err error) {
	totalTime, err = db.CalculateTimeInSeconds(uvTime)
	return
}
