package main

import (
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
