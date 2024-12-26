package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/madhu1992blue/gator/internal/database"
)

func handlerAgg(s *state, c *command) error {
	if len(c.args) < 1 {
		return errors.New("agg: expects an argument of duration like 1s, 1m, 1h")
	}
	timeDurationInput := c.args[0]
	time_between_reqs, err := time.ParseDuration(timeDurationInput)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil

}
func handlerLogin(s *state, cmd *command) error {
	if len(cmd.args) == 0 {
		return errors.New("login: expects a username argument")
	}
	username := cmd.args[0]
	userRecord, err := s.dbQueries.GetUser(context.Background(), username)
	if err != nil {
		return err
	}
	err = s.config.SetUser(userRecord.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Username has been set to %s\n", userRecord.Name)
	return nil

}

func handlerRegister(s *state, cmd *command) error {
	if len(cmd.args) == 0 {
		return errors.New("register: expects a username argument")
	}
	username := cmd.args[0]
	userRecord, err := s.dbQueries.CreateUser(context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      username,
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("User,%s, was created. Switching to %s\n", userRecord.Name, userRecord.Name)
	err = s.config.SetUser(userRecord.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Switched to %s\n", userRecord.Name)
	return nil
}

func handlerListUsers(s *state, cmd *command) error {
	userRecords, err := s.dbQueries.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, userRecord := range userRecords {
		fmt.Print("* ", userRecord.Name)
		if userRecord.Name == s.config.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	return nil
}
func handlerReset(s *state, cmd *command) error {
	err := s.dbQueries.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("DANGER: Dropped all users")
	return nil
}
