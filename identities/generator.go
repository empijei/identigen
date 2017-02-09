package identities

import (
	"errors"
	"math/rand"
	"time"

	"github.com/empijei/identigen/identities/lists"
)

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
