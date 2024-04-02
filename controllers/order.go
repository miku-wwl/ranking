package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetList(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "order info", 1)
}
