package main

import "fmt"

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
