package health

import (
	"fmt"
	"onePunchAkashMonitoring/pkg/node"
	"sync"
)

// Detector stores the health status of nodes.
type Detector struct {
	mu            sync.Mutex
	unhealthy     map[string]bool
	emailNotifier *EmailNotifier
	threshold     int
	failureCounts map[string]int
}

// NewDetector initializes a new Detector with empty state and an email notifier.
func NewDetector(emailNotifier *EmailNotifier, threshold int) *Detector {
	return &Detector{
		unhealthy:     make(map[string]bool),
		emailNotifier: emailNotifier,
		threshold:     threshold,
		failureCounts: make(map[string]int),
	}
}

// MarkUnhealthy marks a node as unhealthy and sends a notification if necessary.
func (d *Detector) MarkUnhealthy(n *node.Node) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.failureCounts[n.Hash]++
	if d.failureCounts[n.Hash] >= d.threshold {
		d.unhealthy[n.Hash] = true
		d.notifyOwner(n)
		delete(d.failureCounts, n.Hash) // reset the counter after notifying
	}
}

// IsUnhealthy checks if a node is marked as unhealthy.
func (d *Detector) IsUnhealthy(n *node.Node) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.unhealthy[n.Hash]
}

// notifyOwner sends an email notification to the owner of the unhealthy node.
func (d *Detector) notifyOwner(n *node.Node) {
	subject := "Node Unhealthy: " + n.IPAddress
	body := "Node with IP " + n.IPAddress + " has been detected as unhealthy. Please take appropriate action."
	err := d.emailNotifier.SendNotification(n.Email, subject, body)
	if err != nil {
		fmt.Printf("Failed to send notification for node %s: %v\n", n.IPAddress, err)
	} else {
		fmt.Printf("Notification sent for node %s\n", n.IPAddress)
	}
}
