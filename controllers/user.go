package controllers

import (
	"ranking/models"
	"ranking/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	name := c.Param("name")
	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 0, name, user, 1)
}

func (u UserController) AddUser(c *gin.Context) {
	name := c.DefaultPostForm("username", "")
	logger.Error(map[string]interface{}{"name=": name})

	id, err := models.AddUser(name)
	if err != nil {
		ReturnError(c, 4002, "保存错误")
	}
	ReturnSuccess(c, 0, "保存成功", id, 1)
}

func (u UserController) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, username)
	ReturnSuccess(c, 0, "更新成功", true, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)

	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4002, "删除错误")
	}
	ReturnSuccess(c, 0, "删除成功", id, 1)
}

func (u UserController) GetList(c *gin.Context) {

	// logger.Write("日志信息", "user")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("捕获异常")
	// 	}
	// }()

	i := 1
	j := 0
	ReturnError(c, 4004+i/j, "没有相关信息")
}

func (u UserController) GetUserListTest(c *gin.Context) {
	users, err := models.GetUserListTest()
	if err != nil {
		ReturnError(c, 4004, "没有相关数据")
	}
	ReturnSuccess(c, 0, "获取成功", users, 1)
}
