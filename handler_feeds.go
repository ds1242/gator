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

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	fmt.Println("Feeds Found")
	fmt.Println("==============================")
	for _, feed := range feeds {
		fmt.Printf("Feed Name:	       %s\n", feed.Name)
		fmt.Printf("Feed URL:            %s\n", feed.Url)
		fmt.Printf("User Name:	       %s\n", feed.Name_2)
		fmt.Println("==============================")
	}

	return nil
}
