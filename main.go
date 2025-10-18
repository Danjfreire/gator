package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Danjfreire/gator/internal/config"
	"github.com/Danjfreire/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}
	db, err := sql.Open("postgres", cfg.DbUrl)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	dbQueries := database.New(db)
	appState := state{cfg: &cfg, db: dbQueries}
	commands := commands{Handlers: make(map[string]func(*state, command) error)}

	commands.register("login", handleLogin)
	commands.register("register", handleRegister)
	commands.register("reset", handleReset)
	commands.register("users", handleListUsers)
	commands.register("agg", handleAgg)
	commands.register("addfeed", middlewareLoggedIn(handleAddFeed))
	commands.register("feeds", handleFeeds)
	commands.register("follow", middlewareLoggedIn(handleFollow))
	commands.register("following", middlewareLoggedIn(handleFollowing))
	commands.register("unfollow", middlewareLoggedIn(handleUnfollow))

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: go run . <command> [args...]")
	}

	cmd := args[1]
	cmdArgs := args[2:]

	err = commands.run(&appState, command{Name: cmd, Args: cmdArgs})

	if err != nil {
		log.Fatal(err)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {

	return func(s *state, cmd command) error {
		userName := s.cfg.CurrentUserName
		ctx := context.Background()

		user, err := s.db.GetUser(ctx, userName)
		if err != nil {
			log.Fatal("Error getting user:", err)
		}

		return handler(s, cmd, user)
	}
}
