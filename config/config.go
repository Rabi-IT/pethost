package config

import (
	"fmt"
	"os"
)

var (
	Port               = requiredEnv("PORT")
	ProductionDatabase = &DatabaseConfig{
		Host:         requiredEnv("DATABASE_HOST"),
		DatabaseName: requiredEnv("DATABASE_NAME"),
		User:         requiredEnv("DATABASE_USER"),
		Password:     requiredEnv("DATABASE_PASSWORD"),
		Port:         requiredEnv("DATABASE_PORT"),
	}
)

type DatabaseConfig struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
	Port         string
}

func requiredEnv(key string) string {
	value := os.Getenv("DATABASE_HOST")

	if value == "" {
		panic(fmt.Sprintf(`Err: missing environment variable: "%s"`, key))
	}

	return value
}
