package identities

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/empijei/identigen/identities/lists"
)

func (p *Person) PartitaIva() (pi string, county string, err error) {
	if p.partitaIva != "" {
		return p.partitaIva, p.partitaIvaCounty, nil
	}
	var location int
	//Getting a random element from the map.
	for location, county = range lists.Cities {
		break
	}
	pi = fmt.Sprintf("%07d%03d", rand.Int()%1000000, location)
	num, _ := strconv.Atoi(pi)
	lastDigit := transformation(num, 10)
	pi = fmt.Sprintf("%s%d", pi, lastDigit)
	p.partitaIva = pi
	p.partitaIvaCounty = county
	p.residence = lists.Addresses[county]
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