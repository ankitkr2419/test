package config

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	configPath = "./conf"
)

var oldString, newString []string

type Conf struct {
	RoomTemperature int64  `json:"room_temperature" validate:"required,lte=30,gte=20"`
	HomingTime      int64  `json:"homing_time" validate:"required,lte=30,gte=16"`
	NumHomingCycles int64  `json:"no_of_homing_cycles" validate:"required,lte=100,gte=0"`
	CycleTime       int64  `json:"cycle_time" validate:"required,lte=30,gte=2"`
	PIDTemperature  int64  `json:"pid_temperature" validate:"required,lte=75,gte=50"`
	PIDMinutes      int64  `json:"pid_minutes" validate:"required,lte=40,gte=20"`
	ReceiverEmail   string `json:"receiver_email"`
	ReceiverName    string `json:"receiver_name"`
}

func SetValues(c Conf) (err error) {
	hT := GetHomingTime()
	hC := GetNumHomingCycles()
	rT := GetRoomTemp()
	cT := GetCycleTime()
	pT := GetPIDTemp()
	mT := GetPIDMinutes()
	rE := GetReceiverEmail()
	rN := GetReceiverName()

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("homing_time: %d", hT),
		fmt.Sprintf("num_homing_cycles: %d", hC),
		fmt.Sprintf("room_temp: %d", int64(rT)),
		fmt.Sprintf("cycle_time: %d", int64(cT)),
		fmt.Sprintf("pid_temp: %d", int64(pT)),
		fmt.Sprintf("pid_time: %d", int64(mT)),
		fmt.Sprintf("receiver_email: %s", rE),
		fmt.Sprintf("receiver_name: %s", rN),
	)
	newString = append(newString,
		fmt.Sprintf("homing_time: %d", c.HomingTime),
		fmt.Sprintf("num_homing_cycles: %d", c.NumHomingCycles),
		fmt.Sprintf("room_temp: %d", c.RoomTemperature),
		fmt.Sprintf("cycle_time: %d", c.CycleTime),
		fmt.Sprintf("pid_temp: %d", c.PIDTemperature),
		fmt.Sprintf("pid_time: %d", c.PIDMinutes),
		fmt.Sprintf("receiver_email: %s", c.ReceiverEmail),
		fmt.Sprintf("receiver_name: %s", c.ReceiverName),
	)

	err = UpdateConfig(configPath, oldString, newString)
	if err != nil {
		return
	}

	SetHomingTime(c.HomingTime)
	SetNumHomingCycles(c.NumHomingCycles)
	SetRoomTemp(c.RoomTemperature)
	SetCycleTime(c.CycleTime)
	SetPIDTemp(c.PIDTemperature)
	SetPIDMinutes(c.PIDMinutes)
	SetReceiverEmail(c.ReceiverEmail)
	SetReceiverName(c.ReceiverName)
	return
}

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil
	}

	// Only search for .yml files
	matched, err := filepath.Match("*.yml", fi.Name())

	if err != nil {
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Replace a bunch of strings
		newContents := string(read)
		for i := 0; i < len(oldString); i++ {
			newContents = strings.Replace(newContents, oldString[i], newString[i], -1)
		}

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateConfig(path string, oldS, newS []string) (err error) {
	oldString = oldS
	newString = newS
	// Below Walk will search for all the files in that path
	err = filepath.Walk(path, visit)
	if err != nil {
		logger.Errorln("Error Updating the Configs: ", err)
	}
	return
}
