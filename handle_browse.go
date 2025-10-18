package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Danjfreire/gator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {

	limit := 2
	if len(cmd.Args) >= 1 {
		// parse limit from cmd.Args[0]
		val, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = val
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	})

	if err != nil {
		log.Fatal("Error fetching posts for user:", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\nURL: %s\nPublished At: %s\n\n", post.Title, post.Url, post.PublishedAt)
		fmt.Println("Description:", post.Description)
		fmt.Println("--------------------------------------------------")
		fmt.Println()
	}

	return nil
}
