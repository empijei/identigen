package main

import (
	"flag"
	"os"

	"github.com/empijei/identigen/identities"
)

var minage = flag.Int("minage", 25, "The minimum age for random people generation. Must be positive and less than maxage.")
var maxage = flag.Int("maxage", 55, "The maximum age for random people generation. Must be positive and more than minage.")
var n = flag.Int("n", 1, "The amount of random people to generate. Must be positive.")
var dateformat = flag.String("dateformat", "eu", "The format of the dates. Supports: 'eu','us','ja'")
var format = flag.String("format", "human", "The comma separated list of formats for the output. Supports: 'json', 'csv', 'human'")
var filter = flag.String("fields", "all", "The comma separated list of fields to print. Use 'all' to print all of them")

func main() {
	flag.Parse()
	args := make(map[string]interface{})
	args["minage"] = *minage
	args["maxage"] = *maxage
	args["n"] = *n
	args["dateformat"] = *dateformat
	args["format"] = *format
	args["filter"] = *filter

	identities.MainModule(args, os.Stdout)
}
