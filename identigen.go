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
	flag.Parse()
	if *minage >= *maxage || *n <= 0 {
		flag.PrintDefaults()
	}
	people := identigen.RandomPeople(*minage, *maxage, *n)
	for _, person := range people {
		b, err := json.MarshalIndent(&person, " ", "\t")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		os.Stdout.Write(b)
		fmt.Println()
	}
}
