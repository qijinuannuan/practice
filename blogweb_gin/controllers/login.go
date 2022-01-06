package controllers

import (
	"blogweb_gin/dao"
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet(ctx *gin.Context) {
	//返回html
	ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(ctx *gin.Context) {
	//获取表单信息
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println("username:", username, ",password:", password)
	userDao := dao.NewUserDao()
	var user *models.User
	user = userDao.QueryUserWithParam(username, utils.MD5(password))
	if user.UserName == username {
		/*
		   设置了session后会将数据处理设置到cookie，然后再浏览器进行网络请求的时候会自动带上cookie
		   因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		*/
		session := sessions.Default(ctx)
		session.Set("loginuser", username)
		session.Save()

		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "登录成功",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"message": "登录失败",
		})
	}
}
