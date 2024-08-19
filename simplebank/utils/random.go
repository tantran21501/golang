package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var globalRand *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	globalRand = rand.New(source)
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + globalRand.Int63n(max - min + 1) // 0 -> max - min
}

// RandomString generate a random string of
func RandomString(n int) string{
	var sb strings.Builder
	k := len(alphabet)

	for  i := 0; i < n; i++ {
		c :=  alphabet[globalRand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Random Owner generate a random owner name
func RandomOwner() string{
	return RandomString(6)
}

// RandomMoney generate a random amount of money 
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate a random currency
func RandomCurrency() string {
	currencies := []string{"USD","EURO","CAD"}
	n := len(currencies)
	return currencies[globalRand.Intn(n)]
}