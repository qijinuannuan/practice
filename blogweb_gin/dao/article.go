package dao

import (
	"blogweb_gin/config"
	"blogweb_gin/models"
	"blogweb_gin/tools"
	"fmt"
)

type ArticleDao struct {
	*tools.Orm
}

func NewArticleDao() *ArticleDao {
	return &ArticleDao{tools.DbEngine}
}

//---------添加文章-----------

func (ad *ArticleDao) AddArticle(article *models.Article) (int64, error) {
	affected, err := ad.Insert(article)
	return affected, err
}

//-----------查询文章---------

// FindArticleWithPage 根据页码查询文章
func (ad *ArticleDao)FindArticleWithPage(page int) ([]*models.Article, error) {
	//从配置文件中获取每页的文章数量
	var articles []*models.Article
	if err := ad.Where(" id > 0 and is_removed = 0 ").Limit(config.NUM, (page-1)*config.NUM).Find(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// QueryArticleRowNum 查询文章的总条数
func (ad *ArticleDao)QueryArticleRowNum() int {
	articles := new(models.Article)
	count, err := ad.Where(" id > ? and is_removed = 0 ", 0).Count(articles)
	if err != nil {
		return 0
	}
	return int(count)
}

//----------查询文章-------------

func (ad *ArticleDao)QueryArticleWithId(id int) *models.Article {
	articles := new(models.Article)
	if _, err := ad.Where(" id = ? and is_removed = 0 ", id).Get(articles); err != nil {
		return nil
	}
	return articles
}

//----------修改数据----------

func (ad *ArticleDao)UpdateArticle(article *models.Article) (int64, error) {
	//数据库操作
	return ad.Id(article.Id).Update(article)
}

func (ad *ArticleDao)DeleteArticle(id int) error {
	article := new(models.Article)
	article.IsRemoved = 1
	if _, err := ad.Id(id).Update(article); err != nil {
		return err
	}
	return nil
}

func (ad *ArticleDao)QueryArticleWithParam(param string) []string {
	article := new(models.Article)
	rows, err := ad.Select(param).Where(" is_removed = 0 ").Rows(article)
	if err != nil {
	}
	var paramList []string
	defer rows.Close()
	for rows.Next() {
		rows.Scan(article)
		paramList = append(paramList, article.Tags)
	}
	return paramList
}

// QueryArticlesWithTag --------------按照标签查询--------------
func (ad *ArticleDao)QueryArticlesWithTag(tag string) ([]*models.Article, error) {
	var articles []*models.Article
	sql := " tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println(sql)
	err := ad.Where(sql).Find(&articles)
	return articles, err
}