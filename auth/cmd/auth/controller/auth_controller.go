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

func RegisterUser(data *gin.Context) {
	token := strings.Split(data.GetHeader("Authorization"), " ")[1]
	fmt.Println(token)

	requestBody := request.AuthRequest{}
	utils.WriteBodyToObject(data.Request.Body, &requestBody)

	responseInfo := response.UserInfoResponse{Status: 200, Message: "User registered successfully", Token: token, Username: requestBody.Username, Role: "admin"}
	data.JSON(http.StatusOK, responseInfo)
}
