package store

import (
	"onePunchAkashMonitoring/pkg/node"
	"testing"
	"time"
)

func TestGossipManager(t *testing.T) {
	routingTable := NewRoutingTable()
	gm := NewGossipManager(routingTable)

	n1 := node.NewNode("192.168.1.1", "node1@example.com", time.Now())
	n2 := node.NewNode("192.168.1.2", "node2@example.com", time.Now().Add(1*time.Minute))

	// Register initial node
	routingTable.RegisterNode(n1)

	// Simulate receiving gossip about a new node
	gm.SimulateReceiveGossip(n2)
	if _, exists := routingTable.GetNode(n2.Hash); !exists {
		t.Errorf("Gossip did not update the node correctly")
	}
}
