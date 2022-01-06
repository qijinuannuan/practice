package dao

import (
	"blogweb_gin/models"
	"blogweb_gin/tools"
)

type UserDao struct {
	*tools.Orm
}

func NewUserDao() *UserDao {
	return &UserDao{tools.DbEngine}
}

func (ud *UserDao) InsertUser(user *models.User) (int64, error) {
	affected, err := ud.Insert(user)
	return affected, err
}

func (ud *UserDao) QueryUserWithUserName(userName string) *models.User {
	var user models.User
	if _, err := ud.Where(" user_name = ? ", userName).Get(&user); err != nil {
		return nil
	}
	return &user
}

func (ud *UserDao) QueryUserWithParam(userName, password string) *models.User {
	var user models.User
	if _, err := ud.Where(" user_name = ? and password = ? ", userName, password).Get(&user); err != nil {
		return nil
	}
	return &user
}