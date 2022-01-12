package controllers

import (
	"blogweb_gin/dao"
	"blogweb_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func UploadPost(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("upload")
	if err != nil {
		responseErr(ctx, err)
		return
	}
	now := time.Now()
	fileType := "other"
	//判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		responseErr(ctx, err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面，此处也可以使用io操作
	ctx.SaveUploadedFile(fileHeader,filePathStr)
	if fileType == "img" {
		album := models.Album{FilePath: filePathStr, FileName: fileName, CreateTime: timeStamp}
		albumDao := dao.NewAlbum()
		if _, err := albumDao.InsertAlbum(&album); err != nil {
			fmt.Println("---", err.Error())
			responseErr(ctx, err)
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 1, "message": "上传成功"})
}

func responseErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": err})
}