package main

import (
	"context"

	"github.com/madhu1992blue/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, c *command, userRecord database.User) error) func(s *state, c *command) error {
	return func(s *state, c *command) error {
		userRecord, err := s.dbQueries.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, c, userRecord)
	}
}
