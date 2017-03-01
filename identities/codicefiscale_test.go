package identities

import (
	"testing"
	"time"
)

/*
Here is the fiscal code of a fictitious Matteo Moretti (male), born in Milan on 9 April 1925:

Surname: MRT
Name: MTT
Birthdate and gender: 25D09
Town of birth: F205
Check character: Z
Fiscal code: MRTMTT25D09F205Z


Here is the fiscal code of a fictitious Samantha Miller (female), born in the USA on 25 September 1982, living in Italy:

Surname: MLL
Name: SNT
Birthdate and gender: 82P65
Municipality of birth: Z404
Check character: U
Fiscal code: MLLSNT82P65Z404U
*/

var cfTests = []struct {
	input    Person
	expected string
}{
	{input: Person{firstName: "Roberto", lastName: "Clapis", town: "CARATE BRIANZA (MB)", townCode: "B729", birthDate: time.Date(1992, time.January, 31, 12, 0, 0, 0, time.UTC)},
		expected: "CLPRRT92A31B729B"},
	{input: Person{firstName: "Samantha", lastName: "Miller", town: "Murica", townCode: "Z404", genderIsFemale: true, birthDate: time.Date(1982, time.September, 25, 12, 0, 0, 0, time.UTC)},
		expected: "MLLSNT82P65Z404U"},
	{input: Person{firstName: "Na", lastName: "OOO", town: "Murica", townCode: "Z404", genderIsFemale: true, birthDate: time.Date(1982, time.September, 25, 12, 0, 0, 0, time.UTC)},
		expected: "OOONAX82P65Z404D"},
}

func TestCodiceFiscale(t *testing.T) {
	for _, tc := range cfTests {
		if cf := tc.input.CodiceFiscale(); cf != tc.expected {
			t.Errorf("Failed test with %v\n, calculated: %v, expected: %v", tc.input, cf, tc.expected)
		}
	}
}
