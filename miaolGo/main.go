package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"miaolGo/controller"
)



func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/miaol/:id", controller.UserGet)
		//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.Run()

}





