package controllers

import (
	"blogweb_gin/dao"
	"blogweb_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// HomeGet 可以通过翻页来获取该网页，也可以通过tag标签获取
//传page参数代表翻页，传tag参数代表标签
//首先判断page有值那么就是翻页，否则判断tag有值就是标签，否则就是默认的第一页
//1. 如果tag有值，page就不会有值。比如：http://127.0.0.1:8081/?pag=web
//2. 如果page有值，那么tag就不会有值。比如：http://127.0.0.1:8081/?page=3
//3. 就是用户直接访问首页，没有传入tag也没有传入page：http://127.0.0.1:8081
func HomeGet(ctx *gin.Context) {
	//获取session，判断用户是否登录
	islogin := GetSession(ctx)

	tag := ctx.Query("tag")
	page, _ := strconv.Atoi(ctx.Query("page"))
	articleDao := dao.NewArticleDao()

	var artList []*models.Article
	var err error
	var hasFooter bool

	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = articleDao.QueryArticlesWithTag(tag)
		hasFooter = false
	} else {
		if page <= 0 {
			page = 1
		}
		artList, err = articleDao.FindArticleWithPage(page)
		if err != nil {
			panic(err)
		}
		hasFooter = true
	}

	html := MakeHomeBlocks(artList, islogin)
	homeFooterPageCode := ConfigHomeFooterPageCode(page)
	fmt.Println("total :", len(artList))
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"IsLogin": islogin,
		"Content": html,
		"HasFooter":hasFooter,
		"PageCode":homeFooterPageCode,
	})
}