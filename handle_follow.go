package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Danjfreire/gator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		log.Fatal("Usage: go run . follow <feed_url>")
	}

	feedURL := cmd.Args[0]

	ctx := context.Background()

	feed, err := s.db.FindFeedByUrl(ctx, feedURL)
	if err != nil {
		log.Fatalf("Feed not found: %v", feedURL)
	}

	now := time.Now()
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		log.Fatal("Error creating feed follow:", err)
	}

	fmt.Printf("User %s is now following feed %s\n", user.Name, feedFollow.FeedName)

	return nil
}
