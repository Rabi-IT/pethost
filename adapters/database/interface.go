package database

import "math"

type Database interface {
	Start() error
}

type PaginateInput struct {
	Page     int
	PageSize int
}

func CalcMaxPages(count int64, pageSize int) int {
	total := float64(count) / float64(pageSize)
	return int(math.Ceil(total))
}
