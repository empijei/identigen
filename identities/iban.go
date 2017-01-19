package identigen

import (
	"strconv"

	"github.com/empijei/identigen/identities/lists"
)

func (p *Person) IBAN() (iban string, err error) {
	//if p.iban != "" {
	//	return p.iban, nil
	//}

	//bbac := fmt.Sprintf(rand.Int()%100000, rand.Int()%100000, rand.Int()%1000000000000)
	bbac := "0542811101000000123456"
	cci := checkDigit(BBAC(bbac))
	return "a", nil
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
