package identities

import (
	"math"
	"math/rand"
	"time"
)

// DateFormat exports the format of the date. Can be 1,2 or anything else.
// 1 (us format) is MM/DD/YYYY
// 2 (ja format) is YYYY/MM/DD
// anything else is formatted in the european style (default) DD/MM/YYYY
type DateFormat int

// LocalizDate is a variable containing the format to use. Set this variable to
// change the way dates are formatted. See "DateFormat" for more.
var LocalizDate DateFormat

// Format returns the date formatted as the DateFormat variable is set
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

// NewDateFormat sets DateFormat as an int from a passed string
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
