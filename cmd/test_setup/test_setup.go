package main

import (
	"log"
	"pethost/config"
	"pethost/fixtures"
	"pethost/frameworks/http/fiber_adapter"
)

func main() {
	server := fiber_adapter.New(fixtures.TestDatabase)

	if err := fixtures.TestDatabase.Start(); err != nil {
		panic(err)
	}

	log.Fatal(server.Start(config.TestPort))
}
