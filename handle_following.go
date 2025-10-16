package main

import (
	"context"
	"fmt"
	"log"
)

func handleFollowing(s *state, cmd command) error {

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		log.Fatal("Error getting feed follows:", err)
	}

	fmt.Println("User %v is following these feeds:", user.Name)
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
