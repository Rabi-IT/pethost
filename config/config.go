package config

import "os"

var (
	Port               = os.Getenv("PORT")
	ProductionDatabase = &DatabaseConfig{
		Host:         os.Getenv("DATABASE_HOST"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		User:         os.Getenv("DATABASE_USER"),
		Password:     os.Getenv("DATABASE_PASSWORD"),
		Port:         os.Getenv("DATABASE_PORT"),
	}
)

type DatabaseConfig struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         string
}
