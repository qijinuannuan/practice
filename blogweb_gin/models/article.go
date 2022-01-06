package models

type Article struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	Title      string `xorm:"varchar(30)" json:"title"`
	Author     string `xorm:"varchar(20)" json:"author"`
	Tags       string `xorm:"varchar(30)" json:"tags"`
	Short      string `xorm:"varchar(255)" json:"short"`
	Content    string `xorm:"longtext" json:"content"`
	CreateTime int64  `xorm:"int" json:"create_time"`
}
