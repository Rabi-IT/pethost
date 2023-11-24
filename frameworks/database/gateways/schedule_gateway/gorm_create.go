package schedule_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"

	"github.com/google/uuid"
)

func (g *GormScheduleGatewayAdapter) Create(input CreateInput) (string, error) {
	id := uuid.NewString()

	result := g.DB.Conn.Create(&models.Schedule{
		ID:        id,
		Status:    input.Status,
		PetIDs:    input.PetIDs,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		Notes:     input.Notes,
		TutorID:   input.TutorID,
		HostID:    input.HostID,
	})

	return id, result.Error
}
