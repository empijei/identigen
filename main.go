package main

import (
	"flag"
	"os"
	"strings"

	"github.com/empijei/identigen/identities"
)

var minage = flag.Int("minage", 25, "The minimum age for random people generation. Must be positive and less than maxage.")
var maxage = flag.Int("maxage", 55, "The maximum age for random people generation. Must be positive, less or equal than 200 and more than minage.")
var number = flag.Int("number", 1, "The amount of random people to generate. Must be positive.")
var dt_fmt = flag.String("dt_fmt", "eu", "The format of the dates. Supports: 'eu','us','ja'")
var format = flag.String("format", "human", "The comma separated list of formats for the output. Supports: 'json', 'csv', 'human'.")
var fields = flag.String("fields", "all", "The comma separated case-sensitive list of fields to print. Use 'all' to print all of them. Supported fields are: "+strings.Join(identities.AllFields, ","))

func main() {
	flag.Parse()
	args := make(map[string]interface{})
	args["dt_fmt"] = *dt_fmt
	args["minage"] = *minage
	args["maxage"] = *maxage
	args["number"] = *number
	args["format"] = *format
	args["fields"] = *fields

	identities.MainModule(args, os.Stdout)
}
