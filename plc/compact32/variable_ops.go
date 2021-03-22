package compact32

import (
	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) IsMachineHomed() bool{
	if temp, ok := homed.Load(d.name); !ok {
		logger.Errorln("homed isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}


func (d *Compact32Deck) IsRunInProgress() bool{
	if temp, ok := runInProgress.Load(d.name); !ok {
		logger.Errorln("runInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}


func (d *Compact32Deck) IsTimerInProgress() bool{
	if temp, ok := timerInProgress.Load(d.name); !ok {
		logger.Errorln("timerInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}
