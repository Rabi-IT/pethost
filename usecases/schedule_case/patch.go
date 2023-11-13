package schedule_case

import (
	"context"
	g "pethost/frameworks/database/gateways/schedule_gateway"
)

type PatchFilter struct {
	ID      *string
	PetID   *string
	Status  *string
	Notes   *string
	TutorID *string
	HostID  *string
	Date    *string
}

type PatchValues struct {
	PetID   string
	Status  string
	Notes   *string
	TutorID string
	HostID  string
	Date    string
}

func (c ScheduleCase) Patch(ctx context.Context, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(
		g.PatchFilter{
			ID:      filter.ID,
			PetID:   filter.PetID,
			Status:  filter.Status,
			Notes:   filter.Notes,
			TutorID: filter.TutorID,
			HostID:  filter.HostID,
			Date:    filter.Date,
		}, g.PatchValues{
			Status: values.Status,
		})
}
