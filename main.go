package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"log"
	"github.com/madhu1992blue/gator/internal/config"
	"github.com/madhu1992blue/gator/internal/database"
)
func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Couldn't read config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	dbQueries := database.New(db)
	
	cmds := commands{}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	st := state{
		config: cfg,
		dbQueries: dbQueries,
	}
	if len(os.Args) < 2 {
		log.Fatalf("Subcommand not specified")
	}
	cmd := command {
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = cmds.run(&st, &cmd)
	if err != nil {
		log.Fatalf("Couldn't execute command: %v", err)
	}
}
