package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/empijei/identigen/identities"
)

var minage = flag.Int("minage", 25, "The minimum age for random people generation. Must be positive and less than maxage.")
var maxage = flag.Int("maxage", 55, "The maximum age for random people generation. Must be positive and more than minage.")
var n = flag.Int("n", 1, "The amount of random people to generate. Must be positive.")
var dateformat = flag.String("dateformat", "eu", "The format of the dates. Supports: 'eu','us','ja'")
var format = flag.String("format", "", "The format of the output. Supports: 'json', 'csv'")
var filter = flag.String("fields", "all", "The comma separated list of fields to print. Use 'all' to print all of them")

func main() {
	flag.Parse()
	if *minage >= *maxage || *n <= 0 {
		flag.PrintDefaults()
	}
	if *filter != "all" {
		tmp := strings.Split(*filter, ",")
		err := identities.SetFilter(tmp)
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
		}
	}

	identities.LocalizDate = identities.NewDateFormat(*dateformat)
	people, err := identities.RandomPeople(*minage, *maxage, *n)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch *format {
	case "json":
		b, err := json.MarshalIndent(&people, "", "\t")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		_, _ = os.Stdout.Write(b)
	case "csv":
		err := identities.MarshalCSV(people, os.Stdout)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
	default:
		for _, person := range people {
			fmt.Println(person)
		}
	}
}
