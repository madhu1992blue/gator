package main

import (
	"context"
	"fmt"
	"errors"
	"github.com/google/uuid"
	"github.com/madhu1992blue/gator/internal/database"
	"time"
)

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
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: username,
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
