package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() error {
	// Check if .env file exists
	if _, err := os.Stat(".env"); err == nil {
		return godotenv.Load()
	}
	
	// If .env doesn't exist, check for environment-specific .env files
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development" // Default environment
	}
	
	envFile := ".env." + env
	if _, err := os.Stat(envFile); err == nil {
		return godotenv.Load(envFile)
	}
	
	// If no .env files found, continue with system environment variables
	return nil
}