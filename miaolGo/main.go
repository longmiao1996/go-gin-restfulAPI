package main

import (
	"miaolGo/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()
	router.Use(Cors())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/miaol/:id", controller.UserGet)
		v1.POST("/user", controller.UserCheck)
		v1.POST("/adduser", controller.UserAdd)
		//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.Run()

}

// 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// defer func() {
		// 	if err := recover(); err != nil {
		// 		core.Logger.Error("Panic info is: %v", err)
		// 		core.Logger.Error("Panic info is: %s", debug.Stack())
		// 	}
		// }()

		c.Next()
	}
}
