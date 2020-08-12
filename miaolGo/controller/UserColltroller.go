package controller

import (
	"log"
	"miaolGo/api"
	"miaolGo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserGet 用户取得
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

// UserCheck 验证用户
func UserCheck(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	pwd := ctx.Request.FormValue("pwd")
	userModel := models.User{}
	ip := ctx.ClientIP()

	var po, err, nickName, imageAddress = userModel.CheckUser(name, pwd, ip)
	if err != nil {
		log.Println("error")
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":          po,
		"nick_name":     nickName,
		"image_address": imageAddress,
	})
}

// UserAdd 用户添加
func UserAdd(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	nickName := ctx.Request.FormValue("nick_name")
	pwd := ctx.Request.FormValue("pwd")
	email := ctx.Request.FormValue("email")
	user := models.User{}

	var err, flag = user.AddUser(name, nickName, pwd, email)
	if err != nil {
		log.Println(err)
	}

	ctx.JSONP(http.StatusOK, gin.H{
		"name":      name,
		"nick_name": nickName,
		"pwd":       pwd,
		"email":     email,
		"flag":      flag,
	})
}

// UserInfoGet 用户信息取得
func UserInfoGet(ctx *gin.Context) {
	username := ctx.Param("username")

	var nickName, imageAddress = api.GetUserInfos(username)
	// if err != nil {
	// 	log.Println("error")
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"nick_name":     nickName,
		"image_address": imageAddress,
	})
}

// ArticleAdd commit article
func ArticleAdd(ctx *gin.Context) {
	userName := ctx.Request.FormValue("userName")
	nickName := ctx.Request.FormValue("nickName")
	title := ctx.Request.FormValue("title")
	article := ctx.Request.FormValue("article")
	imagebase64 := ctx.Request.FormValue("imagebase64")
	imageName := ctx.Request.FormValue("imageName")

	var flag = api.AddArticle(userName, nickName, title, article, imagebase64, imageName)
	ctx.JSONP(http.StatusOK, gin.H{
		"flag": flag,
	})
}

// UserArticleGet 取得用户文章
func UserArticleGet(ctx *gin.Context) {
	username := ctx.Query("username")

	var articleList, err = api.GetUserArticles(username)
	// if err != nil {
	// 	log.Println("error")
	// }
	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ArticleInfos": articleList,
	})

}
