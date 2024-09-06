package main

import (
	"exp-go/internal/database"
	"exp-go/internal/server"
	_ "exp-go/migrations"
	"flag"
	"fmt"

	"github.com/pressly/goose/v3"
)

func main() {
	// Define the flag for running the down migration
	downFlag := flag.Bool("down", false, "Run goose down instead of starting the server")

	// Parse the command-line flags
	flag.Parse()
	server := server.NewServer()
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if *downFlag {
		fmt.Println("Running goose down migration...")
		if err := goose.Down(database.SQL_DB, "migrations"); err != nil {
			panic(fmt.Sprintf("error running goose down: %s", err))
		}
		fmt.Println("Goose down migration completed.")
		return // Exit after running the migration
	}
	if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
		panic(fmt.Sprintf("error in migration: %s", err))
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
