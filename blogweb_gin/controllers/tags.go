package controllers

import (
	"blogweb_gin/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TagsGet(ctx *gin.Context) {
	//获取session
	islogin := GetSession(ctx)

	articleDao := dao.NewArticleDao()
	tags := articleDao.QueryArticleWithParam("tags")
	fmt.Println(HandleTagsListData(tags))

	//返回html
	ctx.HTML(http.StatusOK, "tags.html", gin.H{"Tags": HandleTagsListData(tags),"IsLogin":islogin})
}

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}