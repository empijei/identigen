package identities

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type CartaCredito struct {
	Number  string
	Cvv     string
	Issuer  string
	ExpDate string
}

type ccSeed struct {
	issuer                             string
	base, delta, bodylength, cvvlength int
}

//Source: https://en.wikipedia.org/wiki/Payment_card_number
var ccData = []ccSeed{
	{
		"American Express",
		3400,
		300,
		15,
		4,
	},
	{
		"Maestro",
		6000,
		1000,
		16,
		3,
	},
	{
		"MasterCard",
		5100,
		400,
		16,
		3},
	{
		"Visa",
		4000,
		1000,
		16,
		3,
	},
}

func ccBuilder(seed ccSeed) *CartaCredito {
	cc := &CartaCredito{}
	prefix := rand.Int63n(int64(seed.delta)) + int64(seed.base)
	prefixShift := int64(math.Pow10(seed.bodylength - len(fmt.Sprintf("%d", prefix)) - 1))
	body := rand.Int63n(prefixShift)
	body += prefix * prefixShift
	checkSum := luhn(body)
	cc.Number = fmt.Sprintf("%0"+strconv.Itoa(seed.bodylength-1)+"d%d", body, checkSum)
	if seed.bodylength == 16 {
		cc.Number = ccformatter(cc.Number)
	}
	cc.Issuer = seed.issuer
	cc.Cvv = randString([]rune("0123456789"), seed.cvvlength)
	cc.ExpDate = time.Now().AddDate(6, 6, 6).Format("01/06")
	return cc
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
	p.cc = ccBuilder(ccData[rand.Intn(len(ccData))])
	return p.cc
}

//Calculate the check digit of a credit card number
func luhn(num int64) int {
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
