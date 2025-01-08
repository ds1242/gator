package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/ds1242/gator.git/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var numberOfPosts int
	if len(os.Args) == 1 {
		numberOfPosts = 2
	}
	numberOfPosts, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return err
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
