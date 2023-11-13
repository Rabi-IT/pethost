package schedule_case

import (
	"pethost/app_context"
	"pethost/frameworks/database"
	g "pethost/frameworks/database/gateways/schedule_gateway"
	"pethost/usecases/schedule_case/schedule_status"
)

type PaginateFilter struct {
	Status schedule_status.Status
}

func (c ScheduleCase) Paginate(ctx *app_context.AppContext, input PaginateFilter, paginate database.PaginateInput) (*g.PaginateOutput, error) {
	return c.gateway.Paginate(g.PaginateFilter{
		HostID: ctx.Session.UserID,
		Status: input.Status,
	}, paginate)
}
