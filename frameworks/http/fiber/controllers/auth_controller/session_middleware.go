package auth_controller

import (
	"errors"
	"fmt"
	"pethost/app_context"
	"pethost/usecases/auth_case/role"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Session(c *fiber.Ctx) error {
	token, ok := c.Context().UserValue("user").(*jwt.Token)
	if !ok || !token.Valid {
		return errors.New("INVALID_TOKEN")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("INVALID_CLAIMS")
	}

	c.Context().SetUserValue(
		app_context.SessionKey,
		&app_context.UserSession{
			UserID: fmt.Sprint(claims["user_id"]),
			Name:   fmt.Sprint(claims["name"]),
			Login:  fmt.Sprint(claims["login"]),
			Role:   role.Role(fmt.Sprint(claims["role"])),
		})

	return c.Next()
}
