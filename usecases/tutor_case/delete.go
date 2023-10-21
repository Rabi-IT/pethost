package tutor_case

import (
	"context"
)

func (c TutorCase) Delete(ctx context.Context, id string) (bool, error) {
	return c.gateway.Delete(id)
}
