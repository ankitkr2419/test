package plc

import (
	"encoding/binary"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"time"

	logger "github.com/sirupsen/logrus"
)

/* ****** ALGORITHM *******

1. Validate the temperature: min 20 C and max 120 C (for register it would be 200 and 1200)
2. Validate time duration between 10 sec and 1hr-60sec(3660 secs)
3. Set Temperature
4. Check if aborted before setting heater on
5. Check if syringe module is inDeck, then get it to rest position
6. Heater on
7. check if followup is to be kept on if no then then start heating and the timer and after specified time turn off heater and return.
8. if yes then start heating let it reach to specified temperature and then start timer and after time switch heater off.
9. Switch heater OFF
*/
func (d *Compact32Deck) Heating(ht db.Heating, live bool) (response string, err error) {

	d.setHeaterInProgress()
	defer d.resetHeaterInProgress()

	defer func() {
		if live {
			d.ResetAborted()
		}
		if err != nil {
			logger.Errorln(err)
			if err.Error() == AbortedError {
				d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorOperationAborted, d.name, err.Error())
				return
			}
			d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
			return
		}
		d.WsMsgCh <- "SUCCESS_HeaterRun_HeaterRunSuccess"
	}()

	stopMonitor := make(chan bool, 1)

	// Step 1 : Validation for temperature
	// validation for temperature
	if (ht.Temperature*10) > 1200 || (ht.Temperature*10) <= 200 {
		err = fmt.Errorf("%v not in valid range of 20 to 120", ht.Temperature)
		logger.Errorf("Error Temperature not in valid range: %v", err)
		return "", err
	}

	// Step 2 : Validation for Duration
	//validation for heating duration
	// zero value for duration signifies that the shaker has initiated followTemp.
	if ht.Duration > 3660 || ht.Duration < 10 && ht.Duration != 0 {
		err = fmt.Errorf("%v not in valid range of 10sec to 1hr 60sec", ht.Duration)
		logger.Errorln("Error Duration for heating not in valid range: ", err)
		return "", err
	}

	delay := db.Delay{
		DelayTime: ht.Duration,
	}

	//Step 3: Set Temperature
	//Set Temperature for heater
	result, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][208], uint16(ht.Temperature*10))
	if err != nil {
		logger.Errorln("Error failed to write temperature: ", err)
		return "", err
	}
	logger.Infoln("result from temperature set ", result, ht.Temperature)

	// Step 4 : Check if Aborted
	// first check aborted if yes then exit
	if d.isMachineInAbortedState() {
		err = fmt.Errorf(AbortedError)
		return "", err
	}

	// Step 5 : Syringe To Rest Position
	// Check if syringe module is inDeck, then get it to rest position

	if d.getSyringeModuleState() == InDeck {
		response, err = d.SyringeRestPosition()
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving syringe module before moving the deck. Error: %v", err)
		}
	}

	// Step 6 : Switch heater on
	//Heater on
	response, err = d.switchOnHeater()
	if err != nil {
		logger.Errorln("error in switching heater on ", err)
		return "", err
	}
	logger.Infoln("heating with temperature", ht.Temperature, "started")

	// Step 9:  Switch heater OFF (Called in defer)
	defer d.switchOffHeater()

	d.WsMsgCh <- "PROGRESS_HeaterRun_HeaterRunStarted"

	// Step 7: For not follow Temp
	// first check if not follow up then call delay function.
	// if no then then start heating  after specified time turn off heater and return
	// as we do not need to monitor the temperature here.
	if !ht.FollowTemp {
		logger.Infoln("not follow temperature")
		go d.monitorTemperature(shaker, ht.Temperature, false, stopMonitor)
		defer d.stopMonitorTemperature(stopMonitor)
		response, err = d.AddDelay(delay, false)
		if err != nil {
			return
		}
		response, err = d.switchOffHeater()
		if err != nil {
			logger.Errorln("error in switching heater off ", err)
			return
		}
		return
	}

	// Step 8 : monitor the temperature if follow temp.
	// loop for continous reading of the shaker temp and check if the temperature has reached specified value.
	response, err = d.monitorTemperature(shaker, ht.Temperature, true, stopMonitor)
	if err != nil {
		logger.Errorln("Error in monitor temperature \n", err)
		return "", err
	}

	// for shaker when the heater is needed with follow temp
	if delay.DelayTime == 0 {
		return "SUCCESS", nil
	}

	// After monitoring add delay of specified time period.
	response, err = d.AddDelay(delay, false)
	if err != nil {
		logger.Errorln("error in adding delay ", err.Error())
		return
	}

	logger.Infoln("heating with temperature", ht.Temperature, "completed")

	return
}

func (d *Compact32Deck) monitorTemperature(shakerNo uint16, temperature float64, tempCheck bool, stopMonitor chan bool) (response string, err error) {
	var setTemp, shaker1Temp, shaker2Temp, prevTemp1, prevTemp2 float64
	var heatingFailCounter1, heatingFailCounter2 int

	var registerAddress uint16 = 0
	delay := db.Delay{
		DelayTime: 30,
	}

	for {
		select {
		case n := <-stopMonitor:
			logger.Infoln("stop the montoring :", n)
			return "SUCCESS", nil

		default:
			if d.isMachineInPausedState() {
				response, err = d.waitUntilResumed(d.name)
				if err != nil {
					return
				}
			}

			if d.isMachineInAbortedState() {
				err = fmt.Errorf("operation was ABORTED \n")
				return "ABORTED", err
			}

			time.Sleep(time.Second * 2)
			switch shakerNo {
			case 1, 2:
				if setTemp >= temperature {
					return "SUCCESS", nil
				}
				// here we set the register address according to the shaker
				if shakerNo == 1 {
					registerAddress = MODBUS_EXTRACTION[d.name]["D"][210]
				} else {
					registerAddress = MODBUS_EXTRACTION[d.name]["D"][224]
				}
				results, err := d.DeckDriver.ReadHoldingRegisters(registerAddress, 1)
				if err != nil {
					logger.Errorln("Error failed to read shaker ", shakerNo, "temperature", err)
					return "", err
				}
				setTemp = float64(binary.BigEndian.Uint16(results)) / 10

			case 3:
				if !tempCheck {
					goto skipToMonitor
				}
				// Play of 2 degrees as heater would not heat up that much sometimes
				if (shaker1Temp >= temperature-2) && (shaker2Temp >= temperature-2) {
					return "SUCCESS", nil
				}

			skipToMonitor:

				if (shaker1Temp - prevTemp1) < 1 {
					heatingFailCounter1 += 1
				} else {
					heatingFailCounter1 = 0
				}

				if (shaker2Temp - prevTemp2) < 1 {
					heatingFailCounter2 += 1
				} else {
					heatingFailCounter2 = 0
				}

				prevTemp1 = shaker1Temp
				prevTemp2 = shaker2Temp

				if heatingFailCounter1 >= 15 || heatingFailCounter2 >= 15 {
					err = fmt.Errorf("temperature not upgrading")
					return "", err
				}

				shaker1Temp, shaker2Temp, err = d.readTempValues()
				if err != nil {
					logger.Errorln("Error failed to read temperature values for shaker heaters: ", err)
					return "", err
				}

				response, err = d.AddDelay(delay, false)
				if err != nil {
					logger.Errorln("Error failed to add delay in monitor temperature: ", err)
					return "", err
				}
			}
		}
	}
}

func (d *Compact32Deck) readTempValues() (shaker1Temp, shaker2Temp float64, err error) {

	defer func() {
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.FetchHeaterTempError)
			d.WsErrCh <- err
		}
	}()

	results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][210], 1)
	if err != nil {
		logger.Errorln("Error failed to read shaker 1 temperature")
		return
	}
	shaker1Temp = float64(binary.BigEndian.Uint16(results)) / 10

	logger.Infoln("temp 1 reading", shaker1Temp)

	results, err = d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][224], 1)
	if err != nil {
		logger.Errorln("Error failed to read shaker 2 temperature")
		return
	}
	shaker2Temp = float64(binary.BigEndian.Uint16(results)) / 10
	logger.Infoln("temp 2 reading", shaker2Temp)
	return
}

func (d *Compact32Deck) stopMonitorTemperature(stop chan bool) {
	logger.Infoln("stop monitor temperature")
	stop <- true
}

func (d *Compact32Deck) HeaterData() (err error) {
	retryCounter := 0
	for {
		time.Sleep(5 * time.Second)
		if d.isEngineerOrAdminLogged() {
			err = d.sendHeaterData()
			if err != nil {
				retryCounter++
				if retryCounter > 3 {
					return err
				}
				logger.WithFields(logger.Fields{
					"heater":  err,
					"attempt": retryCounter,
				}).Warnln("Attempt failed. Heater Value couldn't be read. Retrying...")
				time.Sleep(10 * time.Second) // sleep it off for a bit
			} else {
				retryCounter = 0
			}
		}
	}
}
