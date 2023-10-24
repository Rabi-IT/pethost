package app_context

import (
	"context"
	"pethost/usecases/auth_case/role"
)

type sessionKey string

const SessionKey sessionKey = "session"

type UserSession struct {
	UserID         string
	Name           string
	Login          string
	OriginalUserID string
	Role           role.Role
}

func getSession(ctx context.Context) *UserSession {
	session, ok := ctx.Value(SessionKey).(*UserSession)
	if !ok {
		return &UserSession{}
	}

	return session
}
