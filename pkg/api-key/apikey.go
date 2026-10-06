package apikey

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateApiKey returns a random API key represented as a 64-character
// hexadecimal string.
func GenerateApiKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
