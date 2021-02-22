package compact32

import (
	"encoding/binary"
	"fmt"
	"time"
)

/* ****** ALGORITHM *******

1. Validate the temperature: min 20 C and max 120 C (for register it would be 200 and 1200)
2. Validate time duration between 10 sec and 1hr-60sec(3660 secs)
3. Validate shaker no. value as not empty and between 1-3
4. Heater on
5. Select shaker/s for heating
6. check if followup is to be kept on
7. if yes then write respective values on register
8. if no the write the counter values on register
9. according to the value of follow up start the time .
10. before sending remove decimal point from temperature and send time in sec
11.continously read the register to check if the temperature is set after each 300 ms
12. after reading apply decimal point to 1 decimal places
*/

func (d *Compact32Deck) Heat(temperature, shaker uint16, followup bool, heatTime time.Duration) (response string, err error) {

	heatTime = heatTime * time.Second
	var setTemp, setTemp1, setTemp2 uint16 = 0, 0, 0
	var t *time.Timer
	// validation for temperature
	if temperature > uint16(1200) || temperature <= uint16(200) {
		err = fmt.Errorf("%v not in valid range of 20 to 120", temperature)
		fmt.Println("Error 1: ", err)
		return "", err
	}

	//validation for heatTime
	if heatTime > time.Duration(3660*time.Second) || heatTime < time.Duration(10*time.Second) {
		err = fmt.Errorf("%v not in valid range of 10sec to 1hr 60sec", heatTime)
		fmt.Println("Error 2: ", err)
		return "", err
	}

	//validation for shaker
	if shaker > 3 || shaker < 1 {
		err = fmt.Errorf("%v not in valid range of 1-3", shaker)
		fmt.Println("Error 3: ", err)
		return "", err
	}

	//Heater on
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], ON)
	if err != nil {
		fmt.Println("err 4: ", err)
		return "", err
	}

	// //Pid calibration for LH
	// fmt.Println("pid for lh on ")
	// err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][4], ON)
	// if err != nil {
	// 	fmt.Println("err pid: ", err)
	// 	return "", err
	// }

	//select shaker for heating
	result, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][222], shaker)
	if err != nil {
		fmt.Println("err 5: ", err)
		return "", err
	} //Pid calibration

	fmt.Printf("shaker result %v", result)

	//Set Temperature for shakers
	result, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][208], uint16(650))
	if err != nil {
		fmt.Println("err 6: ", err)
		return "", err
	}
	fmt.Printf("temperature result %v", result)
	result, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][208], uint16(1))
	if err != nil {
		fmt.Println("err 6: ", err)
		return "", err
	}
	fmt.Printf("temperature result reading %v", result)

	if !followup {
		t = time.AfterFunc(heatTime, d.SwitchOffTheHeater)
	}

	go func() {
		for {
			result, err = d.DeckDriver.ReadHoldingRegisters(uint16(0x11F8), uint16(1))
			if err != nil {
				fmt.Println("err 6: ", err)

			}
			fmt.Println("D504 reading: ", result)
			time.Sleep(time.Second)
		}
	}()

shakerSelectionLoop:
	for {
		switch shaker {
		case 1:
			if setTemp >= temperature {
				break shakerSelectionLoop
			}
			results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][210], 1)
			if err != nil {
				fmt.Println("err :shaker 1 ", err)
				return "", err
			}
			setTemp = binary.BigEndian.Uint16(results)
			fmt.Println(setTemp)

		case 2:
			if setTemp >= temperature {
				break shakerSelectionLoop
			}
			results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][224], 1)
			if err != nil {
				fmt.Println("err :shaker 2 ", err)
				return "", err
			}
			setTemp = binary.BigEndian.Uint16(results)
			fmt.Println(setTemp)

		case 3:
			if (setTemp1 >= temperature) && (setTemp2 >= temperature) {
				break shakerSelectionLoop
			}
			results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][210], 1)
			if err != nil {
				fmt.Println("err :shaker 3.1 ", err)
				return "", err
			}
			setTemp1 = binary.BigEndian.Uint16(results)
			fmt.Println(setTemp1)

			results, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][224], 1)
			if err != nil {
				fmt.Println("err :shaker 3.2 ", err)
				return "", err
			}
			setTemp2 = binary.BigEndian.Uint16(results)
			fmt.Println(setTemp2)

		}

	}

	time.Sleep(time.Minute * 15)
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][8], OFF)
	if err != nil {
		fmt.Println("err pid: ", err)
		return "", err
	}

	if followup {
		time.Sleep(heatTime)
		d.SwitchOffHeater()
		return

	}

	for {
		select {
		case n := <-t.C:
			fmt.Printf("time expired %v", n)
			result, err := d.DeckDriver.ReadSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3])
			if err != nil {
				fmt.Println("err : ", err)

			}
			fmt.Printf("result %v", result)

		}
	}

	return
}

func (d *Compact32Deck) SwitchOffTheHeater() {
	err := d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
}
