package utils

import (
	"math/rand"
	"time"

	"github.com/ZeroTechh/UserExtraService/core/types"
)

// generateRandStr generates a random string
func generateRandStr(length int) string {
	charset := "1234567890abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// MockData returns mock user extra data for testing
func MockData() types.Extra {
	randomStr := generateRandStr(10)
	return types.Extra{
		UserID:      randomStr,
		FirstName:   randomStr,
		LastName:    randomStr,
		Gender:      "male",
		BirthdayUTC: int64(864466669),
	}
}
