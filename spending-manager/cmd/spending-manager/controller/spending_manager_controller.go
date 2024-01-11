package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "spending-manager/cmd/spending-manager/model/request"
	"spending-manager/cmd/spending-manager/service"
	"spending-manager/cmd/spending-manager/util"
)

func AddSpentItem(c *gin.Context) {
	requestBody := SpendingRequest{}
	util.WriteBodyToObject(c.Request.Body, &requestBody)

	// TODO: Get username from token (from auth module)
	//

	responseBody := service.AddSpentItem(requestBody, "Oleksio")

	c.JSON(http.StatusOK, responseBody)
}
