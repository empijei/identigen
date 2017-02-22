package identities

import (
	"math/big"
	"strconv"
	"unicode"
)

const country = "IT" // iso for italy

func (p *Person) IBAN() (iban string, err error) {
	if p.iban != "" {
		return p.iban, nil
	}

	abi := "05428"
	cab := "11101"
	cc := "000000123456"
	// bbac := "0542811101000000123456" // cci should be X and CCI 60
	// bbac := "0326848670052319093140"   // cci should be A and CCI 07
	//bbac := "0301503200000003564232"   // cci should be W and CCI 24
	iban = GenerateIban(abi, cab, cc)
	p.iban = iban
	return
}

func GenerateIban(abi, cab, cc string) string { // Expecting ABI, CAB and cc to be padded with zeros
	cin := cin(abi, cab, cc)
	bban := cin + abi + cab + cc
	cd := checkDigit(bban)
	iban := country + cd + cin + abi + cab + cc

	return iban

}

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

	// Idk what i'm doing
	bf_number := new(big.Int)
	mod := new(big.Int)
	result := new(big.Int)
	div := big.NewInt(97)
	sub := big.NewInt(98)
	// SO MANY BIGINT

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

func cin(abi, cab, cc string) string {
	w_even := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28}
	w_odd := []int{1, 0, 5, 7, 9, 13, 15, 17, 19, 21, 2, 4, 18, 20, 11, 3, 6, 8, 12, 14, 16, 10, 22, 25, 24, 23, 27, 28, 26}

	var p int
	iban_tmp := abi + cab + cc
	total_weight := 0

	for i, elem := range iban_tmp {

		if unicode.IsLetter(elem) {
			p = int(elem - 65)
		} else {
			p = int(elem - 48)
		}

		if (i % 2) == 0 { // if i is even i+1 is odd
			total_weight += w_odd[p]
		} else {
			total_weight += w_even[p]
		}
	}
	return string((total_weight % 26) + 65)
}
