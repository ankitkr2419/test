package plc

import (
	"fmt"
	"mylab/cpagent/db"
	logger "github.com/sirupsen/logrus"

	"time"
)

/*AddDelay :
---------------ALGORITHM
1. First set the timer in progress to signify it is a timer operation
2. Start Timer
3. Wait for the timer to finsh in a loop through channel.
4. If aborted variable is set then return.
5. If the operation is paused then stop the timer ,calculate the time elasped
   and time remaining.
6. If not then loop till the process is resumed.
7. When the process resumes then again start a new timer with the remaining time.
8. At last when the timer is done then it would return success.

*/
func (d *Compact32Deck) AddDelay(delay db.Delay, recipeRun bool) (response string, err error) {
	var t *time.Timer

	var timeElapsedVar int64 = 0
	timeElapsed := &timeElapsedVar

	// set the timer in progress variable to specify that it is not a motor operation.
	d.setTimerInProgress()
	defer d.resetTimerInProgress()
	if recipeRun {
		// Calling in defer cause this will need to get executed despite failure
		defer d.resetRunRecipeData()
	}

skipToStartTimer:
	// start the timer
	t = time.NewTimer(time.Duration(delay.DelayTime-*timeElapsed) * time.Second)
	time1 := time.Now()
	for {
		select {
		// wait for the timer to finish
		case n := <-t.C:
			logger.Infoln("delay time over ", n)
			if d.isUVLightInProgress() {
				// Send 100 % Progress
				d.sendWSData(time1, timeElapsed, delay.DelayTime, uvlightProgress)
				// Send Success
				d.sendWSData(time1, timeElapsed, delay.DelayTime, uvlightSuccess)
				d.ResetRunInProgress()
			}
			if recipeRun {
				// Send 100 % Progress
				d.sendWSData(time1, timeElapsed, delay.DelayTime, recipeProgress)
				// Send Success
				d.sendWSData(time1, timeElapsed, delay.DelayTime, recipeSuccess)
				d.ResetRunInProgress()
			}
			return "SUCCESS", nil
		default:
			// delay of 500 ms for checking the delay over time to avoid too much loop
			time.Sleep(time.Millisecond * 500)
			if d.isMachineInAbortedState() {
				t.Stop()
				if d.isUVLightInProgress() {
					d.resetAborted()
				}
				err = fmt.Errorf("Operation was ABORTED!")
				return "", err
			}
			// When UV Light is in progress nothing else is so no special handling below
			if d.isUVLightInProgress() {
				d.sendWSData(time1, timeElapsed, delay.DelayTime, uvlightProgress)
			}
			if recipeRun {
				d.sendWSData(time1, timeElapsed, delay.DelayTime, recipeProgress)
			}
			// if paused then
			// when timer was paused go again to timer start
			_, wasTimerPaused, err := d.checkPausedState(t, time1, delay.DelayTime, timeElapsed)
			if err != nil {
				return "", err
			}
			if wasTimerPaused {
				goto skipToStartTimer
			}
		}
	}
}

func (d *Compact32Deck) checkPausedState(t *time.Timer, time1 time.Time, delay int64, timeElapsed *int64) (response string, wasTimerPaused bool, err error) {
	var remainingTime int64

	if d.isMachineInPausedState() {
		wasTimerPaused = true
		// stop the time
		t.Stop()
		//note the time when paused is hit
		time2 := time.Now()
		// calculate the time elapsed
		*timeElapsed += int64(time2.Sub(time1) / time.Second)
		// calculate the remaining time
		remainingTime = delay - *timeElapsed

		logger.Infof("time : %v %v %v", remainingTime, *timeElapsed, delay)

		// else wait for the process to be resumed
		response, err = d.waitUntilResumed(d.name)
	}
	return
}
