package identities

import (
	"math/big"
	"strconv"
	"unicode"
)

const country = "IT" // iso for italy

//Returns an Italian valid IBAN with random bank details.
func (p *Person) IBAN() (iban string) {
	if p.iban != "" {
		return p.iban
	}

	abi := randString([]rune("1234567890"), 5) //"05428"
	cab := randString([]rune("1234567890"), 5) //"11101"
	cc := randString([]rune("1234567890"), 12) //"000000123456"

	cin := cin(abi, cab, cc)
	bban := cin + abi + cab + cc
	cd := checkDigit(bban)
	iban = country + cd + cin + abi + cab + cc

	p.iban = iban
	return
}

//Returns the EU standard check digits for the IBAN.
func checkDigit(bban string) string {
	iban_tmp := bban + "IT00"
	var ret string
	aux := ""

	for _, elem := range iban_tmp {
		if unicode.IsLetter(elem) {
			var tmp string
			tmp = strconv.Itoa(int(elem - 65 + 10))
			if len(tmp) == 1 {
				tmp = "0" + tmp
			}
			aux += tmp

		} else {
			aux += string(elem)
		}
	}

	bf_number := new(big.Int)
	mod := new(big.Int)
	result := new(big.Int)
	div := big.NewInt(97)
	sub := big.NewInt(98)

	bf_number.SetString(aux, 10)

	mod.Mod(bf_number, div)
	result.Sub(sub, mod)
	ret = result.String()

	// Padding with 0
	if len(ret) == 1 {
		ret = "0" + ret
	}
	return ret
}

//Returns custom Italian check digit for the IBAN.
func cin(abi, cab, cc string) string {
	w_odd := []int{1, 0, 5, 7, 9, 13, 15, 17, 19, 21, 2, 4, 18, 20, 11, 3, 6, 8, 12, 14, 16, 10, 22, 25, 24, 23, 27, 28, 26}

	iban_tmp := abi + cab + cc
	total_weight := 0

	for i, elem := range iban_tmp {
		var p int
		if unicode.IsLetter(elem) {
			p = int(elem - 65)
		} else {
			p = int(elem - 48)
		}

		if (i % 2) == 0 { // if i is even i+1 is odd
			total_weight += w_odd[p]
		} else {
			total_weight += p
		}
	}
	return string((total_weight % 26) + 65)
}
