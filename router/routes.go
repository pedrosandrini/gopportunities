package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrosandrini/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Show opening job
		v1.GET("/opening", handler.ShowOpeningHandler)

		// Create
		v1.POST("/opening", handler.CreateOpeningHandler)

		// Delete
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		// Update
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		// Get all opening jobs
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
