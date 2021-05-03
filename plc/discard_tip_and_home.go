package plc

import (
	"encoding/json"
	"fmt"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) DiscardTipAndHome(discard bool) (response string, err error) {

	//Machine Should be in aborted state
	if !d.isMachineInAbortedState() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		logger.Errorln("previous run already in progress... wait or abort it", d.name)
		return "", err
	}

	d.resetAborted()

	if discard {
		response, err = d.tipDiscard()
		if err != nil {
			logger.Errorln("error in discarding tip for deck --->", d.name)
			return
		}
	}

	//home
	response, err = d.Homing()
	if err != nil {
		logger.Errorln("error in homing for deck --->", d.name)
		return
	}

	// send success ws data
	successWsData := WSData{
		Progress: 100,
		Deck:     d.name,
		Status:   "SUCCESS_DISCARDANDHOMED",
		OperationDetails: OperationDetails{
			Message: fmt.Sprintf("successfully completed tip discard and/or homing for deck %v", d.name),
		},
	}
	wsData, err := json.Marshal(successWsData)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
		d.WsErrCh <- err
		return
	}
	d.WsMsgCh <- fmt.Sprintf("success_discardAndHomed_%v", string(wsData))

	return "SUCCESS", nil
}
