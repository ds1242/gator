package main

import (
	"fmt"

	"github.com/ds1242/gator.git/internal/config"
)

type State struct {
	config *config.ConfigJSON
}

type Command struct {
	commandName string
	arguments   []string
}

type Commands struct {
	handlers map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.handlers[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	if handler, ok := c.handlers[cmd.commandName]; ok {
		err := handler(s, cmd)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("command does not exist")
	}
	return nil
}
