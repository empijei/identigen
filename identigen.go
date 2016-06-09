package main

import (
	"flag"
	"fmt"

	"github.com/empijei/identigen/identities"
)

//TODO provide flags

var minage = flag.Int("minage", 25, "The minimum age for random people generation")
var maxage = flag.Int("maxage", 50, "The maximum age for random people generation")
var n = flag.Int("n", 1, "The amount of random people to generate")

func main() {
	//	test := identigen.Person{FirstName: "Roberto", Surname: "Clapis", Town: "CARATE BRIANZA (MB)", TownCode: "B729"}
	//	test.BirthDate = time.Date(1992, time.January, 31, 12, 0, 0, 0, time.UTC)
	//	cf, err := test.CodiceFiscale()
	//	if err != nil {
	//	}
	flag.Parse()
	ppl := identigen.RandomPeople(*minage, *maxage, *n)
	for _, ppl := range ppl {
		cf, _ := ppl.CodiceFiscale()
		fmt.Printf("%v, %v\n", ppl, cf)
	}
}
