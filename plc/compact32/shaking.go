package compact32

import (
	"encoding/binary"
	"errors"
	"fmt"
	"mylab/cpagent/db"
	"time"

	logger "github.com/sirupsen/logrus"
)

// Shaking : function
/* Algorithm ******************
1. Validate that rpm 2 and time 2 value is not set before setting rpm 1 and time 1
2. Switch off the shaker bit first and reset the completion bit to avoid any inconsistency.
3. Set the shaker , here in this case it is both the shaker.
4. Set the rpm 1 value.
5. If withTemp is true then operate with temp according to follow up or not follow up.
6. If follow up then wait for the temperature to reach that certain value and then start shaking.
7. Else if not follow up then just start the heater and then start the shaker.
8. If withTemp is false then proceed with the normal flow by starting the shaker.
9. Let the shaker run at the specified rpm1 till the time1 duration is completed.
10. After this run the shaker with rpm 2 till the time1 duration is completed if rpm 2
	is specified.
11. After all this process is done switch the shaker and the heater off.
*/
func (d *Compact32Deck) Shaking(shakerData db.Shaker) (result string, err error) {

	var shakerNo = 3

	//check if aborted
	if d.isMachineInAbortedState() {
		err = fmt.Errorf("Operation was ABORTED!")
		return "", err
	}

	//validate that rpm 1 is definately set and futher
	if shakerData.RPM1 == 0 || shakerData.Time1 == 0 {
		if shakerData.RPM2 != 0 || shakerData.Time2 != 0 {
			err = errors.New("please check value for rpm 1 data")
			return
		}
	} else {
		if shakerData.RPM2 == 0 && shakerData.Time2 != 0 {
			err = errors.New("please check rpm 2 value")
			return
		}
		if shakerData.RPM2 != 0 && shakerData.Time2 == 0 {
			err = errors.New("please check rpm  2 time value")
			return
		}
	}

	//switch off the shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], OFF)
	if err != nil {
		fmt.Println("err starting shaker: ", err)
		return "", err
	}

	//reset completion bit
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if err != nil {
		fmt.Println("err starting shaker: ", err)
		return "", err
	}

	//restart process motor
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		fmt.Println("err starting shaker: ", err)
		return "", err
	}

	//set shaker selection register
	results, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][220], uint16(shakerNo))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	logger.Infof("selected shaker %v", results)

	//set shaker register with rpm 1
	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][218], uint16(shakerData.RPM1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	if shakerData.WithTemp {
		//check temperature limit
		if shakerData.Temperature > 1200 || shakerData.Temperature <= 200 {
			err = fmt.Errorf("%v not in valid range of 20 to 120", shakerData.Temperature)
			fmt.Println("Error Temperature not in valid range: ", err)
			return "", err

		}

		result, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][222], uint16(shakerNo))
		if err != nil {
			fmt.Println("failed to select heater shaker: ", err)
			return "", err
		}
		logger.Infof("selected shaker %v", result)

		//set heater value on selected shaker
		result, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][208], uint16(shakerData.Temperature))
		if err != nil {
			fmt.Println("Error failed to write temperature: ", err)
			return "", err
		}

		//heater on
		err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], ON)
		if err != nil {
			fmt.Printf("failed to switch on heater err : %v \n ", err)
			return "", err
		}

		if shakerData.FollowTemp {
			//continously monitor the temp until it reaches that temp then proceed
			result, err := d.MonitorTemperature(uint16(shakerNo), uint16(shakerData.Temperature))
			if err != nil {
				fmt.Printf("failed to read heater value for shaker err : %v \n ", err)
				return "", err
			}
			logger.Debugf("follow temp done %v", result)
		}
	}

	//start shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], ON)
	if err != nil {
		fmt.Println("err starting shaker: ", err)
		return "", err
	}

	//read shaker value
	result1, err := d.DeckDriver.ReadSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5])
	if err != nil {
		fmt.Println("err starting shaker: ", err)
		return "", err
	}

	fmt.Printf("shaker started %v ", result1)
	// TODO : move this to single method for every time it goes to sleep

	//wait for time 1 duration
	t := time.NewTimer(shakerData.Time1 * time.Second)

skipToShakerRpm2:
	for {
		select {
		case n := <-t.C:
			fmt.Printf("time expired %v", n)
			//switch off shaker
			d.switchOffShaker()
			//switch off heater
			d.switchOffHeater()
			break skipToShakerRpm2
		default:
			if d.isMachineInAbortedState() {
				err = fmt.Errorf("Operation was ABORTED!")
				return "", err
			}
			// delay of 300 ms for checking the expired time to avoid too much loop
			time.Sleep(time.Millisecond * 300)
		}
	}

	//set shaker value with rpm 2 if it exists
	if shakerData.RPM2 != 0 {

		//set shaker register with rpm 2
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][218], uint16(shakerData.RPM2))
		if err != nil {
			fmt.Println("err : ", err)
			return "", err
		}
		//switch on the shaker
		err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], ON)
		if err != nil {
			fmt.Println("err starting shaker: ", err)
			return "", err
		}
		//wait for time 2 duration
		t := time.NewTimer(shakerData.Time2 * time.Second)

	skipToShakerOff:
		for {
			select {
			case n := <-t.C:
				fmt.Printf("time expired %v", n)
				//switch off shaker
				d.switchOffShaker()
				//switch off heater
				d.switchOffHeater()
				break skipToShakerOff
			default:
				if d.isMachineInAbortedState() {
					err = fmt.Errorf("Operation was ABORTED!")
					return "", err
				}
				// delay of 300 ms for checking the expired time to avoid too much loop
				time.Sleep(time.Millisecond * 300)
			}
		}

	}

	//switch off both shaker and heater
	_, err = d.switchOffHeater()
	if err != nil {
		fmt.Printf("err in switching off heater---> error: %v\n ", err)
		return "", err
	}
	_, err = d.switchOffShaker()
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	return "Suceess", nil
}

func (d *Compact32Deck) switchOffShaker() (response string, err error) {

	// Switch off shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], OFF)
	if err != nil {
		fmt.Println("err Switching off the shaker: ", err)
		return "", err
	}
	fmt.Println("Switched off the shaker--> for deck ", d.name)
	return "SUCCESS", nil

}

func (d *Compact32Deck) MonitorTemperature(shakerNo, temperature uint16) (result string, err error) {
	var setTemp, setTemp1, setTemp2 uint16 = 0, 0, 0

	var done = false
	var registerAddress uint16 = 0

	for {
		if d.isMachineInAbortedState() {
			err = fmt.Errorf("operation was ABORTED \n")
			return "aborted", err
		}
		if done {
			return "Operation was successful \n", nil
		}
		time.Sleep(time.Second * 5)
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
			setTemp = binary.BigEndian.Uint16(results)

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
			setTemp1 = binary.BigEndian.Uint16(results)

			time.Sleep(time.Second * 5)
			results, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][224], 1)
			if err != nil {
				fmt.Printf("Error failed to read shaker 2 temperature ----%v \n", err)
				return "", err
			}
			setTemp2 = binary.BigEndian.Uint16(results)
		}

	}

}
