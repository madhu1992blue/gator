package main

import (
	"context"
	"fmt"

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
		fmt.Println(item.Title)
	}

	return nil
}
