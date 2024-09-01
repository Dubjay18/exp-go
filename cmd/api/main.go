package main

import (
	"exp-go/internal/database"
	"exp-go/internal/server"
	_ "exp-go/migrations"
	"fmt"
	"github.com/pressly/goose/v3"
)

func main() {

	server := server.NewServer()
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
		panic(fmt.Sprintf("error in migration: %s", err))
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
