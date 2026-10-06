package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateExpireDate(t *testing.T) {
	expireDate := GenerateExpireDate()
	now := time.Now()

	assert.GreaterOrEqual(t, expireDate.Year(), now.Year())

	assert.LessOrEqual(t, expireDate.Year(), now.Year()+5)

	assert.GreaterOrEqual(t, int(expireDate.Month()), 1)
	assert.LessOrEqual(t, int(expireDate.Month()), 12)

	assert.Equal(t, 1, expireDate.Day())
	assert.Equal(t, 0, expireDate.Hour())
	assert.Equal(t, 0, expireDate.Minute())
	assert.Equal(t, 0, expireDate.Second())
	assert.Equal(t, time.UTC, expireDate.Location())
}
