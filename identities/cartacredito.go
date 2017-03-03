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

func (cc *CartaCredito) String() string {
	return cc.Issuer + " " + cc.Number + ", " + cc.Cvv + ", " + cc.ExpDate
}

type ccSeed struct {
	//The issuer name
	issuer string
	//The base value for a CC prefix for the given issuer (e.g. 4000 for Visas)
	base int
	//The range to add to base to obtain a random issuer prefix
	//(e.g. 1000 for Visas, to range from 4000 to 4999)
	delta int
	//The TOTAL length of the credit card number
	bodylength int
	//Length of CVV
	cvvlength int
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
	//Compute the first digits, these determine the issuer
	prefix := rand.Int63n(int64(seed.delta)) + int64(seed.base)
	//Compute how much the prefix needs to be shifted (base10) to the
	//left in order to fall right on the left of the random body.
	prefixShift := int64(math.Pow10(seed.bodylength - len(fmt.Sprintf("%d", prefix)) - 1))
	//Generate random sequence between prefix and checksum
	body := rand.Int63n(prefixShift)
	//Add the prefix on the left
	body += prefix * prefixShift
	//Compute luhn checksum
	checkSum := luhn(body)
	//Concat the checksum and cast to string
	cc.Number = fmt.Sprintf("%0"+strconv.Itoa(seed.bodylength-1)+"d%d", body, checkSum)
	//Beautify CC number
	cc.Number = ccformatter(cc.Number)
	//Set Issuer
	cc.Issuer = seed.issuer
	//Set CVV
	cc.Cvv = randString([]rune("0123456789"), seed.cvvlength)
	//Set expiration date
	cc.ExpDate = time.Now().AddDate(6, 6, 6).Format("01/06")
	return cc
}

func ccformatter(cc string) string {
	switch len(cc) {
	case 16:
		cc = strings.Join([]string{cc[:4], cc[4:8], cc[8:12], cc[12:16]}, "-")
	case 15:
		cc = strings.Join([]string{cc[:5], cc[5:10], cc[10:15]}, "-")
	default:
	}
	return cc
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
			if multiplied >= 10 {
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
