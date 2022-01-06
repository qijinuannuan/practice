package controllers

import (
	"blogweb_gin/dao"
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*type UserRegister struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}*/

func RegisterGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

func RegisterPost(ctx *gin.Context) {
	//获取表单信息
	username := ctx.PostForm("username")

	password := ctx.PostForm("password")
	repassword := ctx.PostForm("repassword")
	if password != repassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":0,
			"message":"密码需要保持一致",
		})
	}
	fmt.Printf("user info: user_name: %s, password: %s\n", username, password)
	userDao := dao.NewUserDao()
	var user *models.User
	user = userDao.QueryUserWithUserName(username)
	if user.UserName == username {
		fmt.Printf("user name: %s is exist, register failed !\n", user.UserName)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"message": "用户名已经存在",
		})
		return
	}
	password = utils.MD5(password)
	user.UserName = username
	user.Password = password
	user.CreateTime = time.Now().Unix()
	if _, err := userDao.InsertUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":0,
			"message":"注册失败",
		})
	}else{
		ctx.JSON(http.StatusOK, gin.H{
			"code":1,
			"message":"注册成功",
		})
	}
}
