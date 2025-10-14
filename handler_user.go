package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Danjfreire/gator/internal/database"
	"github.com/google/uuid"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}

	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		log.Fatal("User not found:", name)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User set to", user.Name)
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("name is required")
	}
	name := cmd.Args[0]

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, name)
	if err == nil {
		log.Fatal("User already exists:", name)
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	s.cfg.SetUser(user.Name)
	fmt.Println("User registered successfully:", user.Name)
	fmt.Println("user:", user)
	return nil
}

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatal("Error deleting users:", err)
	}

	return nil
}

func handleListUsers(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		log.Fatal("Error listing users:", err)
	}

	for _, u := range users {
		if u.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", u.Name)

		} else {
			fmt.Printf("* %v\n", u.Name)
		}
	}

	return nil
}
