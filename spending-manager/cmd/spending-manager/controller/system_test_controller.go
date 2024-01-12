package controller

import (
	stresponse "spending-manager/cmd/spending-manager/model/response"

	"github.com/gin-gonic/gin"
	"net/http"
)

func SystemTest(c *gin.Context) {
	systemTestResponse := stresponse.SystemTestResponse{Status: 200, Message: "Application is working successfully"}
	c.JSON(http.StatusOK, systemTestResponse)
}
