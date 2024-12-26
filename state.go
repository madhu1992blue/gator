package main
import (
	"github.com/madhu1992blue/gator/internal/config"
	"github.com/madhu1992blue/gator/internal/database"
)
type state struct {
	config *config.Config
	dbQueries *database.Queries
}
