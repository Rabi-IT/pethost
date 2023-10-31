package schedule_case

import (
	"context"
	g "pethost/frameworks/database/gateways/schedule_gateway"
)

type CreateInput struct {
	HostID  string
	Date    string
	PetID   string
	Status  string
	Notes   string
	TutorID string
}

func (c ScheduleCase) Create(ctx context.Context, input *CreateInput) (string, error) {
	return c.gateway.Create(g.CreateInput{
		HostID:  input.HostID,
		Date:    input.Date,
		PetID:   input.PetID,
		Status:  input.Status,
		Notes:   input.Notes,
		TutorID: input.TutorID,
	})
}
