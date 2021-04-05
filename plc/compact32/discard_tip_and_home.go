package compact32

import (
	"fmt"
	"math"
	"time"
)

func (d *Compact32Deck) DiscardTipAndHome() (response string, err error) {

	// Machine Should be in aborted state
	if d.isMachineInAbortedState() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}

}