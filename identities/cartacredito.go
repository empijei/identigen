package identities

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type CartaCredito struct {
	Number  string
	Cvv     string
	Issuer  string
	ExpDate string
}

type ccGen func() int

//Source: https://en.wikipedia.org/wiki/Payment_card_number
var ccGens = map[string]ccGen{
	"American Express": func() int {
		return rand.Intn(300) + 3400
	},
	"Maestro": func() int {
		return rand.Intn(1000) + 6000
	},
	"MasterCard": func() int {
		return rand.Intn(400) + 5100
	},
	"Visa": func() int {
		return rand.Intn(1000) + 4000
	},
}

func ccformatter(cc string) string {
	var toret string
	toret = strings.Join([]string{cc[:4], cc[4:8], cc[8:12], cc[12:16]}, "-")
	return toret
}

//Returns a valid CartaCredito object with credit card number, cvv, issuer and expiration date.
func (p *Person) CartaCredito() *CartaCredito {
	if p.cc != nil {
		return p.cc
	}

	cc := &CartaCredito{}
	num := rand.Int63n(10e11)
	//FIXME this is not really random
	for emit, val := range ccGens {
		cc.Issuer = emit
		num += int64(val()) * 10e11
		break
	}
	lastDigit := transform(num)
	cc.Number = fmt.Sprintf("%015d%d", num, lastDigit)
	cc.Number = ccformatter(cc.Number)

	//This generates a 4 chars long CVV for Amex, 3 in all other cases
	cc.Cvv = randString([]rune("0123456789"),
		map[bool]int{true: 4, false: 3}[cc.Issuer == "American Express"])

	//CC expires in 6 years, 6 months 6 days from now
	cc.ExpDate = time.Now().AddDate(6, 6, 6).Format("01/06")
	p.cc = cc
	return cc
}

//Calculate the check digit of a credit card number
func transform(num int64) int {
	cclen := fmt.Sprintf("%d", num)
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
		return 10 - nthdigit(int64(summed), 0)
	}
	return 0
}
