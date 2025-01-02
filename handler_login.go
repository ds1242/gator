package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("A username is required")
	}

	err := s.config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
