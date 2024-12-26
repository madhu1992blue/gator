package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/madhu1992blue/gator/internal/database"
)

func handlerGetFeedFollowsForUser(s *state, _ *command, userRecord database.User) error {
	feedFollowsForUserRows, err := s.dbQueries.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	fmt.Println("Username\t\tFeedname")
	for _, row := range feedFollowsForUserRows {
		fmt.Printf("%s\t\t%s", row.UserName, row.FeedName)
	}
	return nil
}

func handlerFollowFeed(s *state, c *command, user database.User) error {
	if len(c.args) < 1 {
		return errors.New("follow: expects a URL to follow")
	}
	feedUrl := c.args[0]
	createFeedFollowRow, err := s.dbQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  s.config.CurrentUserName,
		FeedUrl:   feedUrl,
	})
	if err != nil {
		return err
	}
	type RespDTO struct {
		FeedName string `xml:"feed_name"`
		UserName string `xml:"username"`
	}

	resp := RespDTO{
		FeedName: createFeedFollowRow.FeedName,
		UserName: createFeedFollowRow.UserName,
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	fmt.Println(string(respBytes))
	return nil
}
