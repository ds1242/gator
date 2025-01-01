package config

import (
	"fmt"
	"os"
)

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("an error occurred getting the home dir: ")
		return Config{}, err
	}
	fmt.Println(homeDir)

	return Config{}, nil

}
