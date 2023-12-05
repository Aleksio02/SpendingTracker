package main

import (
	"auth/cmd/auth/config"
	"auth/cmd/auth/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	r := gin.New()
	v1 := r.Group("/" + config.Config.Application.Name)
	{
		v1.GET("/system/test", controller.SystemTest)
		v1.GET("/getUserByToken", controller.GetUserByToken)
	}
	r.Run(fmt.Sprintf(":%v", config.Config.Application.Port))
}
