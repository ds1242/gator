package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("A username is required")
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("username does not exist")
		return err
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		fmt.Println("unable to set user")
		return err
	}

	fmt.Println("User has been set")

	return nil
}
