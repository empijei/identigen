package identigen

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"./lists"
)

type Person struct {
	//TODO protect these fields, use accessor and cache data
	firstName, surname string
	genderIsFemale     bool
	birthDate          time.Time
	town, townCode     string
	fiscalCode         string
	phone              string
	id                 string
}

func (p *Person) FirstName() string {
	return p.firstName
}
func (p *Person) Surname() string {
	return p.surname
}
func (p *Person) Gender() string {
	if p.genderIsFemale {
		return "Female"
	} else {
		return "Male"
	}
}
func (p *Person) BirthDate() time.Time {
	return p.birthDate
}
func (p *Person) BirthTown() string {
	return p.town
}
func (p *Person) Phone() string {
	return p.phone
}
func (p *Person) ID() string {
	return p.id
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (p Person) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString(p.firstName)
	_, _ = buf.WriteString(" ")
	_, _ = buf.WriteString(p.surname)
	_, _ = buf.WriteString(",")
	_, _ = buf.WriteString(p.Gender())
	_, _ = buf.WriteString(fmt.Sprintf(" %d/%d/%d ", p.birthDate.Day(), int(p.birthDate.Month()), p.birthDate.Year()))
	_, _ = buf.WriteString(p.town)
	return buf.String()
}

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
		person.surname = lists.ItalianSurnames[rand.Int()%len(lists.ItalianSurnames)]
		townAndCode := strings.Split(lists.Comuni[rand.Int()%len(lists.Comuni)], "|")
		person.town = townAndCode[0]
		person.townCode = townAndCode[1]
		people = append(people, person)
		count--
	}
	return
}
