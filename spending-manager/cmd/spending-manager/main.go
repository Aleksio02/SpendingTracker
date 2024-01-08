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
	}
	r.Run(fmt.Sprintf(":%v", config.Config.Application.Port))
}
