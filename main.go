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
	db.CreateDatabase()
	db.Connect()
	db.Migrate()

	httpServer := http.New(db)

	log.Fatal(httpServer.Start(config.Port))
}
