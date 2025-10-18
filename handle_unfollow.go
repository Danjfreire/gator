package main

import (
	"context"
	"log"

	"github.com/Danjfreire/gator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		log.Fatal("Usage: go run . unfollow <feed_url>")
	}

	feedUrl := cmd.Args[0]

	ctx := context.Background()
	feed, err := s.db.FindFeedByUrl(ctx, feedUrl)
	if err != nil {
		log.Fatalf("Feed not found: %v", feedUrl)
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		log.Fatal("Error deleting feed follow:", err)
	}

	log.Printf("User %s has unfollowed feed %s\n", user.Name, feed.Name)
	return nil
}
