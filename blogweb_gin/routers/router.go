package routers

import (
	"blogweb_gin/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	//设置session midddleware
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))

	//首页
	router.GET("/", controllers.HomeGet)

	//注册：
	router.GET("/register",controllers.RegisterGet)
	router.POST("/register",controllers.RegisterPost)

	//登录
	router.GET("/login",controllers.LoginGet)
	router.POST("/login",controllers.LoginPost)

	//退出
	router.GET("/exit", controllers.ExistGet)

	//路由组
	v1 := router.Group("/article")
	{
		//写文章
		v1.GET("/add", controllers.AddArticleGet)
		v1.POST("/add", controllers.AddArticlePost)

		//更新文章
		v1.GET("/update",controllers.UpdateArticleGet)
		v1.POST("/update",controllers.UpdateArticlePost)

		// 删除文章
		v1.GET("delete", controllers.DeleteArticle)
		//相册
		router.GET("/album",controllers.AlbumGet)

		//文件上传
		router.POST("/upload",controllers.UploadPost)
	}

	//显示文章内容
	router.GET("/show/:id", controllers.ShowArticleGet)
	//标签
	router.GET("/tags",controllers.TagsGet)
	//关于我
	router.GET("/aboutme",controllers.AboutMeGet)

	return router
}
