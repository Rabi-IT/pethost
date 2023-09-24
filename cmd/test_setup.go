package main

import (
	"log"
	"pethost/adapters/http"
	"pethost/config"
	"pethost/fixtures"
)

func main() {
	server := http.New(fixtures.TestDatabase)

	if err := fixtures.TestDatabase.CreateDatabase(); err != nil {
		panic(err)
	}

	if err := fixtures.TestDatabase.Connect(); err != nil {
		panic(err)
	}

	if err := fixtures.TestDatabase.Migrate(); err != nil {
		panic(err)
	}

	log.Fatal(server.Start(config.Port))
}
