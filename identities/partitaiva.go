package identigen

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func (p *Person) PartitaIva() (pi string, err error) {
	//if p.partitaiva != "" {
	//	return p.partitaiva, nil
	//}

	pi = fmt.Sprintf("%d%3d", rand.Int()%10000000, rand.Int()%100+1)
	fmt.Println(pi)
	num, _ := strconv.Atoi(pi)
	lastDigit := transformation(num, 10)
	fmt.Println(lastDigit)
	return
}

func nthdigit(pos, num int) int {
	return int(float64(num)/math.Pow10(pos)) % 10
}

func transformation(num, len int) int {
	var digit, Z, evenSum, oddSum int
	for pos := 0; pos < len; pos++ {
		if (pos+1)%2 == 0 {
			digit = nthdigit(num, pos)
			tmp := digit * 2
			if tmp > 9 {
				tmp -= 9
			}
			evenSum += tmp
			if digit >= 5 {
				Z += 1
			}
		} else {
			oddSum += nthdigit(num, pos)
		}
	}
	fmt.Println(evenSum, oddSum, Z)
	T := (Z + evenSum + oddSum) % 10
	return (10 - T) % 10
}
