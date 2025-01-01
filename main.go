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

	state := &State{
		config: configFile,
	}

}
