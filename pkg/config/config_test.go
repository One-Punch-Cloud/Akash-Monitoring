package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Setup test environment variables
	os.Setenv("PING_INTERVAL_INTERNAL", "15")
	os.Setenv("PING_INTERVAL_EXTERNAL", "45")
	os.Setenv("UNHEALTHY_THRESHOLD", "3")

	config, err := LoadConfig(".")
	if err != nil {
		t.Fatalf("Failed to load configs: %s", err)
	}

	if config.PingIntervalInternal != 15 {
		t.Errorf("Expected PingIntervalInternal to be '15', got '%d'", config.PingIntervalInternal)
	}

	if config.PingIntervalExternal != 45 {
		t.Errorf("Expected PingIntervalExternal to be '45', got '%d'", config.PingIntervalExternal)
	}

	if config.UnhealthyThreshold != 3 {
		t.Errorf("Expected UnhealthyThreshold to be '3', got '%d'", config.UnhealthyThreshold)
	}
}
