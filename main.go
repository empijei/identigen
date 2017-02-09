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
var dateformat = flag.String("dateformat", "eu", "The format of the dates. Supports: 'eu','us','ja'")
var f = flag.String("f", "json", "The format of the output. Supports: 'json', 'csv'")

func main() {
	flag.Parse()
	if *minage >= *maxage || *n <= 0 {
		flag.PrintDefaults()
	}
	identities.LocalizDate = identities.NewDateFormat(*dateformat)
	people, err := identities.RandomPeople(*minage, *maxage, *n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *f == "json" {
		b, err := json.MarshalIndent(&people, "", "\t")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		_, _ = os.Stdout.Write(b)
	} else if *f == "csv" {
		err := identities.MarshalCSV(people, os.Stdout)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	}
	fmt.Println()
}
