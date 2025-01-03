package main

import (
	"context"
	"fmt"
)

func handlerUserReset(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("too many arguments")
	}
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Println("error resetting users")
		return err
	}
	return nil
}
