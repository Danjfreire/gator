package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {
	ctx := context.Background()

	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	printRSSFeed(*feed)

	return nil
}

func printRSSFeed(feed RSSFeed) {

	fmt.Println("Feed Title:", feed.Channel.Title)
	fmt.Println("Feed Link:", feed.Channel.Link)
	fmt.Println("Feed Description:", feed.Channel.Description)
	fmt.Println()

	for _, item := range feed.Channel.Item {
		fmt.Println("Item Title:", item.Title)
		fmt.Println("Item Link:", item.Link)
		fmt.Println("Item Description:", item.Description)
		fmt.Println("Item PubDate:", item.PubDate)
		fmt.Println()
	}
}
