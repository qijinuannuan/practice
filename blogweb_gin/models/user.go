package models

type User struct {
	Id         int    `xorm:"pk autoincr" json:"id"`
	UserName   string `xorm:"varchar(64)" json:"user_name"`
	Password   string `xorm:"varchar(64)" json:"password"`
	Status     int    `xorm:"int" json:"status"`
	CreateTime int64  `xorm:"int" json:"create_time"`
}
