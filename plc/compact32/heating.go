package compact32

import (
	"encoding/binary"
	"fmt"
	"mylab/cpagent/db"
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
9. if no then then start heating and the timer and after specified time turn off heater and return.
10. if yes then start heating let it reach to specified temperature and then start timer and after time switch heater off
11. before sending remove decimal point from temperature and send time in sec
12. continously read the register to check if the temperature is set after each 300 ms
13. after reading apply decimal point to 1 decimal places for any further use.
*/
func (d *Compact32Deck) Heating(ht db.Heating) (response string, err error) {

	// here we are hardcoding the shaker no in future this is to be fetched dynamically.
	// 3 is the value that needs to be passed for heating both the shakers.
	shaker := uint16(3)

	delay := db.Delay{
		DelayTime: ht.Duration,
	}

	// validation for temperature
	if (ht.Temperature*10) > 1200 || (ht.Temperature*10) <= 200 {
		err = fmt.Errorf("%v not in valid range of 20 to 120", ht.Temperature)
		fmt.Println("Error Temperature not in valid range: ", err)
		return "", err
	}

	//validation for heating duration
	if ht.Duration > 3660 || ht.Duration < 10 {
		err = fmt.Errorf("%v not in valid range of 10sec to 1hr 60sec", ht.Duration)
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
		fmt.Println("Error failed to write temperature:", err)
		return "", err
	}

	//Set Temperature for heater
	result, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][208], uint16(ht.Temperature*10))
	if err != nil {
		fmt.Println("Error failed to write temperature: ", err)
		return "", err
	}
	fmt.Printf("Set Temperature %v", result)

	// first check aborted if yes then exit
	if d.isMachineInAbortedState() {
		err = fmt.Errorf("Operation was ABORTED!")
		return "", err
	}

	//Heater on
	response, err = d.switchOnHeater()
	if err != nil {
		fmt.Println("error in switching heater on ", err)
		return "", err
	}

	// first check if not follow up then call delay function.

	// when the timer is expired turn off the heater and return.

	// as we do not need to monitor the temperature here.
	if !ht.FollowTemp {
		fmt.Printf("inside not follow up")
		d.setHeaterInProgress()
		defer d.resetHeaterInProgress()
		response, err = d.AddDelay(delay)
		if err != nil {
			return
		}
		response, err = d.switchOffHeater()
		if err != nil {
			fmt.Println("error in switching heater off ", err)
			return
		}

		return
	}

	// loop for continous reading of the shaker temp and check if the temperature has reached specified value.
	d.monitorTemperature(shaker, ht.Temperature)

	// if follow up is true then first let it reach the temperature then add delay
	// specified duration and then turn the heater off

	response, err = d.AddDelay(delay)
	if err != nil {
		fmt.Printf("error in adding delay %v", err.Error())
		return
	}

	response, err = d.switchOffHeater()
	if err != nil {
		fmt.Printf("error in adding delay %v", err.Error())
		return
	}
	return
}

func (d *Compact32Deck) monitorTemperature(shakerNo uint16, temperature float64) (response string, err error) {
	var setTemp, setTemp1, setTemp2 float64

	var registerAddress uint16 = 0

	for {
		if d.isMachineInPausedState() {
			response, err = d.waitUntilResumed(d.name)
			if err != nil {
				return
			}
		}

		if d.isMachineInAbortedState() {
			err = fmt.Errorf("operation was ABORTED \n")
			return "aborted", err
		}

		time.Sleep(time.Second * 2)
		switch shakerNo {
		case 1, 2:
			if setTemp >= temperature {
				return "success", nil
			}
			// here we set the register address according to the shaker
			if shakerNo == 1 {
				registerAddress = MODBUS_EXTRACTION[d.name]["D"][210]
			} else {
				registerAddress = MODBUS_EXTRACTION[d.name]["D"][224]
			}
			results, err := d.DeckDriver.ReadHoldingRegisters(registerAddress, 1)
			if err != nil {
				fmt.Printf("Error failed to read shaker %v temperature ---%v \n", shakerNo, err)
				return "", err
			}
			setTemp = float64(binary.BigEndian.Uint16(results)) / 10

		case 3:
			if (setTemp1 >= temperature) && (setTemp2 >= temperature) {
				return "success", nil
			}
			time.Sleep(time.Second * 5)
			results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][210], 1)
			if err != nil {
				fmt.Printf("Error failed to read shaker 1 temperature ----%v \n", err)
				return "", err
			}
			setTemp1 = float64(binary.BigEndian.Uint16(results)) / 10

			time.Sleep(time.Second * 5)
			results, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][224], 1)
			if err != nil {
				fmt.Printf("Error failed to read shaker 2 temperature ----%v \n", err)
				return "", err
			}
			setTemp2 = float64(binary.BigEndian.Uint16(results)) / 10
		}

	}

}
