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
	ReturnError(c, 4004, "没有相关信息")
}
