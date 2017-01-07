package identigen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/empijei/identigen/identities/lists"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Person struct {
	//TODO protect these fields, use accessor and cache data
	firstName, lastName string
	genderIsFemale      bool
	birthDate           time.Time
	town, townCode      string
	residence           string
	fiscalCode          string
	phone               string
	id                  string
}

func (p *Person) FirstName() string {
	return p.firstName
}
func (p *Person) LastName() string {
	return p.lastName
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
func (p *Person) Town() string {
	//TODO
	return p.residence
}
func (p *Person) Phone() string {
	//TODO
	return p.phone
}
func (p *Person) ID() string {
	//TODO
	if p.id != "" {
		return p.id
	}
	p.id = fmt.Sprintf("A%s%d", string("QWERTYUIOPASDFGHJKLZXCVBNM"[rand.Int()%26]), rand.Int()%10000000)
	return p.id
}

//TODO test
func (p Person) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString(p.firstName)
	_, _ = buf.WriteString(" ")
	_, _ = buf.WriteString(p.lastName)
	_, _ = buf.WriteString(", ")
	_, _ = buf.WriteString(p.Gender())
	_, _ = buf.WriteString(fmt.Sprintf(" %d/%d/%d ", p.birthDate.Day(), int(p.birthDate.Month()), p.birthDate.Year()))
	_, _ = buf.WriteString(p.town)
	return buf.String()
}

func (p *Person) MarshalJSON() (b []byte, err error) {
	cf, err := p.CodiceFiscale()
	pi, county, err := p.PartitaIva()
	cc, err := p.CartaCredito()
	bd := fmt.Sprintf("%02d/%02d/%04d", p.birthDate.Day(), int(p.birthDate.Month()), p.birthDate.Year())
	//bd := fmt.Sprintf("%02d/%02d/%04d", p.BirthDate().Day, p.BirthDate().Month, p.BirthDate().Year)
	if err != nil {
		return
	}

	var wrapper = struct {
		Nome             string
		Cognome          string
		Gender           string
		PaeseDiNascita   string
		NumeroDiTelefono string
		DataDiNascita    string
		CodiceFiscale    string
		PartitaIva       string
		ComunePartitaIva string
		Documento        string
		CartaCredito     string
	}{
		p.FirstName(),
		p.LastName(),
		p.Gender(),
		p.BirthTown(),
		p.Phone(),
		bd,
		cf,
		pi,
		county,
		p.ID(),
		cc,
	}
	return json.Marshal(wrapper)
}

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
		people = append(people, person)
		count--
	}
	return
}
