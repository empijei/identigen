package identigen

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func (p *Person) PartitaIva() (pi string, err error) {
	pi = fmt.Sprintf("%07d%03d", rand.Int()%1000000, rand.Int()%100+1)
	num, _ := strconv.Atoi(pi)
	lastDigit := transformation(num, 10)
	pi = fmt.Sprintf("%s%d", pi, lastDigit)
	return
}

func nthdigit(num, pos int) int {
	return int(float64(num)/math.Pow10(pos)) % 10
}

func transformation(num, len int) int {
	var digit, evenSum, oddSum int
	for pos := 0; pos < len; pos++ {
		digit = nthdigit(num, pos)
		if pos%2 == 0 {
			tmp := digit * 2
			if tmp > 9 {
				tmp -= 9
			}
			evenSum += tmp
		} else {
			oddSum += digit
		}
	}
	T := (evenSum + oddSum) % 10
	return 10 - T
}
