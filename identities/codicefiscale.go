package identities

import (
	"bytes"
	"strconv"
	"strings"
)

// CodiceFiscale generates a valid italian codice fiscale
func (p *Person) CodiceFiscale() (cf string) {
	if p.fiscalCode != "" {
		return p.fiscalCode
	}

	composer := bytes.NewBuffer(make([]byte, 0, 16))
	composer.WriteString(threePad(getConsonants(p.lastName), p.lastName))
	composer.WriteString(threePad(fixFirstNameConsonants(getConsonants(p.firstName)), p.firstName))
	composer.WriteString(birthDayStringCalc(p))
	composer.WriteString(p.townCode)
	composer.WriteString(checksum(composer.String()))
	p.fiscalCode = composer.String()
	return p.fiscalCode
}

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

func getVowels(input string) (output string) {
	input = strings.ToUpper(input)
	for _, c := range input {
		c := string(c)
		if strings.ContainsAny(c, "AEIOU") {
			output += c
		}
		if i := strings.Index("ÀÌÈÉÒÙ", c); i >= 0 {
			switch c {
			case "À":
				output += "A"
			case "Ì":
				output += "I"
			case "È":
				fallthrough
			case "É":
				output += "E"
			case "Ò":
				output += "O"
			case "Ù":
				output += "U"
			}
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

func fixFirstNameConsonants(seq string) string {
	if len(seq) > 3 {
		seq = seq[0:1] + seq[2:]
	}
	return seq
}

func birthDayStringCalc(p *Person) (seq string) {
	seq += strconv.Itoa(p.birthDate.Year())
	seq = seq[2:]
	seq += string("ABCDEHLMPRST"[p.birthDate.Month()-1])
	tmp := p.birthDate.Day()
	if p.genderIsFemale {
		tmp += 40
	}
	seq += strconv.Itoa(tmp)
	return
}
