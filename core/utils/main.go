package utils

import (
	"math/rand"
	"time"

	"github.com/ZeroTechh/UserExtraService/core/types"
)

func randStr(length int) string {
	charset := "1234567890abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Mock returns mock user extra data for testing
func Mock() types.Extra {
	randomStr := randStr(10)
	return types.Extra{
		UserID:      randomStr,
		FirstName:   randomStr,
		LastName:    randomStr,
		Gender:      "male",
		BirthdayUTC: int64(864466669),
	}
}
