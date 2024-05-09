package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Config struct {
	NodeIP               string `json:"NodeIP"`
	NodeEmail            string `json:"NodeEmail"`
	ServerPort           string `json:"ServerPort"`
	InternalPingInterval string `json:"InternalPingInterval"`
	ExternalPingInterval string `json:"ExternalPingInterval"`
	UnhealthyThreshold   int    `json:"UnhealthyThreshold"`
}

func LoadConfig(path string) (*Config, error) {
	if path != "" {
		return loadConfigFromFile(path)
	}
	return loadConfigFromEnv()
}

func loadConfigFromFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	if err = json.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadConfigFromEnv() (*Config, error) {
	cfg := &Config{
		NodeIP:               os.Getenv("NODE_IP"),
		NodeEmail:            os.Getenv("NODE_EMAIL"),
		ServerPort:           os.Getenv("SERVER_PORT"),
		InternalPingInterval: os.Getenv("INTERNAL_PING_INTERVAL"),
		ExternalPingInterval: os.Getenv("EXTERNAL_PING_INTERVAL"),
		UnhealthyThreshold:   parseEnvInt("UNHEALTHY_THRESHOLD"),
	}
	return cfg, nil
}

func parseEnvInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Printf("Error parsing integer from environment for %s: %v", key, err)
		return 0 // Default or handle error differently
	}
	return value
}
