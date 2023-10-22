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
	if err := db.Start(); err != nil {
		panic(err)
	}

	httpServer := http.New(db)

	log.Fatal(httpServer.Start(config.Port))
}
