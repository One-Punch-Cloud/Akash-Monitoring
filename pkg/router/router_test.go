package router

import (
	"net/http"
	"net/http/httptest"
	"onePunchAkashMonitoring/pkg/node"
	"testing"
	"time"
)

func TestRouter(t *testing.T) {
	n := node.NewNode("127.0.0.1", "node@example.com", time.Now())
	router := NewRouter(n)

	// Test the /health endpoint
	req, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status OK for /health, got %v", response.Code)
	}

	// Test the /update endpoint
	req, _ = http.NewRequest("POST", "/update", nil)
	response = httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusAccepted {
		t.Errorf("Expected status Accepted for /update, got %v", response.Code)
	}
}
