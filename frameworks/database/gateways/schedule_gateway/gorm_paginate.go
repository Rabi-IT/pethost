package schedule_gateway

import (
	"pethost/frameworks/database"
	"pethost/frameworks/database/gorm_adapter"
	"pethost/frameworks/database/gorm_adapter/models"
)

func (g GormScheduleGatewayAdapter) Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error) {
	data := []PaginateData{}

	query := g.DB.Conn.Model(&models.Schedule{}).Where(
		"status = ?", filter.Status,
	).Where(
		"host_id = ?", filter.HostID,
	)

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
