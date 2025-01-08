package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ds1242/gator.git/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var numberOfPosts int
	if len(cmd.Args) == 0 {
		numberOfPosts = 2
	}

	if len(cmd.Args) == 1 {
		postInput, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		numberOfPosts = postInput
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(numberOfPosts),
	})
	if err != nil {
		return err
	}

	fmt.Println("===================")
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.Url)
		fmt.Println("===================")
	}

	return nil

}
