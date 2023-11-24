package schedule_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/frameworks/database/gateways/schedule_gateway/ports"
	"pethost/usecases/schedule_case/schedule_status"
	"time"
)

type PatchFilter struct {
	ID      string
	TutorID *string
	HostID  *string
}

type PatchValues struct {
	PetID   string
	Status  schedule_status.Status
	Notes   string
	TutorID string
	HostID  string
}

func (c ScheduleCase) Patch(ctx *app_context.AppContext, filter PatchFilter, values PatchValues) (bool, error) {
	switch values.Status {
	case schedule_status.Accepted:
		return c.accept(ctx, filter, values)

	case schedule_status.Canceled:
		return c.cancel(ctx, filter, values)

	case schedule_status.Paid, schedule_status.PaidAccepted:
		return c.pay(ctx, filter, values)

	default:
		return false, nil
	}
}

func (c ScheduleCase) accept(ctx *app_context.AppContext, filter PatchFilter, values PatchValues) (bool, error) {
	schedule, err := c.gateway.GetByID(filter.ID)
	if err != nil || schedule == nil {
		return false, err
	}

	var newStatus schedule_status.Status
	if schedule_status.Open == schedule.Status {
		newStatus = schedule_status.Accepted
	} else if schedule_status.Paid == schedule.Status {
		newStatus = schedule_status.PaidAccepted
	} else {
		return false, nil
	}

	return c.gateway.Patch(
		g.PatchFilter{
			ID:     filter.ID,
			HostID: &ctx.Session.UserID,
			Status: &schedule.Status,
		}, g.PatchValues{
			Status: &newStatus,
			History: &ports.ScheduleHistory{
				UserID:    ctx.Session.GetOriginalUser(),
				NewStatus: newStatus,
				Date:      time.Now(),
				Notes:     values.Notes,
			},
		},
	)
}

func (c ScheduleCase) cancel(ctx *app_context.AppContext, filter PatchFilter, values PatchValues) (bool, error) {
	return c.gateway.Patch(
		g.PatchFilter{
			ID:     filter.ID,
			HostID: &ctx.Session.UserID,
			StatusOR: []schedule_status.Status{
				schedule_status.Open,
				schedule_status.Paid,
				schedule_status.PaidAccepted,
			},
		}, g.PatchValues{
			Status: &schedule_status.Canceled,
			History: &ports.ScheduleHistory{
				UserID:    ctx.Session.GetOriginalUser(),
				NewStatus: schedule_status.Canceled,
				Date:      time.Now(),
				Notes:     values.Notes,
			},
		},
	)
}

func (c ScheduleCase) pay(ctx *app_context.AppContext, filter PatchFilter, values PatchValues) (bool, error) {
	if ctx.Session.Role.IsUser() {
		return false, nil
	}

	schedule, err := c.gateway.GetByID(filter.ID)
	if err != nil || schedule == nil {
		return false, err
	}

	var newStatus schedule_status.Status
	if schedule_status.Open == schedule.Status {
		newStatus = schedule_status.Paid
	} else if schedule_status.Accepted == schedule.Status {
		newStatus = schedule_status.PaidAccepted
	} else {
		return false, nil
	}

	return c.gateway.Patch(
		g.PatchFilter{
			ID:     filter.ID,
			Status: &schedule.Status,
		}, g.PatchValues{
			Status: &newStatus,
			History: &ports.ScheduleHistory{
				UserID:    ctx.Session.GetOriginalUser(),
				NewStatus: newStatus,
				Date:      time.Now(),
				Notes:     values.Notes,
			},
		},
	)
}
