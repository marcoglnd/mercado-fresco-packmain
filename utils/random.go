package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i += 1 {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomCategory() string {
	return RandomString(6)
}

func RandomCode() int {
	return int(RandomInt(0, 1000))
}

func RandomInt64() int64 {
	return RandomInt(0, 1000)
}

func RandomFloat64() float64 {
	return float64(RandomInt(0, 1000))
}