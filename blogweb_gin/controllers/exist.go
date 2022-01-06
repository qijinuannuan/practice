package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExistGet(ctx *gin.Context) {
	//清除该用户登录状态的数据
	session := sessions.Default(ctx)
	session.Delete("loginuser")
	session.Save()
	//session.Clear()

	//清除session后，重定位到"/"路径上(首页)。
	fmt.Println("delete session...",session.Get("loginuser"))
	ctx.Redirect(http.StatusMovedPermanently, "/")
}
