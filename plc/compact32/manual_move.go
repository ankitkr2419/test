package compact32

import (
	"encoding/binary"
	"fmt"
)

func (d *Compact32Deck) ManualMovement(motorNum, direction, pulses uint16) (response string, err error) {

	if temp, ok := runInProgress.Load(d.name); !ok {
		err = fmt.Errorf("runInProgress isn't loaded!")
		return
	} else if temp.(bool) {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}
	aborted.Store(d.name, false)
	runInProgress.Store(d.name, true)
	defer d.ResetRunInProgress()

	response, err = d.SetupMotor(uint16(2000), pulses, uint16(100), direction, motorNum)
	if err != nil {
		return "", fmt.Errorf("there was some issue doing manual movement")
	}

	return
}

func (d *Compact32Deck) Pause() (response string, err error) {

	// If machine is already PAUSED OR
	if temp, ok := paused.Load(d.name); !ok {
		err = fmt.Errorf("paused isn't loaded!")
		return
	} else if temp.(bool) {
		err = fmt.Errorf("Machine is already in PAUSED state")
		return "", err
	}

	// run is not in Progress
	if temp, ok := runInProgress.Load(d.name); !ok {
		err = fmt.Errorf("runInProgress isn't loaded!")
		return
	} else if !temp.(bool) {
		err = fmt.Errorf("Machine is already in IDLE state")
		return "", err
	}

	if temp, ok := timerInProgress.Load(d.name); !ok {
		err = fmt.Errorf("timerInProgress isn't loaded!")
		return
	} else if !temp.(bool) {
		response, err = d.SwitchOffMotor()
		if err != nil {
			return "", err
		}
	}

	paused.Store(d.name, true)

	return "Operation PAUSED Successfully", nil
}

func (d *Compact32Deck) Resume() (response string, err error) {

	// if paused only then resume
	if temp, ok := paused.Load(d.name); !ok {
		err = fmt.Errorf("paused isn't loaded!")
		return
	} else if !temp.(bool) {
		err = fmt.Errorf("System is already running, or done with the run")
		return "", err
	}

	if temp, ok := timerInProgress.Load(d.name); !ok {
		err = fmt.Errorf("timerInProgress isn't loaded!")
		return
	} else if !temp.(bool) {
		response, err = d.ReadExecutedPulses()
		if err != nil {
			fmt.Println("err : ", err)
			return "", err
		}

		if temp1, ok := wrotePulses.Load(d.name); !ok {
			err = fmt.Errorf("wrotePulses isn't loaded!")
		} else if temp2, ok := executedPulses.Load(d.name); !ok {
			err = fmt.Errorf("executedPulses isn't loaded!")
		} else if temp1.(int) <= temp2.(int) {
			err = fmt.Errorf("executedPulses is greater than wrote Pulses that means nothing to resume.")
			wrotePulses.Store(d.name, 0)
			executedPulses.Store(d.name, 0)
		} else {
			// calculating wrotePulses.[d.name] - executedPulses.[d.name]
			response, err = d.ResumeMotorWithPulses(uint16(temp1.(int) - temp2.(int)))
		}
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}

	paused.Store(d.name, false)

	return "Operation RESUMED Successfully.", nil
}

func (d *Compact32Deck) Abort() (response string, err error) {

	fmt.Println("aborting the operation....")

	fmt.Println("switching motor off....")
	response, err = d.SwitchOffMotor()
	if err != nil {
		fmt.Println("From deck ", d.name, err)
		return "", err
	}

	response, err = d.SwitchOffHeater()
	if err != nil {
		fmt.Println("From deck ", d.name, err)
		return "", err
	}

	//  Switch off UV Light
	response, err = d.switchOffUVLight()
	if err != nil {
		return
	}

	aborted.Store(d.name, true)
	wrotePulses.Store(d.name, 0)
	paused.Store(d.name, false)
	homed.Store(d.name, false)

	// If runInProgress and no timer is in progress, that means we need to read pulses

	if temp1, ok := runInProgress.Load(d.name); !ok {
		err = fmt.Errorf("runInProgress isn't loaded!")
		return
	} else if temp2, ok := timerInProgress.Load(d.name); !ok {
		err = fmt.Errorf("timerInProgress isn't loaded!")
		return
	} else if temp1.(bool) && !temp2.(bool) {
		response, err = d.ReadExecutedPulses()
		if err != nil {
			fmt.Println("err : ", err)
			return "", fmt.Errorf("Operation is ABORTED but current position was lost, please home the machine")
		}
	}

	runInProgress.Store(d.name, false)
	timerInProgress.Store(d.name, false)

	return "ABORT SUCCESS", nil
}

func (d *Compact32Deck) ResumeMotorWithPulses(pulses uint16) (response string, err error) {

	results, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][202], pulses)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Pulse. res : ", binary.BigEndian.Uint16(results))
	wrotePulses.Store(d.name, pulses)
	fmt.Println("Wrote Pulses ---> ", pulses)

	results, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][202], uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Printf("read ReadHoldingRegisters_Pulse : %+v \n", binary.BigEndian.Uint16(results))

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	return "RESUMED with pulses.", nil
}
