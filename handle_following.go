package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Danjfreire/gator/internal/database"
)

func handleFollowing(s *state, cmd command, user database.User) error {

	ctx := context.Background()

	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		log.Fatal("Error getting feed follows:", err)
	}

	fmt.Printf("User %v is following these feeds:\n", user.Name)
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
