package database

type Database interface {
	Connect() error
	CreateDatabase() error
	Migrate() error
}
