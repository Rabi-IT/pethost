package http

import (
	"pethost/adapters/database"
)

func New(d database.Database) HTTPServer {
	return newFiber(d)
}
