package models

import (
	"ranking/dao"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (User) TableName() string {
	return "user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ? ", id).First(&user).Error
	return user, err
}
