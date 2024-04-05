package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, "success 666"+id+name, "user info", 1)
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
