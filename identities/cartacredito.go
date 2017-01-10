package identigen

import (
	"fmt"
	"math/rand"
	"strconv"
)

func (p *Person) CartaCredito() (cc string, err error) {
	if p.cc != "" {
		return p.cc, nil
	}
	num := rand.Intn(10000000000)
	lastDigit := transform(num)

	//cc = strconv.Atoi(fmt.Sprintf("%s%d", strconv.Itoa(num), lastDigit))
	cc = fmt.Sprintf("%s%d", strconv.Itoa(num), lastDigit)
	p.cc = cc
	return
}

func transform(num int) int {
	cclen := strconv.Itoa(num)
	summed := 0
	multiplied := 0

	for i := 0; i < len(cclen); i++ {
		dig := nthdigit(num, i)
		multiplied = dig
		if i%2 == 0 {
			multiplied = dig * 2
			if multiplied > 10 {
				multiplied = multiplied - 9
			}
		}
		summed = summed + multiplied
	}
	if summed%10 != 0 {
		return 10 - nthdigit(summed, 0)
	}
	return 0
}
