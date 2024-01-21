package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrosandrini/gopportunities/handler"

	docs "github.com/pedrosandrini/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
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

	// initialize swagger
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
