package compact32

import (
	"encoding/binary"
	"fmt"
	"mylab/cpagent/db"
	"time"
)

var wrotePulses = map[string]uint16{
	"A": 0,
	"B": 0,
}
var executedPulses = map[string]uint16{
	"A": 0,
	"B": 0,
}
var sensorHasCut = map[string]bool{
	"A": false,
	"B": false,
}
var aborted = map[string]bool{
	"A": false,
	"B": false,
}

const (
	K1_Syringe_Module_LH = iota + 1
	K2_Syringe_Module_RH
	K3_Syringe_LH
	K4_Syringe_RH
	K5_Deck
	K6_Magnet_Up_Down
	K7_Magnet_Rev_Fwd
	K8_Shaker
	K9_Syringe_Module_LHRH
	K10_Syringe_LHRH
)

type DeckNumber struct {
	Deck   string
	Number int
}

var motors = make(map[DeckNumber]map[string]uint16)

func SelectAllMotors(store db.Storer) (err error) {
	allMotors, err := store.GetAllMotors()
	if err != nil {
		return
	}

	for _, motor := range allMotors {
		deckNumber := DeckNumber{Deck: motor.Deck, Number: motor.Number}
		motors[deckNumber] = make(map[string]uint16)
		motors[deckNumber]["ramp"] = uint16(motor.Ramp)
		motors[deckNumber]["steps"] = uint16(motor.Steps)
		motors[deckNumber]["slow"] = uint16(motor.Slow)
		motors[deckNumber]["fast"] = uint16(motor.Fast)
	}
	return
}

func (d *Compact32Deck) SetupMotor(speed, pulse, ramp, direction, motorNum uint16) (response string, err error) {

	if aborted[d.name] {
		err := fmt.Errorf("Machine in ABORTED STATE")
		return "", err
	}

	var results []byte
	statusChannel := make(chan int)
	//*** statusChannel return Values ***
	// 1. Pulses were completed successfully
	// 2. Aborted
	// 3. Sensor has cut
	// 4. Sensor has uncut and motor travelled 2000 pulses reverse

	// Switch OFF The motor
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][202], pulse)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Pulse. res : ", results)
	wrotePulses[d.name] = pulse

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][200], speed)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Speed. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][204], ramp)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Ramp. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][206], direction)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote direction. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][226], motorNum)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote motorNum. res : ", results)

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][0], uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Printf("Read On/Off Coil. res : %+v \n", results)

	go func() {
		for {
			if aborted[d.name] {
				statusChannel <- 2
				return
			}
			results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][1], uint16(1))
			if err != nil {
				fmt.Println("err : ", err)
			}
			if len(results) > 0 {
				if int(results[0]) == 1 {
					fmt.Println("Completion returned ---> ", results)
					statusChannel <- 1
					return
				}
			}

			if direction == uint16(0) && pulse != uint16(19999) {
				goto skipSensor
			}
			results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][2], uint16(1))
			if err != nil {
				fmt.Println("err : ", err)
			}
			fmt.Println("Sensor returned ---> ", results)
			if len(results) > 0 {
				if int(results[0]) == 3 && pulse != uint16(19999) {
					fmt.Println("Sensor returned ---> ", results[0])
					statusChannel <- 3
					sensorHasCut[d.name] = true
					return
				} else if int(results[0]) == 2 && pulse == uint16(19999) {
					fmt.Println("Sensor returned ---> ", results[0])
					d.SwitchOffMotor()
					sensorHasCut[d.name] = false
					time.Sleep(100 * time.Millisecond)
					response, err = d.SetupMotor(uint16(2000), uint16(2000), uint16(100), REV, motorNum)
					if err != nil {
						return
					}
					statusChannel <- 4
					return
				}
			}

		skipSensor:
			switch pulse {
			case 29999, 59199, 19999, 26666:
				time.Sleep(100 * time.Millisecond)
			case 2999:
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
		case status := <-statusChannel:
			fmt.Println("received ", status)
			break forLoop1
			// Go ON
		}
	}

	fmt.Println("After Blocking")

	// OFF The motor
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	wrotePulses[d.name] = 0
	executedPulses[d.name] = 0
	return "RUN Completed", nil
}

func (d *Compact32Deck) SwitchOffMotor() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	return "SUCCESS", nil
}

func (d *Compact32Deck) IsRunInProgress() (response string, err error) {

	if d.IsMotorOff() == false {
		err = fmt.Errorf("Previous run is already in Progress. Abort it or let it finish.")
		return "", err
	}

	// check if D212 has any value and completion bit is Off
	// This means that Run In Progres but PAUSED.

	response, err = d.ReadExecutedPulses()
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	if d.IsCompletionBitOff() && wrotePulses[d.name] > 0 {
		err = fmt.Errorf("Previous RUN is in PAUSED state. RESUME it or ABORT it at first.")
		return "", err
	}

	return "Your RUN is GOOD to GO", nil
}

func (d *Compact32Deck) IsMotorOff() bool {

	results, err := d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][0], uint16(1))
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

func (d *Compact32Deck) ReadExecutedPulses() (response string, err error) {

	results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][212], uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	fmt.Printf("Read D212AddressBytesUint16. res : %+v \n", results)
	if len(results) > 0 {
		executedPulses[d.name] = binary.BigEndian.Uint16(results)
	} else {
		err = fmt.Errorf("couldn't read D212")
		return "", err
	}
	fmt.Println("Read D212 Pulses -> ", executedPulses[d.name])

	return "D212 Reading SUCESS", nil

}

func (d *Compact32Deck) IsCompletionBitOff() bool {

	results, err := d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][1], uint16(1))
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

func (d *Compact32Deck) Homing() (response string, err error) {

	aborted[d.name] = false
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

	fmt.Println("Homing Magnet")
	response, err = d.MagnetHoming()
	if err != nil {
		return
	}

	fmt.Println("Homing Completed Successfully")

	return "HOMING SUCCESS", nil
}

func (d *Compact32Deck) SyringeHoming() (response string, err error) {

	fmt.Println("Syringe is moving down until sensor not cut")
	// NOTE: Syringe UP means going to sensor DOWN
	// response, err = d.SetupMotor(uint16(2000), uint16(26666), uint16(100), UP, uint16(10))

	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}]["fast"], uint16(26666), motors[DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}]["ramp"], UP, uint16(K10_Syringe_LHRH))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Aspiring and getting cut then aspiring 2000")
	//response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), DOWN, uint16(10))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}]["fast"], uint16(19999), motors[DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}]["ramp"], DOWN, uint16(K10_Syringe_LHRH))
	if err != nil {
		return
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Syringe dispencing again")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), UP, uint16(10))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}]["slow"], uint16(2999), motors[DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}]["ramp"], UP, uint16(K10_Syringe_LHRH))
	if err != nil {
		return
	}

	fmt.Println("Syringe homing is completed")

	return "SYRINGE HOMING COMPLETED", nil
}

func (d *Compact32Deck) SyringeModuleHoming() (response string, err error) {

	fmt.Println("Syringe Module moving Up")
	// response, err = d.SetupMotor(uint16(2000), uint16(29999), uint16(100), UP, uint16(9))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}]["fast"], uint16(29999), motors[DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}]["ramp"], UP, uint16(K9_Syringe_Module_LHRH))
	if err != nil {
		return
	}

	fmt.Println("++++++++After First Fast Moving Up and getting Cut++++++++")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("+++++++++Syringe Module moving Down 20 mm or More!!+++++++++++")
	// response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), DOWN, uint16(9))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}]["fast"], uint16(19999), motors[DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}]["ramp"], DOWN, uint16(K9_Syringe_Module_LHRH))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Syringe Module moving Up")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), UP, uint16(9))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}]["slow"], uint16(2999), motors[DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}]["ramp"], UP, uint16(K9_Syringe_Module_LHRH))
	if err != nil {
		return
	}

	fmt.Println("++++++++After Final Slow Moving Up and getting Cut++++++++")

	return "SYRINGE HOMING SUCCESS", nil
}

func (d *Compact32Deck) DeckHoming() (response string, err error) {

	fmt.Println("Deck is moving forward")
	// response, err = d.SetupMotor(uint16(2000), uint16(59199), uint16(100), FWD, uint16(5))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K5_Deck}]["fast"], uint16(59199), motors[DeckNumber{Deck: d.name, Number: K5_Deck}]["ramp"], FWD, uint16(K5_Deck))
	if err != nil {
		return
	}

	//	sensorHasCut[d.name] = false
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Deck is moving back by and after not cut -> 2000")
	// response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), REV, uint16(5))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K5_Deck}]["fast"], uint16(19999), motors[DeckNumber{Deck: d.name, Number: K5_Deck}]["ramp"], REV, uint16(K5_Deck))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Deck is moving forward again by 2999")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), FWD, uint16(5))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K5_Deck}]["slow"], uint16(2999), motors[DeckNumber{Deck: d.name, Number: K5_Deck}]["ramp"], FWD, uint16(K5_Deck))

	fmt.Println("Deck homing is completed.")

	return "DECK HOMING SUCCESS", nil
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

	// response, err = d.SetupMotor(uint16(2000), uint16(10000), uint16(100), REV, uint16(7))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["fast"], uint16(10000), motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["ramp"], REV, uint16(K7_Magnet_Rev_Fwd))
	if err != nil {
		return
	}

	return "MAGNET HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetUpDownHoming() (response string, err error) {

	fmt.Println("Magnet is moving up")
	// response, err = d.SetupMotor(uint16(2000), uint16(29999), uint16(100), UP, uint16(6))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}]["fast"], uint16(29999), motors[DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}]["ramp"], UP, uint16(6))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Magnet is moving down by and after not cut -> 2000")
	// response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), DOWN, uint16(6))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}]["fast"], uint16(19999), motors[DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}]["ramp"], DOWN, uint16(6))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Magnet is moving up again by 2999 till sensor cuts")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), UP, uint16(6))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}]["slow"], uint16(2999), motors[DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}]["ramp"], UP, uint16(6))

	fmt.Println("Magnet Up/Down homing is completed.")

	return "MAGNET UP/DOWN HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetFwdRevHoming() (response string, err error) {

	fmt.Println("Magnet is moving forward")
	// response, err = d.SetupMotor(uint16(2000), uint16(29999), uint16(100), FWD, uint16(7))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["fast"], uint16(29999), motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["ramp"], FWD, uint16(K7_Magnet_Rev_Fwd))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Magnet is moving back by and after not cut -> 2000")
	// response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), REV, uint16(7))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["fast"], uint16(19999), motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["ramp"], REV, uint16(K7_Magnet_Rev_Fwd))
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Magnet is moving forward again by 2999")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), FWD, uint16(7))
	response, err = d.SetupMotor(motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["slow"], uint16(2999), motors[DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}]["ramp"], FWD, uint16(K7_Magnet_Rev_Fwd))

	fmt.Println("Magnet Up/Down homing is completed.")

	return "MAGNET FWD/REV HOMING SUCCESS", nil
}
