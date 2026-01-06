package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome returns demo API welcome message
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Demo API",
		"version": "1.0.0",
	})
}

// HealthCheck returns API health status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"service": "demo-api",
	})
}

// GetDocs returns dummy documentation data
func GetDocs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"api": "Demo API",
		"endpoints": []string{"/", "/health", "/docs"},
		"description": "Simple demo API for ArgoCD deployment",
	})
}