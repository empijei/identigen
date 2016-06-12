package main

import (
	"flag"
	"fmt"

	"github.com/empijei/identigen/identities"
)

//TODO provide flags

var minage = flag.Int("minage", 25, "The minimum age for random people generation. Must be positive and less than maxage.")
var maxage = flag.Int("maxage", 50, "The maximum age for random people generation. Must be positive and more than minage.")
var n = flag.Int("n", 1, "The amount of random people to generate. Must be positive.")

func main() {
	flag.Parse()
	if *minage >= *maxage || *n <= 0 {
		flag.PrintDefaults()
	}
	ppl := identigen.RandomPeople(*minage, *maxage, *n)
	for _, ppl := range ppl {
		cf, _ := ppl.CodiceFiscale()
		id := ppl.ID()
		fmt.Printf("%v, %v, %v", ppl, cf, id)
	}
}
