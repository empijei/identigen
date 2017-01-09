package identigen

import "math"

func nthdigit(num, pos int) int {
	return int(float64(num)/math.Pow10(pos)) % 10
}
