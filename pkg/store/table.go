package store

import (
	"onePunchAkashMonitoring/pkg/node"
	"sync"
)

// RoutingTable holds the node data in a map and uses a Mutex for safe concurrent access.
type RoutingTable struct {
	sync.Mutex
	nodes map[string]*node.Node
}

// NewRoutingTable initializes a new RoutingTable.
func NewRoutingTable() *RoutingTable {
	return &RoutingTable{
		nodes: make(map[string]*node.Node),
	}
}

// RegisterNode adds a new node to the table.
func (table *RoutingTable) RegisterNode(n *node.Node) {
	table.Lock()
	defer table.Unlock()
	table.nodes[n.Hash] = n
}

// GetNode retrieves a node by hash from the table.
func (table *RoutingTable) GetNode(hash string) (*node.Node, bool) {
	table.Lock()
	defer table.Unlock()
	n, exists := table.nodes[hash]
	return n, exists
}

// DeleteNode removes a node from the table by hash.
func (table *RoutingTable) DeleteNode(hash string) {
	table.Lock()
	defer table.Unlock()
	delete(table.nodes, hash)
}

// GetAllNodes returns all nodes in the table.
func (table *RoutingTable) GetAllNodes() map[string]*node.Node {
	table.Lock()
	defer table.Unlock()
	copy := make(map[string]*node.Node)
	for k, v := range table.nodes {
		copy[k] = v
	}
	return copy
}
