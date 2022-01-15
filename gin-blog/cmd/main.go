package main

import (
	"fmt"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"github.com/fvbock/endless"
	"syscall"
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

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		logging.Info("Server err: %v", err)
	}

	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr: fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler: router,
	//	ReadTimeout: setting.ReadTimeout,
	//	WriteTimeout: setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}

	////s.ListenAndServe()
	//go func() {
	//	if err := s.ListenAndServe(); err != nil {
	//		log.Printf("Listen: %s\n", err)
	//	}
	//}()
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<- quit
	//log.Println("Shutdown Server ...")
	//ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//defer cancel()
	//if err := s.Shutdown(ctx); err != nil {
	//	logging.Fatal("Server Shutdown:", err)
	//}
	//logging.Info("Server exiting")
}
