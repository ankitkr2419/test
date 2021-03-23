package compact32

import (
	"fmt"
	"mylab/cpagent/db"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) AddDelay(delay db.Delay) (reponse string, err error) {
	var t *time.Timer
	t = time.AfterFunc(delay.DelayTime, func() {
		logger.Printf("Delay time over")
		return
	})
	for {
		select {
		case n := <-t.C:
			fmt.Printf("delay time over %v", n)
			return "SUCCESS", nil
		default:
			if d.isMachineInAbortedState() {
				err = fmt.Errorf("Operation was ABORTED!")
				return "", err
			}
			// delay of 300 ms for checking the delay over time to avoid too much loop
			time.Sleep(time.Millisecond * 300)
		}
	}

}
