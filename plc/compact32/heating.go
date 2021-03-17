package compact32

import (
	// "encoding/binary"
	"fmt"
	"time"
)

/* ****** ALGORITHM *******

1. Validate the temperature: min 20 C and max 120 C (for register it would be 200 and 1200)
2. Validate time duration between 10 sec and 1hr-60sec(3660 secs)
3. Validate shaker no. value as not empty and between 1-3
4. Select shaker/s for heating
5. Set Temperature
6. Check if aborted before setting heater on
7. Heater on
8. check if followup is to be kept on
9. if no then then start heating and the timer and after specified time turn off heater
10. if yes then start heating let it reach to specified temperature and then start timer and after time switch heater off
11. before sending remove decimal point from temperature and send time in sec
12. continously read the register to check if the temperature is set after each 300 ms
13. after reading apply decimal point to 1 decimal places for any further use.
*/
func (d *Compact32Deck) Heating(temperature uint16, follow_temperature bool, heatTime time.Duration) (response string, err error) {

	/*
	// here we are hardcoding the shaker no in future this is to be fetched dynamically.
	// 3 is the value that needs to be passed for heating both the shakers.
	shaker := uint16(3)

	heatTime = heatTime * time.Second
	var setTemp, setTemp1, setTemp2 uint16 = 0, 0, 0
	var t *time.Timer
	var done = false
	var registerAddress uint16 = 0

	// validation for temperature
	if temperature > uint16(1200) || temperature <= uint16(200) {
		err = fmt.Errorf("%v not in valid range of 20 to 120", temperature)
		fmt.Println("Error Temperature not in valid range: ", err)
		return "", err
	}

	//validation for heatTime
	if heatTime > time.Duration(3660*time.Second) || heatTime < time.Duration(10*time.Second) {
		err = fmt.Errorf("%v not in valid range of 10sec to 1hr 60sec", heatTime)
		fmt.Println("Error Duration for heating not in valid range: ", err)
		return "", err
	}

	//validation for shaker
	if shaker > 3 || shaker < 1 {
		err = fmt.Errorf("%v not in valid range of 1-3", shaker)
		fmt.Println("Error shaker number not in valid range: ", err)
		return "", err
	}

	//select shaker for heating
	result, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][222], shaker)
	if err != nil {
		fmt.Println("err 5: ", err)
		return "", err
	}

	//Set Temperature for heater
	result, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][208], temperature)
	if err != nil {
		fmt.Println("Error failed to write temperature: ", err)
		return "", err
	}
	fmt.Printf("Set Temperature %v", result)

	// first check aborted if yes then exit
	if temp, ok := aborted.Load(d.name); !ok {
		err = fmt.Errorf("aborted isn't loaded!")
		return
	} else if temp.(bool) {
		err = fmt.Errorf("Operation was ABORTED!")
		return "", err
	}

	//Heater on
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], ON)
	if err != nil {
		fmt.Println("err 4: ", err)
		return "", err
	}

	// first check if not follow up then inside a go routine start timer(afterFunc does this by default)
	// and start up the heating and when the timer is expired turn off the heater
	if !follow_temperature {
		fmt.Printf("inside not follow up")
		t = time.AfterFunc(heatTime, d.SwitchOffTheHeater)
	}

	// loop for continous reading of the shaker temp and check if the temperature has reached specified value.
shakerSelectionLoop:
	for {

		if temp, ok := aborted.Load(d.name); !ok {
			err = fmt.Errorf("aborted isn't loaded!")
			return
		} else if temp.(bool) {
			err = fmt.Errorf("Operation was ABORTED!")
			return "", err
		}

		if done {
			return "Operation was successful \n", nil
		}

		time.Sleep(time.Millisecond * 300)
		switch shaker {
		case 1, 2:
			if setTemp >= temperature {
				break shakerSelectionLoop
			}
			// here we set the register address according to the shaker
			if shaker == 1 {
				registerAddress = MODBUS_EXTRACTION[d.name]["D"][210]
			} else {
				registerAddress = MODBUS_EXTRACTION[d.name]["D"][224]
			}
			results, err := d.DeckDriver.ReadHoldingRegisters(registerAddress, 1)
			if err != nil {
				fmt.Printf("Error failed to read shaker %v temperature ---%v \n", shaker, err)
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
				fmt.Printf("Error failed to read shaker 1 temperature ----%v \n", err)
				return "", err
			}
			setTemp1 = binary.BigEndian.Uint16(results)
			fmt.Println(setTemp1)

			results, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][224], 1)
			if err != nil {
				fmt.Printf("Error failed to read shaker 2 temperature ----%v \n", err)
				return "", err
			}
			setTemp2 = binary.BigEndian.Uint16(results)
			fmt.Println(setTemp2)

		}

	}

	// if follow up is true then first let it reach the temperature then wait for
	// specified duration and then turn the heater off
	if follow_temperature {
		time.Sleep(heatTime)
		d.SwitchOffHeater()
		return

	}

	for {
		select {
		case n := <-t.C:
			fmt.Printf("time expired %v", n)
			result, err := d.SwitchOffHeater()
			if err != nil {
				fmt.Println("err : ", err)
				return "", err
			}
			fmt.Printf("result %v", result)
			done = true
			fmt.Println("Heating Was Successful")
			return "SUCCESS", nil
		default:
			if temp, ok := aborted.Load(d.name); !ok {
				err = fmt.Errorf("aborted isn't loaded!")
				return
			} else if temp.(bool) {
				err = fmt.Errorf("Operation was ABORTED!")
				return "", err
			}
			// delay of 300 ms for checking the expired time to avoid too much loop
			time.Sleep(time.Millisecond * 300)
		}
	}

	*/
	return "SUCCESS", nil
}

func (d *Compact32Deck) SwitchOffTheHeater() {
	_, err := d.SwitchOffHeater()
	if err != nil {
		fmt.Println("Error failed to switch heater off : ", err)
		return
	}
}
