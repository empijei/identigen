package identigen

import (
	"math/rand"
	"strings"
	"time"

	"github.com/empijei/identigen/identities/lists"
)

//TODO test
func RandomPeople(minage, maxage int, count int) (people []Person) {
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
		age := rand.Int()%(maxage-minage) + minage
		person.birthDate = time.Date(time.Now().Year()-age, time.Month(rand.Int()%12+1), rand.Int()%28+1, 12, 0, 0, 0, time.UTC)
		person.lastName = lists.ItalianSurnames[rand.Int()%len(lists.ItalianSurnames)]
		townAndCode := strings.Split(lists.Comuni[rand.Int()%len(lists.Comuni)], "|")
		person.town = townAndCode[0]
		person.townCode = townAndCode[1]
		person.phone = "3" + randString([]rune("1234567890"), 9)
		people = append(people, person)
		count--
	}
	return
}
