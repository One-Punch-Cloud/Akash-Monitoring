package main

import (
	"log"
	"net/http"
	"onePunchAkashMonitoring/pkg/config"
	"onePunchAkashMonitoring/pkg/node"
	"onePunchAkashMonitoring/pkg/router"
	"onePunchAkashMonitoring/pkg/store"
	"os"
	"time"
)

func main() {
	// Check if configuration path is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./onePunchAkashMonitoring <config_path>")
	}
	configPath := os.Args[1]

	// Load configuration settings from the provided file path
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration from %s: %v", configPath, err)
	}

	// Initialize a new node
	currentNode := node.NewNode(cfg.NodeIP, cfg.NodeEmail, time.Now())

	// Initialize the routing table
	routingTable := store.NewRoutingTable()

	// Register the current node in the routing table
	routingTable.RegisterNode(currentNode)

	// Set up the gossip manager
	gossipManager := store.NewGossipManager(routingTable)

	// Start the gossip protocol in a goroutine to continuously synchronize the nodes
	go gossipManager.StartGossip()

	// Initialize and start the HTTP router with only the current node
	httpRouter := router.NewRouter(currentNode)

	// Set up the HTTP server
	httpServer := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: httpRouter,
	}

	log.Printf("Starting server on %s", httpServer.Addr)

	// Start the HTTP server
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
