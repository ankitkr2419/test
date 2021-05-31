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
	uvlightProgress WebsocketOperation = "PROGRESS_UVLIGHT"
	recipeProgress  WebsocketOperation = "PROGRESS_RECIPE"
	uvlightSuccess  WebsocketOperation = "SUCCESS_UVLIGHT"
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
			RemainingTime: convertToHMS(opRemainingTime),
			TotalTime:     convertToHMS(delayTime),
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

func convertToHMS(secs int64) *TimeHMS {
	var t TimeHMS
	t.Hours = uint8(secs / (60 * 60))
	t.Minutes = uint8(secs/60 - int64(t.Hours)*60)
	t.Seconds = uint8(secs % 60)
	logger.Infoln("Converted time: ", t)
	return &t
}