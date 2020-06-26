package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const FileName = ".jarvis-config.json"

var (
	parsedConfig = false
	UserConfig   Config
)

func init() {
	userHomeDir, _ := os.UserHomeDir()
	workingDir, _ := os.Getwd()
	directories := []string{
		fmt.Sprintf("%s/%s", userHomeDir, FileName),
		fmt.Sprintf("%s/%s", workingDir, FileName),
		fmt.Sprintf("%s/../%s", workingDir, FileName),
	}

	for _, dir := range directories {
		config, err := parseConfig(dir)
		if err != nil {
			log.Println(err)
		} else {
			UserConfig = config
			parsedConfig = true
			break
		}
	}

	if !parsedConfig {
		log.Panic("unable to parse user config")
	}
}

func parseConfig(filePath string) (config Config, err error) {
	if _, fErr := os.Stat(filePath); !os.IsNotExist(fErr) {
		data, rErr := ioutil.ReadFile(filePath)
		if rErr != nil {
			err = fmt.Errorf("unable to read file %s", filePath)
			return
		}

		if jErr := json.Unmarshal(data, &config); jErr != nil {
			err = fmt.Errorf("unable to parse config file %s properly", filePath)
			return
		}

		return
	}

	err = fmt.Errorf("file not found at provided path %s", filePath)
	return
}
