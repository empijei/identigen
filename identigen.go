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
var maxage = flag.Int("maxage", 55, "The maximum age for random people generation. Must be positive and more than minage.")
var n = flag.Int("n", 1, "The amount of random people to generate. Must be positive.")

func main() {
	flag.Parse()
	if *minage >= *maxage || *n <= 0 {
		flag.PrintDefaults()
	}
	people, err := identigen.RandomPeople(*minage, *maxage, *n)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.MarshalIndent(&people, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	_, _ = os.Stdout.Write(b)
	fmt.Println()
}
