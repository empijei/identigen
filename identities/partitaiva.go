package identities

import (
	"fmt"
	"math/rand"
	"strconv"
)

func (p *Person) PartitaIva() (pi string, county string) {
	//Need to know where you are to compute your P.I.
	if p.locationCode == 0 || p.partitaIvaCounty == "" {
		_ = p.Address()
	}
	if p.partitaIva != "" {
		return p.partitaIva, p.partitaIvaCounty
	}
	county = p.partitaIvaCounty
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
