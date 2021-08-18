package plc

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"

	"time"
)

type WebsocketOperation string

const (
	heater          WebsocketOperation = "HEATER_HEATER"
	uvlightProgress WebsocketOperation = "PROGRESS_UVLIGHT"
	pidProgress     WebsocketOperation = "PROGRESS_PID"
	recipeProgress  WebsocketOperation = "PROGRESS_RECIPE"
	uvlightSuccess  WebsocketOperation = "SUCCESS_UVLIGHT"
	pidSuccess      WebsocketOperation = "SUCCESS_PID"
	recipeSuccess   WebsocketOperation = "SUCCESS_RECIPE"
)

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
			RemainingTime: ConvertToHMS(opRemainingTime),
			TotalTime:     ConvertToHMS(delayTime),
		},
	}

	switch ops {
	case uvlightProgress:
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("uv light cleanup in progress for deck %s ", d.name)
	case uvlightSuccess:
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("successfully completed UV Light clean up for deck %v", d.name)
	case pidProgress:
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("pid tuning in progress for deck %s ", d.name)
	case pidSuccess:
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("successfully completed pid tuning for deck %v", d.name)
	case recipeProgress:
		currentStep := getCurrentProcessNumber(d.name)

		wsProgressOp.OperationDetails.RecipeID = deckRecipe[d.name].ID
		wsProgressOp.OperationDetails.TotalProcesses = deckRecipe[d.name].ProcessCount

		// This means all Processes are completed!!
		// But Our Recipe Remaining time was still working!!
		if currentStep == -2 {
			wsProgressOp.OperationDetails.CurrentStep = deckRecipe[d.name].ProcessCount
			wsProgressOp.OperationDetails.Message = fmt.Sprintf("process %v for deck %v in progress", deckRecipe[d.name].ProcessCount, d.name)
			wsProgressOp.OperationDetails.RemainingTime = ConvertToHMS(0)
			wsProgressOp.Progress = 100
			defer d.sendWSData(time1, timeElapsed, delayTime, recipeSuccess)
			break
		} else if currentStep == -1 {
			err = responses.InvalidCurrentStep
			return
		}

		wsProgressOp.OperationDetails.Message = fmt.Sprintf("process %v for deck %v in progress", currentStep+1, d.name)
		wsProgressOp.OperationDetails.CurrentStep = currentStep + 1
		wsProgressOp.OperationDetails.ProcessName = deckProcesses[d.name][currentStep].Name
		wsProgressOp.OperationDetails.ProcessType = deckProcesses[d.name][currentStep].Type

	case recipeSuccess:
		wsProgressOp.Progress = 100
		wsProgressOp.OperationDetails.Message = fmt.Sprintf("process %v for deck %v completed", deckRecipe[d.name].ProcessCount, d.name)
		wsProgressOp.OperationDetails.CurrentStep = deckRecipe[d.name].ProcessCount
		wsProgressOp.OperationDetails.RecipeID = deckRecipe[d.name].ID
		wsProgressOp.OperationDetails.TotalProcesses = deckRecipe[d.name].ProcessCount
		wsProgressOp.OperationDetails.RemainingTime = ConvertToHMS(0)

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

func (d *Compact32Deck) sendHeaterData() (err error) {
	hData := HeaterData{
		Deck:     d.name,
		HeaterOn: d.isHeaterInProgress(),
	}

	var wsData []byte

	hData.Shaker1Temp, hData.Shaker2Temp, err = d.readTempValues()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.FetchHeaterTempError)
		return err
	}

	wsData, err = json.Marshal(hData)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.WebsocketMarshallingError)
		d.WsErrCh <- err
		return err
	}
	d.WsMsgCh <- fmt.Sprintf("%v_%v", heater, string(wsData))

	return
}

func ConvertToHMS(secs int64) *TimeHMS {
	var t TimeHMS
	t.Hours = uint8(secs / (60 * 60))
	t.Minutes = uint8(secs/60 - int64(t.Hours)*60)
	t.Seconds = uint8(secs % 60)
	logger.Debugln("Converted time: ", t)
	return &t
}
