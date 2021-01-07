package compact32

import (
	"encoding/binary"
	"fmt"
	"time"
)

func (d *Compact32Deck) Check() bool {
	return true
}

var wrotePulses uint16 = 0
var completedPulses uint16 = 0
var sensorHasCut = false
var aborted = false

var statusChannel2 = make(chan int)

var pause = make(chan int)

func (d *Compact32Deck) SetupMotor(speed, pulse, ramp, direction, motorNum, onOff, completion uint16) (response string, err error) {
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

	sensorAddressUint16 := completion

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

	err = d.DeckDriver.WriteSingleCoil(onOffAddressUint16, onOff)
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

			if sensorAddressUint16 == uint16(0) {
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
					response, err = d.SetupMotor(uint16(2000), uint16(2000), uint16(100), uint16(0), uint16(motorNum), uint16(0xff00), uint16(0x0000))
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
	var sensorAddressBytes = []byte{0x08, 0x02}
	sensorAddressUint16 := binary.BigEndian.Uint16(sensorAddressBytes)

	sensorHasCut = false
	fmt.Println("Deck is moving forward")
	response, err = d.SetupMotor(uint16(2000), uint16(59199), uint16(100), uint16(1), uint16(5), uint16(0xff00), sensorAddressUint16)
	if err != nil {
		return
	}

	sensorHasCut = false
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Deck is moving back by and after not cut -> 2000")
	response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), uint16(0), uint16(5), uint16(0xff00), sensorAddressUint16)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Deck is moving forward again by 2999")
	response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), uint16(1), uint16(5), uint16(0xff00), sensorAddressUint16)

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

	//response, err = m.readD2000()
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

// func (d *Compact32Deck) DeckHoming() (response string, err error) {
// 	response, err = d.deckHoming()
// 	return
// }
