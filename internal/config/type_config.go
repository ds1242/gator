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

func (c ConfigJSON) SetUser(username string) error {
	c.CurrentUser = username
	dat, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshalling")
	}

	configPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting file path")
	}

	err = os.WriteFile(configPath, dat, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file")
	}
}
