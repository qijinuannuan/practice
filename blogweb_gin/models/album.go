package models

type Album struct {
	Id         int    `xorm:"pk aotuincr" json:"id"`
	FilePath   string `xorm:"varchar(255)" json:"file_path"`
	FileName   string `xorm:"varchar(4)" json:"file_name"`
	Status     int    `xorm:"int" json:"status"`
	CreateTime int64  `xorm:"int" json:"create_time"`
}