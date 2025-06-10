package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	
	"github.com/organization/go-api-template/handlers"
)

func TestHealthCheck(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	
	// Create a response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	// Call the handler
	handlers.HealthCheck(c)
	
	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)
	
	// Parse response
	var response handlers.HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	// Assert no error in parsing
	assert.NoError(t, err)
	
	// Assert response content
	assert.Equal(t, "ok", response.Status)
}

func TestGetHostname(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	
	// Create a response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	
	// Call the handler
	handlers.GetHostname(c)
	
	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)
	
	// Parse response
	var response handlers.HostnameResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	// Assert no error in parsing
	assert.NoError(t, err)
	
	// Get expected hostname
	expectedHostname, err := os.Hostname()
	if err != nil {
		expectedHostname = "unknown"
	}
	
	// Assert response content
	assert.Equal(t, expectedHostname, response.Hostname)
}