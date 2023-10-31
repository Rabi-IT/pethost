package schedule_gateway

import "pethost/frameworks/database/gorm_adapter/models"

func (g GormScheduleGatewayAdapter) Patch(filter PatchFilter, newValues PatchValues) (bool, error) {
	query := g.DB.Conn.Model(&models.Schedule{})

	if filter.ID != nil {
		query = query.Where("id = ?", filter.ID)
	} else {

		if filter.Status != nil {
			query = query.Where("status = ?", filter.Status)
		}

		if filter.Notes != nil {
			query = query.Where("notes = ?", filter.Notes)
		}

		if filter.TutorID != nil {
			query = query.Where("tutor_id = ?", filter.TutorID)
		}

		if filter.HostID != nil {
			query = query.Where("host_id = ?", filter.HostID)
		}

		if filter.Date != nil {
			query = query.Where("date = ?", filter.Date)
		}

		if filter.PetID != nil {
			query = query.Where("pet_id = ?", filter.PetID)
		}

	}

	result := query.Updates(newValues)
	return result.RowsAffected > 0, result.Error
}
