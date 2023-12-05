package controller

import (
	"auth/cmd/auth/model/response"
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
