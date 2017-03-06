package identities

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func MainModule(args map[string]interface{}, out io.Writer) (err error) {
	clamp := func(val, min, max int) int {
		if val < min {
			return min
		}
		if val > max {
			return max
		}
		return val
	}

	dt_fmt := args["dt_fmt"].(string)
	minage := clamp(args["minage"].(int), 1, 200)
	maxage := clamp(args["maxage"].(int), 1, 200)
	number := args["number"].(int)
	format := args["format"].(string)
	fields := args["fields"].(string)

	if number <= 0 {
		return fmt.Errorf("'number' should be positive")
	}
	if fields != "all" {
		tmp := uniqSlice(strings.Split(fields, ","))
		err := SetFilter(tmp)
		if err != nil {
			return err
		}
	}

	LocalizDate = NewDateFormat(dt_fmt)
	people, err := RandomPeople(minage, maxage, number)
	if err != nil {
		return err
	}

	formats := uniqSlice(strings.Split(format, ","))
	for _, f := range formats {
		switch f {
		case "json":
			b, err := json.MarshalIndent(&people, "", "\t")
			if err != nil {
				return err
			}
			_, _ = out.Write(b)
			fmt.Fprintln(out)
		case "xml":
			_, _ = out.Write([]byte("<People>\n"))
			b, err := xml.MarshalIndent(&people, "\t", "\t")
			if err != nil {
				return err
			}
			_, _ = out.Write(b)
			_, _ = out.Write([]byte("\n</People>"))
			fmt.Fprintln(out)
		case "csv":
			err := MarshalCSV(people, out)
			if err != nil {
				return err
			}
			fmt.Fprintln(out)
		default:
			for _, person := range people {
				fmt.Fprintln(out, person)
			}
		}
	}
	return nil
}

func uniqSlice(in []string) []string {
	out := make([]string, 0, len(in))
	supportMap := make(map[string]struct{})
	for _, f := range in {
		if _, ok := supportMap[f]; !ok {
			out = append(out, f)
			supportMap[f] = struct{}{}
		}
	}
	return out
}

func RandomPeople(minage, maxage int, count int) (people []Person, err error) {
	if minage > maxage {
		return nil, fmt.Errorf("maxage(%d) should not be less than minage(%d)", maxage, minage)
	}
	for count > 0 {
		people = append(people, *NewPerson(minage, maxage))
		count--
	}
	return
}
