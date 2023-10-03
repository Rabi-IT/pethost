package host_gateway

import (
	database "pethost/frameworks/database/gorm"
	"pethost/frameworks/database/gorm/models"
)

func (g GormPetHostGatewayAdapter) Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error) {
	data := []PaginateData{}

	query := g.DB.Conn.Model(&models.PetHost{})

	if filter.State != nil {
		query = query.Where("state = ?", filter.State)
	}

	if filter.Name != nil {
		query = query.Where("name = ?", filter.Name)
	}

	if filter.City != nil {
		query = query.Where("city = ?", filter.City)
	}

	if filter.Neighborhood != nil {
		query = query.Where("neighborhood = ?", filter.Neighborhood)
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

	database.Paginate(query, paginate)

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
