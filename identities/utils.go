package identigen

import (
	"math"
	"math/rand"
	"time"
)

type DateFormat int

var LocalizDate DateFormat

func (d DateFormat) Format() string {
	switch int(d) {
	case 1:
		return "01/02/2006"
	case 2:
		return "2006/01/02"
	default:
		return "02/01/2006"
	}
}

func NewDateFormat(fmt string) DateFormat {
	switch fmt {
	case "us":
		return DateFormat(1)
	case "ja":
		return DateFormat(2)
	default:
		return DateFormat(0)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func nthdigit(num int64, pos int) int {
	return int(float64(num)/math.Pow10(pos)) % 10
}

func randString(charSet []rune, length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(b)
}
