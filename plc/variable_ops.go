package plc

import (
	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) NameOfDeck() string {
	return d.Name
}

func (d *Compact32Deck) SetRunInProgress() {
	runInProgress.Store(d.Name, true)
}

func (d *Compact32Deck) ResetRunInProgress() {
	runInProgress.Store(d.Name, false)
}

func (d *Compact32Deck) setTimerInProgress() {
	timerInProgress.Store(d.Name, true)
}

func (d *Compact32Deck) resetTimerInProgress() {
	timerInProgress.Store(d.Name, false)
}

func (d *Compact32Deck) setHeaterInProgress() {
	heaterInProgress.Store(d.Name, true)
}

func (d *Compact32Deck) resetHeaterInProgress() {
	heaterInProgress.Store(d.Name, false)
}

func (d *Compact32Deck) setShakerInProgress() {
	shakerInProgress.Store(d.Name, true)
}

func (d *Compact32Deck) resetShakerInProgress() {
	shakerInProgress.Store(d.Name, false)
}

func (d *Compact32Deck) setAborted() {
	aborted.Store(d.Name, true)
}

func (d *Compact32Deck) resetAborted() {
	aborted.Store(d.Name, false)
}

func (d *Compact32Deck) setPaused() {
	paused.Store(d.Name, true)
}

func (d *Compact32Deck) resetPaused() {
	paused.Store(d.Name, false)
}

func (d *Compact32Deck) setHomed() {
	homed.Store(d.Name, true)
}

func (d *Compact32Deck) resetHomed() {
	homed.Store(d.Name, false)
}

func (d *Compact32Deck) setUVLightInProgress() {
	uvLightInProgress.Store(d.Name, true)
}

func (d *Compact32Deck) resetUVLightInProgress() {
	uvLightInProgress.Store(d.Name, false)
}

func (d *Compact32Deck) setHomingPercent(percent float64) {
	homingPercent.Store(d.Name, percent)
}

func (d *Compact32Deck) IsMachineHomed() bool {
	if temp, ok := homed.Load(d.Name); !ok {
		logger.Errorln("homed isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) IsRunInProgress() bool {
	if temp, ok := runInProgress.Load(d.Name); !ok {
		logger.Errorln("runInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isTimerInProgress() bool {
	if temp, ok := timerInProgress.Load(d.Name); !ok {
		logger.Errorln("timerInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isMachineInAbortedState() bool {
	if temp, ok := aborted.Load(d.Name); !ok {
		logger.Errorln("aborted isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isMachineInPausedState() bool {
	if temp, ok := paused.Load(d.Name); !ok {
		logger.Errorln("paused isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isHeaterInProgress() bool {
	if temp, ok := heaterInProgress.Load(d.Name); !ok {
		logger.Errorln("heaterInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isShakerInProgress() bool {
	if temp, ok := heaterInProgress.Load(d.Name); !ok {
		logger.Errorln("shakerInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}
func (d *Compact32Deck) isUVLightInProgress() bool {
	if temp, ok := uvLightInProgress.Load(d.Name); !ok {
		logger.Errorln("uvLightInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) getMagnetState() int {
	if temp, ok := magnetState.Load(d.Name); !ok {
		logger.Errorln("magnet State isn't loaded!")
		return -1
	} else {
		return temp.(int)
	}
}

func (d *Compact32Deck) getSyringeModuleState() int {
	if temp, ok := syringeModuleState.Load(d.Name); !ok {
		logger.Errorln("Syringe Module State isn't loaded!")
		return -1
	} else {
		return temp.(int)
	}
}

func (d *Compact32Deck) getExecutedPulses() uint16 {
	if temp, ok := executedPulses.Load(d.Name); !ok {
		logger.Errorln("executed Pulses isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getWrotePulses() uint16 {
	if temp, ok := wrotePulses.Load(d.Name); !ok {
		logger.Errorln("wrote Pulses isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getMotorNumReg() uint16 {
	if temp, ok := motorNumReg.Load(d.Name); !ok {
		logger.Errorln("motorNumReg isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getSpeedReg() uint16 {
	if temp, ok := speedReg.Load(d.Name); !ok {
		logger.Errorln("speed Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getDirectionReg() uint16 {
	if temp, ok := directionReg.Load(d.Name); !ok {
		logger.Errorln("direction Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getRampReg() uint16 {
	if temp, ok := rampReg.Load(d.Name); !ok {
		logger.Errorln("ramp Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getPulseReg() uint16 {
	if temp, ok := pulseReg.Load(d.Name); !ok {
		logger.Errorln("pulse Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getOnReg() uint16 {
	if temp, ok := onReg.Load(d.Name); !ok {
		logger.Errorln("onReg isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getHomingDeckName() string {
	if BothDeckHomingInProgress {
		return ""
	}
	return d.Name
}

func (d *Compact32Deck) getHomingPercent() float64 {
	if BothDeckHomingInProgress {
		if tempA, ok := homingPercent.Load("A"); !ok {
			logger.Errorln("homingPercent isn't loaded!")
			return -1
		} else if tempB, ok := homingPercent.Load("B"); !ok {
			logger.Errorln("homingPercent isn't loaded!")
			return -1
		} else {
			return (tempA.(float64) + tempB.(float64)) / 2
		}
	}
	if temp, ok := homingPercent.Load(d.Name); !ok {
		logger.Errorln("homingPercent isn't loaded!")
		return -1
	} else {
		return temp.(float64)
	}
}

func SetBothDeckHomingInProgress() {
	BothDeckHomingInProgress = true
}

func ResetBothDeckHomingInProgress() {
	BothDeckHomingInProgress = false
}

func IsBothDeckHomingInProgress() bool {
	return BothDeckHomingInProgress
}
