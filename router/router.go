package router

import (
	"github.com/gin-gonic/gin"
	"github.com/organization/go-api-template/handlers"
)

// SetupRouter configures the API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// API endpoints
	r.GET("/", handlers.Welcome)
	r.GET("/health", handlers.HealthCheck)
	r.GET("/docs", handlers.GetDocs)
	
	return r
}