package main

import (
	"context"
	"fmt"
	"html"
)

func handlerAgg(s *state, cmd command) error {
	//if len(cmd.Args) < 2 {
	//	return fmt.Errorf("Not enough arguments")
	//}
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
