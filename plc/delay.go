package plc

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"

	"time"
)

type WebsocketOperation string

const (
	uvlightProgress WebsocketOperation = "PROGRESS_UVLIGHT"
	recipeProgress  WebsocketOperation = "PROGRESS_RECIPE"
	uvlightSuccess  WebsocketOperation = "SUCCESS_UVLIGHT"
	recipeSuccess   WebsocketOperation = "SUCCESS_RECIPE"
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

func (d *Compact32Deck) RunRecipeWebsocketData(recipe db.Recipe, processes []db.Process) (err error) {
	deckRecipe[d.name] = recipe
	deckProcesses[d.name] = processes
	if recipe.ProcessCount == 0 {
		return responses.ProcessesAbsentError
	}

	d.SetCurrentProcess(0)
	// TODO: Call Delay
	go d.AddDelay(db.Delay{DelayTime: recipe.TotalTime}, true)
	return
}

// func(d *Compact32Deck) sendWSData(recipeID uuid.UUID, processLength, currentStep int, processName string, processType db.ProcessType) {
func (d *Compact32Deck) sendWSData(time1 time.Time, timeElapsed *int64, delayTime int64, ops WebsocketOperation) (err error) {
	var wsProgressOp WSData
	var wsData []byte

	opTime := time.Now()
	opTimePassed := int64(opTime.Sub(time1).Seconds()) + *timeElapsed
	opRemainingTime := delayTime - opTimePassed
	if opTimePassed > delayTime {
		opTimePassed = delayTime
		opRemainingTime = 0
	}
	progress := (opTimePassed * 100) / delayTime

	wsProgressOp = WSData{
		Progress: float64(progress),
		Deck:     d.name,
		Status:   string(ops),
		OperationDetails: OperationDetails{
			RemainingTime: opRemainingTime,
			TotalTime:     delayTime,
		},
	}

	switch ops {
	case uvlightProgress:
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("uv light cleanup in progress for deck %s ", d.name)
	case uvlightSuccess:
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("successfully completed UV Light clean up for deck %v", d.name)
	case recipeProgress:
		currentStep := d.getCurrentProcess()

		wsProgressOp.OperationDetails.Message = fmt.Sprintf("process %v for deck %v in progress", currentStep+1, d.name)
		wsProgressOp.OperationDetails.CurrentStep = currentStep + 1
		wsProgressOp.OperationDetails.RecipeID = deckRecipe[d.name].ID
		wsProgressOp.OperationDetails.TotalProcesses = deckRecipe[d.name].ProcessCount
		wsProgressOp.OperationDetails.ProcessName = deckProcesses[d.name][currentStep].Name
		wsProgressOp.OperationDetails.ProcessType = deckProcesses[d.name][currentStep].Type

	case recipeSuccess:
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("process %v for deck %v completed", deckRecipe[d.name].ProcessCount, d.name)
		wsProgressOp.OperationDetails.CurrentStep = deckRecipe[d.name].ProcessCount
		wsProgressOp.OperationDetails.RecipeID = deckRecipe[d.name].ID
		wsProgressOp.OperationDetails.TotalProcesses = deckRecipe[d.name].ProcessCount

	default:
		return responses.InvalidOperationWebsocket
	}

	wsData, err = json.Marshal(wsProgressOp)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.WebsocketMarshallingError)
		d.WsErrCh <- err
		return err
	}
	d.WsMsgCh <- fmt.Sprintf("%v_%v", ops, string(wsData))

	return
}
