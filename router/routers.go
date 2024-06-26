package router

import (
	"net/http"
	"ranking/config"
	"ranking/controllers"
	"ranking/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true // 允许所有源，生产环境中请替换为具体的源
	r.Use(cors.Default())
	r.Use(logger.Recover)
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Wolrd!")
	})

	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
	}

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PalyerController{}.GetPlayers)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("/ranking", controllers.PalyerController{}.GetRanking)
	return r

}
