package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

//go:generate mockgen -source=hash.go -destination=mocks/hash.go
type HashPassword interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashPassword string, password string) error
}

type Hashing struct{}

// NewHashingPassword initializes a new Hashing instance for hashing and comparing passwords.
//
// Returns:
//   - HashPassword: The initialized Hashing instance.
func NewHashingPassword() HashPassword {
	return &Hashing{}
}

// HashPassword takes a plaintext password and returns a hashed version of the password.
// The hashed password is a string that can be safely stored in a database or other
// secure storage.
//
// Parameters:
//   - password: The plaintext password to be hashed (string)
//
// Returns:
//   - string: The hashed version of the password (string)
//   - error: Any error encountered while hashing the password (error)
func (h Hashing) HashPassword(password string) (string, error) {
	pw := []byte(password)
	hashedPw, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPw), nil
}

// ComparePassword takes a hashed password and a plaintext password and
// returns an error if the passwords do not match.
//
// Parameters:
//   - hashPassword: The hashed password to compare against (string)
//   - password: The plaintext password to compare to the hashed password (string)
//
// Returns:
//   - error: nil if the passwords match, otherwise an error (error)
func (h Hashing) ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
