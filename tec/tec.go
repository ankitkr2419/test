package tec


/*
int DemoFunc(double, double);
int InitiateTEC();
int checkForErrorState();
int autoTune();
int resetDevice();
#include <stdlib.h>
#include <time.h>
#include <fcntl.h>
#include <termios.h>
#include <unistd.h>
#include <errno.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#cgo CFLAGS  : -std=gnu99 -Wall -g -O3
#cgo LDFLAGS : -pthread -lrt
*/
import "C"
import (
	"time"
logger "github.com/sirupsen/logrus"

)

type TECTempSet struct {
	TargetTemperature      float64 `json:"target_temp" validate:"gte=-20,lte=100"`
	TargetRampRate		   float64 `json:"ramp_rate" validate:"gte=-20,lte=100"`
}

var errorCheckStopped bool 

func InitiateTEC()(err error){
	C.InitiateTEC()
	
	go startErrorCheck()
	
	return nil
}

func startErrorCheck(){
go func(){
	time.Sleep(5 * time.Second)
	for {
		time.Sleep(2 * time.Second)
		var errNum C.int
		errNum = C.checkForErrorState()
		if errNum != 0{
			errorCheckStopped = true
			logger.Errorln("Error Code for TEC: ", errNum)
			return
		}
	}
}()
}

func ConnectTEC(t TECTempSet) (err error) {
	C.DemoFunc(C.double(t.TargetTemperature), C.double(t.TargetRampRate))
	return nil
}

func AutoTune() (err error){
	C.autoTune()
	return nil
}

func ResetDevice() (err error){
	C.resetDevice()

	if errorCheckStopped {
		startErrorCheck()
	}
	return nil
}