package main

import (
	"log"
	"pethost/adapters/http"
	"pethost/config"
	"pethost/factories"
	"time"
)

func main() {
	time.Local = time.UTC

	db := factories.NewProductionDatabase()
	if err := db.CreateDatabase(); err != nil {
		panic(err)
	}

	if err := db.Connect(); err != nil {
		panic(err)
	}

	if err := db.Migrate(); err != nil {
		panic(err)
	}

	httpServer := http.New(db)

	log.Fatal(httpServer.Start(config.Port))
}
