package compact32

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"

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
func (d *Compact32Deck) AddDelay(delay db.Delay) (response string, err error) {
	var t *time.Timer

	var timeElapsedVar int64 = 0
	var progress int64 = 0
	timeElapsed := &timeElapsedVar

	// set the timer in progress variable to specify that it is not a motor operation.
	d.setTimerInProgress()
	defer d.resetTimerInProgress()

skipToStartTimer:
	// start the timer
	t = time.NewTimer(time.Duration(delay.DelayTime) * time.Second)
	time1 := time.Now()
	for {
		select {
		// wait for the timer to finish
		case n := <-t.C:
			fmt.Printf("delay time over %v", n)
			return "SUCCESS", nil
		default:
			// delay of 300 ms for checking the delay over time to avoid too much loop
			time.Sleep(time.Millisecond * 300)
			if d.isMachineInAbortedState() {
				t.Stop()
				err = fmt.Errorf("Operation was ABORTED!")
				return "", err
			}
			if d.isUVLightInProgress() {
				uvtime := time.Now()
				uvTimePassed := uvtime.Sub(time1).Seconds()
				progress = (int64(uvTimePassed) * 100) / delay.DelayTime
				wsProgressOperation := plc.WSData{
					Progress: float64(progress),
					Deck:     d.name,
					Status:   "PROGRESS_UVLIGHT",
					OperationDetails: plc.OperationDetails{
						Message: fmt.Sprintf("progress_uvLight_uv light cleanup in progress for deck %s ", d.name),
					},
				}

				wsData, err := json.Marshal(wsProgressOperation)
				if err != nil {
					logger.Errorf("error in marshalling web socket data %v", err.Error())
					d.WsErrCh <- err
				}
				d.WsMsgCh <- fmt.Sprintf("progress_uvLight_%v", string(wsData))
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
		if err != nil {
			return
		}

	}
	return
}
