package preference_case

import (
	"pethost/app_context"
	g "pethost/frameworks/database/gateways/preference_gateway"
)

func (c *PreferenceCase) GetByFilter(ctx *app_context.AppContext, filter *g.GetByFilterInput) (*g.GetByFilterOutput, error) {
	return c.gateway.GetByFilter(filter)
}
