package compact32

import (
	"encoding/binary"
	"fmt"
)

func (d *Compact32Deck) ManualMovement(motorNum, direction, pulses uint16) (response string, err error) {

	if runInProgress[d.name] {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}
	sensorHasCut[d.name] = false
	aborted[d.name] = false
	runInProgress[d.name] = true
	defer d.ResetRunInProgress()

	response, err = d.SetupMotor(uint16(2000), pulses, uint16(100), direction, motorNum)
	if err != nil {
		return "", fmt.Errorf("there was some issue doing manual movement")
	}

	return
}

func (d *Compact32Deck) NameOfDeck() string {
	return d.name
}

func (d *Compact32Deck) SetRunInProgress() {
	runInProgress[d.name] = true
}

func (d *Compact32Deck) ResetRunInProgress() {
	runInProgress[d.name] = false
}

func (d *Compact32Deck) SetTimerInProgress() {
	timerInProgress[d.name] = true
}

func (d *Compact32Deck) ResetTimerInProgress() {
	timerInProgress[d.name] = false
}

func (d *Compact32Deck) Pause() (response string, err error) {

	// If machine is already PAUSED OR
	// run is not in Progress
	if paused[d.name] || !runInProgress[d.name] {
		fmt.Println("err : ", err)
		err = fmt.Errorf("Machine is already in PAUSED/IDLE state")
		return
	}

	if !timerInProgress[d.name] {
		response, err = d.SwitchOffMotor()
		if err != nil {
			return "", err
		}
	}

	paused[d.name] = true

	return "Operation PAUSED Successfully", nil
}

func (d *Compact32Deck) Resume() (response string, err error) {

	// if paused only then resume
	if !paused[d.name] {
		err = fmt.Errorf("System is already running, or done with the run")
		return "", err
	}

	if !timerInProgress[d.name] {
		response, err = d.ReadExecutedPulses()
		if err != nil {
			fmt.Println("err : ", err)
			return "", err
		}

		if int(wrotePulses[d.name]) <= int(executedPulses[d.name]) {
			err = fmt.Errorf("executedPulses is greater than wrote Pulses that means nothing to resume.")
			wrotePulses[d.name] = 0
			executedPulses[d.name] = 0

			return "", err
		}

		response, err = d.ResumeMotorWithPulses(wrotePulses[d.name] - executedPulses[d.name])
		if err != nil {
			fmt.Println(response, "resumeMotorWithPulse")
			return
		}
	}

	paused[d.name] = false

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

	aborted[d.name] = true
	wrotePulses[d.name] = 0
	paused[d.name] = false

	// If no runInProgress and timer is in progress, that means no need to read pulses
	if runInProgress[d.name] && !timerInProgress[d.name] {
		response, err = d.ReadExecutedPulses()
		if err != nil {
			fmt.Println("err : ", err)
			return "", fmt.Errorf("Operation is ABORTED but current position was lost, please home the machine")
		}
	}
	runInProgress[d.name] = false
	timerInProgress[d.name] = false

	return "ABORT SUCCESS", nil
}

func (d *Compact32Deck) ResumeMotorWithPulses(pulses uint16) (response string, err error) {

	results, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][202], pulses)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Pulse. res : ", binary.BigEndian.Uint16(results))
	wrotePulses[d.name] = pulses
	fmt.Println("Wrote Pulses ---> ", wrotePulses[d.name])

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
