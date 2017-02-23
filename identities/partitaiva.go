package identities

import (
	"fmt"
	"math/rand"
	"strconv"
)

func (p *Person) PartitaIva() (pi string, county string) {
	if p.partitaIva != "" {
		return p.partitaIva, p.partitaIvaCounty
	}
	if p.locationCode == 0 {
		_ = p.Address()
	}
	pi = fmt.Sprintf("%07d%03d", rand.Int()%1000000, p.locationCode)
	num, _ := strconv.Atoi(pi)
	lastDigit := transformation(num, 10)
	pi = fmt.Sprintf("%s%d", pi, lastDigit)
	p.partitaIva = pi
	return
}

func transformation(num, len int) int {
	var digit, evenSum, oddSum int
	for pos := 0; pos < len; pos++ {
		digit = nthdigit(int64(num), pos)
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
