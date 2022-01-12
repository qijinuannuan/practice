package controllers

import (
	"blogweb_gin/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AlbumGet(ctx *gin.Context) {
	//获取session
	islogin := GetSession(ctx)
	albumDao := dao.NewAlbum()
	albums, err := albumDao.FindAllAlbums()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
	}
	ctx.HTML(http.StatusOK, "album.html", gin.H{
		"IsLogin": islogin,
		"Album":albums,
	})
}
