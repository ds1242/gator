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
	fmt.Println(state)
	cmds := &Commands{
		handlers: make(map[string]func(*State, Command) error),
	}

	fmt.Println(os.Args)

	cmds.register("login", handlerLogin)
}
