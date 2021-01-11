package compact32

import (
	"encoding/binary"
	"fmt"
)

func (d *Compact32Deck) ManualMovement(motorNum, direction, pulses uint16) (response string, err error) {
	sensorHasCut[d.name] = false
	aborted[d.name] = false

	return d.SetupMotor(uint16(2000), pulses, uint16(100), direction, motorNum)
}

func (d *Compact32Deck) NameOfDeck() string {
	return d.name
}

func (d *Compact32Deck) Pause() (response string, err error) {

	results, err := d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][0], uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Printf("Read On/Off Coil. res : %+v \n", results)

	var resultsInt int
	if len(results) > 0 {
		resultsInt = int(results[0])
	}

	if resultsInt == 0 {
		err = fmt.Errorf("system is already in paused state")
		return "", err
	}

	response, err = d.SwitchOffMotor()
	if err != nil {
		return "", err
	}

	return "Motor PAUSED Successfully", nil

}

func (d *Compact32Deck) Resume() (response string, err error) {

	// if already on then throw error
	if d.IsMotorOff() == false {
		err = fmt.Errorf("System is already running")
		return "", err
	}

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

	return "RESUMED Successfully.", nil
}

func (d *Compact32Deck) Abort() (response string, err error) {

	fmt.Println("aborting the operation....")

	fmt.Println("switching motor off....")
	response, err = d.SwitchOffMotor()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Write 0 Pulses
	aborted[d.name] = true
	executedPulses[d.name] = 0
	wrotePulses[d.name] = 0
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Simulation : homing the machine.")

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
	fmt.Println("Wrote Pulses ---> ", wrotePulses)

	results, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][202], uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Printf("read ReadHoldingRegisters_Pulse : %+v \n", binary.BigEndian.Uint16(results))

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	return "RESUMED with pulses.", nil
}
