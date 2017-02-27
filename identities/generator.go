package identities

import (
	"encoding/json"
	"errors"
	"flag"
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
	minage := args["minage"].(int)
	maxage := args["maxage"].(int)
	n := args["n"].(int)
	dateformat := args["dateformat"].(string)
	format := args["format"].(string)
	filter := args["filter"].(string)

	if minage >= maxage || n <= 0 {
		flag.PrintDefaults()
	}
	if filter != "all" {
		tmp := strings.Split(filter, ",")
		err := SetFilter(tmp)
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
		}
	}

	LocalizDate = NewDateFormat(dateformat)
	people, err := RandomPeople(minage, maxage, n)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpArray := strings.Split(format, ",")

	//I don't want to have duplicate formats, this is me creating an array with unique items
	tmpMap := make(map[string]struct{})
	var formats []string
	for _, f := range tmpArray {
		if _, ok := tmpMap[f]; !ok {
			formats = append(formats, f)
			tmpMap[f] = struct{}{}
		}
	}
	for _, f := range formats {
		switch f {
		case "json":
			b, err := json.MarshalIndent(&people, "", "\t")
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			_, _ = out.Write(b)
			fmt.Fprintln(out)
		case "csv":
			err := MarshalCSV(people, out)
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			fmt.Fprintln(out)
		default:
			for _, person := range people {
				fmt.Fprintln(out, person)
			}
		}
	}
}

//TODO test
func RandomPeople(minage, maxage int, count int) (people []Person, err error) {
	if minage > maxage {
		return nil, errors.New("maxage should not be less than minage")
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
