package compact32

import (
	"fmt"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) DiscardTipAndHome(discard bool) (response string, err error) {

	var messageWithTipDiscard string
	// Machine Should be in aborted state
	if !d.isMachineInAbortedState() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}

	d.resetAborted()

	if discard {
		response, err = d.tipDiscard()
		if err != nil {
			logger.Errorln("error in discarding tip for deck --->", d.name)
			return
		}
		messageWithTipDiscard = "tip discarded and "
	}

	//home
	response, err = d.Homing()
	if err != nil {
		logger.Errorln("error in homing for deck --->", d.name)
		return
	}

	d.WsMsgCh <- fmt.Sprintf("success_discardAndHomed_%vmachine homed successfully", messageWithTipDiscard)

	return "SUCCESS", nil
}
