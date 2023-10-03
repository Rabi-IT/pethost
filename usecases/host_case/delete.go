package host_case

import (
	"context"
)

func (c PetHostCase) Delete(ctx context.Context, id string) (bool, error) {
	return c.gateway.Delete(id)
}
