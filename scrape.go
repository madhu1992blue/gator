package main

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/madhu1992blue/gator/internal/database"
	"github.com/madhu1992blue/gator/internal/feedsApi"
)

func scrapeFeeds(s *state) error {
	feedRecord, err := s.dbQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	feed, err := feedsApi.FetchFeed(context.Background(), feedRecord.Url)
	if err != nil {
		return err
	}
	s.dbQueries.MarkFeedFetched(context.Background(), feedRecord.ID)
	for _, item := range feed.Channel.Items {
		timestamp, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			log.Printf("Unable to parse PubDate %s: %v", item.PubDate, err)
			return err
		}
		_, err = s.dbQueries.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			FeedID:      feedRecord.ID,
			PublishedAt: timestamp,
		})
		if err != nil {
			var pqError *pq.Error
			matched := errors.As(err, &pqError)
			if matched && strings.Contains(pqError.Message, "duplicate key") {
				continue
			} else {
				log.Printf("Error: %v", err)
			}
		}
	}
	return nil
}
