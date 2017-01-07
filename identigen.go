package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/empijei/identigen/identities"
)

//TODO provide flags

var minage = flag.Int("minage", 25, "The minimum age for random people generation. Must be positive and less than maxage.")
var maxage = flag.Int("maxage", 50, "The maximum age for random people generation. Must be positive and more than minage.")
var n = flag.Int("n", 1, "The amount of random people to generate. Must be positive.")

func main() {
	type toPrint struct {
		CodiceFiscale    string
		PartitaIva       string
		ComunePartitaIva string
		Documento        string
		CartaCredito     string
	}
	var items toPrint

	flag.Parse()
	if *minage >= *maxage || *n <= 0 {
		flag.PrintDefaults()
	}
	ppl := identigen.RandomPeople(*minage, *maxage, *n)
	for _, ppl := range ppl {
		cf, _ := ppl.CodiceFiscale()
		pi, county, _ := ppl.PartitaIva()
		id := ppl.ID()
		cc, _ := ppl.CartaCredito()
		fmt.Printf("%v\n", ppl)

		items = toPrint{
			CodiceFiscale:    cf,
			PartitaIva:       pi,
			ComunePartitaIva: county,
			Documento:        id,
			CartaCredito:     cc,
		}

		b, err := json.MarshalIndent(items, " ", "   ")
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)
		fmt.Println()
	}
}
