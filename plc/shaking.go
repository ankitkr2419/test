package plc

import (
	"errors"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

const secondsInMinutes = 60

// Shaking : function
/* Algorithm ******************
1. Validate that rpm 2 and time 2 value is not set before setting rpm 1 and time 1
2. switch off the shaker
3. Check if syringe module is inDeck, then get it to rest position
4. Let the shaker run at the specified rpm1 till the time1 duration is completed.
5.  Switch Off Heater & Shaker (Call in defer)
6. WithTemp handle
7. Handle Follow Temp
8. After this run the shaker with rpm 2 till the time1 duration is


NOTE: live is to be set only for engineer/admin flow when he is starting it directly
*/
func (d *Compact32Deck) Shaking(shakerData db.Shaker, live bool) (response string, err error) {

	defer func() {
		if live {
			d.ResetAborted()
		}

		if err != nil {
			logger.Errorln(err)
			if err == responses.AbortedError {
				d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorOperationAborted, d.name, err.Error())
				return
			}
			d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
			return
		}
		d.WsMsgCh <- "SUCCESS_ShakerRun_ShakerRunSuccess"
	}()

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

	// 2 switch off the shaker
	_, err = d.switchOffShaker()
	if err != nil {
		logger.Errorln("err switching off shaker: ", err)
		return "", err
	}

	//3. Check if syringe module is inDeck, then get it to rest position
	if d.getSyringeModuleState() == InDeck {
		response, err = d.SyringeRestPosition()
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving syringe module before moving the deck. Error: %v", err)
		}
	}

	// 4. Let the shaker run at the specified rpm1 till the time1 duration is completed.
	response, err = d.switchOnShaker(uint16(float64(shakerData.RPM1) * rpmToPulses))
	if err != nil {
		logger.Errorln("err in switching on shaker---> error: ", err)
		return "", err
	}
	logger.Infoln("shaking with rpm 1", shakerData.RPM1, "started")

	d.WsMsgCh <- "PROGRESS_ShakerRun_ShakerRunStarted"

	// 5:  Switch Off Heater & Shaker (Call in defer)
	defer d.switchOffHeater()
	defer d.switchOffShaker()

	// 6. WithTemp handle
	// 7. Handle Follow Temp
	if shakerData.WithTemp {

		ht := db.Heating{
			Temperature: shakerData.Temperature,
			FollowTemp:  shakerData.FollowTemp,
			Duration:    0,
		}
		response, err = d.Heating(ht, live)
		if err != nil {
			return "", err
		}
		d.switchOnHeater(uint16(shakerData.Temperature * 10))
		logger.Infoln("switched on heater")
	}

	//check if aborted
	err = d.sleepIfPaused()
	if err != nil {
		return
	}

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

	// 8. After this run the shaker with rpm 2 till the time1 duration is
	// completed if rpm 2 is specified.
	//set shaker value with rpm 2 if it exists
	if shakerData.RPM2 != 0 {

		//set shaker register with rpm 2
		//switch on the shaker
		response, err = d.switchOnShaker(uint16(float64(shakerData.RPM2) * rpmToPulses))
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
