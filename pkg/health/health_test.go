package health

import (
	"net/http"
	"net/http/httptest"
	"onePunchAkashMonitoring/pkg/node"
	"testing"
	"time"
)

func TestCheckNodeHealth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	testNode := &node.Node{
		IPAddress:  server.URL[7:], // Strip "http://" from server.URL
		LaunchTime: time.Now(),
	}

	healthy := CheckNodeHealth(testNode)
	if !healthy {
		t.Errorf("Expected node to be healthy")
	}
}
