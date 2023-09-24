package pet_case

import (
	"context"
)

func (c PetCase) Delete(ctx context.Context, id string) (bool, error) {
	return c.gateway.Delete(id)
}
