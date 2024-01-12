package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"spending-manager/cmd/spending-manager/config"
	"spending-manager/cmd/spending-manager/controller"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	r := gin.New()

	v1 := r.Group("/" + config.Config.Application.Name)
	{
		v1.GET("/system/test", controller.SystemTest)
		v1.POST("/addSpentItem", controller.AddSpentItem)
		v1.GET("/getSpendings/all", controller.GetSpendingsForUser)
	}
	r.Run(fmt.Sprintf(":%v", config.Config.Application.Port))
}
