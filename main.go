package main

import (
	"fmt"

	"./identigen"
)

func main() {
	//	test := identigen.Person{FirstName: "Roberto", Surname: "Clapis", Town: "CARATE BRIANZA (MB)", TownCode: "B729"}
	//	test.BirthDate = time.Date(1992, time.January, 31, 12, 0, 0, 0, time.UTC)
	//	cf, err := test.CodiceFiscale()
	//	if err != nil {
	//	}
	ppl := identigen.RandomPeople(25, 50, 4)
	for _, ppl := range ppl {
		cf, _ := ppl.CodiceFiscale()
		fmt.Printf("%v, %v\n", ppl, cf)
	}
}
