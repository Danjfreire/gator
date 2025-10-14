package main

import "fmt"

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
