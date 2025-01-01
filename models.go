package main

import (
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
