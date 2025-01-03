package main

import (
	"fmt"
)

func handlerRegister(s *state, cmd *command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("A name is required")
	}

}
