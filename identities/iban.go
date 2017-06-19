package identities

import (
	"bytes"
	"math/big"
	"math/rand"
	"strconv"
	"unicode"

	"github.com/empijei/identigen/identities/lists"
)

const country = "IT" // iso for italy

// Iban represents an IBAN object
type Iban struct {
	BankName  string
	Iban      string
	abi       string
	cab       string
	cc        string
	cinCache  string
	bbanCache string
}

//Returns custom Italian check digit for the IBAN.
func (i *Iban) cin() string {
	if i.cinCache != "" {
		return i.cinCache
	}
	wOdd := []int{1, 0, 5, 7, 9, 13, 15, 17, 19, 21, 2, 4, 18, 20, 11, 3, 6, 8, 12, 14, 16, 10, 22, 25, 24, 23, 27, 28, 26}

	sb := bytes.NewBuffer(make([]byte, 0, 22))
	sb.WriteString(i.abi)
	sb.WriteString(i.cab)
	sb.WriteString(i.cc)
	ibanTmp := sb.String()

	totalWeight := 0

	for i, elem := range ibanTmp {
		p, _ := strconv.Atoi(string(elem))

		if (i % 2) == 0 {
			totalWeight += wOdd[p]
		} else {
			totalWeight += p
		}
	}
	i.cinCache = string((totalWeight % 26) + 65)
	return i.cinCache
}

func (i *Iban) bban() string {
	if i.bbanCache != "" {
		return i.bbanCache
	}
	sb := bytes.NewBuffer(make([]byte, 0, 23))
	sb.WriteString(i.cin())
	sb.WriteString(i.abi)
	sb.WriteString(i.cab)
	sb.WriteString(i.cc)
	i.bbanCache = sb.String()
	return i.bbanCache
}

// Gives a string representation of the Iban struct
func (i *Iban) String() string {
	return i.Iban + ", " + i.BankName
}

// IBAN returns an Italian valid IBAN with random bank details.
func (p *Person) IBAN() *Iban {
	if p.iban != nil {
		return p.iban
	}
	bank := lists.Banks[rand.Intn(len(lists.Banks))]
	branch := bank.Branches[rand.Intn(len(bank.Branches))]
	i := &Iban{
		BankName: bank.Name,
		abi:      branch.ABI,
		cab:      branch.CAB[rand.Intn(len(branch.CAB))],
		cc:       randString([]rune("1234567890"), 12),
	}
	sb := bytes.NewBuffer(make([]byte, 0, 31))
	sb.WriteString(country)
	sb.WriteString(i.cd())
	sb.WriteString(" ")
	sb.WriteString(i.cin())
	sb.WriteString(" ")
	sb.WriteString(i.abi)
	sb.WriteString(" ")
	sb.WriteString(i.cab)
	sb.WriteString(" ")
	sb.WriteString(i.cc)
	i.Iban = sb.String()
	p.iban = i
	return p.iban
}

// cd returns the EU standard check digits for the IBAN.
func (i *Iban) cd() string {
	ibanTmp := i.bban() + country + "00"
	var ret string
	aux := ""

	for _, elem := range ibanTmp {
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

	bfNumber := new(big.Int)
	mod := new(big.Int)
	result := new(big.Int)
	div := big.NewInt(97)
	sub := big.NewInt(98)

	bfNumber.SetString(aux, 10)

	mod.Mod(bfNumber, div)
	result.Sub(sub, mod)
	ret = result.String()

	// Padding with 0
	if len(ret) == 1 {
		ret = "0" + ret
	}
	return ret
}
