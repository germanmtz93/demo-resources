package main

import (
	"log"
	"os"

	"github.com/organization/go-api-template/config"
	"github.com/organization/go-api-template/router"
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Initialize and start the router
	r := router.SetupRouter()
	log.Printf("Demo API starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}