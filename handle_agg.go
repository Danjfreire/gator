package main

import (
	"context"
	"fmt"
	"log"
	"time"
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
	nextToFetch, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Println("Couldn't get next feed to fetch:", err)
		return
	}

	s.db.MarkFeedFetched(ctx, nextToFetch.ID)

	feed, err := fetchFeed(ctx, nextToFetch.Url)
	if err != nil {
		fmt.Println("Error fetching feed:", err)
		return
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("Found post: %v\n", item.Title)
	}

	fmt.Printf("Feed %v fetched successfully! %v posts found!\n", nextToFetch.Url, len(feed.Channel.Item))
}
