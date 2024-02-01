package tokenauth

import (
	"errors"
	"goth/internal/store"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
)

type TokenAuth struct {
	JWTAuth   *jwtauth.JWTAuth
	secretKey []byte
}

type NewTokenAuthParams struct {
	SecretKey []byte
}

func NewTokenAuth(params NewTokenAuthParams) *TokenAuth {
	jwtAuth := jwtauth.New("HS256", []byte(params.SecretKey), nil)
	return &TokenAuth{
		JWTAuth:   jwtAuth,
		secretKey: params.SecretKey,
	}

}

func (a *TokenAuth) GenerateToken(user store.User) (string, error) {

	payload := map[string]interface{}{
		"email": user.Email,
		"exp": time.Now().Add(time.Hour *
			24).Unix(), // Token expires in 24 hours
	}

	_, tokenString, err := a.JWTAuth.Encode(payload)

	// tokenString, err := token.SignedString(a.SecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *TokenAuth) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
