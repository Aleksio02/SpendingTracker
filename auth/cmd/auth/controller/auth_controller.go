package controller

import (
	"auth/cmd/auth/model"
	"auth/cmd/auth/model/response"
	"auth/cmd/auth/service"
	"auth/cmd/auth/util"
	request "auth/cmd/auth/model/request"
	response "auth/cmd/auth/model/response"
	"auth/cmd/auth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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

func RegisterUser(data *gin.Context) {

	requestBody := request.AuthRequest{}
	utils.WriteBodyToObject(data.Request.Body, &requestBody)

	responseInfo := response.UserInfoResponse{Status: 200, Message: "User registered successfully", Token: "jfdkjkjdkfhdjnkdjokdokspkp", Username: requestBody.Username, Role: "admin"}
	data.JSON(http.StatusOK, responseInfo)
}
