package controllers

import (
	"ranking/models"
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
