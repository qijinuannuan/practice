package controllers

import (
	"blogweb_gin/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HomeGet(ctx *gin.Context) {
	//获取session，判断用户是否登录
	islogin := GetSession(ctx)

	page, _ := strconv.Atoi(ctx.Query("page"))
	if page <= 0 {
		page = 1
	}
	articleDao := dao.NewArticleDao()
	artList, err := articleDao.FindArticleWithPage(page)
	if err != nil {
		panic(err)
	}
	html := MakeHomeBlocks(artList, islogin)
	homeFooterPageCode := ConfigHomeFooterPageCode(page)
	fmt.Println("total :", len(artList))
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"IsLogin": islogin,
		"Content": html,
		"HasFooter":true,
		"PageCode":homeFooterPageCode,
	})
}