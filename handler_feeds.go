package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < len(feeds); i++ {
		fmt.Println(feeds[i])
	}

	return nil
}
