package identities

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/empijei/identigen/identities/lists"
)

func MainModule(args map[string]interface{}, out io.Writer) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(out, "Error occurred: ", r)
		}
	}()
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
		panic("'number' should be positive")
	}
	if fields != "all" {
		tmp := uniqSlice(strings.Split(fields, ","))
		err := SetFilter(tmp)
		if err != nil {
			panic(err)
		}
	}

	LocalizDate = NewDateFormat(dt_fmt)
	people, err := RandomPeople(minage, maxage, number)
	if err != nil {
		panic(err)
	}

	formats := uniqSlice(strings.Split(format, ","))
	for _, f := range formats {
		switch f {
		case "json":
			b, err := json.MarshalIndent(&people, "", "\t")
			if err != nil {
				panic(err)
			}
			_, _ = out.Write(b)
			fmt.Fprintln(out)
		case "xml":
			b, err := xml.MarshalIndent(&people, "", "\t")
			if err != nil {
				panic(err)
			}
			_, _ = out.Write(b)
			fmt.Fprintln(out)
		case "csv":
			err := MarshalCSV(people, out)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(out)
		default:
			for _, person := range people {
				fmt.Fprintln(out, person)
			}
		}
	}
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

//TODO test
func RandomPeople(minage, maxage int, count int) (people []Person, err error) {
	if minage > maxage {
		return nil, errors.New(fmt.Sprintf("maxage (%d) should not be less than minage(%d)", maxage, minage))
	}
	for count > 0 {
		person := Person{}
		person.genderIsFemale = rand.Int()%2 == 0
		var names []string
		if person.genderIsFemale {
			names = lists.ItalianFemaleNames
		} else {
			names = lists.ItalianMaleNames
		}
		person.firstName = names[rand.Int()%len(names)]
		var age int
		if minage == maxage {
			age = minage
		} else {
			age = rand.Int()%(maxage-minage) + minage
		}
		person.birthDate = time.Date(time.Now().Year()-age, time.Month(rand.Int()%12+1), rand.Int()%28+1, 12, 0, 0, 0, time.UTC)
		person.lastName = lists.ItalianSurnames[rand.Int()%len(lists.ItalianSurnames)]
		birthInfo := lists.BirthInfo[rand.Int()%len(lists.BirthInfo)]
		person.town = birthInfo.Paese
		person.townCode = birthInfo.CodiceCatasto
		person.birthDistrict = birthInfo.Provincia
		person.mobilePhone = "3" + randString([]rune("1234567890"), 9)
		people = append(people, person)
		count--
	}
	return
}
