package schedule_gateway

import (
	"pethost/frameworks/database"
	"pethost/frameworks/database/gorm_adapter"
	"pethost/frameworks/database/gorm_adapter/models"
)

var EMPTY_PAGE = &PaginateOutput{
	Data:     []PaginateData{},
	MaxPages: 0,
}

func (g GormScheduleGatewayAdapter) Paginate(filter PaginateFilter, paginate database.PaginateInput) (*PaginateOutput, error) {

	query := g.DB.Conn.Model(&models.Schedule{})

	if filter.Status == nil {
		query.Where("status = ?", filter.Status)
	}

	if filter.HostID != nil {
		query.Where("host_id = ?", filter.HostID)
	}

	if filter.TutorID != nil {
		query.Where("tutor_id = ?", filter.TutorID)
	}

	var count int64
	result := query.Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 || count == 0 {
		return EMPTY_PAGE, nil
	}

	gorm_adapter.Paginate(query, paginate)

	model := []models.Schedule{}
	result = query.Scan(&model)

	if result.Error != nil {
		return nil, result.Error
	}

	data := make([]PaginateData, len(model))
	for i, m := range model {
		data[i] = PaginateData{
			PetIDs:    m.PetIDs,
			TutorID:   m.TutorID,
			StartDate: m.StartDate,
			EndDate:   m.EndDate,
			Status:    m.Status,
			Notes:     m.Notes,
		}
	}

	output := &PaginateOutput{
		Data:     data,
		MaxPages: database.CalcMaxPages(count, paginate.PageSize),
	}

	return output, nil
}
