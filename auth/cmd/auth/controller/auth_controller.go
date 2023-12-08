package controller

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/model"
	"auth/cmd/auth/model/response"
	"auth/cmd/auth/service"
	"auth/cmd/auth/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetUserByToken(c *gin.Context) {
	authorizationHeaderValue := c.GetHeader("Authorization")
	if len(authorizationHeaderValue) > 0 {
		jwtToken := strings.Split(authorizationHeaderValue, " ")[1]
		decodedUser := model.User{}
		token, err := util.ParseJWTToken(jwtToken, &decodedUser)

		if err == nil && token.Valid {
			foundUser, err := service.GetUserByUsername(decodedUser)
			if err == nil {
				userResponse := response.UserInfoResponse{Status: 200,
					Message:  "Success",
					Token:    jwtToken,
					Username: foundUser.Username,
					Role:     foundUser.Role}
				c.JSON(http.StatusOK, userResponse)
				return
			}
		}
	}

	c.JSON(http.StatusUnauthorized, response.UserInfoResponse{Status: 401, Message: "User is not authorized"})
}
