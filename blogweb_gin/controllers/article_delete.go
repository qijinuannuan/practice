package controllers

import (
	"blogweb_gin/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteArticle(ctx *gin.Context) {
	idstr := ctx.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("删除 id : ", id)

	articleDao := dao.NewArticleDao()
	err := articleDao.DeleteArticle(id)
	if err != nil {
		panic(err.Error())
	}
	ctx.Redirect(302, "/")
}
