package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"miaolGo/models"
	"net/http"
	"strconv"
)

func UserGet(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	bookModel := models.Book{}

	var data, err = bookModel.GetUser(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func UserCheck(ctx *gin.Context) {
	name := ctx.Param("name")
	pwd := ctx.Param("pwd")
	userModel := models.User{}

	var po, err = userModel.CheckUser(name, pwd)
	if err != nil {
		log.Println("error")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": po,
	})
}

func UserAdd(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	nick_name := ctx.Request.FormValue("nick_name")
	pwd := ctx.Request.FormValue("pwd")
	email := ctx.Request.FormValue("email")
	user := models.User{}

	err := user.AddUser(name, nick_name, pwd, email)
	if err != nil {
		log.Println(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"name":      name,
		"nick_name": nick_name,
		"pwd":       pwd,
		"email":     email,
	})
}
