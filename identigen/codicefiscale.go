package identigen

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (p *Person) CodiceFiscale() (cf string, err error) {
	var startTime time.Time
	if p.FirstName == "" || p.Surname == "" || p.BirthDate == startTime || p.Town == "" {
		err = fmt.Errorf("Missing fields in person")
		return
	}
	cf += threePad(getConsonants(p.Surname), p.Surname)
	cf += threePad(fixFirstNameConsonants(getConsonants(p.FirstName)), p.FirstName)
	cf += p.birthDayStringCalc()
	if p.TownCode == "" {
		//TODO
	}
	cf += p.TownCode
	cf += checksum(cf)
	return
}

//TODO create a cfbuilder, move function to its methods, use a slice or buffer

func threePad(tmp string, name string) string {
	if delta := len(tmp) - 3; delta != 0 {
		if delta > 0 {
			tmp = tmp[:3]
		}
		if delta < 0 {
			vow := getVowels(name)
			if -delta <= len(vow) {
				tmp += vow[:-delta]
			} else {
				tmp += vow
			}
		}
		for len(tmp) < 3 {
			tmp += "X"
		}
	}
	return tmp
}

func getVowels(input string) (output string) {
	input = strings.ToUpper(input)
	for _, c := range input {
		c := string(c)
		if strings.ContainsAny(c, "AEIOU") {
			output += c
		}
		if i := strings.Index("ÀÌÈÉÒ", c); i >= 0 {
			//TODO remove accent
			output += string("AIEEO"[i])
		}
	}
	return
}

func getConsonants(input string) (output string) {
	input = strings.ToUpper(input)
	for _, c := range input {
		c := string(c)
		if strings.ContainsAny(c, "QWRTYPSDFGHJKLZXCVBNM") {
			output += c
		}
	}
	return
}

func (p *Person) birthDayStringCalc() (seq string) {
	seq += strconv.Itoa(p.BirthDate.Year())
	seq = seq[2:]
	seq += string("ABCDEHLMPRST"[p.BirthDate.Month()-1])
	tmp := p.BirthDate.Day()
	if p.GenderIsFemale {
		tmp += 40
	}
	seq += strconv.Itoa(tmp)
	return
}

func fixFirstNameConsonants(seq string) string {
	if len(seq) > 3 {
		seq = seq[0:1] + seq[2:]
	}
	return seq
}

func checksum(cf string) (chech string) {
	oddcheck := map[string]int{"0": 1, "9": 21, "I": 19, "R": 8, "1": 0, "A": 1, "J": 21, "S": 12, "2": 5, "B": 0, "K": 2, "T": 14, "3": 7, "C": 5, "L": 4, "U": 16, "4": 9, "D": 7, "M": 18, "V": 10, "5": 13, "E": 9, "N": 20, "W": 22, "6": 15, "F": 13, "O": 11, "X": 25, "7": 17, "G": 15, "P": 3, "Y": 24, "8": 19, "H": 17, "Q": 6, "Z": 23}
	evencheck := make(map[string]int)
	for i, c := range "0123456789" {
		evencheck[string(c)] = i
	}
	for i, c := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		evencheck[string(c)] = i
	}
	var sum int
	for i, c := range cf {
		if (i+1)%2 == 0 {
			sum += evencheck[string(c)]
		} else {
			sum += oddcheck[string(c)]
		}
	}
	remainder := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	chech = string(remainder[sum%26])
	return
}
