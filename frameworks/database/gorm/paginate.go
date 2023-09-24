package gorm_adapter

import (
	"math"

	"gorm.io/gorm"
)

type PaginateInput struct {
	Page     int
	PageSize int
}

func Paginate(db *gorm.DB, paginate PaginateInput) {
	offset := paginate.Page * paginate.PageSize
	db.Limit(paginate.PageSize).Offset(offset)
}

func CalcMaxPages(count int64, pageSize int) int {
	total := float64(count) / float64(pageSize)
	return int(math.Ceil(total))
}
