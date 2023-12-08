package model

import "github.com/golang-jwt/jwt/v4"

type User struct {
	jwt.RegisteredClaims
	Id       int
	Username string
	Role     string
}
