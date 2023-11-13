package main

import (
	"log"
	"pethost/config"
	"pethost/fixtures"
	"pethost/frameworks/http/fiber_adapter"
	"time"
)

func main() {
	time.Local = time.UTC
	server := fiber_adapter.New(fixtures.TestDatabase)

	if err := fixtures.TestDatabase.Start(); err != nil {
		panic(err)
	}

	log.Fatal(server.Start(config.TestPort))
}
