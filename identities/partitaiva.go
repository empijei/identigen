package identities

import (
	"fmt"
	"math/rand"
	"strconv"
)

// PartitaIVa returns a valid partita iva number and palce of issue.
func (p *Person) PartitaIva() (pi string, county string) {
	//Need to know where you are to compute your P.I.
	if p.locationCode == 0 || p.partitaIvaCounty == "" {
		_ = p.Address()
	}
	if p.partitaIva != "" {
		return p.partitaIva, p.partitaIvaCounty
	}
	county = p.partitaIvaCounty
	pi = fmt.Sprintf("%07d%03d", rand.Intn(10e6), p.locationCode)
	lastDigit := piCheckDigit(pi)
	pi = fmt.Sprintf("IT%s%d", pi, lastDigit)
	p.partitaIva = pi
	return
}

func piCheckDigit(num string) int {
	var digit, evenSum, oddSum int
	for pos := 0; pos < len(num); pos++ {
		digit, _ = strconv.Atoi(string(num[pos]))
		if pos%2 != 0 {
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
	return (10 - T) % 10
}
