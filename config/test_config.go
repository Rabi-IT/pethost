package config

import (
	"os"
)

var TestDatabase = &DatabaseConfig{
	Host:         testEnv("TEST_DATABASE_HOST", "localhost"),
	User:         testEnv("TEST_DATABASE_USER", "postgres"),
	Password:     testEnv("TEST_DATABASE_PASSWORD", "root"),
	DatabaseName: testEnv("TEST_DATABASE_NAME", "pethost_test"),
	Port:         testEnv("TEST_DATABASE_PORT", "5432"),
}

func testEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
