package auth

import (
	"goth/internal/db"

	"github.com/golang-jwt/jwt"
)

type TokenAuth interface {
	GenerateToken(user db.User) (string, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}
