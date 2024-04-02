package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/url", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Wolrd!")
	})
	r.Run(":9999")
}
