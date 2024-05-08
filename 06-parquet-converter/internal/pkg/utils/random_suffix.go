package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func NewRandomSuffix() (string, error) {
	// Generate 8 random bytes
	randomBytes := make([]byte, 2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("error while generating random suffix name: %w", err)
	}
	// Convert bytes to hexadecimal string
	randomHex := hex.EncodeToString(randomBytes)
	return randomHex, nil
}
