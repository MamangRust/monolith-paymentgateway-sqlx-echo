package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// TestHashPassword_Success tests that hashing a password returns a non-empty hash
func TestHashPassword_Success(t *testing.T) {
	hasher := NewHashingPassword()

	password := "securePassword123"
	hashed, err := hasher.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
	assert.Contains(t, hashed, "$2")
}

// TestComparePassword_Success tests that comparing a correct password returns no error
func TestComparePassword_Success(t *testing.T) {
	hasher := NewHashingPassword()

	password := "myStrongPassword"
	hashed, err := hasher.HashPassword(password)
	assert.NoError(t, err)

	err = hasher.ComparePassword(hashed, password)
	assert.NoError(t, err)
}

// TestComparePassword_Failure tests that comparing an incorrect password returns an error
func TestComparePassword_Failure(t *testing.T) {
	hasher := NewHashingPassword()

	password := "correctPassword"
	wrongPassword := "wrongPassword"

	hashed, err := hasher.HashPassword(password)
	assert.NoError(t, err)

	err = hasher.ComparePassword(hashed, wrongPassword)
	assert.Error(t, err)
	assert.Equal(t, bcrypt.ErrMismatchedHashAndPassword, err)
}
