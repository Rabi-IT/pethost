package user_gateway

import (
	"pethost/frameworks/database"
	"pethost/frameworks/database/gorm_adapter"
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormUserGatewayAdapter) Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error) {
	data := []PaginateData{}

	query := g.DB.Conn.Model(&models.User{})

	if filter.Name != nil {
		query = query.Where("name = ?", filter.Name)
	}

	if filter.City != nil {
		query = query.Where("city = ?", filter.City)
	}

	if filter.State != nil {
		query = query.Where("state = ?", filter.State)
	}

	var count int64
	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return &PaginateOutput{
			Data:     data,
			MaxPages: 0,
		}, nil
	}

	gorm_adapter.Paginate(query, paginate)

	result = query.Scan(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	output := &PaginateOutput{
		Data:     data,
		MaxPages: database.CalcMaxPages(count, paginate.PageSize),
	}

	return output, nil
}
