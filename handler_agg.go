package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("A duration is required")
	}

	time_between_reqs, _ := time.ParseDuration(cmd.Args[0])

	fmt.Println("Collecting feeds every ", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}
