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

func CreateTokenForUser(data model.User) string {

	data.RegisteredClaims = jwt.RegisteredClaims{}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// Create the actual JWT token
	signedString, _ := token.SignedString([]byte(config.Config.Jwt.Secret))

	return signedString
}
