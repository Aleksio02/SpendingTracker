package controller

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/model"
	request "auth/cmd/auth/model/request"
	response "auth/cmd/auth/model/response"
	"auth/cmd/auth/service"
	"auth/cmd/auth/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	util.WriteBodyToObject(data.Request.Body, &requestBody)

	registeredUser, err := service.RegisterUser(requestBody)

	if err == nil {
		userResponse := response.UserInfoResponse{
			Status:   200,
			Message:  "User registered successfully",
			Token:    util.CreateTokenForUser(registeredUser),
			Username: registeredUser.Username,
			Role:     registeredUser.Role}

		data.JSON(http.StatusOK, userResponse)
		return
	}

	data.JSON(http.StatusBadRequest, response.UserInfoResponse{Status: 400, Message: "User with this name is exist"})
}
func Authorize(data *gin.Context) {
	requestBody := request.AuthRequest{}
	util.WriteBodyToObject(data.Request.Body, &requestBody)

	currentUser, err := service.Authorize(requestBody)

	if err != nil {
		data.JSON(http.StatusNotFound, response.UserInfoResponse{Status: 404, Message: "Incorrect username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, currentUser)
	signedString, _ := token.SignedString([]byte(config.Config.Jwt.Secret))

	responseInfo := response.UserInfoResponse{Status: 200, Message: "User authorized successfully", Token: signedString, Username: requestBody.Username, Role: currentUser.Role}
	data.JSON(http.StatusOK, responseInfo)
}
