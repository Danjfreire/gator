package main

import (
	"fmt"
	"os"

	"github.com/Danjfreire/gator/internal/config"
)

type state struct {
	Config *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	Handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {

	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.Handlers[name] = handler
}

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}

	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("User set to", cmd.Args[0])

	return nil
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

	err = commands.run(&appState, command{Name: args[1], Args: args[2:]})

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
