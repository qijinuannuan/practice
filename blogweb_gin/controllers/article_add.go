package controllers

import (
	"blogweb_gin/dao"
	"blogweb_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AddArticleGet 当访问/add路径的时候会触发AddArticleGet方法响应的页面是通过HTML
func AddArticleGet(ctx *gin.Context) {

	//获取session
	islogin := GetSession(ctx)

	ctx.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin})
}

func AddArticlePost(ctx *gin.Context) {

	//获取浏览器传输的数据，通过表单的name属性获取值
	//获取表单信息
	title := ctx.PostForm("title")
	tags := ctx.PostForm("tags")
	short := ctx.PostForm("short")
	author := ctx.PostForm("author")
	content := ctx.PostForm("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中

	art := &models.Article{Title: title, Author: author, Tags: tags, Short: short, Content: content, CreateTime: time.Now().Unix()}
	articleDao := dao.NewArticleDao()
	_, err := articleDao.AddArticle(art)

	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "ok"}
	} else {
		response = gin.H{"code": 0, "message": "error"}
	}

	ctx.JSON(http.StatusOK, response)

}
