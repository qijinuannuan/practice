package main
import (
	"ginws/ws"
	"github.com/gin-gonic/gin"
)
//server
func main() {
	gin.SetMode(gin.ReleaseMode) //线上环境
	go ws.Manager.Start()
	r := gin.Default()
	r.GET("/ws",ws.WsHandler)
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}