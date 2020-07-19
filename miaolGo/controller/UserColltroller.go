package controller

import (
	"log"
	"miaolGo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	name := ctx.Request.FormValue("name")
	pwd := ctx.Request.FormValue("pwd")
	userModel := models.User{}
	ip := ctx.ClientIP()

	var po, err, nick_name, image_address = userModel.CheckUser(name, pwd, ip)
	if err != nil {
		log.Println("error")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":          po,
		"nick_name":     nick_name,
		"image_address": image_address,
	})
}

func UserAdd(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	nick_name := ctx.Request.FormValue("nick_name")
	pwd := ctx.Request.FormValue("pwd")
	email := ctx.Request.FormValue("email")
	user := models.User{}

	var err, flag = user.AddUser(name, nick_name, pwd, email)
	if err != nil {
		log.Println(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"name":      name,
		"nick_name": nick_name,
		"pwd":       pwd,
		"email":     email,
		"flag":      flag,
	})
}
