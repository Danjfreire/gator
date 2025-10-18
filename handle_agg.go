package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Danjfreire/gator/internal/database"
	"github.com/google/uuid"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		log.Fatal("Usage: go run . agg <time_between_reqs>")
	}

	timeBetweenReqsStr := cmd.Args[0]
	duration, err := time.ParseDuration(timeBetweenReqsStr)
	if err != nil {
		log.Fatal("Invalid duration format:", err)
	}

	fmt.Printf("Collecting feeds every %s\n", duration)

	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Println("Couldn't get next feed to fetch:", err)
		return
	}

	s.db.MarkFeedFetched(ctx, feed.ID)

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Println("Error fetching feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Found post: %v\n", item.Title)
		t, err := time.Parse(time.RFC1123, item.PubDate)

		if err != nil {
			fmt.Println("Error parsing publication date:", err)
			continue
		}

		s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: t,
			FeedID:      feed.ID,
		})
	}

	fmt.Printf("Feed %v fetched successfully! %v posts found!\n", feed.Url, len(rssFeed.Channel.Item))
}
