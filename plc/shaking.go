package plc

import (
	"errors"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"

	logger "github.com/sirupsen/logrus"
)

const secondsInMinutes = 60

// Shaking : function
/* Algorithm ******************
1. Validate that rpm 2 and time 2 value is not set before setting rpm 1 and time 1
2. Switch off the shaker bit first and reset the completion bit to avoid any inconsistency.
3. Set the shaker, here in this case it is both the shaker.
4. Set the rpm 1 value.
5. If withTemp is true then operate with temp according to follow up or not follow up.
6. If follow up then wait for the temperature to reach that certain value and then start shaking.
7. Else if not follow up then just start the heater and then start the shaker.
8. If withTemp is false then proceed with the normal flow by starting the shaker.
9. Let the shaker run at the specified rpm1 till the time1 duration is completed.
10. After this run the shaker with rpm 2 till the time1 duration is completed if rpm 2
	is specified.
11. After all this process is done switch the shaker and the heater off(Called in defer)
*/
func (d *Compact32Deck) Shaking(shakerData db.Shaker) (response string, err error) {

	defer func() {
		if err != nil {
			logger.Errorln(err)
			d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
			return
		}
		d.WsMsgCh <- "SUCCESS_ShakerRun_ShakerRunSuccess"
	}()

	var motorNum = K8_Shaker
	var results []byte

	rpmToPulses := float64(config.GetShakerStepsPerRev() / secondsInMinutes)

	// 1. validate that rpm 1 is definately set and futher
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

	// 2.1 switch off the shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], OFF)
	if err != nil {
		logger.Errorln("err starting shaker: ", err)
		return "", err
	}

	// 2.2 reset completion bit
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if err != nil {
		logger.Errorln("err resetting completion bit: ", err)
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
		logger.Errorln("err starting shaker: ", err)
		return "", err
	}

	// 4 set shaker register with rpm 1
	// NOTE: Calculation of RPM involves multiplying it with 13.3
	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][218], uint16(float64(shakerData.RPM1)*rpmToPulses))
	if err != nil {
		logger.Errorln("error in setting rpm 1 value : ", err)
		return "", err
	}

	// Check if syringe module is inDeck, then get it to rest position
	if d.getSyringeModuleState() == InDeck {
		response, err = d.SyringeRestPosition()
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving syringe module before moving the deck. Error: %v", err)
		}
	}

	d.WsMsgCh <- "PROGRESS_ShakerRun_ShakerRunStarted"

	// 5. WithTemp handle
	// 6. Handle Follow Temp
	if shakerData.WithTemp {

		ht := db.Heating{
			Temperature: shakerData.Temperature,
			FollowTemp:  shakerData.FollowTemp,
			Duration:    0,
		}
		response, err = d.Heating(ht)
		if err != nil {
			return "", err
		}
		d.switchOnHeater()
		logger.Infoln("switched on heater")
		d.setHeaterInProgress()
	}

	// Step 11:  Switch Off Heater & Shaker (Call in defer)
	defer d.switchOffHeater()
	defer d.switchOffShaker()

	// 7. Else if not follow up then just start the heater and then start the shaker.
	// 8. If withTemp is false then proceed with the normal flow by starting the shaker.
	//check if aborted
	if d.isMachineInAbortedState() {
		err = fmt.Errorf("Operation was ABORTED!")
		return "", err
	}

	//start shaker
	// 9. Let the shaker run at the specified rpm1 till the time1 duration is completed.
	response, err = d.switchOnShaker()
	if err != nil {
		logger.Errorln("err in switching on shaker---> error: ", err)
		return "", err
	}
	logger.Infoln("shaking with rpm 1", shakerData.RPM1, "started")

	d.setShakerInProgress()
	defer d.resetShakerInProgress()

	// add delay of time1 duration
	delay := db.Delay{
		DelayTime: shakerData.Time1,
	}
	response, err = d.AddDelay(delay, false)
	if err != nil {
		logger.Errorln("error adding delay: ", err)
		return "", err
	}

	logger.Infoln("shaking with rpm 1", shakerData.RPM1, "completed")

	// 10. After this run the shaker with rpm 2 till the time1 duration is
	// completed if rpm 2 is specified.
	//set shaker value with rpm 2 if it exists
	if shakerData.RPM2 != 0 {

		//set shaker register with rpm 2
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][218], uint16(float64(shakerData.RPM2)*rpmToPulses))
		if err != nil {
			logger.Errorln("error in setting rpm 2 value : ", err)
			return "", err
		}
		//switch on the shaker
		response, err = d.switchOnShaker()
		if err != nil {
			logger.Errorln("err in switching on shaker :", err)
			return "", err
		}
		logger.Infoln("shaking with rpm 2", shakerData.RPM2, "started")

		//wait for time 2 duration
		delay.DelayTime = shakerData.Time2
		response, err = d.AddDelay(delay, false)
		if err != nil {
			logger.Errorln("err adding delay: ", err)
			return "", err
		}
		logger.Infoln("shaking with rpm 2", shakerData.RPM2, "completed")

	}

	return "SUCCESS", nil
}
