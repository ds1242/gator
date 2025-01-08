package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ds1242/gator.git/internal/database"
)

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	markedFeed, err := s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feedToFetch.ID,
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	})

	if err != nil {
		return err
	}

	rssfeed, err := fetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return err
	}

	for _, feed := range rssfeed.Channel.Item {
		msg := fmt.Sprintf("%s", feed.Title)
		fmt.Println(msg)
	}

	return nil

}
