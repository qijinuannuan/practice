package controllers

import (
	"blogweb_gin/config"
	"blogweb_gin/dao"
	"blogweb_gin/models"
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

// MakeHomeBlocks ----------首页显示内容---------
func MakeHomeBlocks(articles []*models.Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		//将数据库model转换为首页模板所需要的model
		homeParam := models.HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		//fmt.Println("tag-->", art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		//homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/show/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为穿进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	//fmt.Println("htmlHome-->",htmlHome)
	return template.HTML(htmlHome)
}

//将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tags string) []models.TagLink {
	var tagLink []models.TagLink
	tagsPamar := strings.Split(tags, "&")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, models.TagLink{TagName: tag, TagUrl: "/?tag=" + tag})
	}
	return tagLink
}

//-----------翻页-----------

// ConfigHomeFooterPageCode page是当前的页数
func ConfigHomeFooterPageCode(page int) models.HomeFooterPageCode {
	pageCode := models.HomeFooterPageCode{}
	//查询出总的条数
	num := GetArticleRowsNum()
	//计算出总页数
	allPageNum := (num-1)/config.NUM + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode

}

//------翻页------

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

// GetArticleRowsNum 只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	articleDao := dao.NewArticleDao()
	if artcileRowsNum == 0 {
		artcileRowsNum = articleDao.QueryArticleRowNum()
	}
	return artcileRowsNum
}

// SetArticleRowsNum 设置页数
func SetArticleRowsNum(){
	articleDao := dao.NewArticleDao()
	artcileRowsNum = articleDao.QueryArticleRowNum()
}