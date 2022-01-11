package dao

import (
	"blogweb_gin/config"
	"blogweb_gin/models"
	"blogweb_gin/tools"
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
	if err := ad.Where(" id > 0 ").Limit(config.NUM, (page-1)*config.NUM).Find(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// QueryArticleRowNum 查询文章的总条数
func (ad *ArticleDao)QueryArticleRowNum() int {
	articles := new(models.Article)
	count, err := ad.Where(" id > ? ", 0).Count(articles)
	if err != nil {
		panic(err)
		return 0
	}
	return int(count)
}

//----------查询文章-------------

func (ad *ArticleDao)QueryArticleWithId(id int) *models.Article {
	articles := new(models.Article)
	if _, err := ad.Where(" id = ? ", id).Get(articles); err != nil {
		return nil
	}
	return articles
}