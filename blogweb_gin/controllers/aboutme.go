package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func  AboutMeGet(ctx *gin.Context) {

	//获取session
	islogin := GetSession(ctx)

	ctx.HTML(http.StatusOK, "aboutme.html", gin.H{
		"IsLogin": islogin,
		"wechat":"微信：tia_520",
		"qq":"QQ：79539705",
		"tel":"Tel：13910439137",
	})
}