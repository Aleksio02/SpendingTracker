package controller

import (
	"auth/cmd/auth/model"
	request "auth/cmd/auth/model/request"
	response "auth/cmd/auth/model/response"
	"auth/cmd/auth/service"
	"auth/cmd/auth/util"
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
	//Приходит имя пользователя и пароль в теле запроса(38)

	//Парсим в объект AuthRequest(42)
	requestBody := request.AuthRequest{}
	util.WriteBodyToObject(data.Request.Body, &requestBody)

	registeredUser, err := service.RegisterUser(requestBody)

	//Пользователь зарегистрирован (сохранен в БД) (53)
	if err == nil {
		////Регистрируем пользователя(56-61)
		userResponse := response.UserInfoResponse{
			Status:   200,
			Message:  "User registered successfully",
			Token:    util.CreateTokenForUser(registeredUser),
			Username: registeredUser.Username,
			Role:     registeredUser.Role}

		//Выводим статус 200(64)
		data.JSON(http.StatusOK, userResponse)
		return
	}

	//Пользователь с таким именем найден(64)
	data.JSON(http.StatusBadRequest, response.UserInfoResponse{Status: 400, Message: "User with this name is exist"})
}
