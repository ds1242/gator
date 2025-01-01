package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

type ConfigJSON struct {
	DBUrl       string `json:"db_url"`
	CurrentUser string `json:"current_user_name,omitempty"`
}

func (c ConfigJSON) SetUser() {
	c.CurrentUser = "dshaw12"
	dat, err := json.Marshal(c)
	if err != nil {
		fmt.Print("error marshalling")
		return
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting file path")
		return
	}

	err = os.WriteFile(configPath, dat, 0644)
	if err != nil {
		fmt.Println("error writing to file")
		return
	}
}
