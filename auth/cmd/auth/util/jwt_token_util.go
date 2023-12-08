package util

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/model"
	"github.com/golang-jwt/jwt/v4"
)

func ParseJWTToken(jwtToken string, objectToWrite *model.User) (*jwt.Token, error) {
	return jwt.ParseWithClaims(jwtToken, objectToWrite, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Jwt.Secret), nil
	})
}
