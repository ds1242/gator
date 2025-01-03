package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {

	if len(cmd.Args) > 1 {
		return fmt.Errorf("too many arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("error getting users")
		return err
	}

	for i := 0; i < len(users); i++ {
		if users[i] == s.config.CurrentUserName {
			fmt.Println("%s (current)", users[i])
		} else {
			fmt.Println("%s", users[i])
		}
	}

	return nil
}
