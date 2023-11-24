package service_rating_case

import (
	"context"
)

func (c ServiceRatingCase) Delete(ctx context.Context, id string) (bool, error) {
	return c.gateway.Delete(id)
}
