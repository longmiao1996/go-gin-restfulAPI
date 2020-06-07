package controller

import (
	"github.com/gin-gonic/gin"
	"miaolGo/models"
	"net/http"
	"strconv"
)

func UserGet(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userModel := models.Book{}

	var data, err = userModel.GetUser(id)

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
