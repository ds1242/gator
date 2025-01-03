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
		var msg string
		if users[i] == s.config.CurrentUserName {
			msg = fmt.Sprintf("%s (current)", users[i])
		} else {
			msg = fmt.Sprintf("%s", users[i])
		}
		fmt.Println(msg)
	}

	return nil
}
