package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Danjfreire/gator/internal/config"
)

type state struct {
	Config *config.Config
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	appState := state{Config: &cfg}
	commands := commands{Handlers: make(map[string]func(*state, command) error)}

	commands.register("login", handleLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run . <command> [args...]")
		os.Exit(1)
	}

	cmd := args[1]
	cmdArgs := args[2:]

	err = commands.run(&appState, command{Name: cmd, Args: cmdArgs})

	if err != nil {
		log.Fatal(err)
	}

}
