package identigen

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func nthdigit(num, pos int) int {
	return int(float64(num)/math.Pow10(pos)) % 10
}

func randString(charSet []rune, length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(b)
}
