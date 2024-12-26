package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/madhu1992blue/gator/internal/database"
)

func handlerBrowse(s *state, c *command, userRec database.User) error {
	bNum := 2
	if len(c.args) >= 1 {
		// We have a default, so let's ignore if its not parse-able to int
		bNum, _ = strconv.Atoi(c.args[0])
	}

	posts, err := s.dbQueries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: userRec.ID,
		Limit:  int32(bNum),
	})
	if err != nil {
		return err
	}
	for _, p := range posts {
		fmt.Println(p.Title)
	}
	return nil
}
