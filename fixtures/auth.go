package fixtures

import (
	"pethost/config"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func SystemToken(t *testing.T) string {
	claims := jwt.MapClaims{
		"user_id": "system",
		"name":    "system",
		"email":   "system@system.com",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tk, err := token.SignedString([]byte(config.AuthSecret))
	require.Nil(t, err)

	return tk
}
