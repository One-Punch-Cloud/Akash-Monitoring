package health

import (
	"onePunchAkashMonitoring/pkg/node"
	"testing"
	"time"
)

func TestDetector(t *testing.T) {
	detector := NewDetector()
	n := node.NewNode("192.168.1.1", "node@example.com", time.Now())

	// Initially, the node should not be marked as unhealthy
	if detector.IsUnhealthy(n) {
		t.Errorf("New node should not be unhealthy")
	}

	// Mark the node as unhealthy and check again
	detector.MarkUnhealthy(n)
	if !detector.IsUnhealthy(n) {
		t.Errorf("Node should be marked as unhealthy")
	}
}
