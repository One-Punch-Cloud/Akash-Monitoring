package node

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateNodeHash computes a SHA-256 hash for a node using its IP address and launch time.
func GenerateNodeHash(ip string, launchTime time.Time) string {
	data := fmt.Sprintf("%s-%v", ip, launchTime.Unix())
	hashBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashBytes[:])
}
