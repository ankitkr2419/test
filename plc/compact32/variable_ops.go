package compact32

import (
	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) NameOfDeck() string {
	return d.name
}

func (d *Compact32Deck) SetRunInProgress() {
	runInProgress.Store(d.name, true)
}

func (d *Compact32Deck) ResetRunInProgress() {
	runInProgress.Store(d.name, false)
}

func (d *Compact32Deck) SetTimerInProgress() {
	timerInProgress.Store(d.name, true)
}

func (d *Compact32Deck) ResetTimerInProgress() {
	timerInProgress.Store(d.name, false)
}

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

func (d *Compact32Deck) IsMachineInAbortedState() bool{
	if temp, ok := aborted.Load(d.name); !ok {
		logger.Errorln("aborted isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) IsMachineInPausedState() bool{
	if temp, ok := paused.Load(d.name); !ok {
		logger.Errorln("paused isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}
