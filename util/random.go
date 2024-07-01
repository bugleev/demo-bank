package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alhabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alhabet)

	for i := 0; i < n; i++ {
		c := alhabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomID() int64 {
	return RandomInt(1000, 1000000)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandmCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@example.com", RandomString(6))
}
