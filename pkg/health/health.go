package health

import (
	"net/http"
	"onePunchAkashMonitoring/pkg/node"
	"time"
)

// CheckNodeHealth sends a HTTP GET request to the node's health endpoint and returns if it is healthy or not.
func CheckNodeHealth(n *node.Node) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("http://" + n.IPAddress + "/health")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
