package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Danjfreire/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		log.Fatal("Usage: go run . add_feed <name> <url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	userName := s.cfg.CurrentUserName
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, userName)
	if err != nil {
		log.Fatal("Error getting user:", err)
	}

	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})

	s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		log.Fatal("Error creating feed:", err)
	}

	log.Println("Feed added successfully")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Println("Feed ID:", feed.ID)
	fmt.Println("Feed Name:", feed.Name)
	fmt.Println("Feed URL:", feed.Url)
	fmt.Println("Feed UserID:", feed.UserID)
	fmt.Println("Feed CreatedAt:", feed.CreatedAt)
	fmt.Println("Feed UpdatedAt:", feed.UpdatedAt)
}
