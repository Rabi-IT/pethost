package schedule_gateway

import (
	"pethost/frameworks/database/gorm_adapter/models"
	"pethost/usecases/schedule_case/schedule"
	"strings"

	"github.com/google/uuid"
)

func adaptError(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "unique_schedule_date") {
		return schedule.ErrAlreadyScheduled
	}

	return err
}

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

	return id, adaptError(result.Error)
}
