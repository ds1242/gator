package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (ConfigStruct, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("an error occurred getting the home dir")
		return ConfigStruct{}, err
	}
	dat, err := os.ReadFile(homeDir + "/.gatorconfig.json")
	if err != nil {
		fmt.Println("unable to open file")
		return ConfigStruct{}, err
	}

	config := ConfigStruct{}

	err = json.Unmarshal(dat, &config)
	if err != nil {
		fmt.Println("error unmarshalling json")
		return config, err
	}

	return config, nil

}
