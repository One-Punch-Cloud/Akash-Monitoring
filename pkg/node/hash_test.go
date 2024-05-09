package node

import (
	"testing"
	"time"
)

func TestGenerateNodeHash(t *testing.T) {
	ip := "192.168.1.1"
	launchTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	expected := "ff021f15f236b394f1797a0de89778ac0246448226c95c6ec361f1087b33431b" // Expected hash output for this specific input

	hash := GenerateNodeHash(ip, launchTime)
	if hash != expected {
		t.Errorf("Hash does not match expected value. Expected %s, got %s", expected, hash)
	}
}
