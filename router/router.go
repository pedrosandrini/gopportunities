package router

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000")
}