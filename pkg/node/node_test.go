package node

import (
	"testing"
	"time"
)

func TestNewNode(t *testing.T) {
	ip := "192.168.1.1"
	email := "example@example.com"
	launchTime := time.Now()

	node := NewNode(ip, email, launchTime)

	if node.IPAddress != ip {
		t.Errorf("Expected IP address %s, got %s", ip, node.IPAddress)
	}

	if node.Email != email {
		t.Errorf("Expected email %s, got %s", email, node.Email)
	}

	expectedHash := GenerateNodeHash(ip, launchTime)
	if node.Hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, node.Hash)
	}
}
