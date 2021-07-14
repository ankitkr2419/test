package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	logger "github.com/sirupsen/logrus"

)

var oldString, newString []string 
var replaceCountLength int

type Conf struct {
	RoomTemperature int64 `json:"room_temperature" validate:"required,lte=30,gte=20"`
	HomingTime      int64 `json:"homing_time" validate:"required,lte=30,gte=16"`
	NumHomingCycles int64 `json:"no_of_homing_cycles" validate:"required,lte=100,gte=0"`
}

func SetValues(c Conf) (err error) {
	
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
		for i:=0; i< replaceCountLength; i++{
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
