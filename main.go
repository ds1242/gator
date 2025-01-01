package main

import (
	"fmt"
	"log"

	"github.com/ds1242/gator.git/internal/config"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		log.Fatal("Error getting home environment")
	}
	fmt.Println(configFile)
	configFile.SetUser()

	configFileAgain, err := config.Read()
	if err != nil {
		log.Fatal("error re-reading config file")
	}
	fmt.Println(configFileAgain)

}
