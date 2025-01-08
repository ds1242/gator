package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ds1242/gator.git/internal/database"
	"github.com/google/uuid"
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

		// List of potential layouts
		layouts := []string{
			time.RFC3339,
			time.UnixDate,
			time.RFC1123Z,
			// Add custom layouts as needed
		}

		var timeConverted time.Time
		var err error

		// Attempt parsing with each layout
		for _, layout := range layouts {
			timeConverted, err = time.Parse(layout, feed.PubDate)
			if err == nil {
				break
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     feed.Title,
			Description: sql.NullString{
				String: feed.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  timeConverted,
				Valid: true,
			},
			FeedID: uuid.NullUUID{
				UUID:  feedToFetch.ID,
				Valid: true,
			},
			Url: feed.Link,
		})

		if err != nil {
			fmt.Println(err)
		}

	}

	return nil

}
