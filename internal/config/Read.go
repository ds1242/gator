package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type ConfigJSON struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *ConfigJSON) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func Read() (ConfigJSON, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting home dir")
	}
	dat, err := os.Open(filePath)
	if err != nil {
		fmt.Println("unable to open file")
		return ConfigJSON{}, err
	}
	defer dat.Close()

	decoder := json.NewDecoder(dat)
	j := ConfigJSON{}

	err = decoder.Decode(&j)
	if err != nil {
		return ConfigJSON{}, err
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

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
