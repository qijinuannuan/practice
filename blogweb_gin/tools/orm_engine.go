package tools

import (
	"blogweb_gin/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(cfg *Config) (*Orm, error) {
	fmt.Println("Init Mysql ... ... ")
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		fmt.Println("dao connect failed, err : ", err)
		return nil, err
	}
	engine.ShowSQL(database.ShowSql)
	// 添加表到数据库
	err = engine.Sync2(
		new(models.User),
		new(models.Article),
		new(models.Album))
	if err != nil {
		fmt.Println("add table to dao failed, err : ", err.Error())
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return DbEngine, nil
}