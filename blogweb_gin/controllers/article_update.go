package controllers

import (
	"blogweb_gin/dao"
	"blogweb_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateArticleGet(ctx *gin.Context) {

	//获取session
	islogin := GetSession(ctx)

	idstr := ctx.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("id : ", id)

	//获取id所对应的文章信息
	articleDao := dao.NewArticleDao()
	art := articleDao.QueryArticleWithId(id)
	ctx.HTML(http.StatusOK, "write_article.html", gin.H{
		"IsLogin": islogin,
		"Title": art.Title,
		"Author": art.Author,
		"Tags": art.Tags,
		"Short": art.Short,
		"Content": art.Content,
		"Id": art.Id,
	})
}

// UpdateArticlePost 修改文章
func UpdateArticlePost(ctx *gin.Context) {

	idstr := ctx.PostForm("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("postid:", id)

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := ctx.PostForm("title")
	tags := ctx.PostForm("tags")
	short := ctx.PostForm("short")
	author := ctx.PostForm("author")
	content := ctx.PostForm("content")

	//实例化model，修改数据库
	art := models.Article{Id: id, Title: title, Author: author, Tags: tags, Short: short, Content: content}
	articleDao := dao.NewArticleDao()
	_, err := articleDao.UpdateArticle(&art)

	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "更新成功"}
	} else {
		fmt.Println(err)
		response = gin.H{"code": 0, "message": "更新失败"}
	}

	ctx.JSON(http.StatusOK, response)
}