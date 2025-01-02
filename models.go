package main

import (
	"fmt"

	"github.com/ds1242/gator.git/internal/config"
)

type state struct {
	config *config.ConfigJSON
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.registeredCommands[cmd.Name]; ok {
		err := handler(s, cmd)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("command does not exist")
	}
	return nil
}
