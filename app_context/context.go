package app_context

import (
	"context"

	"github.com/rs/zerolog"
)

type AppContext struct {
	ID      string
	Logger  *zerolog.Logger
	Context context.Context
	Session *UserSession
}

func New(ctx context.Context) *AppContext {
	session := getSession(ctx)
	id := getRequestId(ctx)

	return &AppContext{
		ID:      id,
		Context: ctx,
		Logger:  getLogger(ctx, session),
		Session: session,
	}
}

func getRequestId(ctx context.Context) string {
	id, ok := ctx.Value("requestIdKey").(string)
	if !ok {
		return ""
	}

	return id
}
