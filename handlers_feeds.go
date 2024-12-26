package main

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/madhu1992blue/gator/internal/database"
)

type FeedDTO struct {
	ID        uuid.UUID `xml:"ID"`
	CreatedAt time.Time `xml:"created_at"`
	UpdatedAt time.Time `xml:"updated_at"`
	Name      string    `xml:"name"`
	Url       string    `xml:"url"`
	UserID    uuid.UUID `xml:"user_id"`
}

func handlerListFeeds(s *state, _ *command) error {
	type FeedWithUsernameDTO struct {
		Name     string `xml:"name"`
		Url      string `xml:"url"`
		Username string `xml:"username"`
	}
	feedRecords, err := s.dbQueries.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	resultDTOs := []FeedWithUsernameDTO{}
	for _, fr := range feedRecords {
		resultDTOs = append(resultDTOs, FeedWithUsernameDTO{
			Name:     fr.Name,
			Url:      fr.Url,
			Username: fr.Username,
		})
	}
	dataBytes, err := xml.Marshal(resultDTOs)
	if err != nil {
		return err
	}
	fmt.Println(string(dataBytes))
	return nil
}

func handlerAddFeed(s *state, c *command, userRecord database.User) error {
	if len(c.args) < 2 {
		return errors.New("addfeed: needs 2 arguments, name and url")
	}
	feedRecord, err := s.dbQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      c.args[0],
		Url:       c.args[1],
		UserID:    userRecord.ID,
	})
	if err != nil {
		return err
	}
	_, err = s.dbQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  s.config.CurrentUserName,
		FeedUrl:   feedRecord.Url,
	})
	if err != nil {
		return err
	}
	feedBytes, err := xml.Marshal(FeedDTO{
		ID:        feedRecord.ID,
		CreatedAt: feedRecord.CreatedAt,
		UpdatedAt: feedRecord.UpdatedAt,
		Name:      feedRecord.Name,
		Url:       feedRecord.Name,
		UserID:    feedRecord.UserID,
	})
	if err != nil {
		return err
	}
	fmt.Println(string(feedBytes))
	return nil
}
