package identities

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type CartaCredito struct {
	Number  string
	Cvv     string
	Issuer  string
	ExpDate string
}

type ccChecker func(string) bool

var ccs = map[string]ccChecker{
	"American Express": func(cc string) bool {
		re := regexp.MustCompile("^3(4|7)")
		return re.Match([]byte(cc))
	},
	"Maestro": func(cc string) bool {
		re := regexp.MustCompile("^((5(0|[678]))|6)")
		return re.Match([]byte(cc))
	},
	"MasterCard": func(cc string) bool {
		re := regexp.MustCompile("^(5[1-5])")
		return re.Match([]byte(cc)) || (cc[:4] <= "2720" && "2221" <= cc[:4])
	},
	"Visa": func(cc string) bool {
		return strings.HasPrefix(cc, "4")
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
	num := rand.Int63n(10e15)
	lastDigit := transform(num)
	cc := &CartaCredito{
		Number: fmt.Sprintf("%015d%d", num, lastDigit),
	}
	for iss, chk := range ccs {
		if chk(cc.Number) {
			cc.Issuer = iss
			break
		}
	}
	if cc.Issuer == "" {
		cc.Issuer = "Other"
	}
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
