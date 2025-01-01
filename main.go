package main

import (
	"fmt"
	"log"

	"github.com/ds1242/gator.git/internal/config"
)

func main() {
	config, err := config.Read()
	fmt.Println(config)
	if err != nil {
		log.Fatal("Error getting home environment")
	}
}
