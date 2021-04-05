package compact32

import (
	"fmt"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) DiscardTipAndHome(discard bool) (response string, err error) {

	// Machine Should be in aborted state
	if !d.isMachineInAbortedState() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}

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

	d.WsMsgCh <- "success_discardAndHomed_tips discarded successfully and machine homed"

	return "SUCCESS", nil
}
