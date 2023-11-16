package schedule_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"

	"github.com/google/uuid"
)

func (g *GormScheduleGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.Schedule{
		// TODO FIX
		ID:          id,
		DaysOfMonth: input.Date[0].DaysOfMonth,
		PetID:       input.PetIDs[0],
		Status:      input.Status,
		Notes:       input.Notes,
		TutorID:     input.TutorID,
		HostID:      input.HostID,
		MonthYear:   input.Date[0].MonthYear,
	})

	return id, result.Error
}
