package app_context

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

func getLogger(ctx context.Context, session *UserSession) *zerolog.Logger {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(os.Stderr)

	return &logger
}

func AddTracer(ctx *AppContext, logger *zerolog.Event) *zerolog.Event {
	logger.Str(
		"request_id", ctx.ID,
	).Str(
		"user_id", ctx.Session.UserID,
	)

	if ctx.Session.Role.IsBackoffice() {
		logger.Str("original_user_id", ctx.Session.OriginalUserID)
	}

	return logger
}
