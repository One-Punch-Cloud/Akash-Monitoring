package node

import (
	"time"
)

// Node represents a single node in the monitoring system.
type Node struct {
	IPAddress  string    // IP address of the node
	LaunchTime time.Time // Launch time of the node
	Hash       string    // Unique identifier based on IP and launch time
	Email      string    // Email for sending alerts
}

// NewNode creates a new node and calculates its hash using the GenerateNodeHash function.
func NewNode(ipAddress, email string, launchTime time.Time) *Node {
	node := &Node{
		IPAddress:  ipAddress,
		LaunchTime: launchTime,
		Email:      email,
	}
	node.Hash = GenerateNodeHash(ipAddress, launchTime)
	return node
}
