package router

import (
	"net/http"
	"ranking/controllers"
	"ranking/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Wolrd!")
	})

	user := r.Group("/user/")
	{
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/update", controllers.UserController{}.UpdateUser)
		user.POST("/delete", controllers.UserController{}.DeleteUser)
		user.POST("/list/test", controllers.UserController{}.GetUserListTest)

		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
