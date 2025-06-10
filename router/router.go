package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	_ "github.com/organization/go-api-template/docs" // Import generated Swagger docs
	"github.com/organization/go-api-template/handlers"
)

// SetupRouter configures the API routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// API documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// API endpoints
	r.GET("/", handlers.GetHostname)
	r.GET("/health", handlers.HealthCheck)
	
	return r
}