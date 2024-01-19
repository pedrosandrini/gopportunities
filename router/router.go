package router

import (
	gin "github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router
	r := gin.Default()

	// Initialize routes
	initializeRoutes(r)

	// Run the server
	r.Run(":8000")
}
