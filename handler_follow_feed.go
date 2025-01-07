package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ds1242/gator.git/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("A url is required")
	}

	feed_follow, err := s.db.FollowFeeds(context.Background(), database.FollowFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	return nil
}
