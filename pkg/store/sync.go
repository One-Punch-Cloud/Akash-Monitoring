package store

import (
	"log"
	"math/rand"
	"onePunchAkashMonitoring/pkg/node"
	"time"
)

// GossipManager handles the gossip-based synchronization of nodes.
type GossipManager struct {
	routingTable *RoutingTable
}

// NewGossipManager creates a new instance of GossipManager.
func NewGossipManager(routingTable *RoutingTable) *GossipManager {
	return &GossipManager{
		routingTable: routingTable,
	}
}

// StartGossip initiates the gossip protocol to synchronize node states across the network.
func (gm *GossipManager) StartGossip() {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		gm.gossip()
	}
}

// gossip performs one round of gossip, randomly selecting a node and sharing state information.
func (gm *GossipManager) gossip() {
	nodes := gm.routingTable.GetAllNodes()
	if len(nodes) == 0 {
		return // No nodes to gossip about
	}

	// Select a random node to "gossip" with
	keys := make([]string, 0, len(nodes))
	for k := range nodes {
		keys = append(keys, k)
	}
	randomNodeKey := keys[rand.Intn(len(keys))]
	randomNode := nodes[randomNodeKey]

	// Simulate the gossip: print the chosen node's info
	log.Printf("Gossiping about node %s at %s\n", randomNodeKey, randomNode.IPAddress)
}

// SimulateReceiveGossip simulates receiving gossip from another node.
func (gm *GossipManager) SimulateReceiveGossip(receivedNode *node.Node) {
	// Add or update the node information in the routing table
	gm.routingTable.RegisterNode(receivedNode)
}
