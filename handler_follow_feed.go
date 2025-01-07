package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ds1242/gator.git/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("A url is required")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	feed_follow, err := s.db.FollowFeeds(context.Background(), database.FollowFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})

	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Feed %s followed by %s\n", feed_follow.FeedName, feed_follow.UserName)
	fmt.Println(msg)

	return nil
}
