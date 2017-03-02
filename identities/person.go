package identities

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/empijei/identigen/identities/lists"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Represents a person object. It must be initialized by the generator.
type Person struct {
	firstName, lastName           string
	genderIsFemale                bool
	birthDate                     time.Time
	town, townCode, birthDistrict string
	residence                     string
	drv                           *DrivingLicense
	fiscalCode                    string
	partitaIva                    string
	locationCode                  int
	partitaIvaCounty              string
	cc                            *CartaCredito
	mobilePhone                   string
	id                            string
	iban                          string
	up                            *Credentials
}

func NewPerson(minage, maxage int) *Person {
	person := &Person{}
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
	return person
}

//Person first name
func (p *Person) FirstName() string {
	return p.firstName
}

//Person last name
func (p *Person) LastName() string {
	return p.lastName
}

//Person gender in Italian
func (p *Person) Gender() string {
	if p.genderIsFemale {
		return "Donna"
	} else {
		return "Uomo"
	}
}

//Person birth date formatted using the globally specified format
func (p *Person) BirthDate() string {
	return p.birthDate.Format(LocalizDate.Format())
}

//Person birth town
func (p *Person) BirthTown() string {
	return p.town
}

//The name and label of the city the person was birth in
func (p *Person) BirthDistrict() string {
	return p.birthDistrict
}

//The phone number (Without the +39 italian prefix)
func (p *Person) Phone() string {
	return p.mobilePhone
}

//Identity card number
func (p *Person) ID() string {
	if p.id != "" {
		return p.id
	}
	p.id = fmt.Sprintf("A%s%d", string("QWERTYUIOPASDFGHJKLZXCVBNM"[rand.Int()%26]), rand.Int()%10000000)
	return p.id
}

//String representation, the human readable serialization of a Person object
func (p Person) String() string {
	m := p.toMap()
	re := regexp.MustCompile("([a-z])([A-Z]+)")
	var buf bytes.Buffer
	for _, field := range fields {
		_, _ = buf.WriteString(re.ReplaceAllString(field, "$1 $2"))
		_, _ = buf.WriteString(": ")
		_, _ = buf.WriteString(m[field])
		_, _ = buf.WriteString(",\n")
	}
	return buf.String()
}

//Implementation of encoding/json.Marshaler
func (p *Person) MarshalJSON() (b []byte, err error) {
	return json.Marshal(p.toMap())
}

//Implementation of encoding/xml.Marshaler
func (p *Person) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	_panic := func(e error) {
		if err != nil {
			panic(err)
		}
	}

	_panic(e.EncodeToken(start))

	for key, value := range p.toMap() {
		_panic(e.EncodeToken(xml.StartElement{Name: xml.Name{Local: key}}))
		_panic(e.EncodeToken(xml.CharData(value)))
		_panic(e.EncodeToken(xml.EndElement{Name: xml.Name{Local: key}}))
	}

	_panic(e.EncodeToken(xml.EndElement{start.Name}))

	// flush to ensure tokens are written
	return e.Flush()
}

//Returns a []string that can be passed to an encoding/csv.Writer.Write() call
func (p Person) MarshalCSV() []string {
	m := p.toMap()
	var out []string
	for _, f := range fields {
		out = append(out, m[f])
	}
	return out
}

func (p *Person) toMap() map[string]string {
	toret := make(map[string]string)
	for _, f := range fields {
		toret[f] = printerMap[f](p)
	}
	return toret
}
