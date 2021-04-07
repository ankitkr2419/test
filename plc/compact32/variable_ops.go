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

func (d *Compact32Deck) setTimerInProgress() {
	timerInProgress.Store(d.name, true)
}

func (d *Compact32Deck) resetTimerInProgress() {
	timerInProgress.Store(d.name, false)
}

func (d *Compact32Deck) setHeaterInProgress() {
	heaterInProgress.Store(d.name, true)
}

func (d *Compact32Deck) resetHeaterInProgress() {
	heaterInProgress.Store(d.name, false)
}

func (d *Compact32Deck) setShakerInProgress() {
	shakerInProgress.Store(d.name, true)
}

func (d *Compact32Deck) resetShakerInProgress() {
	shakerInProgress.Store(d.name, false)
}

func (d *Compact32Deck) setUVLightInProgress() {
	uvLightInProgress.Store(d.name, true)
}

func (d *Compact32Deck) resetUVLightInProgress() {
	uvLightInProgress.Store(d.name, false)
}

func (d *Compact32Deck) IsMachineHomed() bool {
	if temp, ok := homed.Load(d.name); !ok {
		logger.Errorln("homed isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) IsRunInProgress() bool {
	if temp, ok := runInProgress.Load(d.name); !ok {
		logger.Errorln("runInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isTimerInProgress() bool {
	if temp, ok := timerInProgress.Load(d.name); !ok {
		logger.Errorln("timerInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isMachineInAbortedState() bool {
	if temp, ok := aborted.Load(d.name); !ok {
		logger.Errorln("aborted isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isMachineInPausedState() bool {
	if temp, ok := paused.Load(d.name); !ok {
		logger.Errorln("paused isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isHeaterInProgress() bool {
	if temp, ok := heaterInProgress.Load(d.name); !ok {
		logger.Errorln("heaterInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) isShakerInProgress() bool {
	if temp, ok := heaterInProgress.Load(d.name); !ok {
		logger.Errorln("shakerInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}
func (d *Compact32Deck) isUVLightInProgress() bool {
	if temp, ok := uvLightInProgress.Load(d.name); !ok {
		logger.Errorln("uvLightInProgress isn't loaded!")
	} else if temp.(bool) {
		return true
	}
	return false
}

func (d *Compact32Deck) getMagnetState() int {
	if temp, ok := magnetState.Load(d.name); !ok {
		logger.Errorln("magnet State isn't loaded!")
		return -1
	} else {
		return temp.(int)
	}
}

func (d *Compact32Deck) getSyringeModuleState() int {
	if temp, ok := syringeModuleState.Load(d.name); !ok {
		logger.Errorln("Syringe Module State isn't loaded!")
		return -1
	} else {
		return temp.(int)
	}
}

func (d *Compact32Deck) getExecutedPulses() uint16 {
	if temp, ok := executedPulses.Load(d.name); !ok {
		logger.Errorln("executed Pulses isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getWrotePulses() uint16 {
	if temp, ok := wrotePulses.Load(d.name); !ok {
		logger.Errorln("wrote Pulses isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getMotorNumReg() uint16 {
	if temp, ok := motorNumReg.Load(d.name); !ok {
		logger.Errorln("motorNumReg isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getSpeedReg() uint16 {
	if temp, ok := speedReg.Load(d.name); !ok {
		logger.Errorln("speed Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getDirectionReg() uint16 {
	if temp, ok := directionReg.Load(d.name); !ok {
		logger.Errorln("direction Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getRampReg() uint16 {
	if temp, ok := rampReg.Load(d.name); !ok {
		logger.Errorln("ramp Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getPulseReg() uint16 {
	if temp, ok := pulseReg.Load(d.name); !ok {
		logger.Errorln("pulse Register isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}

func (d *Compact32Deck) getOnReg() uint16 {
	if temp, ok := onReg.Load(d.name); !ok {
		logger.Errorln("onReg isn't loaded!")
		return highestUint16
	} else {
		return temp.(uint16)
	}
}
