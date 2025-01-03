package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ds1242/gator.git/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("A name is required")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		fmt.Println("unable to set user")
		return err
	}

	fmt.Println("User registered: ", user)
	return nil

}
