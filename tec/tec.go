package tec


/*
int DemoFunc(double, double);
int InitiateTEC();
int checkForErrorState();
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

func InitiateTEC()(err error){
	C.InitiateTEC()
	
	go func(){
		for {
			time.Sleep(1 * time.Second)
			var errNum C.int
			errNum = C.checkForErrorState()
			if errNum != 0{
				logger.Errorln("Error Code for TEC: ", errNum)
				return
			}
		}

	}()
	
	return nil
}


func ConnectTEC(t TECTempSet) (err error) {
	C.DemoFunc(C.double(t.TargetTemperature), C.double(t.TargetRampRate))
	return nil
}
