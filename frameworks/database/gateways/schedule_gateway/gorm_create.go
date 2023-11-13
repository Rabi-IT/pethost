package schedule_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"

	"github.com/google/uuid"
)

func (g *GormScheduleGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.Schedule{
		ID:          id,
		DaysOfMonth: input.DaysOfMonth,
		PetID:       input.PetID,
		Status:      input.Status,
		Notes:       input.Notes,
		TutorID:     input.TutorID,
		HostID:      input.HostID,
		MonthYear:   input.MonthYear,
	})

	return id, result.Error
}
