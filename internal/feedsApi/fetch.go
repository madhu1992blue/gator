package feedsApi

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	request.Header.Set("User-Agent", "gator")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var rssFeed RSSFeed
	decoder := xml.NewDecoder(resp.Body)
	if err := decoder.Decode(&rssFeed); err != nil {
		return nil, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)

	for itemIdx := range rssFeed.Channel.Items {
		rssFeed.Channel.Items[itemIdx].Title = html.UnescapeString(rssFeed.Channel.Items[itemIdx].Title)
		rssFeed.Channel.Items[itemIdx].Description = html.UnescapeString(rssFeed.Channel.Items[itemIdx].Description)
	}

	return &rssFeed, nil
}
