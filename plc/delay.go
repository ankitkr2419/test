package plc

import (
	"encoding/json"
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
			logger.Infoln("delay time over ", n)
			return "SUCCESS", nil
		default:
			// delay of 300 ms for checking the delay over time to avoid too much loop
			time.Sleep(time.Millisecond * 300)
			if d.isMachineInAbortedState() {
				t.Stop()
				if d.isUVLightInProgress() || d.isPIDCalibrationInProgress() {
					d.resetAborted()
				}
				err = fmt.Errorf("Operation was ABORTED!")
				return "", err
			}
			if d.isUVLightInProgress() {
				uvtime := time.Now()
				uvTimePassed := int64(uvtime.Sub(time1).Seconds()) + *timeElapsed
				uvRemainingTime := delay.DelayTime - uvTimePassed
				progress = (uvTimePassed * 100) / delay.DelayTime
				wsProgressOperation := WSData{
					Progress: float64(progress),
					Deck:     d.name,
					Status:   "PROGRESS_UVLIGHT",
					OperationDetails: OperationDetails{
						Message:       fmt.Sprintf("progress_uvLight_uv light cleanup in progress for deck %s ", d.name),
						RemainingTime: uvRemainingTime,
					},
				}

				wsData, err := json.Marshal(wsProgressOperation)
				if err != nil {
					logger.Errorf("error in marshalling web socket data %v", err.Error())
					d.WsErrCh <- err
				}
				d.WsMsgCh <- fmt.Sprintf("progress_uvLight_%v", string(wsData))
			} else if d.isPIDCalibrationInProgress() {
				pidtime := time.Now()
				pidTimePassed := int64(pidtime.Sub(time1).Seconds()) + *timeElapsed
				pidRemainingTime := delay.DelayTime - pidTimePassed
				progress = (pidTimePassed * 100) / delay.DelayTime
				wsProgressOperation := WSData{
					Progress: float64(progress),
					Deck:     d.name,
					Status:   "PROGRESS_PIDCALIBRATION",
					OperationDetails: OperationDetails{
						Message:       fmt.Sprintf("progress_pidCalibration_pid calibration in progress for deck %s ", d.name),
						RemainingTime: pidRemainingTime,
					},
				}

				wsData, err := json.Marshal(wsProgressOperation)
				if err != nil {
					logger.Errorf("error in marshalling web socket data %v", err.Error())
					d.WsErrCh <- err
				}
				d.WsMsgCh <- fmt.Sprintf("progress_pidCalibration_%v", string(wsData))
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
