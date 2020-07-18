package main

import (
	"miaolGo/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/miaol/:id", controller.UserGet)
		v1.GET("/user/:name/:pwd", controller.UserCheck)
		v1.POST("/adduser", controller.UserAdd)
		//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.Run()

}
