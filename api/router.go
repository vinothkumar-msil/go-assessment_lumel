package api

import (
	"github.com/gin-gonic/gin"
	"go-backend-assessment/api/handlers"
)

// SetupRoutes registers all API routes
func SetupRoutes(router *gin.Engine) {
	// Register the refresh route
	router.GET("/refresh", handlers.RefreshDataHandler)

	// Register the revenue route
	router.GET("/revenue", handlers.RevenueHandler)
}
