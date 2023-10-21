package main

import (
	"log"
	"pethost/adapters/http"
	"pethost/config"
	"pethost/fixtures"
)

func main() {
	server := http.New(fixtures.TestDatabase)

	if err := fixtures.TestDatabase.Start(); err != nil {
		panic(err)
	}

	log.Fatal(server.Start(config.TestPort))
}
