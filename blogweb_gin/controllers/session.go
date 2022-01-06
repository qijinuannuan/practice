package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetSession 获取session
func GetSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	loginuser := session.Get("loginuser")
	fmt.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}
