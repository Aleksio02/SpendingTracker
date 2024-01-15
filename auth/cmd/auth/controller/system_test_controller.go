package controller

import (
	stresponse "auth/cmd/auth/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SystemTest(c *gin.Context) {
	systemTestResponse := stresponse.SystemTestResponse{Status: 200, Message: "Application is working successfully"}
	c.JSON(http.StatusOK, systemTestResponse)
}
