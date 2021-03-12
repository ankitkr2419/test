package compact32

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) DiscardBoxCleanup() (response string, err error) {

	var position, distToTravel float64
	var ok bool
	var pulses uint16
	deckAndMotor := DeckNumber{Deck: d.name, Number: K5_Deck}

	if runInProgress[d.name] {
		err = fmt.Errorf("previous run is already in progress... wait or abort it")
		return
	}
	aborted[d.name] = false
	runInProgress[d.name] = true
	defer d.ResetRunInProgress()

	fmt.Println("Deck is moving to discard_box_open_position")

	if position, ok = consDistance["discard_box_open_position"]; !ok {
		err = fmt.Errorf("discard_box_open_position doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distToTravel = position - positions[deckAndMotor]

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

	if runInProgress[d.name] {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}
	aborted[d.name] = false
	runInProgress[d.name] = true
	defer d.ResetRunInProgress()

	fmt.Println("Deck is moving to deck_start")

	if position, ok = consDistance["deck_start"]; !ok {
		err = fmt.Errorf("deck_start doesn't exist for consumable distances")
		fmt.Println("Error: ", err)
		return "", err
	}

	distToTravel = positions[deckAndMotor] - position

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

/*
ALGORITHM
	1.  Start UV Light
	2.  Start Timer
	3.  Monitor for PAUSE and abort
	4.  If Paused then monitor for resumed
*/

func (d *Compact32Deck) UVLight(uvTime string) (response string, err error) {

	// totalTime is UVLight timer time in Seconds
	// timeElapsed is the time from start to pause
	var totalTime, timeElapsed int64
	var t *time.Timer

	if runInProgress[d.name] {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}

	aborted[d.name] = false
	runInProgress[d.name] = true
	defer d.ResetRunInProgress()

	totalTime, err = calculateUVTimeInSeconds(uvTime)
	if err != nil {
		return "", err
	}

	// set the timer in progress variable to specify that it is not a motor operation.
	d.SetTimerInProgress()
	defer d.ResetTimerInProgress()

skipToStartTimer:
	// start the timer
	t = time.NewTimer(time.Duration(totalTime) * time.Second)
	time1 := time.Now()
	for {
		select {
		// wait for the timer to finish
		case n := <-t.C:
			fmt.Printf("delay time over %v", n)
			return "SUCCESS", nil
		// or check for its pause/abort
		default:
			// delay of 300 ms to reduce CPU usage
			time.Sleep(time.Millisecond * 300)
			if aborted[d.name] {
				t.Stop()
				err = fmt.Errorf("Operation was ABORTED!")
				return "", err
			}
			// if paused then
			if paused[d.name] {
				// stop the timer
				t.Stop()
				//note the time when paused was hit
				time2 := time.Now()
				// calculate the time elapsed in Seconds
				timeElapsed = int64(time2.Sub(time1) / time.Second)
				// calculate the remaining time
				totalTime = totalTime - timeElapsed

				logger.Infof("remaining time %v and elapsed time %v", totalTime, timeElapsed)
				// if the remaining time is less than a sec then time is over
				if totalTime < 2 {
					return "SUCCESS", nil
				}
				// else wait for the process to be resumed
				for {
					time.Sleep(time.Millisecond * 300)
					if !paused[d.name] {
						// when resumed go again to timer start
						goto skipToStartTimer
					}
				}

			}
		}
	}

	return "UV Light Completed Successfully", nil
}

func calculateUVTimeInSeconds(uvTime string) (totalTime int64, err error) {

	timeArr := strings.Split(uvTime, ":")
	if len(timeArr) != 3 {
		err = fmt.Errorf("time format isn't of the form HH:MM:SS")
		return 0, err
	}
	hours, err := strconv.ParseInt(timeArr[0], 10, 64)
	if err != nil || hours > 99 || hours < 0 {
		err = fmt.Errorf("please check hours format, valid range: [0,99]")
		return 0, err
	}
	minutes, err := strconv.ParseInt(timeArr[1], 10, 64)
	if err != nil || minutes > 59 || minutes < 0 {
		err = fmt.Errorf("please check minutes format, valid range: [0,59]")
		return 0, err
	}
	seconds, err := strconv.ParseInt(timeArr[2], 10, 64)
	if err != nil || seconds > 59 || seconds < 0 {
		err = fmt.Errorf("please check seconds format, valid range: [0,59]")
		return 0, err
	}

	totalTime = hours*60*60 + minutes*60 + seconds

	if totalTime < minimumUVLightOnTime {
		err = fmt.Errorf("please check your time. minimum time is : ", minimumUVLightOnTime, " seconds")
		return 0, err
	}

	return
}
