package apikey

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateApiKey verifies that GenerateApiKey returns a valid API key.
//
// The test asserts that the API key is 64 hex characters (32 bytes) and that
// it can be decoded with the hex package.
func TestGenerateApiKey(t *testing.T) {
	apiKey, _ := GenerateApiKey()

	assert.Len(t, apiKey, 64, "API key should be 64 hex characters (32 bytes)")
	_, err := hex.DecodeString(apiKey)
	assert.NoError(t, err, "API key should be valid hex")
}
