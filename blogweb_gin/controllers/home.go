package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeGet(ctx *gin.Context) {
	//获取session，判断用户是否登录
	islogin := GetSession(ctx)
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"IsLogin": islogin,
	})
}