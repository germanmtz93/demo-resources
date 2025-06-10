package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}

// HostnameResponse represents the hostname response
type HostnameResponse struct {
	Hostname string `json:"hostname"`
}

// GetHostname returns the hostname of the machine
// @Summary Get hostname
// @Description Returns the hostname of the machine the API is running on
// @Produce json
// @Success 200 {object} HostnameResponse
// @Router / [get]
func GetHostname(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	
	c.JSON(http.StatusOK, HostnameResponse{
		Hostname: hostname,
	})
}

// HealthCheck returns the API health status
// @Summary Health check
// @Description Returns the health status of the API
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status: "ok",
	})
}