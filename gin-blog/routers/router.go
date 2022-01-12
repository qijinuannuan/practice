package routers

import (
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	engine.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 1,
			"message": "test",
		})
	})

	return engine
}