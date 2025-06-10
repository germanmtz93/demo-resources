package main

import (
	"log"
	"os"

	"github.com/organization/go-api-template/config"
	"github.com/organization/go-api-template/router"
)

// @title Go API Template
// @version 1.0
// @description A standard template for Go REST APIs
// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize and start the router
	r := router.SetupRouter()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}