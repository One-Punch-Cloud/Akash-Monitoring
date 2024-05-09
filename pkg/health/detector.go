package health

import (
	"onePunchAkashMonitoring/pkg/node"
	"sync"
)

// Detector stores the health status of nodes.
type Detector struct {
	mu        sync.Mutex
	unhealthy map[string]bool
}

// NewDetector initializes a new Detector with empty state.
func NewDetector() *Detector {
	return &Detector{
		unhealthy: make(map[string]bool),
	}
}

// MarkUnhealthy marks a node as unhealthy.
func (d *Detector) MarkUnhealthy(n *node.Node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.unhealthy[n.Hash] = true
}

// IsUnhealthy checks if a node is marked as unhealthy.
func (d *Detector) IsUnhealthy(n *node.Node) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.unhealthy[n.Hash]
}
