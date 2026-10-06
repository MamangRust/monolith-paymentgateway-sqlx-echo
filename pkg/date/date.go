package date

import (
	"time"

	"math/rand"
)

// GenerateExpireDate generates a random date in the future within the next 5 years as the expiration date.
// The date is in the format "YYYY-MM-DD", and the time is midnight UTC.
func GenerateExpireDate() time.Time {
	now := time.Now()
	year := now.Year() + rand.Intn(5)
	month := time.Month(rand.Intn(12) + 1)
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}
