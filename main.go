package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ds1242/gator.git/internal/config"
	"github.com/ds1242/gator.git/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error getting home environment")
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("error connecting to database")
	}

	dbQueries := database.New(db)

	programState := &state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := &commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerUserReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Not enough arguments entered")
		os.Exit(1)
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	os.Exit(0)

}
