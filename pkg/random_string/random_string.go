package randomstring

import (
	"crypto/rand"
	"math/big"
)

const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// GenerateRandomString generates a random string with the given length
// using the characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".
// If the length is less than 1, it will return an empty string.
// If the length is greater than the number of characters, it will return an error.
func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		result[i] = characters[num.Int64()]
	}
	return string(result), nil
}
