package identigen

import (
	"fmt"

	"github.com/empijei/identigen/identities/lists"
)

func (p *Person) IBAN() (iban string, err error) {
	//if p.iban != "" {
	//	return p.iban, nil
	//}

	//bbac := fmt.Sprintf(rand.Int()%100000, rand.Int()%100000, rand.Int()%1000000000000)
	bbac := [22]int{0, 5, 4, 2, 8, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6}
	fmt.Println(bbac)
	cci := checkDigit(bbac)
	fmt.Println(cci)
	return "a", nil
}

func checkDigit(bbac []int) int {
	fmt.Println(bbac, len(bbac))
	result := 0
	for x := range bbac {
		if x%2 != 0 {
			fmt.Println("--", x, bbac[x])
			result += bbac[x]
		} else {
			fmt.Println(x, lists.CCI[string(bbac[x])])
			result += lists.CCI[string(bbac[x])]
		}
	}
	fmt.Println(result)
	result = result % 26
	return result
}
