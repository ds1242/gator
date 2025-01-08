package main

import (
	"context"
	"fmt"

	"github.com/ds1242/gator.git/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("please provide a feed to unfollow")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	err = s.db.Unfollow(context.Background(), database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return nil
	}

	fmt.Println("unfollowed")

	return nil

}
