package main

import (
	"fmt"
	"log"
	"os"

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

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Not enough arguments entered")
		os.Exit(1)
	}

	cmd := Command{
		commandName: args[1],
		arguments:   args[2:],
	}

	err = cmds.run(state, cmd)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	os.Exit(0)

}
