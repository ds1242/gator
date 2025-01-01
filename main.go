package main

import (
	"fmt"
	"log"

	"github.com/ds1242/gator.git/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error getting home environment")
	}

	state := &State{
		config: &cfg,
	}

	cmds := &Commands{
		handlers: make(map[string]func(*State, Command) error),
	}

	cmds.register("login", handlerLogin)
}
