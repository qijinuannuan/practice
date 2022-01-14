package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	//router := gin.Default()
	//router.GET("/test", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"code": 1,
	//		"message": "test",
	//	})
	//})
	//router.Run()

	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
