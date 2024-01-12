package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spending-manager/cmd/spending-manager/connector"
	. "spending-manager/cmd/spending-manager/model/request"
	"spending-manager/cmd/spending-manager/model/response"
	"spending-manager/cmd/spending-manager/service"
	"spending-manager/cmd/spending-manager/util"
)

func AddSpentItem(c *gin.Context) {
	requestBody := SpendingRequest{}
	util.WriteBodyToObject(c.Request.Body, &requestBody)

	authResponseBody, _ := connector.GetCurrentUser(c.Request.Header)
	currentUser := response.UserInfoResponse{}
	util.WriteBodyToObject(authResponseBody.Body, &currentUser)

	if currentUser.Status == 200 {
		responseBody := service.AddSpentItem(requestBody, currentUser.Username)

		c.JSON(responseBody.Status, responseBody)
		return
	}
	c.JSON(http.StatusUnauthorized, response.SpendingResponse{Status: 401, Message: "You should authorize!"})
}

func GetSpendingsForUser(c *gin.Context) {
	authResponseBody, _ := connector.GetCurrentUser(c.Request.Header)
	currentUser := response.UserInfoResponse{}
	util.WriteBodyToObject(authResponseBody.Body, &currentUser)

	if currentUser.Status == 200 {
		responseBody := service.GetSpendingsForUser(currentUser.Username)

		c.JSON(responseBody.Status, responseBody)
		return
	}
	c.JSON(http.StatusUnauthorized, response.SpendingResponse{Status: 401, Message: "You should authorize!"})
}
