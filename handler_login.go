package main

import "fmt"

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("A username is required")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
