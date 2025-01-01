package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (ConfigJSON, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting home dir")
	}
	dat, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to open file")
		return ConfigJSON{}, err
	}
	j := ConfigJSON{}

	err = json.Unmarshal(dat, &j)
	if err != nil {
		fmt.Println("error unmarshalling json")
		return j, err
	}
	return j, nil

}

func getConfigFilePath() (string, error) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}
	configFilePath := homeDir + configFileName

	return configFilePath, nil
}
