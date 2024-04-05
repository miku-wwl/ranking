package models

import (
	"ranking/dao"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ? ", id).First(&user).Error
	return user, err
}

func GetUserListTest() ([]User, error) {
	var users []User
	err := dao.Db.Where("id < ?", 3).Find(&users).Error

	return users, err
}
func AddUser(username string) (int, error) {
	user := User{Username: username}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func UpdateUser(id int, username string) {
	dao.Db.Model(&User{}).Where("id = ?", true).Update("username", username)
}

func DeleteUser(id int) error {
	err := dao.Db.Delete(&User{}, id).Error
	return err
}
