package main

import (
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/google/uuid"
)

func handleFeeds(s *state, command command) error {
	ctx := context.Background()

	feeds, err := s.db.ListFeeds(ctx)
	if err != nil {
		return err
	}

	userIds := make(map[uuid.UUID]string)
	for _, feed := range feeds {
		userIds[feed.UserID] = ""
	}

	users, err := s.db.FindManyUsersById(ctx, slices.Collect(maps.Keys(userIds)))
	if err != nil {
		return err
	}

	for _, user := range users {
		userIds[user.ID] = user.Name
	}

	for _, feed := range feeds {
		printFeed(feed)
		if userName, ok := userIds[feed.UserID]; ok {
			fmt.Println("User:", userName)
		}

		fmt.Println()
	}

	return nil
}
