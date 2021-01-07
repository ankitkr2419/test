package compact32

import (
	"encoding/binary"
	"fmt"
	"time"
)

const (
	UP   = uint16(1)
	DOWN = uint16(0)
	FWD  = uint16(1)
	REV  = uint16(0)
)

// *** NOTE ***
// For Syringe UP means DOWN and DOWN means UP
// This is because of hardware compatibility
// 1 means towards sensor, 0 means against sensor
// ************

var wrotePulses uint16 = 0
var completedPulses uint16 = 0
var sensorHasCut = false
var aborted = false
var statusChannel2 = make(chan int)

func (d *Compact32Deck) Check() bool {
	return true
}

func (d *Compact32Deck) SetupMotor(speed, pulse, ramp, direction, motorNum uint16) (response string, err error) {
	//var choice int = 0

	if aborted {
		return
	}

	var results []byte
	statusChannel1 := make(chan int)

	//var wg sync.WaitGroup

	var speedAddressBytes = []byte{0x10, 0xC8}
	speedAddressUint16 := binary.BigEndian.Uint16(speedAddressBytes)

	var pulseAddressBytes = []byte{0x10, 0xCA}
	pulseAddressUint16 := binary.BigEndian.Uint16(pulseAddressBytes)

	var rampAddressBytes = []byte{0x10, 0xCC}
	rampAddressUint16 := binary.BigEndian.Uint16(rampAddressBytes)

	var directionAddressBytes = []byte{0x10, 0xCE}
	directionAddressUint16 := binary.BigEndian.Uint16(directionAddressBytes)

	var motorNumAddressBytes = []byte{0x10, 0xE2}
	motorNumAddressUint16 := binary.BigEndian.Uint16(motorNumAddressBytes)

	var onOffAddressBytes = []byte{0x08, 0x00}
	onOffAddressUint16 := binary.BigEndian.Uint16(onOffAddressBytes)

	var completionAddressBytes = []byte{0x08, 0x01}
	completionAddressUint16 := binary.BigEndian.Uint16(completionAddressBytes)

	sensorAddressUint16 := MODBUS_EXTRACTION["A"]["M"][2]

	// M start with 08
	//fmt.Println("2. Off, anyother is On")
	//
	//fmt.Scanln(&choice)
	//switch choice {
	//case 2:
	//	results, err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, onOff)
	//	//results, err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, uint16(0x0000))
	//default:
	//	results, err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, onOff)
	//	//results, err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, uint16(0xff00))
	//}

	// OFF The motor
	err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, uint16(0x0000))

	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Printf("Wrote Off motor. res : %+v \n", results)

	// results, err = d.DeckDriver.WriteSingleCoil(completionAddressUint16, completion)
	err = d.DeckDriver.WriteSingleCoil(completionAddressUint16, uint16(0x0000))
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Printf("Wrote Completion. res : %+v \n", results)

	//time.Sleep(1 * time.Second)

	// D start with 1

	results, err = d.DeckDriver.WriteSingleRegister(pulseAddressUint16, pulse)
	//results, err = d.DeckDriver.WriteSingleRegister(pulseAddressUint16, uint16(500))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Pulse. res : ", results)
	wrotePulses = pulse

	results, err = d.DeckDriver.WriteSingleRegister(speedAddressUint16, speed)
	//results, err = d.DeckDriver.WriteSingleRegister(speedAddressUint16, uint16(1000))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Speed. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(rampAddressUint16, ramp)
	//results, err = d.DeckDriver.WriteSingleRegister(rampAddressUint16, uint16(1000))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Ramp. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(directionAddressUint16, direction)
	//results, err = d.DeckDriver.WriteSingleRegister(dirAddressUint16, uint16(1000))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote direction. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(motorNumAddressUint16, motorNum)
	//results, err = d.DeckDriver.WriteSingleRegister(motorNumAddressUint16, uint16(1000))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote motorNum. res : ", results)

	err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, ON)
	//results, err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, uint16(0xff00))

	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Printf("Wrote On/Off motor. res : %+v \n", results)

	results, err = d.DeckDriver.ReadCoils(onOffAddressUint16, uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Printf("Read On/Off Coil. res : %+v \n", results)

	/*
		results, err = d.DeckDriver.WriteSingleRegister(motorNumAddressUint16, speed)
		//results, err = d.DeckDriver.WriteSingleRegister(motorNumAddressUint16, uint16(1000))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("res : ", results)
	*/
	go func() {
		//count := 0
		for {
			if sensorHasCut || aborted {
				statusChannel1 <- 2
				return
			}
			//count++;
			results, err = d.DeckDriver.ReadCoils(completionAddressUint16, uint16(1))
			if err != nil {
				fmt.Println("err : ", err)
			}
			if len(results) > 0 {
				//fmt.Println("Read Completion. res :", results[0])
				if int(results[0]) == 1 {
					fmt.Println("Completion returned ---> ", results)
					statusChannel1 <- 1
					return
				}
			}

			if direction == uint16(0) {
				goto skipSensor
			}
			results, err = d.DeckDriver.ReadCoils(sensorAddressUint16, uint16(1))

			if err != nil {
				fmt.Println("err : ", err)
			}
			fmt.Println("Sensor returned ---> ", results)
			if len(results) > 0 {
				//fmt.Println("Read Completion. res :", results[0])
				if int(results[0]) == 3 && pulse != uint16(19999) {
					fmt.Println("Sensor returned ---> ", results[0])
					statusChannel1 <- 2
					sensorHasCut = true
					return
				} else if int(results[0]) == 2 && pulse == uint16(19999) {
					fmt.Println("Sensor returned ---> ", results[0])
					d.SwitchOffMotor()
					sensorHasCut = false
					time.Sleep(100 * time.Millisecond)
					response, err = d.SetupMotor(uint16(2000), uint16(2000), uint16(100), uint16(0), uint16(motorNum))
					if err != nil {
						return
					}
					statusChannel1 <- 3
					return
				}
			}

		skipSensor:
			switch pulse {
			case 29999, 59199, 19999, 26666:
				fmt.Println("Homing Fast")
				time.Sleep(100 * time.Millisecond)
			case 2999:
				fmt.Println("Homing Slow")
				time.Sleep(10 * time.Millisecond)
			default:
				time.Sleep(1 * time.Second)
			}

		}
	}()

	fmt.Println("Blocked")

forLoop1:
	for {
		select {
		case status := <-statusChannel1:
			fmt.Println("received ", status)
			break forLoop1
			// Go ON
		}
	}

	fmt.Println("After Blocking")

	// OFF The motor
	err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, uint16(0x0000))

	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Printf("Wrote Off motor. res : %+v \n", results)

	wrotePulses = 0
	completedPulses = 0
	return "RUN Completed", nil
}

func (d *Compact32Deck) DeckHoming() (response string, err error) {

	// M2
	// var sensorAddressBytes = []byte{0x08, 0x02}
	// sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	sensorHasCut = false
	fmt.Println("Deck is moving forward")
	response, err = d.SetupMotor(uint16(2000), uint16(59199), uint16(100), uint16(1), uint16(5))
	if err != nil {
		return
	}

	sensorHasCut = false
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Deck is moving back by and after not cut -> 2000")
	response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), uint16(0), uint16(5))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Deck is moving forward again by 2999")
	response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), uint16(1), uint16(5))

	fmt.Println("Deck homing is completed.")

	return "DECK HOMING SUCCESS", nil
}

func (d *Compact32Deck) SwitchOffMotor() (response string, err error) {
	var onOffAddressBytes = []byte{0x08, 0x00}
	onOffAddressUint16 := binary.BigEndian.Uint16(onOffAddressBytes)

	//var completionAddressBytes = []byte{0x08, 0x01}
	//completionAddressUint16 := binary.BigEndian.Uint16(completionAddressBytes)

	err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, uint16(0x0000))
	//results, err = m.Client.WriteSingleCoil(onOffAddressUint16, uint16(0xff00))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	//fmt.Printf("Wrote Off motor Forcefully. res : %+v \n", results)

	//response, err = m.ReadD2000()
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return "", err
	//}

	//results, err = m.Client.WriteSingleCoil(completionAddressUint16, uint16(0x0000))
	////results, err = m.Client.WriteSingleCoil(onOffAddressUint16, uint16(0xff00))
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return "", err
	//}
	//fmt.Printf("Wrote Off Completion Forcefully. res : %+v \n", results)
	return "SUCCESS", nil

}

func (d *Compact32Deck) IsRunInProgress() (response string, err error) {

	// check if motor is On -> if so then throw error

	if d.IsMotorOff() == false {
		err = fmt.Errorf("Previous run is already in Progress. Abort it or let it finish.")
		return "", err
	}

	// check if D2000 has any value and completion bit is Off
	// This means that Run In Progres but PAUSED.

	response, err = d.ReadD2000()
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	if d.IsCompletionBitOff() && wrotePulses > 0 {
		err = fmt.Errorf("Previous RUN is in PAUSED state. RESUME it or ABORT it at first.")
		return "", err
	}

	return "Your RUN is GOOD to GO", nil
}

func (d *Compact32Deck) IsMotorOff() bool {

	var onOffAddressBytes = []byte{0x08, 0x00}
	onOffAddressUint16 := binary.BigEndian.Uint16(onOffAddressBytes)

	results, err := d.DeckDriver.ReadCoils(onOffAddressUint16, uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return false
	}
	fmt.Printf("Read On/Off Coil. res : %+v \n", results)

	var resultsInt int
	resultsInt = 10 // something unique
	if len(results) > 0 {
		resultsInt = int(results[0])
	}

	if resultsInt == 0 {
		return true
	}

	return false
}

func (d *Compact32Deck) ReadD2000() (response string, err error) {
	var D2000AddressBytes = []byte{0x17, 0xD0}
	D2000AddressBytesUint16 := binary.BigEndian.Uint16(D2000AddressBytes)

	results, err := d.DeckDriver.ReadHoldingRegisters(D2000AddressBytesUint16, uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
	}
	fmt.Printf("Read D2000AddressBytesUint16. res : %+v \n", results)
	//if len(results) > 0 {
	//	completedPulses = uint16(results[0])
	//}

	var D800AddressBytes = []byte{0x10, 0xD4}
	D800AddressBytesUint16 := binary.BigEndian.Uint16(D800AddressBytes)

	results, err = d.DeckDriver.ReadHoldingRegisters(D800AddressBytesUint16, uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	fmt.Printf("Read D800AddressBytesUint16. res : %+v \n", results)
	if len(results) > 0 {
		completedPulses = binary.BigEndian.Uint16(results)
	} else {
		err = fmt.Errorf("couldn't read D800")
		return "", err
	}
	fmt.Println("Read D800 Pulses -> ", completedPulses)

	//fmt.Printf("read ReadHoldingRegisters_Speed : %+v ", )

	return "D800 Reading SUCESS", nil

}

func (d *Compact32Deck) IsCompletionBitOff() bool {

	var completionAddressBytes = []byte{0x08, 0x01}
	completionAddressUint16 := binary.BigEndian.Uint16(completionAddressBytes)

	results, err := d.DeckDriver.ReadCoils(completionAddressUint16, uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return false
	}
	fmt.Printf("Read Completion Coil. res : %+v \n", results)

	var resultsInt int
	resultsInt = 10 // something unique
	if len(results) > 0 {
		resultsInt = int(results[0])
	}

	if resultsInt == 0 {
		return true
	}

	return false
}

func (d *Compact32Deck) SyringeHoming() (response string, err error) {
	// M2
	// var sensorAddressBytes = []byte{0x08, 0x02}
	// sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	sensorHasCut = false
	fmt.Println("Syringe is moving down until sensor not cut")
	response, err = d.SetupMotor(uint16(2000), uint16(26666), uint16(100), uint16(1), uint16(10))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)
	sensorHasCut = false
	fmt.Println("Aspiring and getting cut then aspiring 2000")
	response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), uint16(0), uint16(10))
	if err != nil {
		return
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Syringe dispencing again")
	response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), uint16(1), uint16(10))
	if err != nil {
		return
	}

	fmt.Println("Syringe homing is completed")

	return "SYRINGE HOMING COMPLETED", nil
}

func (d *Compact32Deck) SyringeModuleHoming() (response string, err error) {

	// M2

	// var sensorAddressBytes = []byte{0x08, 0x02}
	// sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	// Make motor go way up fast
	// K9
	sensorHasCut = false
	fmt.Println("Syringe Module moving Up")
	response, err = d.SetupMotor(uint16(2000), uint16(29999), uint16(100), uint16(1), uint16(9))
	if err != nil {
		return
	}

	sensorHasCut = false
	//response, err = m.SwitchOffMotor()
	fmt.Println("++++++++After First Fast Moving Up and getting Cut++++++++")

	time.Sleep(100 * time.Millisecond)

	// Make motor go down 2000 pulses slow
	// K9
	fmt.Println("+++++++++Syringe Module moving Down 20 mm or More!!+++++++++++")
	response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), uint16(0), uint16(9))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	// Make motor go up very slow
	// K9

	fmt.Println("Syringe Module moving Up")
	response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), uint16(1), uint16(9))
	if err != nil {
		return
	}

	//response, err = m.SwitchOffMotor()
	fmt.Println("++++++++After Final Slow Moving Up and getting Cut++++++++")

	return "SYRINGE HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetHoming() (response string, err error) {
	response, err = d.MagnetUpDownHoming()
	if err != nil {
		return
	}
	response, err = d.MagnetFwdRevHoming()
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Moving Magnet Back by 50mm")
	sensorHasCut = false
	response, err = d.SetupMotor(uint16(2000), uint16(10000), uint16(100), uint16(0), uint16(7))
	if err != nil {
		return
	}

	return "MAGNET HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetUpDownHoming() (response string, err error) {

	// M2
	// var sensorAddressBytes = []byte{0x08, 0x02}
	// sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	sensorHasCut = false
	fmt.Println("Magnet is moving up")
	response, err = d.SetupMotor(uint16(2000), uint16(29999), uint16(100), uint16(1), uint16(6))
	if err != nil {
		return
	}

	sensorHasCut = false
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Magnet is moving down by and after not cut -> 2000")
	response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), uint16(0), uint16(6))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Magnet is moving up again by 2999 till sensor cuts")
	response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), uint16(1), uint16(6))

	fmt.Println("Magnet Up/Down homing is completed.")

	return "MAGNET UP/DOWN HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetFwdRevHoming() (response string, err error) {

	// M2
	//var sensorAddressBytes = []byte{0x08, 0x02}
	//sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	sensorHasCut = false
	fmt.Println("Magnet is moving forward")
	response, err = d.SetupMotor(uint16(2000), uint16(29999), uint16(100), uint16(1), uint16(7))
	if err != nil {
		return
	}

	sensorHasCut = false
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Magnet is moving back by and after not cut -> 2000")
	response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), uint16(0), uint16(7))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Magnet is moving forward again by 2999")
	response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), uint16(1), uint16(7))

	fmt.Println("Magnet Up/Down homing is completed.")

	return "MAGNET FWD/REV HOMING SUCCESS", nil
}

func (d *Compact32Deck) Homing() (response string, err error) {
	aborted = false
	// check if run is already running, i.e check if motor is on and completion is off
	response, err = d.IsRunInProgress()
	if err != nil {
		return
	}

	fmt.Println("Moving Syringe DOWN till sensor cuts it")
	response, err = d.SyringeHoming()
	if err != nil {
		return
	}

	fmt.Println("Moving Syringe Module UP till sensor cuts it")
	response, err = d.SyringeModuleHoming()
	if err != nil {
		return
	}

	fmt.Println("Moving deck forward till sensor cuts it")
	response, err = d.DeckHoming()
	if err != nil {
		return
	}
	// Move deck forward till sensor cuts it

	fmt.Println("Homing Magnet")
	response, err = d.MagnetHoming()
	if err != nil {
		return
	}
	//response, err = m.runForward()
	//if err != nil {
	//	return
	//}

	return "HOMING SUCCESS", nil
}
