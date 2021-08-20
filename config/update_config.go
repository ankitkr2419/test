package config

import (
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

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if fi.IsDir() {
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

func UpdateConfig(path string) (err error) {
	// Below Walk will search for all the files in that path
	err = filepath.Walk(path, visit)
	if err != nil {
		logger.Errorln("Error Updating the Configs: ", err)
	}
	return
}
