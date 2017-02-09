package identities

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/empijei/identigen/identities/lists"
)

func (p *Person) IBAN() (iban string, err error) {
	//if p.iban != "" {
	//	return p.iban, nil
	//}

	//bbac := fmt.Sprintf(rand.Int()%100000, rand.Int()%100000, rand.Int()%1000000000000)
	bbac := "0542811101000000123456"
	cci := checkDigit(BBAC(bbac)) + 10 //Convert letter to double digit number
	country := 1930                    //IT converted with the (nth + 10) letter of the alphabet (I=19, T=30)

	//	toCheck := strconv.Itoa(cci) + bbac + strconv.Itoa(country) + strconv.Itoa(checks) + strconv.Itoa(checks)

	toCheck, ok := (&big.Int{}).SetString(strconv.Itoa(cci)+bbac+strconv.Itoa(country)+"00", 10)

	if !ok {
		return "", errors.New("Conversion to bigInt failed")
	}
	//log.Println(toCheck)
	//toCheck.SetString("330542811101000000123456193000", 10)
	checks := 98 - int((&big.Int{}).Mod(toCheck, big.NewInt(97)).Uint64())

	sChecks := fmt.Sprintf("%02d", checks) //Padding
	//log.Println(sChecks)
	return "IT" + sChecks + string(int('A')+cci-10) + bbac, nil
}

type BBAC string

func (b *BBAC) Digit(n int) int {
	d, _ := strconv.Atoi(string((*b)[n]))
	return d
}

func (b *BBAC) ConvertedDigit(n int) int {
	if x := n % 2; x == 0 {
		return lists.CCI[b.Digit(n)]
	} else {
		return b.Digit(n)
	}
}

func checkDigit(bbac BBAC) int {
	result := 0
	for x := range bbac {
		result += bbac.ConvertedDigit(x)
	}
	result = result % 26
	return result
}
