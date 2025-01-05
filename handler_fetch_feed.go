package main

import (
	"context"
	"fmt"
	"html"
)

func handlerFetchFeeds(s *state, cmd command) error {
	//if len(cmd.Args) < 2 {
	//	return fmt.Errorf("Not enough arguments")
	//}
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Println(html.UnescapeString(feed.Channel.Title))
	fmt.Println(html.UnescapeString(feed.Channel.Description))
	for i := 0; i < len(feed.Channel.Item); i++ {
		fmt.Println(html.UnescapeString(feed.Channel.Item[i].Title))
		fmt.Println(html.UnescapeString(feed.Channel.Item[i].Description))
	}

	return nil
}
