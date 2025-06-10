package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	
	"github.com/organization/go-api-template/router"
)

func TestRouterSetup(t *testing.T) {
	// Initialize router
	r := router.SetupRouter()
	
	// Create a test server
	ts := httptest.NewServer(r)
	defer ts.Close()
	
	// Test health endpoint
	resp, err := http.Get(ts.URL + "/health")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	
	// Test root endpoint
	resp, err = http.Get(ts.URL + "/")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	
	// Test Swagger endpoint
	resp, err = http.Get(ts.URL + "/swagger/index.html")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}