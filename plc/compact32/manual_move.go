package compact32

import (
	"encoding/binary"
	"fmt"
)

func (d *Compact32Deck) ManualMovement(motorNum, direction, pulses uint16) (response string, err error) {
	sensorHasCut = false
	aborted = false
	// var sensorAddressBytes = []byte{0x08, 0x02}
	// sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	return d.SetupMotor(uint16(2000), pulses, uint16(100), direction, motorNum)
}

func (d *Compact32Deck) Pause() (response string, err error) {
	// var onOffAddressBytes = []byte{0x08, 0x00}
	onOffAddressUint16 := MODBUS_EXTRACTION["A"]["M"][0] //binary.BigEndian.Uint16(onOffAddressBytes)

	results, err := d.DeckDriver.ReadCoils(onOffAddressUint16, uint16(1))
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

	// Read D800 here and get rid of go routine
	return "Motor PAUSED Successfully", nil

}

func (d *Compact32Deck) Resume() (response string, err error) {

	//m.ResumeMotorWithPulses(wrotePulses - completedPulses)

	// if already on then throw error
	if d.IsMotorOff() == false {
		err = fmt.Errorf("System is already running")
		return "", err
	}

	response, err = d.ReadD2000()
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	if int(wrotePulses) <= int(completedPulses) {
		err = fmt.Errorf("CompletedPulses is greater than wrote Pulses that means nothing to resume.")
		wrotePulses = 0
		completedPulses = 0

		return "", err
	}

	response, err = d.ResumeMotorWithPulses(wrotePulses - completedPulses)
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
	//response, err = m.ResumeMotorWithPulses(uint16(1))
	aborted = true
	completedPulses = 0
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// Just go to homing position.
	// For now just simulating
	fmt.Println("Simulation : homing the machine.")

	return "ABORT SUCCESS", nil
}

func (d *Compact32Deck) ResumeMotorWithPulses(pulses uint16) (response string, err error) {

	// Write Pulses
	//var pulseAddressBytes = []byte{0x10, 0xCA}
	pulseAddressUint16 := MODBUS_EXTRACTION["A"]["M"][202] //binary.BigEndian.Uint16(pulseAddressBytes)

	results, err := d.DeckDriver.WriteSingleRegister(pulseAddressUint16, pulses)
	//results, err = m.Client.WriteSingleRegister(pulseAddressUint16, uint16(500))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Pulse. res : ", binary.BigEndian.Uint16(results))
	wrotePulses = pulses
	fmt.Println("Wrote Pulses ---> ", wrotePulses)

	results, err = d.DeckDriver.ReadHoldingRegisters(pulseAddressUint16, uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Printf("read ReadHoldingRegisters_Pulse : %+v \n", binary.BigEndian.Uint16(results))

	// var onOffAddressBytes = []byte{0x08, 0x00}
	onOffAddressUint16 := MODBUS_EXTRACTION["A"]["M"][0] //binary.BigEndian.Uint16(onOffAddressBytes)

	//var completionAddressBytes = []byte{0x08, 0x01}
	//completionAddressUint16 := binary.BigEndian.Uint16(completionAddressBytes)

	//results, err = m.Client.WriteSingleCoil(completionAddressUint16, uint16(0x0000))
	////results, err = m.Client.WriteSingleCoil(onOffAddressUint16, uint16(0xff00))
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return "", err
	//}
	//fmt.Printf("Wrote Off Completion Forcefully. res : %+v \n", results)

	err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, ON)
	//results, err = m.Client.WriteSingleCoil(onOffAddressUint16, uint16(0xff00))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Printf("Wrote On motor Forcefully. res : %+v \n", results)

	return "RESUMED with pulses.", nil

}
