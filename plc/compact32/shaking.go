package compact32

import (
	"errors"
	"fmt"
	"mylab/cpagent/db"

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
func (d *Compact32Deck) Shaking(shakerData db.Shaker) (response string, err error) {

	var shakerNo = 3

	var motorNum = K8_Shaker

	var results []byte

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

	// write motor number for shaker
	if temp := d.getMotorNumReg(); temp == highestUint16 {
		err = fmt.Errorf("motor Number Register isn't loaded!")
		return
	} else if temp != motorNum {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][226], motorNum)
	}

	if err != nil {
		logger.Errorln("error writing motor num: ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote motorNum. res : ", results)
	motorNumReg.Store(d.name, motorNum)

	//restart process motor
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		fmt.Println("err starting shaker: ", err)
		return "", err
	}

	//set shaker selection register
	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][220], uint16(shakerNo))
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

		switch {
		case shakerData.FollowTemp:
			ht := db.Heating{
				Temperature: shakerData.Temperature,
				FollowTemp:  shakerData.FollowTemp,
				Duration:    0,
			}
			response, err = d.Heating(ht)

		default:
			response, err = d.switchOnHeater()

		}
		if err != nil {
			return "", err
		}
		d.setHeaterInProgress()

	}

	//check if aborted
	if d.isMachineInAbortedState() {
		err = fmt.Errorf("Operation was ABORTED!")
		return "", err
	}

	//start shaker
	response, err = d.switchOnShaker()
	if err != nil {
		fmt.Printf("err in switching on shaker---> error: %v\n ", err)
		return "", err
	}

	// add delay of time1 duration
	delay := db.Delay{
		DelayTime: shakerData.Time1,
	}
	response, err = d.AddDelay(delay)
	if err != nil {
		fmt.Println("err adding delay: ", err)
		return "", err
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
		response, err = d.switchOnShaker()
		if err != nil {
			fmt.Printf("err in switching on shaker---> error: %v\n ", err)
			return "", err
		}

		//wait for time 2 duration
		delay.DelayTime = shakerData.Time2
		response, err = d.AddDelay(delay)
		if err != nil {
			fmt.Println("err adding delay: ", err)
			return "", err
		}
	}

	//switch off both shaker and heater
	response, err = d.switchOffHeater()
	if err != nil {
		fmt.Printf("err in switching off heater---> error: %v\n ", err)
		return "", err
	}
	response, err = d.switchOffShaker()
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	return "SUCCESS", nil
}
