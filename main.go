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

	state := &state{
		config: &cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Not enough arguments entered")
		os.Exit(1)
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(state, cmd)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	os.Exit(0)

}
