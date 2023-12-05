package controller

import (
	request "auth/cmd/auth/model/request"
	response "auth/cmd/auth/model/response"
	"auth/cmd/auth/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetUserByToken(c *gin.Context) {
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	fmt.Println(token)

	userResponse := response.UserInfoResponse{Status: 200, Message: "Success", Token: token, Username: "Supervisor", Role: "admin"}

	c.JSON(http.StatusOK, userResponse)
}

func RegistrationUser(data *gin.Context) {
	token := strings.Split(data.GetHeader("Authorization"), " ")[1]
	fmt.Println(token)

	requestBody := request.AuthInfoRequest{}
	utils.WriteBodyToObject(data.Request.Body, &requestBody)

	if requestBody.Status == 200 {
		responseInfo := response.UserInfoResponse{Status: 200, Message: "User registered successfully", Token: token, Username: "Supervisor", Role: "admin"}
		data.JSON(http.StatusOK, responseInfo)
	} else if requestBody.Status == 400 {
		responseInfo := response.UserInfoResponse{Status: 400, Message: "User with this name is exist"}
		data.JSON(http.StatusBadRequest, responseInfo)
	}
}
